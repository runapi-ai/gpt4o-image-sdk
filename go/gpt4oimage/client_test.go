package gpt4oimage

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
)

type stubHTTPClient struct {
	method   string
	path     string
	body     any
	response json.RawMessage
}

func (s *stubHTTPClient) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method = method
	s.path = path
	if opts != nil {
		s.body = opts.Body
	}
	return s.response, nil
}

func boolPtr(v bool) *bool { return &v }

func TestGenerationsCreate(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"id":"task_gen_123","status":"processing"}`)}
	client := NewClientWithHTTP(stub)
	resp, err := client.Generations.Create(context.Background(), GenerationParams{
		Model:          "gpt4o-image",
		Prompt:         "a still life",
		Size:           "1:1",
		FilesURL:       []string{"https://example.com/input.png"},
		NVariants:      2,
		EnableFallback: boolPtr(true),
		FallbackModel:  "FLUX_MAX",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gpt4o_image/generations" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "gpt4o-image" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
	if body["n_variants"] != float64(2) {
		t.Fatalf("unexpected n_variants: %v", body["n_variants"])
	}
	if resp.ID != "task_gen_123" {
		t.Fatalf("unexpected task ID: %v", resp.ID)
	}
}

func TestGenerationsGet(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"id":"task_gen_456","status":"completed","progress":"1.00","images":[{"url":"https://file.runapi.ai/result.png"}]}`)}
	client := NewClientWithHTTP(stub)
	resp, err := client.Generations.Get(context.Background(), "task_gen_abc")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gpt4o_image/generations/task_gen_abc" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	if resp.ID != "task_gen_456" {
		t.Fatalf("unexpected ID: %v", resp.ID)
	}
}
