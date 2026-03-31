package runtime

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// StringOrInt represents a value that can be either a string or an integer in JSON.
type StringOrInt struct {
	String *string
	Int    *int
}

// NewStringOrIntFromString creates a StringOrInt from a string.
func NewStringOrIntFromString(s string) StringOrInt {
	return StringOrInt{String: &s}
}

// NewStringOrIntFromInt creates a StringOrInt from an int.
func NewStringOrIntFromInt(i int) StringOrInt {
	return StringOrInt{Int: &i}
}

// Value returns the underlying value as an interface{}.
func (s StringOrInt) Value() interface{} {
	if s.String != nil {
		return *s.String
	}
	if s.Int != nil {
		return *s.Int
	}
	return nil
}

// StringValue returns the value as a string regardless of the underlying type.
func (s StringOrInt) StringValue() string {
	if s.String != nil {
		return *s.String
	}
	if s.Int != nil {
		return strconv.Itoa(*s.Int)
	}
	return ""
}

// MarshalJSON implements json.Marshaler.
func (s StringOrInt) MarshalJSON() ([]byte, error) {
	if s.String != nil {
		return json.Marshal(*s.String)
	}
	if s.Int != nil {
		return json.Marshal(*s.Int)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *StringOrInt) UnmarshalJSON(data []byte) error {
	// Handle null first
	if string(data) == "null" {
		s.String = nil
		s.Int = nil
		return nil
	}

	// Try string
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		s.String = &str
		s.Int = nil
		return nil
	}

	// Try integer
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		s.Int = &num
		s.String = nil
		return nil
	}

	// Try float (JSON numbers can be floats)
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		i := int(f)
		s.Int = &i
		s.String = nil
		return nil
	}

	return fmt.Errorf("StringOrInt: cannot unmarshal %s", string(data))
}

// IsZero returns true if neither String nor Int is set.
func (s StringOrInt) IsZero() bool {
	return s.String == nil && s.Int == nil
}
