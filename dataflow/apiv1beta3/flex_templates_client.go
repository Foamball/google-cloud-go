// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package dataflow

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"

	dataflowpb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newFlexTemplatesClientHook clientHook

// FlexTemplatesCallOptions contains the retry settings for each method of FlexTemplatesClient.
type FlexTemplatesCallOptions struct {
	LaunchFlexTemplate []gax.CallOption
}

func defaultFlexTemplatesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("dataflow.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFlexTemplatesCallOptions() *FlexTemplatesCallOptions {
	return &FlexTemplatesCallOptions{
		LaunchFlexTemplate: []gax.CallOption{},
	}
}

func defaultFlexTemplatesRESTCallOptions() *FlexTemplatesCallOptions {
	return &FlexTemplatesCallOptions{
		LaunchFlexTemplate: []gax.CallOption{},
	}
}

// internalFlexTemplatesClient is an interface that defines the methods available from Dataflow API.
type internalFlexTemplatesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	LaunchFlexTemplate(context.Context, *dataflowpb.LaunchFlexTemplateRequest, ...gax.CallOption) (*dataflowpb.LaunchFlexTemplateResponse, error)
}

// FlexTemplatesClient is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides a service for Flex templates. This feature is not ready yet.
type FlexTemplatesClient struct {
	// The internal transport-dependent client.
	internalClient internalFlexTemplatesClient

	// The call options for this service.
	CallOptions *FlexTemplatesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FlexTemplatesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FlexTemplatesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *FlexTemplatesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// LaunchFlexTemplate launch a job with a FlexTemplate.
func (c *FlexTemplatesClient) LaunchFlexTemplate(ctx context.Context, req *dataflowpb.LaunchFlexTemplateRequest, opts ...gax.CallOption) (*dataflowpb.LaunchFlexTemplateResponse, error) {
	return c.internalClient.LaunchFlexTemplate(ctx, req, opts...)
}

// flexTemplatesGRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type flexTemplatesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing FlexTemplatesClient
	CallOptions **FlexTemplatesCallOptions

	// The gRPC API client.
	flexTemplatesClient dataflowpb.FlexTemplatesServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewFlexTemplatesClient creates a new flex templates service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides a service for Flex templates. This feature is not ready yet.
func NewFlexTemplatesClient(ctx context.Context, opts ...option.ClientOption) (*FlexTemplatesClient, error) {
	clientOpts := defaultFlexTemplatesGRPCClientOptions()
	if newFlexTemplatesClientHook != nil {
		hookOpts, err := newFlexTemplatesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := FlexTemplatesClient{CallOptions: defaultFlexTemplatesCallOptions()}

	c := &flexTemplatesGRPCClient{
		connPool:            connPool,
		flexTemplatesClient: dataflowpb.NewFlexTemplatesServiceClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *flexTemplatesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *flexTemplatesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *flexTemplatesGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type flexTemplatesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing FlexTemplatesClient
	CallOptions **FlexTemplatesCallOptions
}

// NewFlexTemplatesRESTClient creates a new flex templates service rest client.
//
// Provides a service for Flex templates. This feature is not ready yet.
func NewFlexTemplatesRESTClient(ctx context.Context, opts ...option.ClientOption) (*FlexTemplatesClient, error) {
	clientOpts := append(defaultFlexTemplatesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultFlexTemplatesRESTCallOptions()
	c := &flexTemplatesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &FlexTemplatesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultFlexTemplatesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://dataflow.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://dataflow.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://dataflow.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *flexTemplatesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *flexTemplatesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *flexTemplatesRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *flexTemplatesGRPCClient) LaunchFlexTemplate(ctx context.Context, req *dataflowpb.LaunchFlexTemplateRequest, opts ...gax.CallOption) (*dataflowpb.LaunchFlexTemplateResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).LaunchFlexTemplate[0:len((*c.CallOptions).LaunchFlexTemplate):len((*c.CallOptions).LaunchFlexTemplate)], opts...)
	var resp *dataflowpb.LaunchFlexTemplateResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.flexTemplatesClient.LaunchFlexTemplate(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LaunchFlexTemplate launch a job with a FlexTemplate.
func (c *flexTemplatesRESTClient) LaunchFlexTemplate(ctx context.Context, req *dataflowpb.LaunchFlexTemplateRequest, opts ...gax.CallOption) (*dataflowpb.LaunchFlexTemplateResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/flexTemplates:launch", req.GetProjectId(), req.GetLocation())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).LaunchFlexTemplate[0:len((*c.CallOptions).LaunchFlexTemplate):len((*c.CallOptions).LaunchFlexTemplate)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.LaunchFlexTemplateResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
