# Agent Instructions

## Standards

Read [`STANDARDS.md`](./STANDARDS.md) before any feature, refactor, architectural decision, or task > 30 minutes. These engineering standards apply to you as they do to any contributor. Skip for: quick bug fixes, typos, documentation-only changes.

Read [`ARCHITECTURE.md`](./ARCHITECTURE.md) before any architectural work. It is
the repository's ownership contract for generated code, handwritten seams,
dependency direction, and required boundary gates.

## Examples

Runnable examples live under [`examples/`](examples/) and are part of this module (`go build ./...` covers them).

- **[`examples/basic/`](examples/basic/)**: validates a marking, aggregates two, and checks access through one published Context.
- **[`examples/go-classify/`](examples/go-classify/)**: scans a directory of text files through a Context and prints a per-file classification report. It sends arbitrary file text, so it needs a runtime that serves the Context.
- **[`examples/server-tour/`](examples/server-tour/)**: lists the Contexts a server publishes and processes a marking through one of them.

Each example's README is the source for its server setup, flag defaults, and expected output; keep facts about an example there, not here. All three read optional `EZC_ACCESS_TOKEN` and `EZC_CONTEXT_TOKEN` for a secured runtime.

Examples have no tests. They are built by `make build` and verified to compile on every CI run.
