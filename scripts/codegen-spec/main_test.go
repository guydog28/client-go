package main

import (
	"os"
	"testing"

	"go.yaml.in/yaml/v3"
)

func decode(t *testing.T, input []byte) *yaml.Node {
	t.Helper()
	var document yaml.Node
	if err := yaml.Unmarshal(input, &document); err != nil {
		t.Fatal(err)
	}
	return &document
}

func at(t *testing.T, node *yaml.Node, names ...string) *yaml.Node {
	t.Helper()
	for _, name := range names {
		node = field(node, name)
		if node == nil {
			t.Fatalf("missing schema field %q", name)
		}
	}
	return node
}

func TestCanonicalCodegenViewPreservesNullableShapes(t *testing.T) {
	input, err := os.ReadFile("../../spec/openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	document := decode(t, input)
	if err := normalizeDocument(document); err != nil {
		t.Fatal(err)
	}
	root := document.Content[0]
	if got := at(t, root, "openapi").Value; got != "3.0.3" {
		t.Fatalf("codegen OpenAPI version = %q", got)
	}
	schemas := at(t, root, "components", "schemas")
	target := at(t, schemas, "DocumentAuditReport", "properties", "wholeValidation")
	if got := at(t, target, "nullable").Value; got != "true" {
		t.Fatalf("target nullable = %q", got)
	}
	ref := at(t, target, "allOf").Content[0]
	if got := at(t, ref, "$ref").Value; got != "#/components/schemas/MarkingValidation" {
		t.Fatalf("target reference = %q", got)
	}
	if field(target, "anyOf") != nil {
		t.Fatal("target retained unsupported nullable anyOf")
	}
	scopes := at(t, schemas, "SchemeScopesResponse")
	if at(t, scopes, "type").Value != "array" || field(scopes, "nullable") != nil {
		t.Fatal("scopes must stay a plain array that a response never answers with null")
	}
	whole := at(t, schemas, "DocumentAuditReport", "properties", "whole")
	if at(t, whole, "type").Value != "string" || at(t, whole, "nullable").Value != "true" {
		t.Fatal("nullable scalar lost its string type or null arm")
	}
	value := at(t, schemas, "RuleCondition", "properties", "value")
	variants := at(t, value, "oneOf")
	if len(variants.Content) != 2 || at(t, variants.Content[0], "type").Value != "string" || at(t, variants.Content[1], "type").Value != "number" {
		t.Fatal("disjoint string/number union was not preserved")
	}
	if field(value, "type") != nil {
		t.Fatal("unsupported primitive type array remained")
	}
	other := at(t, schemas, "CrosswalkForbidMatch")
	if field(other, "anyOf") == nil {
		t.Fatal("ordinary non-null union was changed")
	}
}

func TestCodegenViewRejectsAmbiguousUnions(t *testing.T) {
	for _, schema := range []string{
		"anyOf: [{type: string}, {type: 'null'}, {type: number}]",
		"anyOf: [{type: string}, {type: 'null', description: extra}]",
		"type: [string, boolean]",
		"type: [string, number, 'null']",
		"type: [string, number]\n      oneOf: [{type: string}, {type: number}]",
		"type: [string, number]\n      anyOf: [{type: string}, {type: number}]",
	} {
		document := decode(t, []byte("openapi: 3.1.0\ncomponents:\n  schemas:\n    Test:\n      "+schema+"\n"))
		if err := normalizeDocument(document); err == nil {
			t.Errorf("accepted ambiguous schema %s", schema)
		}
	}
}

func TestCodegenViewLeavesLiteralExamplesUntouched(t *testing.T) {
	document := decode(t, []byte(`openapi: 3.1.0
components:
  schemas:
    Test:
      type: object
      properties:
        value:
          type: string
          example:
            type: [string, number, boolean]
            schema:
              type: [string, number]
      examples:
        - type: [string, number]
      default:
        type: [string, number]
      const:
        type: [string, number]
      enum:
        - type: [string, number]
      x-example:
        type: [string, number]
`))
	if err := normalizeDocument(document); err != nil {
		t.Fatal(err)
	}
	schema := at(t, document.Content[0], "components", "schemas", "Test")
	for _, literal := range []*yaml.Node{
		at(t, schema, "properties", "value", "example"),
		at(t, schema, "properties", "value", "example", "schema"),
		at(t, schema, "examples").Content[0],
		at(t, schema, "default"),
		at(t, schema, "const"),
		at(t, schema, "enum").Content[0],
		at(t, schema, "x-example"),
	} {
		types := at(t, literal, "type")
		if types.Kind != yaml.SequenceNode || len(types.Content) < 2 || field(literal, "oneOf") != nil {
			t.Fatal("literal example or value was rewritten as a schema")
		}
	}
}
