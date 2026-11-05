package ezclassify_test

// Required-member gate.
//
// A 2xx whose body lacks a member the contract requires, at any depth, or
// carries null where the schema does not allow it, is a failure and never a
// success to cache (cache contract C07). The members come from the vendored
// OpenAPI document and from nothing else: each recorded success is walked
// against its operation's response schema, and every member a required list
// names, wherever it sits in the recorded body, is taken away in turn. The
// public call that decoded the recording has to refuse the body with a
// *MalformedResponseError naming that member. Nothing here lists the members
// the client is known to check, so a member the client forgot, or a decoder
// that fills one in with a zero value, fails the member it forgot.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	"go.yaml.in/yaml/v3"

	ez "github.com/ezclassify/sdk-go"
)

// specResponseOperation is one contract operation and the schema of its
// success body.
type specResponseOperation struct {
	id       string
	method   string
	segments []string
	schema   any
}

type specResponses struct {
	components map[string]any
	operations []specResponseOperation
}

// memberWalk is what one walk of a recorded body finds: its required members,
// and the schema shapes it cannot judge, so a contract that starts using one
// fails the gate instead of quietly narrowing it.
type memberWalk struct {
	found       []requiredMember
	unsupported []string
}

var loadSpecResponses = sync.OnceValues(func() (*specResponses, error) {
	raw, err := os.ReadFile(requestSchemaSpecPath)
	if err != nil {
		return nil, err
	}
	var root yaml.Node
	if err := yaml.Unmarshal(raw, &root); err != nil {
		return nil, err
	}
	decoded, err := yamlToJSON(&root)
	if err != nil {
		return nil, err
	}
	components, _ := lookup(decoded, "components", "schemas").(map[string]any)
	paths, _ := lookup(decoded, "paths").(map[string]any)
	if components == nil || paths == nil {
		return nil, errors.New("the OpenAPI document has no components.schemas or paths")
	}
	spec := &specResponses{components: components}
	for path, item := range paths {
		operations, _ := item.(map[string]any)
		for method, operation := range operations {
			responses, _ := lookup(operation, "responses").(map[string]any)
			for status, response := range responses {
				schema := lookup(response, "content", "application/json", "schema")
				if !strings.HasPrefix(status, "2") || schema == nil {
					continue
				}
				id, _ := lookup(operation, "operationId").(string)
				spec.operations = append(spec.operations, specResponseOperation{
					id:       id,
					method:   strings.ToUpper(method),
					segments: strings.Split(strings.Trim(path, "/"), "/"),
					schema:   schema,
				})
			}
		}
	}
	return spec, nil
})

// operationOf returns the contract operation a recorded request addressed.
func (s *specResponses) operationOf(recorded wireCase) (specResponseOperation, error) {
	target, err := url.Parse(recorded.Request.Path)
	if err != nil {
		return specResponseOperation{}, err
	}
	segments := strings.Split(strings.Trim(target.Path, "/"), "/")
	var matches []specResponseOperation
	for _, operation := range s.operations {
		if operation.method != strings.ToUpper(recorded.Request.Method) || len(operation.segments) != len(segments) {
			continue
		}
		matched := true
		for i, template := range operation.segments {
			if !strings.HasPrefix(template, "{") && template != segments[i] {
				matched = false
				break
			}
		}
		if matched {
			matches = append(matches, operation)
		}
	}
	if len(matches) != 1 {
		return specResponseOperation{}, fmt.Errorf("%s: %d operations match %s %s", recorded.ID, len(matches), recorded.Request.Method, target.Path)
	}
	return matches[0], nil
}

func (s *specResponses) resolve(schema any) map[string]any {
	for {
		object, _ := schema.(map[string]any)
		reference, ok := object["$ref"].(string)
		if !ok {
			return object
		}
		schema = s.components[strings.TrimPrefix(reference, componentSchemaPrefix)]
	}
}

// requiredMember is one required member present in a recorded body: the path
// of keys and indexes that reaches it, and its JSON pointer.
type requiredMember struct {
	tokens   []any
	pointer  string
	nullable bool
	// rivals are the members other alternatives of the member's union require
	// that the recorded value does not carry. Without the member, the body can
	// fit one of those alternatives as well as its own.
	rivals []string
}

// requiredMembers lists every member of value that schema requires, at any
// depth, in a stable order, and the shapes the walk could not judge.
func (s *specResponses) requiredMembers(schema, value any) ([]requiredMember, []string) {
	walk := &memberWalk{}
	s.collect(schema, value, nil, "", nil, walk)
	sort.Slice(walk.found, func(i, j int) bool { return walk.found[i].pointer < walk.found[j].pointer })
	return walk.found, walk.unsupported
}

