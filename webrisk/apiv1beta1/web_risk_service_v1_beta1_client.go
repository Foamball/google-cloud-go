// Copyright 2022 Google LLC
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

package webrisk

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	webriskpb "google.golang.org/genproto/googleapis/cloud/webrisk/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newWebRiskServiceV1Beta1ClientHook clientHook

// WebRiskServiceV1Beta1CallOptions contains the retry settings for each method of WebRiskServiceV1Beta1Client.
type WebRiskServiceV1Beta1CallOptions struct {
	ComputeThreatListDiff []gax.CallOption
	SearchUris            []gax.CallOption
	SearchHashes          []gax.CallOption
}

func defaultWebRiskServiceV1Beta1GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("webrisk.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("webrisk.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://webrisk.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultWebRiskServiceV1Beta1CallOptions() *WebRiskServiceV1Beta1CallOptions {
	return &WebRiskServiceV1Beta1CallOptions{
		ComputeThreatListDiff: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchUris: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchHashes: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultWebRiskServiceV1Beta1RESTCallOptions() *WebRiskServiceV1Beta1CallOptions {
	return &WebRiskServiceV1Beta1CallOptions{
		ComputeThreatListDiff: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		SearchUris: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		SearchHashes: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalWebRiskServiceV1Beta1Client is an interface that defines the methods available from Web Risk API.
type internalWebRiskServiceV1Beta1Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ComputeThreatListDiff(context.Context, *webriskpb.ComputeThreatListDiffRequest, ...gax.CallOption) (*webriskpb.ComputeThreatListDiffResponse, error)
	SearchUris(context.Context, *webriskpb.SearchUrisRequest, ...gax.CallOption) (*webriskpb.SearchUrisResponse, error)
	SearchHashes(context.Context, *webriskpb.SearchHashesRequest, ...gax.CallOption) (*webriskpb.SearchHashesResponse, error)
}

// WebRiskServiceV1Beta1Client is a client for interacting with Web Risk API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Web Risk v1beta1 API defines an interface to detect malicious URLs on your
// website and in client applications.
type WebRiskServiceV1Beta1Client struct {
	// The internal transport-dependent client.
	internalClient internalWebRiskServiceV1Beta1Client

	// The call options for this service.
	CallOptions *WebRiskServiceV1Beta1CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *WebRiskServiceV1Beta1Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *WebRiskServiceV1Beta1Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *WebRiskServiceV1Beta1Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ComputeThreatListDiff gets the most recent threat list diffs.
func (c *WebRiskServiceV1Beta1Client) ComputeThreatListDiff(ctx context.Context, req *webriskpb.ComputeThreatListDiffRequest, opts ...gax.CallOption) (*webriskpb.ComputeThreatListDiffResponse, error) {
	return c.internalClient.ComputeThreatListDiff(ctx, req, opts...)
}

// SearchUris this method is used to check whether a URI is on a given threatList.
func (c *WebRiskServiceV1Beta1Client) SearchUris(ctx context.Context, req *webriskpb.SearchUrisRequest, opts ...gax.CallOption) (*webriskpb.SearchUrisResponse, error) {
	return c.internalClient.SearchUris(ctx, req, opts...)
}

// SearchHashes gets the full hashes that match the requested hash prefix.
// This is used after a hash prefix is looked up in a threatList
// and there is a match. The client side threatList only holds partial hashes
// so the client must query this method to determine if there is a full
// hash match of a threat.
func (c *WebRiskServiceV1Beta1Client) SearchHashes(ctx context.Context, req *webriskpb.SearchHashesRequest, opts ...gax.CallOption) (*webriskpb.SearchHashesResponse, error) {
	return c.internalClient.SearchHashes(ctx, req, opts...)
}

// webRiskServiceV1Beta1GRPCClient is a client for interacting with Web Risk API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type webRiskServiceV1Beta1GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing WebRiskServiceV1Beta1Client
	CallOptions **WebRiskServiceV1Beta1CallOptions

	// The gRPC API client.
	webRiskServiceV1Beta1Client webriskpb.WebRiskServiceV1Beta1Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewWebRiskServiceV1Beta1Client creates a new web risk service v1 beta1 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Web Risk v1beta1 API defines an interface to detect malicious URLs on your
// website and in client applications.
func NewWebRiskServiceV1Beta1Client(ctx context.Context, opts ...option.ClientOption) (*WebRiskServiceV1Beta1Client, error) {
	clientOpts := defaultWebRiskServiceV1Beta1GRPCClientOptions()
	if newWebRiskServiceV1Beta1ClientHook != nil {
		hookOpts, err := newWebRiskServiceV1Beta1ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := WebRiskServiceV1Beta1Client{CallOptions: defaultWebRiskServiceV1Beta1CallOptions()}

	c := &webRiskServiceV1Beta1GRPCClient{
		connPool:                    connPool,
		disableDeadlines:            disableDeadlines,
		webRiskServiceV1Beta1Client: webriskpb.NewWebRiskServiceV1Beta1Client(connPool),
		CallOptions:                 &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *webRiskServiceV1Beta1GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *webRiskServiceV1Beta1GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *webRiskServiceV1Beta1GRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type webRiskServiceV1Beta1RESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing WebRiskServiceV1Beta1Client
	CallOptions **WebRiskServiceV1Beta1CallOptions
}

// NewWebRiskServiceV1Beta1RESTClient creates a new web risk service v1 beta1 rest client.
//
// Web Risk v1beta1 API defines an interface to detect malicious URLs on your
// website and in client applications.
func NewWebRiskServiceV1Beta1RESTClient(ctx context.Context, opts ...option.ClientOption) (*WebRiskServiceV1Beta1Client, error) {
	clientOpts := append(defaultWebRiskServiceV1Beta1RESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultWebRiskServiceV1Beta1RESTCallOptions()
	c := &webRiskServiceV1Beta1RESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &WebRiskServiceV1Beta1Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultWebRiskServiceV1Beta1RESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://webrisk.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://webrisk.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://webrisk.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *webRiskServiceV1Beta1RESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *webRiskServiceV1Beta1RESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *webRiskServiceV1Beta1RESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *webRiskServiceV1Beta1GRPCClient) ComputeThreatListDiff(ctx context.Context, req *webriskpb.ComputeThreatListDiffRequest, opts ...gax.CallOption) (*webriskpb.ComputeThreatListDiffResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ComputeThreatListDiff[0:len((*c.CallOptions).ComputeThreatListDiff):len((*c.CallOptions).ComputeThreatListDiff)], opts...)
	var resp *webriskpb.ComputeThreatListDiffResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webRiskServiceV1Beta1Client.ComputeThreatListDiff(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *webRiskServiceV1Beta1GRPCClient) SearchUris(ctx context.Context, req *webriskpb.SearchUrisRequest, opts ...gax.CallOption) (*webriskpb.SearchUrisResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SearchUris[0:len((*c.CallOptions).SearchUris):len((*c.CallOptions).SearchUris)], opts...)
	var resp *webriskpb.SearchUrisResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webRiskServiceV1Beta1Client.SearchUris(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *webRiskServiceV1Beta1GRPCClient) SearchHashes(ctx context.Context, req *webriskpb.SearchHashesRequest, opts ...gax.CallOption) (*webriskpb.SearchHashesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SearchHashes[0:len((*c.CallOptions).SearchHashes):len((*c.CallOptions).SearchHashes)], opts...)
	var resp *webriskpb.SearchHashesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webRiskServiceV1Beta1Client.SearchHashes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ComputeThreatListDiff gets the most recent threat list diffs.
func (c *webRiskServiceV1Beta1RESTClient) ComputeThreatListDiff(ctx context.Context, req *webriskpb.ComputeThreatListDiffRequest, opts ...gax.CallOption) (*webriskpb.ComputeThreatListDiffResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/threatLists:computeDiff")

	params := url.Values{}
	if req.GetConstraints().GetMaxDatabaseEntries() != 0 {
		params.Add("constraints.maxDatabaseEntries", fmt.Sprintf("%v", req.GetConstraints().GetMaxDatabaseEntries()))
	}
	if req.GetConstraints().GetMaxDiffEntries() != 0 {
		params.Add("constraints.maxDiffEntries", fmt.Sprintf("%v", req.GetConstraints().GetMaxDiffEntries()))
	}
	if req.GetConstraints().GetSupportedCompressions() != nil {
		params.Add("constraints.supportedCompressions", fmt.Sprintf("%v", req.GetConstraints().GetSupportedCompressions()))
	}
	params.Add("threatType", fmt.Sprintf("%v", req.GetThreatType()))
	if req.GetVersionToken() != nil {
		params.Add("versionToken", fmt.Sprintf("%v", req.GetVersionToken()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).ComputeThreatListDiff[0:len((*c.CallOptions).ComputeThreatListDiff):len((*c.CallOptions).ComputeThreatListDiff)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &webriskpb.ComputeThreatListDiffResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
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

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// SearchUris this method is used to check whether a URI is on a given threatList.
func (c *webRiskServiceV1Beta1RESTClient) SearchUris(ctx context.Context, req *webriskpb.SearchUrisRequest, opts ...gax.CallOption) (*webriskpb.SearchUrisResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/uris:search")

	params := url.Values{}
	if req.GetThreatTypes() != nil {
		params.Add("threatTypes", fmt.Sprintf("%v", req.GetThreatTypes()))
	}
	params.Add("uri", fmt.Sprintf("%v", req.GetUri()))

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).SearchUris[0:len((*c.CallOptions).SearchUris):len((*c.CallOptions).SearchUris)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &webriskpb.SearchUrisResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
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

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// SearchHashes gets the full hashes that match the requested hash prefix.
// This is used after a hash prefix is looked up in a threatList
// and there is a match. The client side threatList only holds partial hashes
// so the client must query this method to determine if there is a full
// hash match of a threat.
func (c *webRiskServiceV1Beta1RESTClient) SearchHashes(ctx context.Context, req *webriskpb.SearchHashesRequest, opts ...gax.CallOption) (*webriskpb.SearchHashesResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/hashes:search")

	params := url.Values{}
	if req.GetHashPrefix() != nil {
		params.Add("hashPrefix", fmt.Sprintf("%v", req.GetHashPrefix()))
	}
	if req.GetThreatTypes() != nil {
		params.Add("threatTypes", fmt.Sprintf("%v", req.GetThreatTypes()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).SearchHashes[0:len((*c.CallOptions).SearchHashes):len((*c.CallOptions).SearchHashes)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &webriskpb.SearchHashesResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
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

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
