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

package conversions

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	conversionspb "cloud.google.com/go/shopping/merchant/conversions/apiv1beta/conversionspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newConversionSourcesClientHook clientHook

// ConversionSourcesCallOptions contains the retry settings for each method of ConversionSourcesClient.
type ConversionSourcesCallOptions struct {
	CreateConversionSource   []gax.CallOption
	UpdateConversionSource   []gax.CallOption
	DeleteConversionSource   []gax.CallOption
	UndeleteConversionSource []gax.CallOption
	GetConversionSource      []gax.CallOption
	ListConversionSources    []gax.CallOption
}

func defaultConversionSourcesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("merchantapi.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("merchantapi.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("merchantapi.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultConversionSourcesCallOptions() *ConversionSourcesCallOptions {
	return &ConversionSourcesCallOptions{
		CreateConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UndeleteConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListConversionSources: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultConversionSourcesRESTCallOptions() *ConversionSourcesCallOptions {
	return &ConversionSourcesCallOptions{
		CreateConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		UpdateConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeleteConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		UndeleteConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetConversionSource: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListConversionSources: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalConversionSourcesClient is an interface that defines the methods available from Merchant API.
type internalConversionSourcesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateConversionSource(context.Context, *conversionspb.CreateConversionSourceRequest, ...gax.CallOption) (*conversionspb.ConversionSource, error)
	UpdateConversionSource(context.Context, *conversionspb.UpdateConversionSourceRequest, ...gax.CallOption) (*conversionspb.ConversionSource, error)
	DeleteConversionSource(context.Context, *conversionspb.DeleteConversionSourceRequest, ...gax.CallOption) error
	UndeleteConversionSource(context.Context, *conversionspb.UndeleteConversionSourceRequest, ...gax.CallOption) (*conversionspb.ConversionSource, error)
	GetConversionSource(context.Context, *conversionspb.GetConversionSourceRequest, ...gax.CallOption) (*conversionspb.ConversionSource, error)
	ListConversionSources(context.Context, *conversionspb.ListConversionSourcesRequest, ...gax.CallOption) *ConversionSourceIterator
}

// ConversionSourcesClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing conversion sources for a merchant account.
type ConversionSourcesClient struct {
	// The internal transport-dependent client.
	internalClient internalConversionSourcesClient

	// The call options for this service.
	CallOptions *ConversionSourcesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConversionSourcesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConversionSourcesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ConversionSourcesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateConversionSource creates a new conversion source.
func (c *ConversionSourcesClient) CreateConversionSource(ctx context.Context, req *conversionspb.CreateConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	return c.internalClient.CreateConversionSource(ctx, req, opts...)
}

// UpdateConversionSource updates information of an existing conversion source. Available only for
// Merchant Center Destination conversion sources.
func (c *ConversionSourcesClient) UpdateConversionSource(ctx context.Context, req *conversionspb.UpdateConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	return c.internalClient.UpdateConversionSource(ctx, req, opts...)
}

// DeleteConversionSource archives an existing conversion source. If the conversion source is a
// Merchant Center Destination, it will be recoverable for 30 days. If the
// conversion source is a Google Analytics Link, it will be deleted
// immediately and can be restored by creating a new one.
func (c *ConversionSourcesClient) DeleteConversionSource(ctx context.Context, req *conversionspb.DeleteConversionSourceRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteConversionSource(ctx, req, opts...)
}

// UndeleteConversionSource re-enables an archived conversion source. Only Available for Merchant
// Center Destination conversion sources.
func (c *ConversionSourcesClient) UndeleteConversionSource(ctx context.Context, req *conversionspb.UndeleteConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	return c.internalClient.UndeleteConversionSource(ctx, req, opts...)
}

// GetConversionSource fetches a conversion source.
func (c *ConversionSourcesClient) GetConversionSource(ctx context.Context, req *conversionspb.GetConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	return c.internalClient.GetConversionSource(ctx, req, opts...)
}

// ListConversionSources retrieves the list of conversion sources the caller has access to.
func (c *ConversionSourcesClient) ListConversionSources(ctx context.Context, req *conversionspb.ListConversionSourcesRequest, opts ...gax.CallOption) *ConversionSourceIterator {
	return c.internalClient.ListConversionSources(ctx, req, opts...)
}

// conversionSourcesGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type conversionSourcesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ConversionSourcesClient
	CallOptions **ConversionSourcesCallOptions

	// The gRPC API client.
	conversionSourcesClient conversionspb.ConversionSourcesServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewConversionSourcesClient creates a new conversion sources service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing conversion sources for a merchant account.
func NewConversionSourcesClient(ctx context.Context, opts ...option.ClientOption) (*ConversionSourcesClient, error) {
	clientOpts := defaultConversionSourcesGRPCClientOptions()
	if newConversionSourcesClientHook != nil {
		hookOpts, err := newConversionSourcesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ConversionSourcesClient{CallOptions: defaultConversionSourcesCallOptions()}

	c := &conversionSourcesGRPCClient{
		connPool:                connPool,
		conversionSourcesClient: conversionspb.NewConversionSourcesServiceClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *conversionSourcesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *conversionSourcesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *conversionSourcesGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type conversionSourcesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing ConversionSourcesClient
	CallOptions **ConversionSourcesCallOptions
}

// NewConversionSourcesRESTClient creates a new conversion sources service rest client.
//
// Service for managing conversion sources for a merchant account.
func NewConversionSourcesRESTClient(ctx context.Context, opts ...option.ClientOption) (*ConversionSourcesClient, error) {
	clientOpts := append(defaultConversionSourcesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultConversionSourcesRESTCallOptions()
	c := &conversionSourcesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &ConversionSourcesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultConversionSourcesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://merchantapi.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://merchantapi.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://merchantapi.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *conversionSourcesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *conversionSourcesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *conversionSourcesRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *conversionSourcesGRPCClient) CreateConversionSource(ctx context.Context, req *conversionspb.CreateConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateConversionSource[0:len((*c.CallOptions).CreateConversionSource):len((*c.CallOptions).CreateConversionSource)], opts...)
	var resp *conversionspb.ConversionSource
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionSourcesClient.CreateConversionSource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conversionSourcesGRPCClient) UpdateConversionSource(ctx context.Context, req *conversionspb.UpdateConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "conversion_source.name", url.QueryEscape(req.GetConversionSource().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateConversionSource[0:len((*c.CallOptions).UpdateConversionSource):len((*c.CallOptions).UpdateConversionSource)], opts...)
	var resp *conversionspb.ConversionSource
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionSourcesClient.UpdateConversionSource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conversionSourcesGRPCClient) DeleteConversionSource(ctx context.Context, req *conversionspb.DeleteConversionSourceRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteConversionSource[0:len((*c.CallOptions).DeleteConversionSource):len((*c.CallOptions).DeleteConversionSource)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.conversionSourcesClient.DeleteConversionSource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *conversionSourcesGRPCClient) UndeleteConversionSource(ctx context.Context, req *conversionspb.UndeleteConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UndeleteConversionSource[0:len((*c.CallOptions).UndeleteConversionSource):len((*c.CallOptions).UndeleteConversionSource)], opts...)
	var resp *conversionspb.ConversionSource
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionSourcesClient.UndeleteConversionSource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conversionSourcesGRPCClient) GetConversionSource(ctx context.Context, req *conversionspb.GetConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetConversionSource[0:len((*c.CallOptions).GetConversionSource):len((*c.CallOptions).GetConversionSource)], opts...)
	var resp *conversionspb.ConversionSource
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionSourcesClient.GetConversionSource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conversionSourcesGRPCClient) ListConversionSources(ctx context.Context, req *conversionspb.ListConversionSourcesRequest, opts ...gax.CallOption) *ConversionSourceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListConversionSources[0:len((*c.CallOptions).ListConversionSources):len((*c.CallOptions).ListConversionSources)], opts...)
	it := &ConversionSourceIterator{}
	req = proto.Clone(req).(*conversionspb.ListConversionSourcesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*conversionspb.ConversionSource, string, error) {
		resp := &conversionspb.ListConversionSourcesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.conversionSourcesClient.ListConversionSources(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetConversionSources(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// CreateConversionSource creates a new conversion source.
func (c *conversionSourcesRESTClient) CreateConversionSource(ctx context.Context, req *conversionspb.CreateConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetConversionSource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/conversions/v1beta/%v/conversionSources", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateConversionSource[0:len((*c.CallOptions).CreateConversionSource):len((*c.CallOptions).CreateConversionSource)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &conversionspb.ConversionSource{}
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

// UpdateConversionSource updates information of an existing conversion source. Available only for
// Merchant Center Destination conversion sources.
func (c *conversionSourcesRESTClient) UpdateConversionSource(ctx context.Context, req *conversionspb.UpdateConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetConversionSource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/conversions/v1beta/%v", req.GetConversionSource().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask[1:len(updateMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "conversion_source.name", url.QueryEscape(req.GetConversionSource().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateConversionSource[0:len((*c.CallOptions).UpdateConversionSource):len((*c.CallOptions).UpdateConversionSource)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &conversionspb.ConversionSource{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
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

// DeleteConversionSource archives an existing conversion source. If the conversion source is a
// Merchant Center Destination, it will be recoverable for 30 days. If the
// conversion source is a Google Analytics Link, it will be deleted
// immediately and can be restored by creating a new one.
func (c *conversionSourcesRESTClient) DeleteConversionSource(ctx context.Context, req *conversionspb.DeleteConversionSourceRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/conversions/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
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

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// UndeleteConversionSource re-enables an archived conversion source. Only Available for Merchant
// Center Destination conversion sources.
func (c *conversionSourcesRESTClient) UndeleteConversionSource(ctx context.Context, req *conversionspb.UndeleteConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/conversions/v1beta/%v:undelete", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UndeleteConversionSource[0:len((*c.CallOptions).UndeleteConversionSource):len((*c.CallOptions).UndeleteConversionSource)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &conversionspb.ConversionSource{}
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

// GetConversionSource fetches a conversion source.
func (c *conversionSourcesRESTClient) GetConversionSource(ctx context.Context, req *conversionspb.GetConversionSourceRequest, opts ...gax.CallOption) (*conversionspb.ConversionSource, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/conversions/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetConversionSource[0:len((*c.CallOptions).GetConversionSource):len((*c.CallOptions).GetConversionSource)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &conversionspb.ConversionSource{}
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

// ListConversionSources retrieves the list of conversion sources the caller has access to.
func (c *conversionSourcesRESTClient) ListConversionSources(ctx context.Context, req *conversionspb.ListConversionSourcesRequest, opts ...gax.CallOption) *ConversionSourceIterator {
	it := &ConversionSourceIterator{}
	req = proto.Clone(req).(*conversionspb.ListConversionSourcesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*conversionspb.ConversionSource, string, error) {
		resp := &conversionspb.ListConversionSourcesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/conversions/v1beta/%v/conversionSources", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetShowDeleted() {
			params.Add("showDeleted", fmt.Sprintf("%v", req.GetShowDeleted()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
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
			return nil, "", e
		}
		it.Response = resp
		return resp.GetConversionSources(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
