// Command basic runs three core operations against one published Context:
// validate, aggregate, and access check.
//
// Start the Cedar Valley replay server, which replays real runtime responses
// for cedar-valley@1.0.0, then run the command:
//
//	npx @ezclassify/conformance-server --scenario cedar-valley --port 4010
//	go run ./examples/basic
//
// The defaults target that server, and its recordings cover the inputs below.
// To switch to your runtime, pass --server, --ctx, and --ctx-version, set
// EZC_ACCESS_TOKEN and EZC_CONTEXT_TOKEN if your runtime requires them, and use
// markings your Context accepts.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ezclassify/sdk-go"
)

func main() {
	server := flag.String("server", "http://127.0.0.1:4010", "EzClassify REST API base URL")
	ctxID := flag.String("ctx", "cedar-valley", "Context ID to use")
	ctxVersion := flag.String("ctx-version", "1.0.0", "Context version (or \"latest\")")
	flag.Parse()

	// The replay server asks for no credentials. A runtime usually does.
	credentials := ezclassify.RestClientCredentials{
		AccessToken:  os.Getenv("EZC_ACCESS_TOKEN"),
		ContextToken: os.Getenv("EZC_CONTEXT_TOKEN"),
	}
	client := ezclassify.New(*server, ezclassify.WithCredentials(credentials))
	ctx := context.Background()
	cx := ezclassify.ConnectWith(client, *ctxID, *ctxVersion)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Validate is the reader's call: it answers the facts a marking states and
	// whether they stand, which is what an access check needs.
	fmt.Println("=== validate INTERNAL ===")
	validateResp, err := cx.Scheme.Validate(ctx, ezclassify.CtxSchemeValidateRequest{Input: ptr("INTERNAL")})
	if err != nil {
		log.Fatalf("validate: %v", err)
	}
	if err := enc.Encode(validateResp); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== aggregate two district markings ===")
	aggResp, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{
		Inputs: ptr([]string{
			"INTERNAL//OAKWOOD/PINECREST//STUDENT TRANSPORTATION SERVICES//DO NOT FORWARD",
			"CONFIDENTIAL//FINANCIAL INFORMATION//OAKWOOD/PINECREST//REDACT BEFORE RELEASE",
		}),
	})
	if err != nil {
		log.Fatalf("aggregate: %v", err)
	}
	fmt.Printf("rendered: %s\n", aggResp.Result.Rendered)
	// Notices ride along on the result, so surfacing them takes no second call.
	for _, n := range aggResp.Result.Notices {
		fmt.Printf("notice (%s, required=%v): %s\n", n.Type, n.Required, n.Text)
	}

	fmt.Println("\n=== access check ===")
	attrs := map[string]interface{}{
		"clearance":               "CONFIDENTIAL",
		"recipient_organizations": []string{"CEDAR_VALLEY_DISTRICT"},
	}
	decision, err := cx.Access.Check(ctx, ezclassify.CtxAccessCheckRequest{
		Marking:        ptr("CONFIDENTIAL"),
		UserAttributes: attrs,
	})
	if err != nil {
		log.Fatalf("access check: %v", err)
	}
	fmt.Printf("granted: %v\n", decision.Result.Granted)
}

func ptr[T any](v T) *T { return &v }
