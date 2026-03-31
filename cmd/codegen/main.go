package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// ============================================================
// API type mismatch overrides — real API returns different types than spec
// ============================================================

var fieldTypeOverrides = map[string]string{
	// spec says integer, API returns float (e.g. 620.8)
	"priceWithSellerFee":   "*float64",
	"roblox_credit_balance": "*float64",
	// spec says string, API returns dict {"730": "CS2 Prime"}
	"steam_bans": "interface{}",
	// spec says boolean, API returns object {"type": "...", ...}
	"guarantee": "interface{}",
	// spec says array, API returns dict
	"cs2PremierElo": "interface{}",
	// spec says integer, API returns string "none"
	"discordNitroType":  "interface{}",
	"discord_nitro_type": "interface{}",
	// spec says string, API returns integer
	"instagram_id": "interface{}",
	// spec says list of objects, API returns JSON string
	"socialclub_games": "interface{}",
	// dict or list — PHP API is unpredictable
	"thread_tags":        "interface{}",
	"Skin":               "interface{}",
	"WeaponSkins":        "interface{}",
	"supercellBrawlers":  "interface{}",
	"r6Skins":            "interface{}",
	"tags":               "interface{}",
	"values":             "interface{}",
	"base_params":        "interface{}",
	// These fields return [] (empty JSON array) instead of {} or null
	"feedback_data":    "interface{}",
	"imap_data":        "interface{}",
	"restore_data":     "interface{}",
	"telegram_client":  "interface{}",
	"backgrounds":      "interface{}",
	// Market category search: game inventories can be [] or {}
	"battlenetGames":        "interface{}",
	"eg_games":              "interface{}",
	"lolInventory":          "interface{}",
	"wotTopPremiumTanks":    "interface{}",
	"wotTopTanks":           "interface{}",
	"valorantInventory":     "interface{}",
	// Forum Users Contents: content_id can be string or int
	"content_id":       "interface{}",
}

// modelFieldTypeOverrides provides overrides for specific model+field combinations
// where the field name is too generic for a global override.
// Key format: "ModelName.json_field_name"
var modelFieldTypeOverrides = map[string]string{
	// Forum: GroupedResponse.data returns [] when empty
	"GroupedResponse.data":                             "interface{}",
	// Forum: EditResponse.message returns string for profile post edit
	"EditResponse.message":                             "interface{}",
	// Market: SteamValueResponseData.items returns [] when empty
	"SteamValueResponseData.items":                     "interface{}",
	// Market: SteamResponseItemsItemSteamFullGames.list returns [] when empty
	"SteamResponseItemsItemSteamFullGames.list":         "interface{}",
	// Market: FeeResponse.calculator returns [] when empty
	"FeeResponse.calculator":                            "interface{}",
	// Market: ListResponse.payments returns [] when empty
	"ListResponse.payments":                             "interface{}",
	// Market: BalanceExchange.from and .to return [] when empty
	"BalanceExchange.from":                              "interface{}",
	"BalanceExchange.to":                                "interface{}",
	// Market: PayoutServicesResponseSystemsItem.providers returns [] when empty
	"PayoutServicesResponseSystemsItem.providers":       "interface{}",
}

// ============================================================
// OpenAPI 3.1 types (subset we need)
// ============================================================

type OpenAPISpec struct {
	Paths      map[string]PathItem            `json:"paths"`
	Components Components                     `json:"components"`
}

type Components struct {
	Schemas    map[string]*Schema             `json:"schemas"`
	Parameters map[string]*Parameter          `json:"parameters"`
	Responses  map[string]*Response           `json:"responses"`
}

type PathItem map[string]json.RawMessage // method -> Operation | other

type Operation struct {
	OperationID string                      `json:"operationId"`
	Tags        []string                    `json:"tags"`
	Summary     string                      `json:"summary"`
	Description string                      `json:"description"`
	Parameters  []Parameter                 `json:"parameters"`
	RequestBody *RequestBody                `json:"requestBody"`
	Responses   map[string]json.RawMessage  `json:"responses"`
}

type Parameter struct {
	Ref         string  `json:"$ref"`
	Name        string  `json:"name"`
	In          string  `json:"in"`
	Description string  `json:"description"`
	Required    bool    `json:"required"`
	Schema      *Schema `json:"schema"`
	Style       string  `json:"style"`
	Explode     *bool   `json:"explode"`
}

type RequestBody struct {
	Content map[string]MediaType `json:"content"`
}

type Response struct {
	Ref         string               `json:"$ref"`
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content"`
}

type MediaType struct {
	Schema *Schema `json:"schema"`
}

type Schema struct {
	Ref                  string             `json:"$ref"`
	Type                 interface{}        `json:"type"` // string or []string
	Format               string             `json:"format"`
	Description          string             `json:"description"`
	Enum                 []interface{}      `json:"enum"`
	Properties           map[string]*Schema `json:"properties"`
	Items                *Schema            `json:"items"`
	Required             []string           `json:"required"`
	AdditionalProperties interface{}        `json:"additionalProperties"`
	OneOf                []*Schema          `json:"oneOf"`
	AnyOf                []*Schema          `json:"anyOf"`
	AllOf                []*Schema          `json:"allOf"`
	Title                string             `json:"title"`
}

// SchemaType returns the type as a string. Handles ["string","integer"] arrays.
func (s *Schema) SchemaType() string {
	if s == nil {
		return ""
	}
	switch t := s.Type.(type) {
	case string:
		return t
	case []interface{}:
		// Multi-type: check if it's ["string", "integer"] etc
		var types []string
		for _, v := range t {
			if str, ok := v.(string); ok {
				if str != "null" {
					types = append(types, str)
				}
			}
		}
		if len(types) == 1 {
			return types[0]
		}
		if len(types) >= 2 {
			return "multi:" + strings.Join(types, "|")
		}
	}
	return ""
}

// ============================================================
// Generator context
// ============================================================

type Generator struct {
	spec       *OpenAPISpec
	pkgName    string
	outputDir  string

	// Collected data
	models     map[string]*ModelDef // name -> model
	enums      map[string]*EnumDef  // type name -> enum
	services   map[string]*ServiceDef

	// For cycle protection in ref resolution
	resolving map[string]bool
}

type ModelDef struct {
	Name   string
	Fields []FieldDef
	Doc    string
}

type FieldDef struct {
	Name     string // Go name
	JSONName string // original JSON name
	Type     string // Go type
	Doc      string
}

type EnumDef struct {
	TypeName string
	Values   []EnumValue
}

type EnumValue struct {
	ConstName string
	Value     string
}

