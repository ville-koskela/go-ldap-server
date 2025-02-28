package test

import (
	"fmt"
	"reflect"
	"testing"
)

func Assert(t *testing.T, expected any, actual any, msg ...string) bool {
	if expected == nil {
		return handleNil(t, actual, msg...)
	}

	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		errorMsg := fmt.Sprintf("Expected %v, got %v", expected, actual)
		if len(msg) > 0 {
			errorMsg = fmt.Sprintf("%s: %s", msg[0], errorMsg)
		}
		t.Error(errorMsg)
		return false
	}

	switch expected.(type) {
	case bool:
		return handleBool(t, expected.(bool), actual.(bool), msg...)
	case string:
		return handleString(t, expected.(string), actual.(string), msg...)
	case int, int8, int16, int32, int64:
		return handleNumeric(t, reflect.ValueOf(expected).Int(), reflect.ValueOf(actual).Int(), msg...)
	case uint, uint8, uint16, uint32, uint64:
		return handleNumeric(t, reflect.ValueOf(expected).Uint(), reflect.ValueOf(actual).Uint(), msg...)
	case float32, float64:
		return handleNumeric(t, reflect.ValueOf(expected).Float(), reflect.ValueOf(actual).Float(), msg...)
	case error:
		return handleError(t, expected.(error), actual.(error), msg...)
	default:
		errorMsg := fmt.Sprintf("Type not handled: %v", expected)
		if len(msg) > 0 {
			errorMsg = fmt.Sprintf("%s: %s", msg[0], errorMsg)
		}
		t.Error(errorMsg)
		return false
	}
}

func handleBool(t *testing.T, expected bool, actual bool, msg ...string) bool {
	if expected != actual {
		errorMsg := fmt.Sprintf("Expected %v, got %v", expected, actual)
		if len(msg) > 0 {
			errorMsg = fmt.Sprintf("%s: %s", msg[0], errorMsg)
		}
		t.Error(errorMsg)
		return false
	}

	return true
}

func handleNil(t *testing.T, actual any, msg ...string) bool {
	if actual == nil {
		return true
	}

	errorMsg := fmt.Sprintf("Expected nil, got %v", actual)
	if len(msg) > 0 {
		errorMsg = fmt.Sprintf("%s: %s", msg[0], errorMsg)
	}
	t.Error(errorMsg)
	return false
}

func handleString(t *testing.T, expected string, actual string, msg ...string) bool {
	if expected != actual {
		errorMsg := fmt.Sprintf("Expected %s, got %s", expected, actual)
		if len(msg) > 0 {
			errorMsg = fmt.Sprintf("%s: %s", msg[0], errorMsg)
		}
		t.Error(errorMsg)
		return false
	}

	return true
}

func handleNumeric(t *testing.T, expected any, actual any, msg ...string) bool {
	if expected != actual {
		errorMsg := fmt.Sprintf("Expected %v, got %v", expected, actual)
		if len(msg) > 0 {
			errorMsg = fmt.Sprintf("%s: %s", msg[0], errorMsg)
		}
		t.Error(errorMsg)
		return false
	}

	return true
}

func handleError(t *testing.T, expected error, actual error, msg ...string) bool {
	if expected.Error() != actual.Error() {
		errorMsg := fmt.Sprintf("Expected error %v, got %v", expected, actual)
		if len(msg) > 0 {
			errorMsg = fmt.Sprintf("%s: %s", msg[0], errorMsg)
		}
		t.Error(errorMsg)
		return false
	}

	return true
}
