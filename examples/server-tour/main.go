// Command server-tour lists the Contexts a server publishes, then processes a
// marking through one of them.
//
// Start the Cedar Valley replay server, which replays real runtime responses
// for cedar-valley@1.0.0, then run the command:
//
//	npx @ezclassify/conformance-server --scenario cedar-valley --port 4010
//	go run ./examples/server-tour
//
// The defaults target that server. To switch to your runtime, pass --server,
// --ctx, and --ctx-version, and set EZC_ACCESS_TOKEN and EZC_CONTEXT_TOKEN if
// your runtime requires them. Pass --marking with a marking your Context
// accepts.
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
	version := flag.String("ctx-version", "1.0.0", "Context version to bind")
	marking := flag.String("marking", "CONFIDENTIAL", "Marking string to process")
	flag.Parse()

	// The replay server asks for no credentials. A runtime usually does.
	credentials := ezclassify.RestClientCredentials{
		AccessToken:  os.Getenv("EZC_ACCESS_TOKEN"),
		ContextToken: os.Getenv("EZC_CONTEXT_TOKEN"),
	}
	client := ezclassify.New(*server, ezclassify.WithCredentials(credentials))
	ctx := context.Background()

	contexts, err := client.ListContexts(ctx)
	if err != nil {
		log.Fatalf("list contexts: %v", err)
	}
	fmt.Println("=== Available Contexts ===")
	for _, id := range contexts.Contexts {
		fmt.Printf("  %s\n", id)
	}

	cx := ezclassify.ConnectWith(client, *ctxID, *version)
	result, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: *marking})
	if err != nil {
		log.Fatalf("process %s@%s: %v", *ctxID, *version, err)
	}

	fmt.Printf("\n=== %s@%s process(%q) ===\n", *ctxID, *version, *marking)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(result); err != nil {
		log.Fatalf("encode: %v", err)
	}
	if resolved := cx.ResolvedVersion(); resolved != "" {
		fmt.Printf("Resolved version: %s\n", resolved)
	}
}
