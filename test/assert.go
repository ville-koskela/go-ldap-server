package test

import (
	"reflect"
	"testing"
)

func Assert(t *testing.T, expected any, actual any) bool {
	if expected == nil {
		return handleNil(t, actual)
	}

	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
		return false
	}

	switch expected.(type) {
	case bool:
		return handleBool(t, expected.(bool), actual.(bool))
	case string:
		return handleString(t, expected.(string), actual.(string))
	case int, int8, int16, int32, int64:
		return handleNumeric(t, reflect.ValueOf(expected).Int(), reflect.ValueOf(actual).Int())
	case uint, uint8, uint16, uint32, uint64:
		return handleNumeric(t, reflect.ValueOf(expected).Uint(), reflect.ValueOf(actual).Uint())
	case float32, float64:
		return handleNumeric(t, reflect.ValueOf(expected).Float(), reflect.ValueOf(actual).Float())
	case error:
		return handleError(t, expected.(error), actual.(error))
	default:
		t.Errorf("Type not handled: %v", expected)
		return false
	}

}

func handleBool(t *testing.T, expected bool, actual bool) bool {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
		return false
	}

	return true
}

func handleNil(t *testing.T, actual any) bool {
	if actual == nil {
		return true
	}

	t.Errorf("Expected nil, got %v", actual)
	return false
}

func handleString(t *testing.T, expected string, actual string) bool {
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
		return false
	}

	return true
}

func handleNumeric(t *testing.T, expected any, actual any) bool {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
		return false
	}

	return true
}

func handleError(t *testing.T, expected error, actual error) bool {
	if expected.Error() != actual.Error() {
		t.Errorf("Expected error %v, got %v", expected, actual)
		return false
	}

	return true
}
