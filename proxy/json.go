package proxy

import (
	"fmt"
)

type UnmarshalWrongTypeError struct {
	Field, Expected string
	Got             interface{}
}

func (e *UnmarshalWrongTypeError) Error() string {
	return fmt.Sprintf("Wrong type for %s field, expected %s, but got %T", e.Field, e.Expected, e.Got)
}

// For parsing json objects. Look up (or create) a child object in the given map.
func lookupObject(m map[string]interface{}, key string) (map[string]interface{}, error) {
	iface, ok := m[key]
	if !ok || iface == nil {
		result := map[string]interface{}{}
		m[key] = result
		return result, nil
	}

	result, ok := iface.(map[string]interface{})
	if !ok {
		return nil, &UnmarshalWrongTypeError{key, "object", iface}
	}

	return result, nil
}

// For parsing json objects. Look up (or create) a string in the given map.
func lookupString(m map[string]interface{}, key string) (string, error) {
	iface, ok := m[key]
	if !ok || iface == nil {
		return "", nil
	}

	result, ok := iface.(string)
	if !ok {
		return "", &UnmarshalWrongTypeError{key, "string", iface}
	}

	return result, nil
}

// For parsing json objects. Look up (or create) a string array in the given map.
func lookupStringArray(m map[string]interface{}, key string) ([]string, error) {
	iface, ok := m[key]
	if !ok || iface == nil {
		return []string{}, nil
	}

	switch o := iface.(type) {
	case []string:
		return o, nil
	case []interface{}:
		result := []string{}
		for _, s := range o {
			if s, ok := s.(string); ok {
				result = append(result, s)
			} else {
				return nil, &UnmarshalWrongTypeError{key, "array of strings", iface}
			}
		}
		return result, nil
	}
	return nil, &UnmarshalWrongTypeError{key, "array of strings", iface}
}