type ServiceDef struct {
	Name    string // Go-friendly tag name
	Methods []MethodDef
}

type MethodDef struct {
	Name           string
	HTTPMethod     string
	Path           string
	Doc            string
	PathParams     []ParamDef
	QueryParams    []ParamDef
	BodyParams     []ParamDef
	ResponseType   string
	ParamsStruct   string // name of optional params struct, empty if none
	IsMultipart    bool
	IsSearch       bool
	ContentType    string
	IsArrayBody    bool
	ResponseIsText bool
}

type ParamDef struct {
	Name        string // Go name
	JSONName    string // original API name
	Type        string // Go type
	Required    bool
	Description string
	IsDeepObj   bool
	IsArray     bool
}

func NewGenerator(spec *OpenAPISpec, pkgName, outputDir string) *Generator {
	return &Generator{
		spec:      spec,
		pkgName:   pkgName,
		outputDir: outputDir,
		models:    make(map[string]*ModelDef),
		enums:     make(map[string]*EnumDef),
		services:  make(map[string]*ServiceDef),
		resolving: make(map[string]bool),
	}
}

// ============================================================
// Ref resolution
// ============================================================

func (g *Generator) resolveSchema(s *Schema) *Schema {
	if s == nil {
		return nil
	}
	if s.Ref != "" {
		return g.resolveRef(s.Ref)
	}
	return s
}

func (g *Generator) resolveRef(ref string) *Schema {
	// #/components/schemas/Foo
	parts := strings.Split(ref, "/")
	if len(parts) < 4 {
		return nil
	}
	kind := parts[2] // "schemas" or "parameters"
	name := parts[3]

	if kind == "schemas" {
		if s, ok := g.spec.Components.Schemas[name]; ok {
			return s
		}
	}
	return nil
}

func (g *Generator) resolveParam(p Parameter) Parameter {
	if p.Ref != "" {
		parts := strings.Split(p.Ref, "/")
		if len(parts) >= 4 && parts[2] == "parameters" {
			name := parts[3]
			if cp, ok := g.spec.Components.Parameters[name]; ok {
				return *cp
			}
		}
	}
	return p
}

// resolveResponse resolves a response-level $ref to the actual Response object.
func (g *Generator) resolveResponse(resp *Response) *Response {
	if resp == nil {
		return nil
	}
	if resp.Ref != "" {
		// e.g. "#/components/responses/SaveChanges"
		parts := strings.Split(resp.Ref, "/")
		if len(parts) >= 4 && parts[2] == "responses" {
			name := parts[3]
			if r, ok := g.spec.Components.Responses[name]; ok {
				return r
			}
		}
		return nil
	}
	return resp
}

