package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func failRequestEditor(_ context.Context, _ *http.Request) error {
	return errors.New("editor error")
}

// The editor path is generator infrastructure shared by every operation; one
// discovery GET and one Context-bound POST cover both generated call shapes.
func TestRequestEditorErrors(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("request reached transport after editor failure")
	}))
	t.Cleanup(srv.Close)
	c, err := NewClientWithResponses(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	editor := RequestEditorFn(failRequestEditor)
	if _, err := c.CtxListWithResponse(context.Background(), nil, editor); err == nil {
		t.Fatal("CtxList accepted failing editor")
	}
	if _, err := c.CtxSchemeProcessWithResponse(
		context.Background(), "ctx", "latest", nil, CtxSchemeProcessJSONRequestBody{}, editor,
	); err == nil {
		t.Fatal("CtxSchemeProcess accepted failing editor")
	}
}
