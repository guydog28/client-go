module github.com/ezclassify/sdk-go

go 1.26.10

require (
	github.com/oapi-codegen/runtime v1.7.0
	go.yaml.in/yaml/v3 v3.0.5 // reads spec/openapi.yaml in tests and in scripts/codegen-spec
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

// Tool dependencies (used only by go generate, not imported by the SDK).
// Codegen: oapi-codegen v2 regenerates internal/api/api.gen.go
// from the vendored OpenAPI spec at spec/openapi.yaml.
// tools/codegen/go.mod owns its isolated dependency graph; see generate.go.