// refName extracts the schema name from a $ref string
func refName(ref string) string {
	parts := strings.Split(ref, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

// ============================================================
// Naming helpers
// ============================================================

var goReserved = map[string]bool{
	"break": true, "default": true, "func": true, "interface": true, "select": true,
	"case": true, "defer": true, "go": true, "map": true, "struct": true,
	"chan": true, "else": true, "goto": true, "package": true, "switch": true,
	"const": true, "fallthrough": true, "if": true, "range": true, "type": true,
	"continue": true, "for": true, "import": true, "return": true, "var": true,
}

func safeGoName(name string) string {
	if name == "" {
		return "Field"
	}
	// If it starts with a digit, prefix with "Field"
	if name[0] >= '0' && name[0] <= '9' {
		name = "Field" + name
	}
	if goReserved[strings.ToLower(name)] {
		return name + "_"
	}
	return name
}

// toPascalCase converts snake_case, kebab-case, dot.case to PascalCase
func toPascalCase(s string) string {
	if s == "" {
		return s
	}
	// Common acronyms
	acronyms := map[string]string{
		"id": "ID", "url": "URL", "uri": "URI", "html": "HTML",
		"css": "CSS", "http": "HTTP", "https": "HTTPS", "api": "API",
		"ip": "IP", "json": "JSON", "xml": "XML", "sql": "SQL",
		"sda": "SDA", "mfa": "MFA", "imap": "IMAP", "ea": "EA",
		"wot": "WOT",
	}

	// Split on non-alphanumeric
	parts := regexp.MustCompile(`[^a-zA-Z0-9]+`).Split(s, -1)
	var result strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		lower := strings.ToLower(part)
		if acronym, ok := acronyms[lower]; ok {
			result.WriteString(acronym)
		} else {
			// Preserve existing casing, just ensure first letter is upper
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			result.WriteString(string(runes))
		}
	}
	return result.String()
}

// toCamelCase like toPascalCase but first letter lowercase.
// Result is also made safe for Go reserved words.
func toCamelCase(s string) string {
	p := toPascalCase(s)
	if p == "" {
		return p
	}
	var result string
	// Handle acronyms at start (e.g. "ID" -> "id", "URL" -> "url")
	runes := []rune(p)
	if len(runes) > 1 && unicode.IsUpper(runes[1]) {
		// All-upper prefix: lowercase all leading uppercase
		i := 0
		for i < len(runes) && unicode.IsUpper(runes[i]) {
			i++
		}
		if i == len(runes) {
			result = strings.ToLower(p)
		} else {
			// lowercase all but last uppercase
			for j := 0; j < i-1; j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			result = string(runes)
		}
	} else {
		runes[0] = unicode.ToLower(runes[0])
		result = string(runes)
	}
	if goReserved[result] {
		result = result + "_"
	}
	// If starts with digit, prefix
	if len(result) > 0 && result[0] >= '0' && result[0] <= '9' {
		result = "v" + result
	}
	return result
}

// cleanTagName converts an API tag to a Go-friendly service name
func cleanTagName(tag string) string {
	// "Post comments" -> "PostComments", "Profile Post Comments" -> "ProfilePostComments"
	return toPascalCase(tag)
}

// operationMethodName extracts a method name from operationId
// e.g. "Threads.List" -> "List", "OAuth.Token" -> "Token"
func operationMethodName(opID string) string {
	parts := strings.Split(opID, ".")
	if len(parts) >= 2 {
		return toPascalCase(parts[len(parts)-1])
	}
	return toPascalCase(opID)
}

// ============================================================
// Dynamic dict detection
// ============================================================

var numericKeyRe = regexp.MustCompile(`^\d+$`)

// isDynamicDict returns true if the schema is an object whose property keys
// are ALL purely numeric — indicating example data from a dynamic dictionary
// (keyed by IDs), not a real typed schema.
func isDynamicDict(s *Schema) bool {
	if s == nil {
		return false
	}
	if s.Properties == nil || len(s.Properties) == 0 {
		return false
	}
	for key := range s.Properties {
		if !numericKeyRe.MatchString(key) {
			return false
		}
	}
	return true
}

// ============================================================
// Go type mapping
// ============================================================

func (g *Generator) goTypeForSchema(s *Schema, nameHint string, nullable bool) string {
	if s == nil {
		return "interface{}"
	}

	// Handle $ref
	if s.Ref != "" {
		rn := refName(s.Ref)
		goName := cleanSchemaName(rn)
		resolved := g.resolveRef(s.Ref)
		if resolved != nil {
			rt := resolved.SchemaType()
			// If it's a simple type alias (not object), use native type
			if rt == "integer" || rt == "number" || rt == "string" || rt == "boolean" {
				// Check if it's an enum
				if len(resolved.Enum) > 0 {
					return goName
				}
				// Check for multi-type
				if strings.HasPrefix(rt, "multi:") {
					return "StringOrInt"
				}
				return g.primitiveType(rt, resolved.Format)
			}
			if rt == "object" {
				// Dynamic dict detection: all-numeric property keys → map
				if isDynamicDict(resolved) {
					return "map[string]interface{}"
				}
				g.ensureModel(goName, resolved)
				if nullable {
					return "*" + goName
				}
				return goName
			}
		}
		// For multi-type refs like UserIDModel
		if resolved != nil {
			rt := resolved.SchemaType()
			if strings.HasPrefix(rt, "multi:") {
				return "StringOrInt"
			}
		}
		return "interface{}"
	}

	schemaType := s.SchemaType()

	// Multi-type: ["string", "integer"]
	if strings.HasPrefix(schemaType, "multi:") {
		return "StringOrInt"
	}

	// Handle oneOf / anyOf — use interface{}
	if len(s.OneOf) > 0 || len(s.AnyOf) > 0 {
		return "interface{}"
	}

	// Handle allOf — merge properties
	if len(s.AllOf) > 0 {
		merged := g.mergeAllOf(s.AllOf)
		return g.goTypeForSchema(merged, nameHint, nullable)
	}

	switch schemaType {
	case "string":
		if s.Format == "binary" {
			return "[]byte"
		}
		if len(s.Enum) > 0 && nameHint != "" {
			return "string" // enums are collected separately, use string for fields
		}
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		itemsResolved := g.resolveSchema(s.Items)
		// If array items are a dynamic-dict object, the whole field should be
		// flexible because PHP APIs serialize numeric-keyed arrays as either
		// JSON objects or JSON arrays unpredictably.
		if isDynamicDict(itemsResolved) {
			return "interface{}"
		}
		itemType := g.goTypeForSchema(itemsResolved, nameHint+"Item", false)
		return "[]" + itemType
	case "object":
		// Dynamic dict detection: all-numeric property keys → map
		if isDynamicDict(s) {
			return "map[string]interface{}"
		}
		if s.Properties != nil && len(s.Properties) > 0 {
			// Create an inline struct
			if nameHint != "" {
				g.ensureModel(nameHint, s)
				if nullable {
					return "*" + nameHint
				}
				return nameHint
			}
		}
		// additionalProperties or empty object
		if s.AdditionalProperties != nil {
			switch ap := s.AdditionalProperties.(type) {
			case map[string]interface{}:
				// Has a schema for values
				apBytes, _ := json.Marshal(ap)
				var apSchema Schema
				json.Unmarshal(apBytes, &apSchema)
				valType := g.goTypeForSchema(&apSchema, "", false)
				return "map[string]" + valType
			case bool:
				if ap {
					return "map[string]interface{}"
				}
			}
		}
		return "map[string]interface{}"
	default:
		if schemaType == "" {
			// No type specified; check properties
			if s.Properties != nil && len(s.Properties) > 0 {
				if nameHint != "" {
					g.ensureModel(nameHint, s)
					if nullable {
						return "*" + nameHint
					}
					return nameHint
				}
			}
			return "interface{}"
		}
		return "interface{}"
	}
}

func (g *Generator) primitiveType(schemaType, format string) string {
	switch schemaType {
	case "string":
		if format == "binary" {
			return "[]byte"
		}
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	}
	return "interface{}"
}

func (g *Generator) mergeAllOf(schemas []*Schema) *Schema {
	merged := &Schema{
		Type:       "object",
		Properties: make(map[string]*Schema),
	}
	for _, s := range schemas {
		resolved := g.resolveSchema(s)
		if resolved == nil {
			continue
		}
		for k, v := range resolved.Properties {
			merged.Properties[k] = v
		}
		merged.Required = append(merged.Required, resolved.Required...)
	}
	return merged
}

func cleanSchemaName(name string) string {
	// Resp_UserModel -> RespUserModel, etc.
	name = strings.ReplaceAll(name, "_", "")
	name = strings.ReplaceAll(name, " ", "")
	// Ensure first letter is upper
	if len(name) > 0 {
		runes := []rune(name)
		runes[0] = unicode.ToUpper(runes[0])
		name = string(runes)
	}
	return name
}

// ============================================================
// Model generation from schemas
// ============================================================

func (g *Generator) ensureModel(name string, s *Schema) {
	if _, exists := g.models[name]; exists {
		return
	}
	if g.resolving[name] {
		return // cycle protection
	}
	g.resolving[name] = true
	defer delete(g.resolving, name)

	s = g.resolveSchema(s)
	if s == nil {
		return
	}

	// Handle allOf
	if len(s.AllOf) > 0 {
		s = g.mergeAllOf(s.AllOf)
	}

	// Dynamic dict detection: skip model creation for schemas with all-numeric keys
	if isDynamicDict(s) {
		return
	}

	props := s.Properties
	if props == nil || len(props) == 0 {
		return
	}

	model := &ModelDef{
		Name: name,
		Doc:  s.Description,
	}

	// Sort property names for deterministic output
	var propNames []string
	for k := range props {
		propNames = append(propNames, k)
	}
	sort.Strings(propNames)

	requiredSet := make(map[string]bool)
	for _, r := range s.Required {
		requiredSet[r] = true
	}

	seenFields := make(map[string]bool)
	for _, propName := range propNames {
		propSchema := props[propName]
		resolved := g.resolveSchema(propSchema)

		// Determine Go type
		hint := name + toPascalCase(propName)
		goType := g.goTypeForSchema(resolved, hint, false)

		// Use the original $ref name for objects if applicable
		if propSchema != nil && propSchema.Ref != "" {
			rn := refName(propSchema.Ref)
			goRefName := cleanSchemaName(rn)
			if resolved != nil && (resolved.SchemaType() == "object" || (resolved.Properties != nil && len(resolved.Properties) > 0)) {
				// Skip dynamic dict schemas — use map instead of typed model
				if isDynamicDict(resolved) {
					goType = "map[string]interface{}"
				} else {
					g.ensureModel(goRefName, resolved)
					goType = "*" + goRefName
				}
			}
		}

		// Apply field-level type overrides for API mismatches
		// Check model-qualified override first (ModelName.field_name)
		qualifiedKey := name + "." + propName
		if override, ok := modelFieldTypeOverrides[qualifiedKey]; ok {
			goType = override
		} else if override, ok := fieldTypeOverrides[propName]; ok {
			goType = override
		} else {
			// Make pointer if not already a pointer, slice, or map
			goType = pointerize(goType)
		}

		goFieldName := safeGoName(toPascalCase(propName))

		// Skip duplicate field names (different JSON keys mapping to same Go name)
		if seenFields[goFieldName] {
			continue
		}
		seenFields[goFieldName] = true

		doc := ""
		if resolved != nil {
			doc = resolved.Description
		}

		model.Fields = append(model.Fields, FieldDef{
			Name:     goFieldName,
			JSONName: propName,
			Type:     goType,
			Doc:      doc,
		})
	}

	g.models[name] = model
}

// pointerize wraps a type in a pointer unless it's already a pointer, slice, map, or interface{}
func pointerize(goType string) string {
	if strings.HasPrefix(goType, "*") ||
		strings.HasPrefix(goType, "[]") ||
		strings.HasPrefix(goType, "map[") ||
		goType == "interface{}" ||
		goType == "StringOrInt" {
		return goType
	}
	return "*" + goType
}

// ============================================================
// Enum collection
// ============================================================

func (g *Generator) collectEnum(contextName string, paramName string, values []interface{}, schemaType string) string {
	typeName := contextName + toPascalCase(paramName)
	if _, exists := g.enums[typeName]; exists {
		return typeName
	}

	edef := &EnumDef{TypeName: typeName}
	seen := make(map[string]bool)
	for _, v := range values {
		strVal := fmt.Sprintf("%v", v)
		suffix := toPascalCase(strVal)
		if suffix == "" {
			suffix = "Empty"
		}
		constName := typeName + suffix
		// Prefix numeric-leading names
		if len(constName) > 0 && constName[0] >= '0' && constName[0] <= '9' {
			constName = typeName + "Value" + strVal
		}
		// Ensure no collision with type name itself
		if constName == typeName {
			constName = typeName + "Default"
		}
		// Deduplicate
		base := constName
		for i := 2; seen[constName]; i++ {
			constName = fmt.Sprintf("%s%d", base, i)
		}
		seen[constName] = true
		edef.Values = append(edef.Values, EnumValue{
			ConstName: constName,
			Value:     strVal,
		})
	}
	g.enums[typeName] = edef
	return typeName
}

// ============================================================
// Process the spec
// ============================================================

func (g *Generator) Process() {
	// 1. Process component schemas (sorted for deterministic output)
	var schemaNames []string
	for name := range g.spec.Components.Schemas {
		schemaNames = append(schemaNames, name)
	}
	sort.Strings(schemaNames)

	for _, name := range schemaNames {
		schema := g.spec.Components.Schemas[name]
		goName := cleanSchemaName(name)
		resolved := g.resolveSchema(schema)
		if resolved == nil {
			continue
		}
		rt := resolved.SchemaType()

		// Enum component schemas
		if len(resolved.Enum) > 0 && (rt == "string" || rt == "integer") {
			g.collectEnum("", goName, resolved.Enum, rt)
			continue
		}

		// Multi-type (like UserIDModel: ["string","integer"]) - handled by StringOrInt
		if strings.HasPrefix(rt, "multi:") {
			continue
		}

		// Simple type aliases (integer, string) - skip as models
		if rt == "integer" || rt == "string" || rt == "number" || rt == "boolean" {
			continue
		}

		if resolved.Properties != nil && len(resolved.Properties) > 0 {
			// Skip dynamic dict schemas (all-numeric property keys)
			if !isDynamicDict(resolved) {
				g.ensureModel(goName, resolved)
			}
		}
	}

	// 2. Process all paths (sorted for deterministic output)
	var pathKeys []string
	for path := range g.spec.Paths {
		pathKeys = append(pathKeys, path)
	}
	sort.Strings(pathKeys)

	for _, path := range pathKeys {
		pathItem := g.spec.Paths[path]
		// Sort methods within each path for deterministic output
		var methods []string
		for method := range pathItem {
			methods = append(methods, method)
		}
		sort.Strings(methods)

		for _, method := range methods {
			rawOp := pathItem[method]
			if method == "$ref" || method == "summary" || method == "description" || method == "parameters" || method == "servers" {
				continue
			}

			var op Operation
			if err := json.Unmarshal(rawOp, &op); err != nil {
				log.Printf("WARN: cannot parse operation %s %s: %v", method, path, err)
				continue
			}

			g.processOperation(method, path, &op)
		}
	}
}

func (g *Generator) processOperation(httpMethod, path string, op *Operation) {
	if op.OperationID == "" {
		return
	}

	tag := "Default"
	if len(op.Tags) > 0 {
		tag = op.Tags[0]
	}
	serviceName := cleanTagName(tag)

	if _, ok := g.services[serviceName]; !ok {
		g.services[serviceName] = &ServiceDef{Name: serviceName}
	}

	methodName := operationMethodName(op.OperationID)

	// Check for duplicate method names within the same service
	for _, existingMethod := range g.services[serviceName].Methods {
		if existingMethod.Name == methodName {
			// Append the HTTP method to disambiguate
			methodName = methodName + toPascalCase(httpMethod)
			break
		}
	}

	// Determine response type
	respType := g.processResponseSchema(op, methodName)
	// Detect text/html responses
	responseIsText := g.isTextHTMLResponse(op)

	// Process parameters
	var pathParams, queryParams, bodyParams []ParamDef

	for _, rawParam := range op.Parameters {
		p := g.resolveParam(rawParam)
		if p.Name == "" || p.Schema == nil {
			continue
		}

		resolved := g.resolveSchema(p.Schema)
		goType := g.goTypeForSchema(resolved, "", false)
		if goType == "interface{}" && resolved != nil {
			st := resolved.SchemaType()
			if st != "" {
				goType = g.primitiveType(st, resolved.Format)
			}
		}

		isArray := false
		if resolved != nil && resolved.SchemaType() == "array" {
			isArray = true
		}

		// Collect enums from parameters
		if resolved != nil && len(resolved.Enum) > 0 {
			enumTypeName := g.collectEnum(serviceName, p.Name, resolved.Enum, resolved.SchemaType())
			goType = enumTypeName
		}

		// Strip [] suffix from param name for Go name
		cleanName := strings.TrimSuffix(p.Name, "[]")

		pd := ParamDef{
			Name:        safeGoName(toPascalCase(cleanName)),
			JSONName:    p.Name,
			Type:        goType,
			Required:    p.Required,
			Description: p.Description,
			IsDeepObj:   p.Style == "deepObject",
			IsArray:     isArray,
		}

		switch p.In {
		case "path":
			pathParams = append(pathParams, pd)
		case "query":
			queryParams = append(queryParams, pd)
		}
	}

	// Process request body
	isMultipart := false
	contentType := ""
	isArrayBody := false
	if op.RequestBody != nil {
		// Pick content type deterministically: prefer multipart, then JSON, then first sorted key
		var chosenCT string
		for ct := range op.RequestBody.Content {
			if strings.Contains(ct, "multipart") {
				chosenCT = ct
				break
			}
		}
		if chosenCT == "" {
			for ct := range op.RequestBody.Content {
				if strings.Contains(ct, "json") {
					chosenCT = ct
					break
				}
			}
		}
		if chosenCT == "" {
			var ctKeys []string
			for ct := range op.RequestBody.Content {
				ctKeys = append(ctKeys, ct)
			}
			sort.Strings(ctKeys)
			if len(ctKeys) > 0 {
				chosenCT = ctKeys[0]
			}
		}
		if chosenCT != "" {
			mt := op.RequestBody.Content[chosenCT]
			contentType = chosenCT
			if strings.Contains(chosenCT, "multipart") {
				isMultipart = true
			}
			if mt.Schema != nil {
				resolved := g.resolveSchema(mt.Schema)
				// Check if body schema is an array (e.g. batch endpoints)
				if resolved != nil && resolved.SchemaType() == "array" {
					isArrayBody = true
				}
				if resolved != nil && resolved.Properties != nil {
					for propName, propSchema := range resolved.Properties {
						propResolved := g.resolveSchema(propSchema)
						goType := g.goTypeForSchema(propResolved, "", false)
						if propResolved != nil && propResolved.Format == "binary" {
							goType = "[]byte"
						}
						// Collect enums from body params
						if propResolved != nil && len(propResolved.Enum) > 0 {
							goType = g.collectEnum(serviceName+methodName, propName, propResolved.Enum, propResolved.SchemaType())
						}

						isRequired := false
						for _, r := range resolved.Required {
							if r == propName {
								isRequired = true
								break
							}
						}

						desc := ""
						if propResolved != nil {
							desc = propResolved.Description
						}

						bodyParams = append(bodyParams, ParamDef{
							Name:        safeGoName(toPascalCase(propName)),
							JSONName:    propName,
							Type:        goType,
							Required:    isRequired,
							Description: desc,
						})
					}
				}
				// Also handle oneOf for body (like OAuth)
				if len(resolved.OneOf) > 0 {
					// Collect all properties from all oneOf variants
					seen := make(map[string]bool)
					for _, variant := range resolved.OneOf {
						vResolved := g.resolveSchema(variant)
						if vResolved == nil || vResolved.Properties == nil {
							continue
						}
						for propName, propSchema := range vResolved.Properties {
							if seen[propName] {
								continue
							}
							seen[propName] = true
							propResolved := g.resolveSchema(propSchema)
							goType := g.goTypeForSchema(propResolved, "", false)
							if propResolved != nil && len(propResolved.Enum) > 0 {
								goType = g.collectEnum(serviceName+methodName, propName, propResolved.Enum, propResolved.SchemaType())
							}

							desc := ""
							if propResolved != nil {
								desc = propResolved.Description
							}

							bodyParams = append(bodyParams, ParamDef{
								Name:        safeGoName(toPascalCase(propName)),
								JSONName:    propName,
								Type:        goType,
								Description: desc,
							})
						}
					}
				}
			}
		}
	}

	// Sort body params for deterministic output
	sort.Slice(bodyParams, func(i, j int) bool {
		return bodyParams[i].JSONName < bodyParams[j].JSONName
	})

	// Determine if search endpoint
	isSearch := strings.Contains(path, "/search")

	// Build params struct name
	paramsStructName := ""
	// Collect optional params (query non-required + body non-required)
	var optionalParams []ParamDef
	for _, p := range queryParams {
		if !p.Required {
			optionalParams = append(optionalParams, p)
		} else {
			// Required query params also go as function args
			pathParams = append(pathParams, p)
		}
	}
	for _, p := range bodyParams {
		if !p.Required {
			optionalParams = append(optionalParams, p)
		} else {
			pathParams = append(pathParams, p)
		}
	}

	if len(optionalParams) > 0 {
		paramsStructName = methodName + "Params"
		// Create params model
		pModel := &ModelDef{Name: paramsStructName}
		for _, p := range optionalParams {
			goType := pointerize(p.Type)
			pModel.Fields = append(pModel.Fields, FieldDef{
				Name:     p.Name,
				JSONName: p.JSONName,
				Type:     goType,
				Doc:      p.Description,
			})
		}
		g.models[paramsStructName] = pModel
	}

	doc := op.Summary
	if doc == "" {
		doc = op.Description
	}

	if responseIsText {
		respType = "string"
	}

	md := MethodDef{
		Name:           methodName,
		HTTPMethod:     strings.ToUpper(httpMethod),
		Path:           path,
		Doc:            doc,
		PathParams:     pathParams,
		QueryParams:    queryParams,
		BodyParams:     bodyParams,
		ResponseType:   respType,
		ParamsStruct:   paramsStructName,
		IsMultipart:    isMultipart,
		IsSearch:       isSearch,
		ContentType:    contentType,
		IsArrayBody:    isArrayBody,
		ResponseIsText: responseIsText,
	}

	g.services[serviceName].Methods = append(g.services[serviceName].Methods, md)
}

func (g *Generator) processResponseSchema(op *Operation, methodName string) string {
	// Sort response codes for deterministic output
	var respCodes []string
	for code := range op.Responses {
		respCodes = append(respCodes, code)
	}
	sort.Strings(respCodes)

	for _, code := range respCodes {
		rawResp := op.Responses[code]
		if code != "200" && code != "201" {
			continue
		}

		var resp Response
		if err := json.Unmarshal(rawResp, &resp); err != nil {
			continue
		}

		// Resolve response-level $ref (e.g. "$ref": "#/components/responses/SaveChanges")
		resolved := g.resolveResponse(&resp)
		if resolved == nil {
			continue
		}

		// Sort content types for deterministic output
		var respCTKeys []string
		for ct := range resolved.Content {
			respCTKeys = append(respCTKeys, ct)
		}
		sort.Strings(respCTKeys)

		for _, ctKey := range respCTKeys {
			mt := resolved.Content[ctKey]
			if mt.Schema == nil {
				continue
			}
			s := mt.Schema
			// If it's a ref to a component schema
			if s.Ref != "" {
				rn := refName(s.Ref)
				goName := cleanSchemaName(rn)
				resolvedSchema := g.resolveRef(s.Ref)
				if resolvedSchema != nil {
					g.ensureModel(goName, resolvedSchema)
				}
				return goName
			}

			// Inline response schema: use response component name if from a $ref
			respTypeName := methodName + "Response"
			if resp.Ref != "" {
				rn := refName(resp.Ref)
				respTypeName = cleanSchemaName(rn)
			}

			resolvedSchema := g.resolveSchema(s)
			if resolvedSchema != nil && (resolvedSchema.Properties != nil || len(resolvedSchema.AllOf) > 0) {
				g.ensureModel(respTypeName, resolvedSchema)
				return respTypeName
			}

			// Simple type response
			return g.goTypeForSchema(resolvedSchema, respTypeName, false)
		}
	}
	// No response schema found
	return "map[string]interface{}"
}

// isTextHTMLResponse checks if the operation's 200 response is text/html (not JSON).
func (g *Generator) isTextHTMLResponse(op *Operation) bool {
	for code, rawResp := range op.Responses {
		if code != "200" {
			continue
		}
		var resp Response
		if err := json.Unmarshal(rawResp, &resp); err != nil {
			continue
		}
		resolved := g.resolveResponse(&resp)
		if resolved == nil {
			continue
		}
		_, hasTextHTML := resolved.Content["text/html"]
		_, hasJSON := resolved.Content["application/json"]
		if hasTextHTML && !hasJSON {
			return true
		}
	}
	return false
}

// ============================================================
// Code writers
// ============================================================

func (g *Generator) writeModels() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("package %s\n\n", g.pkgName))

	// Write imports
	b.WriteString("import (\n")
	b.WriteString("\t\"encoding/json\"\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"strconv\"\n")
	b.WriteString(")\n\n")

	// StringOrInt type
	b.WriteString(`// StringOrInt handles JSON fields that can be either a string or an integer.
type StringOrInt struct {
	StringValue *string
	IntValue    *int
}

func (s StringOrInt) MarshalJSON() ([]byte, error) {
	if s.IntValue != nil {
		return json.Marshal(*s.IntValue)
	}
	if s.StringValue != nil {
		return json.Marshal(*s.StringValue)
	}
	return []byte("null"), nil
}

func (s *StringOrInt) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch v := raw.(type) {
	case float64:
		i := int(v)
		s.IntValue = &i
	case string:
		s.StringValue = &v
	case nil:
		// both remain nil
	default:
		return fmt.Errorf("StringOrInt: unexpected type %T", raw)
	}
	return nil
}

func (s StringOrInt) String() string {
	if s.IntValue != nil {
		return strconv.Itoa(*s.IntValue)
	}
	if s.StringValue != nil {
		return *s.StringValue
	}
	return ""
}

`)

	// Sort model names for deterministic output
	var names []string
	for name := range g.models {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		model := g.models[name]
		if model.Doc != "" {
			b.WriteString(fmt.Sprintf("// %s %s\n", name, cleanComment(model.Doc)))
		}
		b.WriteString(fmt.Sprintf("type %s struct {\n", name))
		for _, f := range model.Fields {
			if f.Doc != "" {
				b.WriteString(fmt.Sprintf("\t// %s\n", cleanComment(f.Doc)))
			}
			b.WriteString(fmt.Sprintf("\t%s %s `json:\"%s,omitempty\"`\n", f.Name, f.Type, f.JSONName))
		}
		b.WriteString("}\n\n")
	}

	return b.String()
}

