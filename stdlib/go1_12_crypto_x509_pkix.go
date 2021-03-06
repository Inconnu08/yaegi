// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports crypto/x509/pkix'. DO NOT EDIT.

import (
	"crypto/x509/pkix"
	"reflect"
)

func init() {
	Symbols["crypto/x509/pkix"] = map[string]reflect.Value{
		// function, constant and variable definitions

		// type definitions
		"AlgorithmIdentifier":          reflect.ValueOf((*pkix.AlgorithmIdentifier)(nil)),
		"AttributeTypeAndValue":        reflect.ValueOf((*pkix.AttributeTypeAndValue)(nil)),
		"AttributeTypeAndValueSET":     reflect.ValueOf((*pkix.AttributeTypeAndValueSET)(nil)),
		"CertificateList":              reflect.ValueOf((*pkix.CertificateList)(nil)),
		"Extension":                    reflect.ValueOf((*pkix.Extension)(nil)),
		"Name":                         reflect.ValueOf((*pkix.Name)(nil)),
		"RDNSequence":                  reflect.ValueOf((*pkix.RDNSequence)(nil)),
		"RelativeDistinguishedNameSET": reflect.ValueOf((*pkix.RelativeDistinguishedNameSET)(nil)),
		"RevokedCertificate":           reflect.ValueOf((*pkix.RevokedCertificate)(nil)),
		"TBSCertificateList":           reflect.ValueOf((*pkix.TBSCertificateList)(nil)),

		// interface wrapper definitions

	}
}
