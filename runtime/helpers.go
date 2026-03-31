package runtime

import (
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

// StructToQuery converts a struct to a map[string]interface{} using struct tags.
// It looks for `query:"name"` tags. Fields with nil pointer values are omitted.
func StructToQuery(v interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	if v == nil {
		return result
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return result
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return result
	}

	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fv := rv.Field(i)

		// Get query tag
		tag := field.Tag.Get("query")
		if tag == "" || tag == "-" {
			// Fall back to json tag
			tag = field.Tag.Get("json")
			if tag == "" || tag == "-" {
				continue
			}
		}

		// Strip options from tag (e.g., ",omitempty")
		if idx := strings.Index(tag, ","); idx != -1 {
			tag = tag[:idx]
		}

		if tag == "" {
			continue
		}

		// Skip nil pointers
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}

		// Skip zero values of non-pointer types if omitempty
		result[tag] = fv.Interface()
	}

	return result
}

// BuildQueryString builds a URL query string from a params map.
// It handles:
// - Simple values (string, int, float, bool)
// - Arrays/slices: key=val1&key=val2
// - Maps (deepObject style): key[subkey]=val
// - Nil values are stripped
// - Boolean values are encoded as "1"/"0" (Lolzteam API convention)
func BuildQueryString(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}

	values := url.Values{}

	// Sort keys for deterministic output
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		addQueryParam(values, key, params[key])
	}

	return values.Encode()
}

func addQueryParam(values url.Values, key string, val interface{}) {
	if val == nil {
		return
	}

	rv := reflect.ValueOf(val)

	// Handle pointers
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			addQueryParam(values, key, rv.Index(i).Interface())
		}
	case reflect.Map:
		// deepObject style: key[subkey]=val
		mapKeys := make([]string, 0, rv.Len())
		for _, k := range rv.MapKeys() {
			mapKeys = append(mapKeys, fmt.Sprintf("%v", k.Interface()))
		}
		sort.Strings(mapKeys)
		for _, mk := range mapKeys {
			mv := rv.MapIndex(reflect.ValueOf(mk))
			subKey := fmt.Sprintf("%s[%s]", key, mk)
			addQueryParam(values, subKey, mv.Interface())
		}
	case reflect.Bool:
		if rv.Bool() {
			values.Add(key, "1")
		} else {
			values.Add(key, "0")
		}
	case reflect.Invalid:
		// nil, skip
	default:
		s := fmt.Sprintf("%v", rv.Interface())
		if s != "" {
			values.Add(key, s)
		}
	}
}