func (g *Generator) writeEnums() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("package %s\n\n", g.pkgName))

	// Sort enum type names
	var names []string
	for name := range g.enums {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		edef := g.enums[name]
		b.WriteString(fmt.Sprintf("type %s string\n\n", edef.TypeName))
		b.WriteString("const (\n")
		for _, v := range edef.Values {
			b.WriteString(fmt.Sprintf("\t%s %s = \"%s\"\n", v.ConstName, edef.TypeName, v.Value))
		}
		b.WriteString(")\n\n")
	}

	return b.String()
}

func (g *Generator) writeClient() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("package %s\n\n", g.pkgName))

	// Imports
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n")
	b.WriteString("\t\"encoding/json\"\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"strings\"\n")
	b.WriteString(")\n\n")

	// Requester interface and types
	clientName := toPascalCase(g.pkgName) + "Client"

	b.WriteString(`// Requester is the interface for making HTTP requests.
type Requester interface {
	Request(ctx context.Context, method, path string, opts RequestOptions) (json.RawMessage, error)
}

// RequestOptions holds options for a request.
type RequestOptions struct {
	Params         map[string]interface{}
	JSON           map[string]interface{}
	JSONBody       interface{}
	Files          map[string]FileUpload
	Search         bool
	ForceMultipart bool
}

// FileUpload represents a file to upload.
type FileUpload struct {
	Filename string
	Data     []byte
}

`)

	// Sort services
	var serviceNames []string
	for name := range g.services {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)

	// Client struct
	b.WriteString(fmt.Sprintf("// %s is the API client.\n", clientName))
	b.WriteString(fmt.Sprintf("type %s struct {\n", clientName))
	b.WriteString("\tr Requester\n")
	for _, sName := range serviceNames {
		b.WriteString(fmt.Sprintf("\t%s *%sService\n", sName, sName))
	}
	b.WriteString("}\n\n")

	// New function
	b.WriteString(fmt.Sprintf("// New creates a new %s.\n", clientName))
	b.WriteString(fmt.Sprintf("func New(r Requester) *%s {\n", clientName))
	b.WriteString(fmt.Sprintf("\tc := &%s{r: r}\n", clientName))
	for _, sName := range serviceNames {
		b.WriteString(fmt.Sprintf("\tc.%s = &%sService{r: r}\n", sName, sName))
	}
	b.WriteString("\treturn c\n")
	b.WriteString("}\n\n")

	// Service structs and methods
	for _, sName := range serviceNames {
		svc := g.services[sName]

		b.WriteString(fmt.Sprintf("// %sService handles %s operations.\n", sName, sName))
		b.WriteString(fmt.Sprintf("type %sService struct {\n", sName))
		b.WriteString("\tr Requester\n")
		b.WriteString("}\n\n")

		// Sort methods by name
		sort.Slice(svc.Methods, func(i, j int) bool {
			return svc.Methods[i].Name < svc.Methods[j].Name
		})

		for _, m := range svc.Methods {
			g.writeMethod(&b, sName, &m)
		}
	}

	// Helper to suppress unused import warnings
	b.WriteString("// Ensure imports are used.\n")
	b.WriteString("var _ = fmt.Sprintf\n")
	b.WriteString("var _ = strings.Replace\n")
	b.WriteString("var _ = json.Marshal\n")

	return b.String()
}

