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

package dialogflow

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2beta1"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newSessionsClientHook clientHook

// SessionsCallOptions contains the retry settings for each method of SessionsClient.
type SessionsCallOptions struct {
	DetectIntent          []gax.CallOption
	StreamingDetectIntent []gax.CallOption
	GetLocation           []gax.CallOption
	ListLocations         []gax.CallOption
	CancelOperation       []gax.CallOption
	GetOperation          []gax.CallOption
	ListOperations        []gax.CallOption
}

func defaultSessionsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSessionsCallOptions() *SessionsCallOptions {
	return &SessionsCallOptions{
		DetectIntent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		StreamingDetectIntent: []gax.CallOption{},
		GetLocation:           []gax.CallOption{},
		ListLocations:         []gax.CallOption{},
		CancelOperation:       []gax.CallOption{},
		GetOperation:          []gax.CallOption{},
		ListOperations:        []gax.CallOption{},
	}
}

func defaultSessionsRESTCallOptions() *SessionsCallOptions {
	return &SessionsCallOptions{
		DetectIntent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		StreamingDetectIntent: []gax.CallOption{},
		GetLocation:           []gax.CallOption{},
		ListLocations:         []gax.CallOption{},
		CancelOperation:       []gax.CallOption{},
		GetOperation:          []gax.CallOption{},
		ListOperations:        []gax.CallOption{},
	}
}

