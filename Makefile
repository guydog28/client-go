.PHONY: regen test test-verbose build clean distclean conformance-server smoke api-surface check-api-surface notices check-notices check-dependency-age docs

# regen: generate Go types from the vendored OpenAPI spec (spec/openapi.yaml).
# There is deliberately no local re-vendor target. spec/ is bot-owned: the
# bundle is mirrored out of ezclassify/contracts by the spec bot and lands as
# ezclassify-spec-bot. .github/workflows/vendor-readonly.yml rejects a spec/
# commit from any other identity. To advance the spec, update contracts and let
# the spec bot propagate it, then run `make regen` here.
regen:
	bash scripts/regen.sh

build:
	go build ./...

# api-surface: regenerate the 1.0-locked public API surface golden (api/1.0.txt)
#   after a reviewed, deliberate API change. check-api-surface is the CI gate:
#   it diffs the live exported surface against that file and fails on drift.
api-surface:
	go run ./tools/apisurface -write

check-api-surface:
	go run ./tools/apisurface

# notices: regenerate THIRD-PARTY-NOTICES.md from the vendored production
#   dependency closure after a dependency change. check-notices is the CI
#   gate: it fails on any license outside the allowlist and on a stale file.
notices:
	go run ./tools/notices -write

check-notices:
	go run ./tools/notices

# check-dependency-age: refuse third-party modules in go.sum published inside
#   the last seven days (EZCLASSIFY_DEPENDENCY_AGE_DAYS overrides the window).
check-dependency-age:
	go run ./tools/dependencyage

test:
	go test ./...

test-verbose:
	go test -v -p 1 ./...

clean:
	go clean ./...

# distclean: same as `clean`; Go keeps no toolchain cache in the tree.
distclean: clean

conformance-server: test-conformance-runner
	EZC_SMOKE_URL=$(URL) bash scripts/conformance-server.sh

# smoke: drive the SDK against a live rest-api over the wire (catches a
#   URL-prefix or stale-spec break the replayer cannot).
#   URL is required. Usage: make smoke URL=http://localhost:3000
smoke:
	EZC_SMOKE_URL=$(URL) go test -tags live -run TestSmoke -v .

# Pure helper tests for the explicit Python-backed live conformance lane.
.PHONY: test-conformance-runner
test-conformance-runner:
	PYTHONDONTWRITEBYTECODE=1 python3 -m unittest discover -s scripts -p test_conformance_server.py

# Native reference for every first-party package in this module. The package
# list also includes internal generated wire types and repository tools.
docs:
	rm -rf dist/api
	mkdir -p dist/api
	go list ./... > dist/api/packages.txt
	@while IFS= read -r package; do \
		relative=$${package#github.com/ezclassify/sdk-go}; \
		mkdir -p "dist/api$$relative"; \
		go doc -all "$$package" > "dist/api$$relative/index.txt" || exit; \
	done < dist/api/packages.txt