func (g *Generator) writeMethod(b *strings.Builder, serviceName string, m *MethodDef) {
	// Doc comment
	if m.Doc != "" {
		b.WriteString(fmt.Sprintf("// %s %s\n", m.Name, cleanComment(m.Doc)))
	}

	// Build function signature
	var args []string
	args = append(args, "ctx context.Context")

	// For array body schemas (e.g. batch), add a jobs parameter
	if m.IsArrayBody {
		args = append(args, "jobs []map[string]interface{}")
	}

	// Path params and required params as direct arguments
	for _, p := range m.PathParams {
		goType := p.Type
		// Path params should not be pointers
		goType = strings.TrimPrefix(goType, "*")
		args = append(args, fmt.Sprintf("%s %s", toCamelCase(p.Name), goType))
	}

	// Optional params struct
	if m.ParamsStruct != "" {
		args = append(args, fmt.Sprintf("params *%s", m.ParamsStruct))
	}

	respType := m.ResponseType
	if m.ResponseIsText {
		b.WriteString(fmt.Sprintf("func (s *%sService) %s(%s) (string, error) {\n",
			serviceName, m.Name, strings.Join(args, ", ")))
	} else if respType == "" || respType == "map[string]interface{}" {
		b.WriteString(fmt.Sprintf("func (s *%sService) %s(%s) (map[string]interface{}, error) {\n",
			serviceName, m.Name, strings.Join(args, ", ")))
	} else {
		b.WriteString(fmt.Sprintf("func (s *%sService) %s(%s) (*%s, error) {\n",
			serviceName, m.Name, strings.Join(args, ", "), respType))
	}

	// Build path with substitutions
	if strings.Contains(m.Path, "{") {
		// Find path params
		re := regexp.MustCompile(`\{([^}]+)\}`)
		matches := re.FindAllStringSubmatch(m.Path, -1)
		b.WriteString(fmt.Sprintf("\tpath := \"%s\"\n", m.Path))
		for _, match := range matches {
			paramName := match[1]
			goArgName := toCamelCase(toPascalCase(paramName))
			// Find the param type
			paramType := "string"
			for _, p := range m.PathParams {
				if strings.EqualFold(p.Name, toPascalCase(paramName)) {
					paramType = strings.TrimPrefix(p.Type, "*")
					break
				}
			}
			if paramType == "int" || paramType == "int64" {
				b.WriteString(fmt.Sprintf("\tpath = strings.Replace(path, \"{%s}\", fmt.Sprintf(\"%%d\", %s), 1)\n", paramName, goArgName))
			} else {
				b.WriteString(fmt.Sprintf("\tpath = strings.Replace(path, \"{%s}\", fmt.Sprintf(\"%%v\", %s), 1)\n", paramName, goArgName))
			}
		}
	} else {
		b.WriteString(fmt.Sprintf("\tpath := \"%s\"\n", m.Path))
	}

	// Build request options
	b.WriteString("\topts := RequestOptions{}\n")

	// Array body (e.g. batch endpoints)
	if m.IsArrayBody {
		b.WriteString("\topts.JSONBody = jobs\n")
	}

	if m.IsSearch {
		b.WriteString("\topts.Search = true\n")
	}

	// Build params map (query params from required args)
	hasQueryOrBody := false
	for _, p := range m.PathParams {
		if p.IsDeepObj {
			hasQueryOrBody = true
			break
		}
		// Check if this was originally a query/body param (not a path param)
		// Path params in URL don't need to go into opts.Params
	}

	// Determine if we use JSON body or query params
	// For multipart, non-binary fields go into opts.JSON which the runtime
	// writes as multipart form fields (not query string).
	useJSON := m.HTTPMethod == "POST" || m.HTTPMethod == "PUT" || m.HTTPMethod == "DELETE"

	// Required query/body params (they are in PathParams after our processing)
	hasRequiredQueryBody := false
	for _, p := range m.PathParams {
		// Skip actual path params (they go in URL)
		isActualPathParam := false
		for _, qp := range m.QueryParams {
			if qp.Name == p.Name {
				isActualPathParam = false
				break
			}
		}
		for _, bp := range m.BodyParams {
			if bp.Name == p.Name {
				isActualPathParam = false
				break
			}
		}
		// If this param is in the original path template, it's a URL param
		if strings.Contains(m.Path, "{"+strings.ToLower(p.JSONName)+"}") || strings.Contains(m.Path, "{"+p.JSONName+"}") {
			isActualPathParam = true
		}
		if !isActualPathParam {
			hasRequiredQueryBody = true
			break
		}
	}

	if hasRequiredQueryBody || hasQueryOrBody || m.ParamsStruct != "" {
		mapField := "Params"
		if useJSON {
			mapField = "JSON"
		}

		b.WriteString(fmt.Sprintf("\topts.%s = make(map[string]interface{})\n", mapField))

		// Required params that aren't path params
		for _, p := range m.PathParams {
			if strings.Contains(m.Path, "{"+p.JSONName+"}") {
				continue // actual path param, skip
			}
			// Also check with original name variants
			found := false
			re := regexp.MustCompile(`\{([^}]+)\}`)
			for _, match := range re.FindAllStringSubmatch(m.Path, -1) {
				if match[1] == p.JSONName || toPascalCase(match[1]) == p.Name {
					found = true
					break
				}
			}
			if found {
				continue
			}

			goArgName := toCamelCase(p.Name)
			if p.IsDeepObj {
				b.WriteString(fmt.Sprintf("\tfor k, v := range %s {\n", goArgName))
				b.WriteString(fmt.Sprintf("\t\topts.%s[fmt.Sprintf(\"%s[%%s]\", k)] = v\n", mapField, strings.TrimSuffix(p.JSONName, "[]")))
				b.WriteString("\t}\n")
			} else {
				b.WriteString(fmt.Sprintf("\topts.%s[\"%s\"] = %s\n", mapField, p.JSONName, goArgName))
			}
		}

		// Optional params struct
		if m.ParamsStruct != "" {
			b.WriteString("\tif params != nil {\n")
			if model, ok := g.models[m.ParamsStruct]; ok {
				for _, f := range model.Fields {
					if strings.HasPrefix(f.Type, "*") {
						b.WriteString(fmt.Sprintf("\t\tif params.%s != nil {\n", f.Name))
						b.WriteString(fmt.Sprintf("\t\t\topts.%s[\"%s\"] = *params.%s\n", mapField, f.JSONName, f.Name))
						b.WriteString("\t\t}\n")
					} else if strings.HasPrefix(f.Type, "[]") {
						b.WriteString(fmt.Sprintf("\t\tif params.%s != nil {\n", f.Name))
						b.WriteString(fmt.Sprintf("\t\t\topts.%s[\"%s\"] = params.%s\n", mapField, f.JSONName, f.Name))
						b.WriteString("\t\t}\n")
					} else if strings.HasPrefix(f.Type, "map[") {
						b.WriteString(fmt.Sprintf("\t\tif params.%s != nil {\n", f.Name))
						// deepObject serialization for maps
						b.WriteString(fmt.Sprintf("\t\t\tfor k, v := range params.%s {\n", f.Name))
						baseKey := strings.TrimSuffix(f.JSONName, "[]")
						b.WriteString(fmt.Sprintf("\t\t\t\topts.%s[fmt.Sprintf(\"%s[%%s]\", k)] = v\n", mapField, baseKey))
						b.WriteString("\t\t\t}\n")
						b.WriteString("\t\t}\n")
					} else {
						b.WriteString(fmt.Sprintf("\t\topts.%s[\"%s\"] = params.%s\n", mapField, f.JSONName, f.Name))
					}
				}
			}
			b.WriteString("\t}\n")
		}

		// Handle multipart file uploads
		if m.IsMultipart {
			hasBinary := false
			for _, p := range m.BodyParams {
				if p.Type == "[]byte" {
					hasBinary = true
					goArgName := toCamelCase(p.Name)
					// Remove binary field from JSON map (already added above)
					b.WriteString(fmt.Sprintf("\tdelete(opts.JSON, \"%s\")\n", p.JSONName))
					b.WriteString("\topts.Files = make(map[string]FileUpload)\n")
					b.WriteString(fmt.Sprintf("\topts.Files[\"%s\"] = FileUpload{Filename: \"%s\", Data: %s}\n",
						p.JSONName, p.JSONName, goArgName))
					break
				}
			}
			if !hasBinary {
				// Multipart without binary files (e.g. oauth/token) — force multipart encoding
				b.WriteString("\topts.ForceMultipart = true\n")
			}
		}
	}

	// Make request
	if m.ResponseIsText {
		b.WriteString(fmt.Sprintf("\traw, err := s.r.Request(ctx, \"%s\", path, opts)\n", m.HTTPMethod))
		b.WriteString("\tif err != nil {\n")
		b.WriteString("\t\treturn \"\", err\n")
		b.WriteString("\t}\n")
		b.WriteString("\treturn string(raw), nil\n")
	} else if respType == "" || respType == "map[string]interface{}" {
		b.WriteString(fmt.Sprintf("\traw, err := s.r.Request(ctx, \"%s\", path, opts)\n", m.HTTPMethod))
		b.WriteString("\tif err != nil {\n")
		b.WriteString("\t\treturn nil, err\n")
		b.WriteString("\t}\n")
		b.WriteString("\tvar result map[string]interface{}\n")
		b.WriteString("\tif err := json.Unmarshal(raw, &result); err != nil {\n")
		b.WriteString("\t\treturn nil, err\n")
		b.WriteString("\t}\n")
		b.WriteString("\treturn result, nil\n")
	} else {
		b.WriteString(fmt.Sprintf("\traw, err := s.r.Request(ctx, \"%s\", path, opts)\n", m.HTTPMethod))
		b.WriteString("\tif err != nil {\n")
		b.WriteString("\t\treturn nil, err\n")
		b.WriteString("\t}\n")
		b.WriteString(fmt.Sprintf("\tvar result %s\n", respType))
		b.WriteString("\tif err := json.Unmarshal(raw, &result); err != nil {\n")
		b.WriteString("\t\treturn nil, err\n")
		b.WriteString("\t}\n")
		b.WriteString("\treturn &result, nil\n")
	}

	b.WriteString("}\n\n")
}