func (s *specResponses) collect(schemaNode, value any, tokens []any, pointer string, rivals []string, walk *memberWalk) {
	schema := s.resolve(schemaNode)
	if schema == nil || value == nil {
		return
	}
	if _, composed := schema["allOf"]; composed {
		walk.unsupported = append(walk.unsupported, pointer+": allOf")
		return
	}
	if alternatives := alternativesOf(schema); alternatives != nil {
		// Walk the alternative the value satisfies by name: every member it
		// requires is present and every constant it fixes is the one the value
		// carries.
		for index, alternative := range alternatives {
			resolved := s.resolve(alternative)
			if !s.describes(resolved, value) {
				continue
			}
			s.collect(resolved, value, tokens, pointer, s.rivalMembers(alternatives, index, value), walk)
			return
		}
		switch value.(type) {
		case map[string]any, []any:
			walk.unsupported = append(walk.unsupported, pointer+": no union alternative describes the recorded value")
		}
		return
	}
	switch typed := value.(type) {
	case map[string]any:
		properties, _ := schema["properties"].(map[string]any)
		required, _ := schema["required"].([]any)
		for _, name := range required {
			key := name.(string)
			if _, present := typed[key]; present {
				walk.found = append(walk.found, requiredMember{
					tokens:   append(append([]any(nil), tokens...), key),
					pointer:  pointer + "/" + escapeJSONPointer(key),
					nullable: s.allowsNull(properties[key]),
					rivals:   rivals,
				})
			}
		}
		for key, member := range typed {
			memberSchema, declared := properties[key]
			if !declared {
				memberSchema, _ = schema["additionalProperties"].(map[string]any)
			}
			s.collect(memberSchema, member, append(append([]any(nil), tokens...), key), pointer+"/"+escapeJSONPointer(key), nil, walk)
		}
	case []any:
		for i, item := range typed {
			s.collect(schema["items"], item, append(append([]any(nil), tokens...), i), pointer+"/"+strconv.Itoa(i), nil, walk)
		}
	}
}

// rivalMembers lists the members the alternatives other than chosen require
// that value does not carry.
func (s *specResponses) rivalMembers(alternatives []any, chosen int, value any) []string {
	object, _ := value.(map[string]any)
	var rivals []string
	for index, alternative := range alternatives {
		if index == chosen {
			continue
		}
		required, _ := s.resolve(alternative)["required"].([]any)
		for _, name := range required {
			if _, present := object[name.(string)]; !present {
				rivals = append(rivals, name.(string))
			}
		}
	}
	return rivals
}

func alternativesOf(schema map[string]any) []any {
	if alternatives, ok := schema["anyOf"].([]any); ok {
		return alternatives
	}
	alternatives, _ := schema["oneOf"].([]any)
	return alternatives
}

func (s *specResponses) describes(alternative map[string]any, value any) bool {
	object, isObject := value.(map[string]any)
	if !isObject {
		_, isArray := value.([]any)
		return isArray && alternative["type"] == "array"
	}
	required, _ := alternative["required"].([]any)
	for _, name := range required {
		if _, present := object[name.(string)]; !present {
			return false
		}
	}
	properties, _ := alternative["properties"].(map[string]any)
	for name, property := range properties {
		if constant, fixed := s.resolve(property)["const"]; fixed && fmt.Sprint(object[name]) != fmt.Sprint(constant) {
			return false
		}
	}
	return true
}

// allowsNull reports whether a member schema accepts a JSON null: a type list
// naming null, or an alternative that is null.
func (s *specResponses) allowsNull(schemaNode any) bool {
	schema := s.resolve(schemaNode)
	if types, ok := schema["type"].([]any); ok {
		for _, name := range types {
			if name == "null" {
				return true
			}
		}
	}
	for _, alternative := range alternativesOf(schema) {
		if s.allowsNull(alternative) || s.resolve(alternative)["type"] == "null" {
			return true
		}
	}
	return false
}

// namesRival reports whether named is a member a rival alternative of
// member's union requires, at the same depth, that the body lacks.
func namesRival(member requiredMember, named string) bool {
	parent := member.pointer[:strings.LastIndex(member.pointer, "/")]
	for _, rival := range member.rivals {
		if named == parent+"/"+escapeJSONPointer(rival) {
			return true
		}
	}
	return false
}

func escapeJSONPointer(token string) string {
	return strings.ReplaceAll(strings.ReplaceAll(token, "~", "~0"), "/", "~1")
}

