// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports net/http/httptest'. DO NOT EDIT.

import (
	"net/http/httptest"
	"reflect"
)

func init() {
	Symbols["net/http/httptest"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DefaultRemoteAddr":  reflect.ValueOf(httptest.DefaultRemoteAddr),
		"NewRecorder":        reflect.ValueOf(httptest.NewRecorder),
		"NewRequest":         reflect.ValueOf(httptest.NewRequest),
		"NewServer":          reflect.ValueOf(httptest.NewServer),
		"NewTLSServer":       reflect.ValueOf(httptest.NewTLSServer),
		"NewUnstartedServer": reflect.ValueOf(httptest.NewUnstartedServer),

		// type definitions
		"ResponseRecorder": reflect.ValueOf((*httptest.ResponseRecorder)(nil)),
		"Server":           reflect.ValueOf((*httptest.Server)(nil)),

		// interface wrapper definitions

	}
}
