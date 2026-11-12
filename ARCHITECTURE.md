# Architecture

This repository contains the EzClassify Go SDK. It adapts the public REST
contract to Go; EzClassify Runtime evaluates marking operations and composes
Contexts. It ships one public package,
`github.com/ezclassify/sdk-go`. Generated wire code stays internal; the root
package exposes aliases and small Go-native transport conveniences.

## Sources of truth

- The [ezclassify/contracts](https://github.com/ezclassify/contracts) repository owns
  canonical wire types, JSON Schemas, OpenAPI, and conformance fixtures.
  EzClassify Runtime owns engine behavior and Context composition.
- [`spec/`](spec/) is this repository's read-only contract mirror. The spec bot is
  its sole writer. [`spec/MANIFEST.json`](spec/MANIFEST.json) records `contractsSha` provenance,
  contract digests, the Go generator version and module hash, and the
  committed generated path. Do not edit this directory during ordinary SDK
  work.
- [`spec/openapi.yaml`](spec/openapi.yaml) is the wire-generation input. Plans, recordings,
  fixtures, the replayer contract, and semantic cases under [`spec/`](spec/) define the
  cross-language behavior the adapter must preserve.
- [`api/1.0.txt`](api/1.0.txt) is the reviewed public API golden for the root `ezclassify`
  package. It records generated aliases without copying generated fields.

## Generated code

- [`scripts/regen.sh`](scripts/regen.sh) is the canonical regeneration entrypoint. [`generate.go`](generate.go)
  carries the equivalent `go:generate` directive, and [`oapi-codegen.yaml`](oapi-codegen.yaml)
  configures generation of both models and the typed HTTP client.
- [`tools/codegen/go.mod`](tools/codegen/go.mod) and [`tools/codegen/go.sum`](tools/codegen/go.sum) own the isolated native generator tool
  graph, including patched transitive dependencies. [`scripts/regen.sh`](scripts/regen.sh) uses that
  graph in read-only module mode without adding build tools to SDK dependencies.
- The generator is `oapi-codegen` at the version and module hash pinned by
  [`scripts/regen.sh`](scripts/regen.sh) and [`spec/MANIFEST.json`](spec/MANIFEST.json). The sole committed output is
  [`internal/api/api.gen.go`](internal/api/api.gen.go).
- [`internal/api/api.gen.go`](internal/api/api.gen.go) is the only source file that declares wire DTOs and
  generated HTTP request/response wrappers. Its first line records the OpenAPI
  digest, generator version, generator module hash, and `DO NOT EDIT` marker.
- Run `make regen` from the module root after an authorized OpenAPI or generator
  change. A fresh generation must leave [`internal/api/api.gen.go`](internal/api/api.gen.go) unchanged on
  an otherwise clean checkout.

## Handwritten seams

- The module root is the public `ezclassify` package, which is idiomatic for a
  Go library. [`types.go`](types.go) re-exports selected [`internal/api/`](internal/api/) declarations as
  type aliases; it does not copy wire structs. [`client.go`](client.go) embeds the generated
  response client and owns the `New` constructor, transport injection, the
  health call, and version compatibility warnings. [`options.go`](options.go) owns the
  functional options (credentials, timeout, HTTP client, response handling).
- [`context_client.go`](context_client.go) owns a thin Context-bound facade, operation grouping,
  resolved-version capture, strict response handling, and typed Go errors. Its
  Scheme, Guide, Access, and Document methods call generated HTTP operations;
  they do not evaluate those operations locally.
- [`offsets.go`](offsets.go) owns checked conversion between Go rune indices and UTF-8 byte
  offsets. [`spec_version.go`](spec_version.go) is a release-stamped compatibility constant.
- [`internal/conformancerunner/`](internal/conformancerunner/), [`internal/replayer/`](internal/replayer/), and
  [`internal/subsetmatch/`](internal/subsetmatch/) are test infrastructure for replaying canonical
  recordings and comparing expected payloads. They are not runtime engine
  implementations. [`tools/`](tools/) contains repository maintenance commands.
- Runnable demonstrations remain `main` packages under [`examples/`](examples/), so
  `go build ./...` compiles them with the module. The SDK ships no command, so
  there is no `cmd/` package. Committed third-party Go sources stay under
  [`vendor/`](vendor/) and are never mixed with first-party packages.

## Forbidden ownership

- Do not implement parsing, normalization, validation, aggregation, scanning,
  rendering, suggestion, diffing, guide evaluation, access decisions, document
  policy, or any other engine behavior in this module. Send those operations
  to the runtime-owned REST API.
- Do not compose schemes, layers, guides, access policy, or Contexts here. The
  SDK may address a published Context and expose its server-resolved version;
  composition and evaluation remain upstream.
- Do not add handwritten wire DTOs, shadow generated request or response
  structs, patch [`internal/api/api.gen.go`](internal/api/api.gen.go), or expose [`internal/api/`](internal/api/) as a public
  import path. Change the upstream contract and let the spec bot
  propagate it, or update the generator and regenerate the single output.
- Do not move generated code into the public package or make [`internal/api/`](internal/api/)
  depend on root-package glue. Dependency direction is public adapters to the
  internal generated layer.
- Do not turn conformance replay helpers into production fallback behavior.
  Recorded responses exist only to test transport fidelity when a live server
  is not part of the test run.

## Required gates

- Generation: run `make regen`, then
  `git diff --exit-code -- internal/api/api.gen.go`. CI also verifies the
  OpenAPI digest and downloads the manifest-pinned generator module to verify
  its hash.
- Ownership: run `go test ./...`. `TestGeneratedWireOwnership` requires the
  manifest to name one generated output, checks its ownership header, proves
  every declaration in [`types.go`](types.go) is an alias to [`internal/api/`](internal/api/), and rejects
  unreviewed root-package structs that could shadow wire DTOs.
- Vendored spec: run `go test ./...`. `TestVendorGateShape` and
  `TestVendorGateBehavior` extract the literal shell body of
  [`.github/workflows/vendor-readonly.yml`](.github/workflows/vendor-readonly.yml) and drive it against throwaway git
  history, so the gate that keeps hand edits out of [`spec/`](spec/) is itself
  exercised rather than only reviewed.
- Public API and behavior: run `make check-api-surface` and `go test ./...`.
  The test suite includes generated-client coverage, public adapter tests,
  engine-slice conformance replay, replayer-contract tests, document
  semantic replay, and the request-schema gate, which sends every field each
  OpenAPI request schema declares through the public client and validates the
  bodies against [`spec/openapi.yaml`](spec/openapi.yaml).
- Go quality and distribution: run `go build ./...`, `go vet ./...`, and
  `test -z "$(gofmt -l . | grep -v '^vendor/')"`. The build includes every
  example command.
- Dependencies and notices: run `go mod verify`,
  `go list -mod=vendor -deps .`, and `make check-notices`.
- Review obligation: reject any local engine, policy, or Context-composition
  behavior even when all automated gates pass. Transport fidelity is the
  purpose of this repository; an alternate evaluator is not.

The CI-only SBOM generator has a separate native tool module in [`tools/sbom/`](tools/sbom/).
Its transitive security fixes do not enter the SDK runtime module or vendor tree.
