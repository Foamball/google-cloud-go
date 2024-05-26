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

package appengine

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newAuthorizedDomainsClientHook clientHook

// AuthorizedDomainsCallOptions contains the retry settings for each method of AuthorizedDomainsClient.
type AuthorizedDomainsCallOptions struct {
	ListAuthorizedDomains []gax.CallOption
}

func defaultAuthorizedDomainsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("appengine.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("appengine.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("appengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAuthorizedDomainsCallOptions() *AuthorizedDomainsCallOptions {
	return &AuthorizedDomainsCallOptions{
		ListAuthorizedDomains: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultAuthorizedDomainsRESTCallOptions() *AuthorizedDomainsCallOptions {
	return &AuthorizedDomainsCallOptions{
		ListAuthorizedDomains: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalAuthorizedDomainsClient is an interface that defines the methods available from App Engine Admin API.
type internalAuthorizedDomainsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListAuthorizedDomains(context.Context, *appenginepb.ListAuthorizedDomainsRequest, ...gax.CallOption) *AuthorizedDomainIterator
}

// AuthorizedDomainsClient is a client for interacting with App Engine Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages domains a user is authorized to administer. To authorize use of a
// domain, verify ownership via
// Webmaster Central (at https://www.google.com/webmasters/verification/home).
type AuthorizedDomainsClient struct {
	// The internal transport-dependent client.
	internalClient internalAuthorizedDomainsClient

	// The call options for this service.
	CallOptions *AuthorizedDomainsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AuthorizedDomainsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AuthorizedDomainsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AuthorizedDomainsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListAuthorizedDomains lists all domains the user is authorized to administer.
func (c *AuthorizedDomainsClient) ListAuthorizedDomains(ctx context.Context, req *appenginepb.ListAuthorizedDomainsRequest, opts ...gax.CallOption) *AuthorizedDomainIterator {
	return c.internalClient.ListAuthorizedDomains(ctx, req, opts...)
}

// authorizedDomainsGRPCClient is a client for interacting with App Engine Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type authorizedDomainsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AuthorizedDomainsClient
	CallOptions **AuthorizedDomainsCallOptions

	// The gRPC API client.
	authorizedDomainsClient appenginepb.AuthorizedDomainsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewAuthorizedDomainsClient creates a new authorized domains client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages domains a user is authorized to administer. To authorize use of a
// domain, verify ownership via
// Webmaster Central (at https://www.google.com/webmasters/verification/home).
func NewAuthorizedDomainsClient(ctx context.Context, opts ...option.ClientOption) (*AuthorizedDomainsClient, error) {
	clientOpts := defaultAuthorizedDomainsGRPCClientOptions()
	if newAuthorizedDomainsClientHook != nil {
		hookOpts, err := newAuthorizedDomainsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AuthorizedDomainsClient{CallOptions: defaultAuthorizedDomainsCallOptions()}

	c := &authorizedDomainsGRPCClient{
		connPool:                connPool,
		authorizedDomainsClient: appenginepb.NewAuthorizedDomainsClient(connPool),
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
func (c *authorizedDomainsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *authorizedDomainsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *authorizedDomainsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type authorizedDomainsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing AuthorizedDomainsClient
	CallOptions **AuthorizedDomainsCallOptions
}

// NewAuthorizedDomainsRESTClient creates a new authorized domains rest client.
//
// Manages domains a user is authorized to administer. To authorize use of a
// domain, verify ownership via
// Webmaster Central (at https://www.google.com/webmasters/verification/home).
func NewAuthorizedDomainsRESTClient(ctx context.Context, opts ...option.ClientOption) (*AuthorizedDomainsClient, error) {
	clientOpts := append(defaultAuthorizedDomainsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultAuthorizedDomainsRESTCallOptions()
	c := &authorizedDomainsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &AuthorizedDomainsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultAuthorizedDomainsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://appengine.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://appengine.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://appengine.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *authorizedDomainsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *authorizedDomainsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *authorizedDomainsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *authorizedDomainsGRPCClient) ListAuthorizedDomains(ctx context.Context, req *appenginepb.ListAuthorizedDomainsRequest, opts ...gax.CallOption) *AuthorizedDomainIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListAuthorizedDomains[0:len((*c.CallOptions).ListAuthorizedDomains):len((*c.CallOptions).ListAuthorizedDomains)], opts...)
	it := &AuthorizedDomainIterator{}
	req = proto.Clone(req).(*appenginepb.ListAuthorizedDomainsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.AuthorizedDomain, string, error) {
		resp := &appenginepb.ListAuthorizedDomainsResponse{}
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
			resp, err = c.authorizedDomainsClient.ListAuthorizedDomains(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDomains(), resp.GetNextPageToken(), nil
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

// ListAuthorizedDomains lists all domains the user is authorized to administer.
func (c *authorizedDomainsRESTClient) ListAuthorizedDomains(ctx context.Context, req *appenginepb.ListAuthorizedDomainsRequest, opts ...gax.CallOption) *AuthorizedDomainIterator {
	it := &AuthorizedDomainIterator{}
	req = proto.Clone(req).(*appenginepb.ListAuthorizedDomainsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.AuthorizedDomain, string, error) {
		resp := &appenginepb.ListAuthorizedDomainsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/authorizedDomains", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
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
		return resp.GetDomains(), resp.GetNextPageToken(), nil
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
