// Package gpt4oimage provides the Gpt4o Image API client.
package gpt4oimage

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const generationsPath = "/api/v1/gpt4o_image/generations"

type Client struct {
	Generations *Generations
}

func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{Generations: &Generations{http: httpClient}}
}

type Generations struct{ http core.HTTPClient }

func (r *Generations) Create(ctx context.Context, params GenerationParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, generationsPath, core.CompactParams(params), requestOptions)
}

func (r *Generations) Get(ctx context.Context, id string, opts ...option.RequestOption) (*GenerationResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[GenerationResponse](ctx, r.http, core.ResourcePath(generationsPath, id), requestOptions)
}

func (r *Generations) Run(ctx context.Context, params GenerationParams, opts ...option.RequestOption) (*GenerationResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*GenerationResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}
