// Command go-classify is a runnable EzClassify demo that scans a directory
// of text files for marking strings, validates each marking, and produces a
// per-file classification report on stdout.
//
// Usage:
//
//	go run ./examples/go-classify --server http://localhost:3000 --ctx cedar-valley-protection --dir ./docs
//
// The command:
//  1. Walks every *.txt and *.md file under --dir.
//  2. Calls the Context-bound scan operation on each file's text.
//  3. Calls the Context-bound aggregate operation across valid matches to compute the
//     file's aggregate sensitivity level.
//  4. Prints a structured summary table to stdout.
//
// This command sends arbitrary file text, which the Cedar Valley replay server
// has no recording for, so it needs a runtime that serves the Context. Set
// EZC_ACCESS_TOKEN and EZC_CONTEXT_TOKEN if your runtime requires them. The
// example's README has the setup.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/ezclassify/sdk-go"
)

func main() {
	server := flag.String("server", "http://localhost:3000", "EzClassify REST API base URL")
	ctxID := flag.String("ctx", "cedar-valley-protection", "Context ID to use")
	ctxVersion := flag.String("ctx-version", "latest", "Context version to bind")
	dir := flag.String("dir", ".", "Directory to scan for text files")
	timeout := flag.Duration("timeout", 30*time.Second, "HTTP request timeout")
	flag.Parse()
	if *timeout <= 0 {
		log.Fatal("--timeout must be positive")
	}

	// The replay server asks for no credentials. A runtime usually does.
	credentials := ezclassify.RestClientCredentials{
		AccessToken:  os.Getenv("EZC_ACCESS_TOKEN"),
		ContextToken: os.Getenv("EZC_CONTEXT_TOKEN"),
	}
	client := ezclassify.New(*server, ezclassify.WithTimeout(*timeout), ezclassify.WithCredentials(credentials))
	ctx := context.Background()
	cx := ezclassify.ConnectWith(client, *ctxID, *ctxVersion)

	// Collect all scannable files.
	var files []string
	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".txt" || ext == ".md" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("walk %s: %v", *dir, err)
	}
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No .txt or .md files found in %s\n", *dir)
		os.Exit(1)
	}

	type fileReport struct {
		Path      string
		Matches   int
		Valid     int
		Invalid   int
		Aggregate string
		Error     string
	}

	reports := make([]fileReport, 0, len(files))

	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			reports = append(reports, fileReport{Path: path, Error: err.Error()})
			continue
		}

		// Scan the file. Each match carries what validate answers for it,
		// facts and validity included, so no second call is needed.
		scanResp, err := cx.Scheme.Scan(ctx, ezclassify.CtxSchemeScanRequest{Text: string(data)})
		if err != nil {
			reports = append(reports, fileReport{Path: path, Error: fmt.Sprintf("scan: %v", err)})
			continue
		}
		// The scan answers with the markings and how far it got. A short
		// list from a scan that stopped is not the same fact as a clean
		// file, so the limit is read rather than inferred from the length.
		matches := scanResp.Result.Results
		if string(scanResp.Result.Limits.Status) != "complete" {
			reports = append(reports, fileReport{
				Path:  path,
				Error: fmt.Sprintf("scan did not read the whole file: %s", scanResp.Result.Limits.Status),
			})
			continue
		}
		if len(matches) == 0 {
			reports = append(reports, fileReport{Path: path, Matches: 0, Aggregate: "(none)"})
			continue
		}

		validCount, invalidCount := 0, 0
		var renderedMarkings []string
		for _, m := range matches {
			if m.Result.Valid {
				validCount++
				renderedMarkings = append(renderedMarkings, m.Result.Rendered)
			} else {
				invalidCount++
			}
		}

		// Aggregate all valid markings to determine the file's sensitivity level.
		aggRendered := "(none)"
		if len(renderedMarkings) > 0 {
			aggResp, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{Inputs: &renderedMarkings})
			if err != nil {
				aggRendered = fmt.Sprintf("(error: %v)", err)
			} else {
				aggRendered = aggResp.Result.Rendered
			}
		}

		reports = append(reports, fileReport{
			Path:      path,
			Matches:   len(matches),
			Valid:     validCount,
			Invalid:   invalidCount,
			Aggregate: aggRendered,
		})
	}

	// Print the report table.
	fmt.Printf("\n=== EzClassify File Scan Report ===\n")
	fmt.Printf("Server: %s  Context: %s@%s  Dir: %s\n\n", *server, *ctxID, *ctxVersion, *dir)

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "FILE\tMATCHES\tVALID\tINVALID\tAGGREGATE")
	fmt.Fprintln(tw, "----\t-------\t-----\t-------\t---------")

	for _, r := range reports {
		if r.Error != "" {
			fmt.Fprintf(tw, "%s\t-\t-\t-\tERROR: %s\n", r.Path, r.Error)
		} else {
			fmt.Fprintf(tw, "%s\t%d\t%d\t%d\t%s\n",
				r.Path, r.Matches, r.Valid, r.Invalid, r.Aggregate)
		}
	}
	if err := tw.Flush(); err != nil {
		log.Fatalf("flush tabwriter: %v", err)
	}

	// Print summary.
	totalMatches, totalValid, totalInvalid := 0, 0, 0
	for _, r := range reports {
		totalMatches += r.Matches
		totalValid += r.Valid
		totalInvalid += r.Invalid
	}
	fmt.Printf("\nSummary: %d file(s) scanned, %d marking(s) found (%d valid, %d invalid)\n",
		len(files), totalMatches, totalValid, totalInvalid)
}