// rewriteMember returns body with the member at tokens removed, or set to
// null when toNull is true.
func rewriteMember(body []byte, tokens []any, toNull bool) ([]byte, error) {
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	var tree any
	if err := decoder.Decode(&tree); err != nil {
		return nil, err
	}
	node := tree
	for _, token := range tokens[:len(tokens)-1] {
		switch key := token.(type) {
		case string:
			node = node.(map[string]any)[key]
		case int:
			node = node.([]any)[key]
		}
	}
	parent := node.(map[string]any)
	last := tokens[len(tokens)-1].(string)
	if toNull {
		parent[last] = nil
	} else {
		delete(parent, last)
	}
	return json.Marshal(tree)
}

// decodedSuccesses returns the recorded successes a public call decodes.
func decodedSuccesses(t *testing.T) []wireCase {
	t.Helper()
	corpus, err := loadWireCorpus("spec")
	if err != nil {
		t.Fatal(err)
	}
	var cases []wireCase
	for _, recorded := range corpus.Cases {
		if wireDrivers[recorded.ID] == publicDriver && recorded.Response.Status >= 200 && recorded.Response.Status < 300 {
			cases = append(cases, recorded)
		}
	}
	return cases
}

// replayAnswer drives recorded through its public driver against a replay that
// answers with body instead of the recorded response body.
func replayAnswer(recorded wireCase, body []byte) error {
	recorded.Response.Body = body
	replay := &wireReplay{Case: recorded}
	client := ez.New("http://wire.invalid", ez.WithHTTPClient(&http.Client{Transport: wireTransport{replay: replay}}), ez.WithCredentials(wireCredentials(recorded)))
	_, err := runWireDriver(context.Background(), client, recorded)
	if replay.Failure != nil {
		return replay.Failure
	}
	return err
}

func TestEveryRequiredMemberIsRefusedWhenAbsentOrNull(t *testing.T) {
	spec, err := loadSpecResponses()
	if err != nil {
		t.Fatal(err)
	}
	for _, recorded := range decodedSuccesses(t) {
		recorded := recorded
		t.Run(recorded.ID, func(t *testing.T) {
			operation, err := spec.operationOf(recorded)
			if err != nil {
				t.Fatal(err)
			}
			var tree any
			if err := json.Unmarshal(recorded.Response.Body, &tree); err != nil {
				t.Fatal(err)
			}
			if err := replayAnswer(recorded, recorded.Response.Body); err != nil {
				t.Fatalf("the recorded answer itself is refused: %v", err)
			}
			members, unsupported := spec.requiredMembers(operation.schema, tree)
			if len(unsupported) != 0 {
				t.Fatalf("the response schema uses shapes the walk cannot judge: %v", unsupported)
			}
			for _, member := range members {
				for _, toNull := range []bool{false, true} {
					if toNull && member.nullable {
						continue
					}
					body, err := rewriteMember(recorded.Response.Body, member.tokens, toNull)
					if err != nil {
						t.Fatal(err)
					}
					change := "absent"
					if toNull {
						change = "null"
					}
					err = replayAnswer(recorded, body)
					var malformed *ez.MalformedResponseError
					switch {
					case !errors.As(err, &malformed):
						t.Errorf("%s %s: got %T(%v), want a *MalformedResponseError", member.pointer, change, err, err)
					case malformed.Member == member.pointer:
					case !toNull && namesRival(member, malformed.Member):
						// The generated types carry which members each union
						// alternative requires but not the constants that tell
						// the alternatives apart. Without the one member that
						// separates them, the body fits a rival alternative as
						// well as its own, and the refusal names a member that
						// rival requires and the body lacks.
					default:
						t.Errorf("%s %s: refused naming %q", member.pointer, change, malformed.Member)
					}
				}
			}
		})
	}
}

// TestTheCorpusExercisesEverySuccessSchema keeps the gate above honest: every
// operation the contract answers with a JSON body has a recorded success that
// a public call decodes, so no response schema escapes the walk.
func TestTheCorpusExercisesEverySuccessSchema(t *testing.T) {
	spec, err := loadSpecResponses()
	if err != nil {
		t.Fatal(err)
	}
	covered := map[string]bool{}
	for _, recorded := range decodedSuccesses(t) {
		operation, err := spec.operationOf(recorded)
		if err != nil {
			t.Fatal(err)
		}
		covered[operation.id] = true
	}
	for _, operation := range spec.operations {
		if _, notSent := specOperationsNotSent[operation.id]; !covered[operation.id] && !notSent {
			t.Errorf("no recorded success of %s is decoded through a public call", operation.id)
		}
	}
}
