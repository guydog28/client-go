# Developing the EzClassify Go SDK

Environment setup, the build and test commands, and the spec-vendoring workflow
for working on this repo. For the contribution process (sign-off, issues, pull
requests, commit format), see [CONTRIBUTING.md](CONTRIBUTING.md).

## Local development

Go 1.26.10+, matching [`go.mod`](go.mod), with the standard toolchain. The `Makefile` wraps the same commands.

```sh
go build ./...
go test ./...
```

| Action | Command | Notes |
|--------|---------|-------|
| Build | `go build ./...` | or `make build`. Covers the library and the [`examples/`](examples/). |
| Generate API reference | `make docs` | Native text reference at `dist/api/index.txt`, with per-package references listed in `dist/api/packages.txt`. |
| Test | `go test ./...` | or `make test` (`make test-verbose` for `-v`). |
| Conformance only | `go test -run TestConformance ./...` | Replays the shared spec vectors; no server needed. |
| Smoke test against a server | `make smoke URL=http://localhost:3000` | or `EZC_SMOKE_URL=… go test -tags live -run TestSmoke .`. |
| Vet (static checks) | `go vet ./...` | Ships with the Go toolchain; the project adds no third-party linter. |
| Format | `gofmt -w .` | Standard formatter. |
| Regenerate generated sources | `make regen` | See "Spec and conformance" below. [`spec/`](spec/) itself is bot-owned; you cannot refresh it from here. |
| Clean | `make clean` | |

The public API surface is frozen at 1.0 in [`api/1.0.txt`](api/1.0.txt); `make check-api-surface`
is the CI gate, and `make api-surface` regenerates it for a deliberate, reviewed
change committed alongside it.

## Spec and conformance

- **The mirror lives at [`spec/`](spec/), not `vendor/spec/`.** Go's [`vendor/`](vendor/) is the
  module vendor tree, and `go mod vendor` deletes anything in it that is not a
  module, so the spec bundle sits at the repository root instead.
- **The wire layer is generated.** Client methods come from
  [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) over the spec
  vendored at [`spec/openapi.yaml`](spec/openapi.yaml). That mirror is written exclusively by the
  EzClassify spec bot from [ezclassify/contracts](https://github.com/ezclassify/contracts), landing as
  `ezclassify-spec-bot`; `vendor-readonly.yml` rejects a [`spec/`](spec/) commit from
  anyone else. There is no local target that refreshes it.
  Don't hand-edit generated code; regenerate it with the pinned generator (no
  separate install needed):

  ```sh
  make regen
  ```

  The CI build reruns the generator and fails if the committed tree differs, so
  drift gets caught either way. Wrapper-only changes don't touch [`spec/`](spec/). Spec
  changes belong upstream in the [ezclassify/contracts](https://github.com/ezclassify/contracts) repository; report them through the
  central product tracker.
  Hand-written ergonomics (the `EzClassifyContext` handle) live alongside, not
  inside, the generated files.
- **Conformance plans define parity.** The conformance runner replays the same
  vectors as the TypeScript engine. It lives in
  [`internal/conformancerunner/`](internal/conformancerunner/); it seeds an in-process HTTP replayer with the
  recorded responses in [`spec/`](spec/) and replays every engine-slice plan case against
  it, subset-matching each response, so no live server or Docker is required:

  ```sh
  go test -run TestConformance ./...
  ```

  A change that breaks a conformance plan needs the plan changed in the [ezclassify/contracts](https://github.com/ezclassify/contracts) repository first.
- **Recorded replay and live acceptance cover different boundaries.** The
  replayer matches the normalized path and request body against a recording;
  an unrecorded path or body fails. The live-server smoke test additionally
  verifies a running deployment: point `EZC_SMOKE_URL` at a rest-api and it
  drives `Ping` (which hits `/v1/health`) plus a real `process()` over the wire.
  It is built only with the `live` tag and fails when the variable is unset.

  ```sh
  make smoke URL=http://localhost:3000
  # or: EZC_SMOKE_URL=http://localhost:3000 go test -tags live -run TestSmoke .
  ```

## What to Know Before Changing Things

- **Engine behavior belongs to EzClassify Runtime.** This SDK handles HTTP
  transport and decoding. Report suspected evaluation defects through the
  [central product tracker](https://github.com/ezclassify/ezclassify/issues/new/choose),
  with a synthetic marking, the scheme and Context versions, and sanitized
  request/response details. Runtime source access is not required.

## Source reference

`make docs` runs the standard `go doc -all` command for every first-party package
returned by `go list ./...`. Start with `dist/api/index.txt` for the public SDK.
The matching subdirectories contain the generated wire layer, conformance
helpers, examples and maintenance tools. These are source references, not
additional supported import paths. Third-party vendored sources and separate
code-generation/SBOM tool modules are excluded: the [`tools/codegen/`](tools/codegen/) and [`tools/sbom/`](tools/sbom/) directories contain only native module
manifests and dependency checksums, with no first-party Go source. Their third-party
generator implementations are documented upstream. Shell and Python scripts keep
their command contracts in this guide and their source headers. Output is ignored
by Git.

For interactive terminal lookup, use `go doc Client`, `go doc EzClassifyContext`
or `go doc -all .`. The module's Go doc comments also provide the reference on [pkg.go.dev](https://pkg.go.dev/github.com/ezclassify/sdk-go).
Generated wire declarations remain owned by `make regen`.
