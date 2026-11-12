package ezclassify

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

// singleBodyTransport answers every request with 200 and one fixed body.
type singleBodyTransport struct{ body io.ReadCloser }

func (t singleBodyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {"application/json"}}, Body: t.body, Request: req}, nil
}

type closeErrorBody struct {
	io.Reader
	closes atomic.Int32
	err    error
}

func (b *closeErrorBody) Close() error {
	b.closes.Add(1)
	return b.err
}

func TestResponseBodyClosePreservesErrorAndRunsOnce(t *testing.T) {
	want := errors.New("close failed")
	body := &closeErrorBody{Reader: strings.NewReader(`{"contexts":[]}`), err: want}
	client := New("http://ezclassify.invalid",
		WithHTTPClient(&http.Client{Transport: singleBodyTransport{body}}))
	response, err := client.wire.CtxList(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.ReadAll(response.Body); err != nil {
		t.Fatal(err)
	}
	for range 2 {
		if err := response.Body.Close(); !errors.Is(err, want) {
			t.Errorf("Close error = %v, want %v", err, want)
		}
	}
	if got := body.closes.Load(); got != 1 {
		t.Errorf("underlying Close called %d times, want once", got)
	}
}
