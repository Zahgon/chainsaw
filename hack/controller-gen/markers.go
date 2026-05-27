package main

import (
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type OneOf struct {
	Value any
}

func (m OneOf) ApplyToSchema(schema *apiext.JSONSchemaProps) error {
	_ = "STUB: not implemented"
	return nil
}

type Not struct {
	Value any
}

func (m Not) ApplyToSchema(schema *apiext.JSONSchemaProps) error {
	_ = "STUB: not implemented"
	return nil
}
