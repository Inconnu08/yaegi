// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports go/importer'. DO NOT EDIT.

import (
	"go/importer"
	"reflect"
)

func init() {
	Symbols["go/importer"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Default":     reflect.ValueOf(importer.Default),
		"For":         reflect.ValueOf(importer.For),
		"ForCompiler": reflect.ValueOf(importer.ForCompiler),

		// type definitions
		"Lookup": reflect.ValueOf((*importer.Lookup)(nil)),

		// interface wrapper definitions

	}
}