// internalSessionsClient is an interface that defines the methods available from Dialogflow API.
type internalSessionsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	DetectIntent(context.Context, *dialogflowpb.DetectIntentRequest, ...gax.CallOption) (*dialogflowpb.DetectIntentResponse, error)
	StreamingDetectIntent(context.Context, ...gax.CallOption) (dialogflowpb.Sessions_StreamingDetectIntentClient, error)
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// SessionsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service used for session interactions.
//
// For more information, see the API interactions
// guide (at https://cloud.google.com/dialogflow/docs/api-overview).
type SessionsClient struct {
	// The internal transport-dependent client.
	internalClient internalSessionsClient

	// The call options for this service.
	CallOptions *SessionsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SessionsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SessionsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *SessionsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// DetectIntent processes a natural language query and returns structured, actionable data
// as a result. This method is not idempotent, because it may cause contexts
// and session entity types to be updated, which in turn might affect
// results of future queries.
//
// If you might use
// Agent Assist (at https://cloud.google.com/dialogflow/docs/#aa)
// or other CCAI products now or in the future, consider using
// AnalyzeContent
// instead of DetectIntent. AnalyzeContent has additional
// functionality for Agent Assist and other CCAI products.
//
// Note: Always use agent versions for production traffic.
// See Versions and
// environments (at https://cloud.google.com/dialogflow/es/docs/agents-versions).
func (c *SessionsClient) DetectIntent(ctx context.Context, req *dialogflowpb.DetectIntentRequest, opts ...gax.CallOption) (*dialogflowpb.DetectIntentResponse, error) {
	return c.internalClient.DetectIntent(ctx, req, opts...)
}

// StreamingDetectIntent processes a natural language query in audio format in a streaming fashion
// and returns structured, actionable data as a result. This method is only
// available via the gRPC API (not REST).
//
// If you might use
// Agent Assist (at https://cloud.google.com/dialogflow/docs/#aa)
// or other CCAI products now or in the future, consider using
// StreamingAnalyzeContent
// instead of StreamingDetectIntent. StreamingAnalyzeContent has
// additional functionality for Agent Assist and other CCAI products.
//
// Note: Always use agent versions for production traffic.
// See Versions and
// environments (at https://cloud.google.com/dialogflow/es/docs/agents-versions).
func (c *SessionsClient) StreamingDetectIntent(ctx context.Context, opts ...gax.CallOption) (dialogflowpb.Sessions_StreamingDetectIntentClient, error) {
	return c.internalClient.StreamingDetectIntent(ctx, opts...)
}

// GetLocation gets information about a location.
func (c *SessionsClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *SessionsClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *SessionsClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *SessionsClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *SessionsClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// sessionsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type sessionsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SessionsClient
	CallOptions **SessionsCallOptions

	// The gRPC API client.
	sessionsClient dialogflowpb.SessionsClient

	operationsClient longrunningpb.OperationsClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSessionsClient creates a new sessions client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service used for session interactions.
//
// For more information, see the API interactions
// guide (at https://cloud.google.com/dialogflow/docs/api-overview).
func NewSessionsClient(ctx context.Context, opts ...option.ClientOption) (*SessionsClient, error) {
	clientOpts := defaultSessionsGRPCClientOptions()
	if newSessionsClientHook != nil {
		hookOpts, err := newSessionsClientHook(ctx, clientHookParams{})
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
	client := SessionsClient{CallOptions: defaultSessionsCallOptions()}

	c := &sessionsGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		sessionsClient:   dialogflowpb.NewSessionsClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		locationsClient:  locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *sessionsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *sessionsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *sessionsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type sessionsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing SessionsClient
	CallOptions **SessionsCallOptions
}

// NewSessionsRESTClient creates a new sessions rest client.
//
// A service used for session interactions.
//
// For more information, see the API interactions
// guide (at https://cloud.google.com/dialogflow/docs/api-overview).
func NewSessionsRESTClient(ctx context.Context, opts ...option.ClientOption) (*SessionsClient, error) {
	clientOpts := append(defaultSessionsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultSessionsRESTCallOptions()
	c := &sessionsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &SessionsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultSessionsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://dialogflow.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://dialogflow.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *sessionsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *sessionsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *sessionsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *sessionsGRPCClient) DetectIntent(ctx context.Context, req *dialogflowpb.DetectIntentRequest, opts ...gax.CallOption) (*dialogflowpb.DetectIntentResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 220000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DetectIntent[0:len((*c.CallOptions).DetectIntent):len((*c.CallOptions).DetectIntent)], opts...)
	var resp *dialogflowpb.DetectIntentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionsClient.DetectIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sessionsGRPCClient) StreamingDetectIntent(ctx context.Context, opts ...gax.CallOption) (dialogflowpb.Sessions_StreamingDetectIntentClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	var resp dialogflowpb.Sessions_StreamingDetectIntentClient
	opts = append((*c.CallOptions).StreamingDetectIntent[0:len((*c.CallOptions).StreamingDetectIntent):len((*c.CallOptions).StreamingDetectIntent)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionsClient.StreamingDetectIntent(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sessionsGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sessionsGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
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
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
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

func (c *sessionsGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *sessionsGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sessionsGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

// DetectIntent processes a natural language query and returns structured, actionable data
// as a result. This method is not idempotent, because it may cause contexts
// and session entity types to be updated, which in turn might affect
// results of future queries.
//
// If you might use
// Agent Assist (at https://cloud.google.com/dialogflow/docs/#aa)
// or other CCAI products now or in the future, consider using
// AnalyzeContent
// instead of DetectIntent. AnalyzeContent has additional
// functionality for Agent Assist and other CCAI products.
//
// Note: Always use agent versions for production traffic.
// See Versions and
// environments (at https://cloud.google.com/dialogflow/es/docs/agents-versions).
func (c *sessionsRESTClient) DetectIntent(ctx context.Context, req *dialogflowpb.DetectIntentRequest, opts ...gax.CallOption) (*dialogflowpb.DetectIntentResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta1/%v:detectIntent", req.GetSession())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).DetectIntent[0:len((*c.CallOptions).DetectIntent):len((*c.CallOptions).DetectIntent)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dialogflowpb.DetectIntentResponse{}
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

// StreamingDetectIntent processes a natural language query in audio format in a streaming fashion
// and returns structured, actionable data as a result. This method is only
// available via the gRPC API (not REST).
//
// If you might use
// Agent Assist (at https://cloud.google.com/dialogflow/docs/#aa)
// or other CCAI products now or in the future, consider using
// StreamingAnalyzeContent
// instead of StreamingDetectIntent. StreamingAnalyzeContent has
// additional functionality for Agent Assist and other CCAI products.
//
// Note: Always use agent versions for production traffic.
// See Versions and
// environments (at https://cloud.google.com/dialogflow/es/docs/agents-versions).
func (c *sessionsRESTClient) StreamingDetectIntent(ctx context.Context, opts ...gax.CallOption) (dialogflowpb.Sessions_StreamingDetectIntentClient, error) {
	return nil, fmt.Errorf("StreamingDetectIntent not yet supported for REST clients")
}

// GetLocation gets information about a location.
func (c *sessionsRESTClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &locationpb.Location{}
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

// ListLocations lists information about the supported locations for this service.
func (c *sessionsRESTClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v2beta1/%v/locations", req.GetName())

		params := url.Values{}
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
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
			return nil, "", e
		}
		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
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

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *sessionsRESTClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta1/%v:cancel", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
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

// GetOperation is a utility method from google.longrunning.Operations.
func (c *sessionsRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

// ListOperations is a utility method from google.longrunning.Operations.
func (c *sessionsRESTClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v2beta1/%v/operations", req.GetName())

		params := url.Values{}
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
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
			return nil, "", e
		}
		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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
