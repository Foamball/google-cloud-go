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

package aiplatform

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newSpecialistPoolClientHook clientHook

// SpecialistPoolCallOptions contains the retry settings for each method of SpecialistPoolClient.
type SpecialistPoolCallOptions struct {
	CreateSpecialistPool []gax.CallOption
	GetSpecialistPool    []gax.CallOption
	ListSpecialistPools  []gax.CallOption
	DeleteSpecialistPool []gax.CallOption
	UpdateSpecialistPool []gax.CallOption
}

func defaultSpecialistPoolGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSpecialistPoolCallOptions() *SpecialistPoolCallOptions {
	return &SpecialistPoolCallOptions{
		CreateSpecialistPool: []gax.CallOption{},
		GetSpecialistPool:    []gax.CallOption{},
		ListSpecialistPools:  []gax.CallOption{},
		DeleteSpecialistPool: []gax.CallOption{},
		UpdateSpecialistPool: []gax.CallOption{},
	}
}

// internalSpecialistPoolClient is an interface that defines the methods availaible from Vertex AI API.
type internalSpecialistPoolClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateSpecialistPool(context.Context, *aiplatformpb.CreateSpecialistPoolRequest, ...gax.CallOption) (*CreateSpecialistPoolOperation, error)
	CreateSpecialistPoolOperation(name string) *CreateSpecialistPoolOperation
	GetSpecialistPool(context.Context, *aiplatformpb.GetSpecialistPoolRequest, ...gax.CallOption) (*aiplatformpb.SpecialistPool, error)
	ListSpecialistPools(context.Context, *aiplatformpb.ListSpecialistPoolsRequest, ...gax.CallOption) *SpecialistPoolIterator
	DeleteSpecialistPool(context.Context, *aiplatformpb.DeleteSpecialistPoolRequest, ...gax.CallOption) (*DeleteSpecialistPoolOperation, error)
	DeleteSpecialistPoolOperation(name string) *DeleteSpecialistPoolOperation
	UpdateSpecialistPool(context.Context, *aiplatformpb.UpdateSpecialistPoolRequest, ...gax.CallOption) (*UpdateSpecialistPoolOperation, error)
	UpdateSpecialistPoolOperation(name string) *UpdateSpecialistPoolOperation
}

// SpecialistPoolClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for creating and managing Customer SpecialistPools.
// When customers start Data Labeling jobs, they can reuse/create Specialist
// Pools to bring their own Specialists to label the data.
// Customers can add/remove Managers for the Specialist Pool on Cloud console,
// then Managers will get email notifications to manage Specialists and tasks on
// CrowdCompute console.
type SpecialistPoolClient struct {
	// The internal transport-dependent client.
	internalClient internalSpecialistPoolClient

	// The call options for this service.
	CallOptions *SpecialistPoolCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SpecialistPoolClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SpecialistPoolClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SpecialistPoolClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateSpecialistPool creates a SpecialistPool.
func (c *SpecialistPoolClient) CreateSpecialistPool(ctx context.Context, req *aiplatformpb.CreateSpecialistPoolRequest, opts ...gax.CallOption) (*CreateSpecialistPoolOperation, error) {
	return c.internalClient.CreateSpecialistPool(ctx, req, opts...)
}

// CreateSpecialistPoolOperation returns a new CreateSpecialistPoolOperation from a given name.
// The name must be that of a previously created CreateSpecialistPoolOperation, possibly from a different process.
func (c *SpecialistPoolClient) CreateSpecialistPoolOperation(name string) *CreateSpecialistPoolOperation {
	return c.internalClient.CreateSpecialistPoolOperation(name)
}

// GetSpecialistPool gets a SpecialistPool.
func (c *SpecialistPoolClient) GetSpecialistPool(ctx context.Context, req *aiplatformpb.GetSpecialistPoolRequest, opts ...gax.CallOption) (*aiplatformpb.SpecialistPool, error) {
	return c.internalClient.GetSpecialistPool(ctx, req, opts...)
}

// ListSpecialistPools lists SpecialistPools in a Location.
func (c *SpecialistPoolClient) ListSpecialistPools(ctx context.Context, req *aiplatformpb.ListSpecialistPoolsRequest, opts ...gax.CallOption) *SpecialistPoolIterator {
	return c.internalClient.ListSpecialistPools(ctx, req, opts...)
}

// DeleteSpecialistPool deletes a SpecialistPool as well as all Specialists in the pool.
func (c *SpecialistPoolClient) DeleteSpecialistPool(ctx context.Context, req *aiplatformpb.DeleteSpecialistPoolRequest, opts ...gax.CallOption) (*DeleteSpecialistPoolOperation, error) {
	return c.internalClient.DeleteSpecialistPool(ctx, req, opts...)
}

// DeleteSpecialistPoolOperation returns a new DeleteSpecialistPoolOperation from a given name.
// The name must be that of a previously created DeleteSpecialistPoolOperation, possibly from a different process.
func (c *SpecialistPoolClient) DeleteSpecialistPoolOperation(name string) *DeleteSpecialistPoolOperation {
	return c.internalClient.DeleteSpecialistPoolOperation(name)
}

// UpdateSpecialistPool updates a SpecialistPool.
func (c *SpecialistPoolClient) UpdateSpecialistPool(ctx context.Context, req *aiplatformpb.UpdateSpecialistPoolRequest, opts ...gax.CallOption) (*UpdateSpecialistPoolOperation, error) {
	return c.internalClient.UpdateSpecialistPool(ctx, req, opts...)
}

// UpdateSpecialistPoolOperation returns a new UpdateSpecialistPoolOperation from a given name.
// The name must be that of a previously created UpdateSpecialistPoolOperation, possibly from a different process.
func (c *SpecialistPoolClient) UpdateSpecialistPoolOperation(name string) *UpdateSpecialistPoolOperation {
	return c.internalClient.UpdateSpecialistPoolOperation(name)
}

// specialistPoolGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type specialistPoolGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SpecialistPoolClient
	CallOptions **SpecialistPoolCallOptions

	// The gRPC API client.
	specialistPoolClient aiplatformpb.SpecialistPoolServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSpecialistPoolClient creates a new specialist pool service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for creating and managing Customer SpecialistPools.
// When customers start Data Labeling jobs, they can reuse/create Specialist
// Pools to bring their own Specialists to label the data.
// Customers can add/remove Managers for the Specialist Pool on Cloud console,
// then Managers will get email notifications to manage Specialists and tasks on
// CrowdCompute console.
func NewSpecialistPoolClient(ctx context.Context, opts ...option.ClientOption) (*SpecialistPoolClient, error) {
	clientOpts := defaultSpecialistPoolGRPCClientOptions()
	if newSpecialistPoolClientHook != nil {
		hookOpts, err := newSpecialistPoolClientHook(ctx, clientHookParams{})
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
	client := SpecialistPoolClient{CallOptions: defaultSpecialistPoolCallOptions()}

	c := &specialistPoolGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		specialistPoolClient: aiplatformpb.NewSpecialistPoolServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *specialistPoolGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *specialistPoolGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *specialistPoolGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *specialistPoolGRPCClient) CreateSpecialistPool(ctx context.Context, req *aiplatformpb.CreateSpecialistPoolRequest, opts ...gax.CallOption) (*CreateSpecialistPoolOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateSpecialistPool[0:len((*c.CallOptions).CreateSpecialistPool):len((*c.CallOptions).CreateSpecialistPool)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.specialistPoolClient.CreateSpecialistPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateSpecialistPoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *specialistPoolGRPCClient) GetSpecialistPool(ctx context.Context, req *aiplatformpb.GetSpecialistPoolRequest, opts ...gax.CallOption) (*aiplatformpb.SpecialistPool, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSpecialistPool[0:len((*c.CallOptions).GetSpecialistPool):len((*c.CallOptions).GetSpecialistPool)], opts...)
	var resp *aiplatformpb.SpecialistPool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.specialistPoolClient.GetSpecialistPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *specialistPoolGRPCClient) ListSpecialistPools(ctx context.Context, req *aiplatformpb.ListSpecialistPoolsRequest, opts ...gax.CallOption) *SpecialistPoolIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListSpecialistPools[0:len((*c.CallOptions).ListSpecialistPools):len((*c.CallOptions).ListSpecialistPools)], opts...)
	it := &SpecialistPoolIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListSpecialistPoolsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.SpecialistPool, string, error) {
		resp := &aiplatformpb.ListSpecialistPoolsResponse{}
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
			resp, err = c.specialistPoolClient.ListSpecialistPools(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSpecialistPools(), resp.GetNextPageToken(), nil
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

func (c *specialistPoolGRPCClient) DeleteSpecialistPool(ctx context.Context, req *aiplatformpb.DeleteSpecialistPoolRequest, opts ...gax.CallOption) (*DeleteSpecialistPoolOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteSpecialistPool[0:len((*c.CallOptions).DeleteSpecialistPool):len((*c.CallOptions).DeleteSpecialistPool)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.specialistPoolClient.DeleteSpecialistPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteSpecialistPoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *specialistPoolGRPCClient) UpdateSpecialistPool(ctx context.Context, req *aiplatformpb.UpdateSpecialistPoolRequest, opts ...gax.CallOption) (*UpdateSpecialistPoolOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "specialist_pool.name", url.QueryEscape(req.GetSpecialistPool().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateSpecialistPool[0:len((*c.CallOptions).UpdateSpecialistPool):len((*c.CallOptions).UpdateSpecialistPool)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.specialistPoolClient.UpdateSpecialistPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateSpecialistPoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateSpecialistPoolOperation manages a long-running operation from CreateSpecialistPool.
type CreateSpecialistPoolOperation struct {
	lro *longrunning.Operation
}

// CreateSpecialistPoolOperation returns a new CreateSpecialistPoolOperation from a given name.
// The name must be that of a previously created CreateSpecialistPoolOperation, possibly from a different process.
func (c *specialistPoolGRPCClient) CreateSpecialistPoolOperation(name string) *CreateSpecialistPoolOperation {
	return &CreateSpecialistPoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateSpecialistPoolOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.SpecialistPool, error) {
	var resp aiplatformpb.SpecialistPool
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateSpecialistPoolOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.SpecialistPool, error) {
	var resp aiplatformpb.SpecialistPool
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateSpecialistPoolOperation) Metadata() (*aiplatformpb.CreateSpecialistPoolOperationMetadata, error) {
	var meta aiplatformpb.CreateSpecialistPoolOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateSpecialistPoolOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateSpecialistPoolOperation) Name() string {
	return op.lro.Name()
}

// DeleteSpecialistPoolOperation manages a long-running operation from DeleteSpecialistPool.
type DeleteSpecialistPoolOperation struct {
	lro *longrunning.Operation
}

// DeleteSpecialistPoolOperation returns a new DeleteSpecialistPoolOperation from a given name.
// The name must be that of a previously created DeleteSpecialistPoolOperation, possibly from a different process.
func (c *specialistPoolGRPCClient) DeleteSpecialistPoolOperation(name string) *DeleteSpecialistPoolOperation {
	return &DeleteSpecialistPoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteSpecialistPoolOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteSpecialistPoolOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteSpecialistPoolOperation) Metadata() (*aiplatformpb.DeleteOperationMetadata, error) {
	var meta aiplatformpb.DeleteOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteSpecialistPoolOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteSpecialistPoolOperation) Name() string {
	return op.lro.Name()
}

// UpdateSpecialistPoolOperation manages a long-running operation from UpdateSpecialistPool.
type UpdateSpecialistPoolOperation struct {
	lro *longrunning.Operation
}

// UpdateSpecialistPoolOperation returns a new UpdateSpecialistPoolOperation from a given name.
// The name must be that of a previously created UpdateSpecialistPoolOperation, possibly from a different process.
func (c *specialistPoolGRPCClient) UpdateSpecialistPoolOperation(name string) *UpdateSpecialistPoolOperation {
	return &UpdateSpecialistPoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateSpecialistPoolOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.SpecialistPool, error) {
	var resp aiplatformpb.SpecialistPool
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateSpecialistPoolOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.SpecialistPool, error) {
	var resp aiplatformpb.SpecialistPool
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateSpecialistPoolOperation) Metadata() (*aiplatformpb.UpdateSpecialistPoolOperationMetadata, error) {
	var meta aiplatformpb.UpdateSpecialistPoolOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateSpecialistPoolOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateSpecialistPoolOperation) Name() string {
	return op.lro.Name()
}

// SpecialistPoolIterator manages a stream of *aiplatformpb.SpecialistPool.
type SpecialistPoolIterator struct {
	items    []*aiplatformpb.SpecialistPool
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.SpecialistPool, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SpecialistPoolIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SpecialistPoolIterator) Next() (*aiplatformpb.SpecialistPool, error) {
	var item *aiplatformpb.SpecialistPool
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SpecialistPoolIterator) bufLen() int {
	return len(it.items)
}

func (it *SpecialistPoolIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