func cleanComment(s string) string {
	// Remove newlines and excessive whitespace
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.TrimSpace(s)
	// Truncate very long comments
	if len(s) > 200 {
		s = s[:200] + "..."
	}
	return s
}

// ============================================================
// Write files
// ============================================================

func (g *Generator) Generate() error {
	if err := os.MkdirAll(g.outputDir, 0o755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	files := map[string]string{
		"models.go": g.writeModels(),
		"enums.go":  g.writeEnums(),
		"client.go": g.writeClient(),
	}

	for name, content := range files {
		path := filepath.Join(g.outputDir, name)
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return fmt.Errorf("writing %s: %w", name, err)
		}

		// Run gofmt
		cmd := exec.Command("gofmt", "-s", "-w", path)
		if out, err := cmd.CombinedOutput(); err != nil {
			log.Printf("WARN: gofmt failed for %s: %v\n%s", name, err, string(out))
		}
	}

	return nil
}

// ============================================================
// Main
// ============================================================

func main() {
	schemaPath := flag.String("schema", "", "Path to OpenAPI JSON schema")
	outputDir := flag.String("output", "", "Output directory")
	pkgName := flag.String("package", "", "Go package name")
	flag.Parse()

	if *schemaPath == "" || *outputDir == "" || *pkgName == "" {
		log.Fatal("Usage: codegen -schema <path> -output <dir> -package <name>")
	}

	// Read and parse schema
	data, err := os.ReadFile(*schemaPath)
	if err != nil {
		log.Fatalf("Reading schema: %v", err)
	}

	var spec OpenAPISpec
	if err := json.Unmarshal(data, &spec); err != nil {
		log.Fatalf("Parsing schema: %v", err)
	}

	gen := NewGenerator(&spec, *pkgName, *outputDir)
	gen.Process()

	if err := gen.Generate(); err != nil {
		log.Fatalf("Generation failed: %v", err)
	}

	log.Printf("Generated %d models, %d enums, %d services in %s",
		len(gen.models), len(gen.enums), len(gen.services), *outputDir)
}
