// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commurz/v1/commurz.proto

package commurzv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CommurzServiceName is the fully-qualified name of the CommurzService service.
	CommurzServiceName = "commurz.v1.CommurzService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CommurzServiceListUsersProcedure is the fully-qualified name of the CommurzService's ListUsers
	// RPC.
	CommurzServiceListUsersProcedure = "/commurz.v1.CommurzService/ListUsers"
	// CommurzServiceFindUserByIDProcedure is the fully-qualified name of the CommurzService's
	// FindUserByID RPC.
	CommurzServiceFindUserByIDProcedure = "/commurz.v1.CommurzService/FindUserByID"
	// CommurzServiceListAppProductsProcedure is the fully-qualified name of the CommurzService's
	// ListAppProducts RPC.
	CommurzServiceListAppProductsProcedure = "/commurz.v1.CommurzService/ListAppProducts"
	// CommurzServiceListBackofficeProductsProcedure is the fully-qualified name of the CommurzService's
	// ListBackofficeProducts RPC.
	CommurzServiceListBackofficeProductsProcedure = "/commurz.v1.CommurzService/ListBackofficeProducts"
	// CommurzServiceCreateProductProcedure is the fully-qualified name of the CommurzService's
	// CreateProduct RPC.
	CommurzServiceCreateProductProcedure = "/commurz.v1.CommurzService/CreateProduct"
	// CommurzServiceAddProductStockProcedure is the fully-qualified name of the CommurzService's
	// AddProductStock RPC.
	CommurzServiceAddProductStockProcedure = "/commurz.v1.CommurzService/AddProductStock"
	// CommurzServiceReduceProductStockProcedure is the fully-qualified name of the CommurzService's
	// ReduceProductStock RPC.
	CommurzServiceReduceProductStockProcedure = "/commurz.v1.CommurzService/ReduceProductStock"
	// CommurzServiceAddProductToCartProcedure is the fully-qualified name of the CommurzService's
	// AddProductToCart RPC.
	CommurzServiceAddProductToCartProcedure = "/commurz.v1.CommurzService/AddProductToCart"
	// CommurzServiceCheckoutAllProcedure is the fully-qualified name of the CommurzService's
	// CheckoutAll RPC.
	CommurzServiceCheckoutAllProcedure = "/commurz.v1.CommurzService/CheckoutAll"
)

// CommurzServiceClient is a client for the commurz.v1.CommurzService service.
type CommurzServiceClient interface {
	// user
	ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error)
	FindUserByID(context.Context, *connect_go.Request[v1.FindByIDRequest]) (*connect_go.Response[v1.User], error)
	// app
	ListAppProducts(context.Context, *connect_go.Request[v1.ListAppProductsRequest]) (*connect_go.Response[v1.ListAppProductsResponse], error)
	// backoffice
	ListBackofficeProducts(context.Context, *connect_go.Request[v1.ListBackofficeProductsRequest]) (*connect_go.Response[v1.ListBackofficeProductsResponse], error)
	// product
	CreateProduct(context.Context, *connect_go.Request[v1.CreateProductRequest]) (*connect_go.Response[v1.Product], error)
	AddProductStock(context.Context, *connect_go.Request[v1.ChangeProductStockRequest]) (*connect_go.Response[v1.Product], error)
	ReduceProductStock(context.Context, *connect_go.Request[v1.ChangeProductStockRequest]) (*connect_go.Response[v1.Product], error)
	// cart
	AddProductToCart(context.Context, *connect_go.Request[v1.AddProductToCartRequest]) (*connect_go.Response[v1.Cart], error)
	CheckoutAll(context.Context, *connect_go.Request[v1.CheckoutAllRequest]) (*connect_go.Response[v1.Order], error)
}

// NewCommurzServiceClient constructs a client for the commurz.v1.CommurzService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCommurzServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CommurzServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &commurzServiceClient{
		listUsers: connect_go.NewClient[v1.ListUsersRequest, v1.ListUsersResponse](
			httpClient,
			baseURL+CommurzServiceListUsersProcedure,
			opts...,
		),
		findUserByID: connect_go.NewClient[v1.FindByIDRequest, v1.User](
			httpClient,
			baseURL+CommurzServiceFindUserByIDProcedure,
			opts...,
		),
		listAppProducts: connect_go.NewClient[v1.ListAppProductsRequest, v1.ListAppProductsResponse](
			httpClient,
			baseURL+CommurzServiceListAppProductsProcedure,
			opts...,
		),
		listBackofficeProducts: connect_go.NewClient[v1.ListBackofficeProductsRequest, v1.ListBackofficeProductsResponse](
			httpClient,
			baseURL+CommurzServiceListBackofficeProductsProcedure,
			opts...,
		),
		createProduct: connect_go.NewClient[v1.CreateProductRequest, v1.Product](
			httpClient,
			baseURL+CommurzServiceCreateProductProcedure,
			opts...,
		),
		addProductStock: connect_go.NewClient[v1.ChangeProductStockRequest, v1.Product](
			httpClient,
			baseURL+CommurzServiceAddProductStockProcedure,
			opts...,
		),
		reduceProductStock: connect_go.NewClient[v1.ChangeProductStockRequest, v1.Product](
			httpClient,
			baseURL+CommurzServiceReduceProductStockProcedure,
			opts...,
		),
		addProductToCart: connect_go.NewClient[v1.AddProductToCartRequest, v1.Cart](
			httpClient,
			baseURL+CommurzServiceAddProductToCartProcedure,
			opts...,
		),
		checkoutAll: connect_go.NewClient[v1.CheckoutAllRequest, v1.Order](
			httpClient,
			baseURL+CommurzServiceCheckoutAllProcedure,
			opts...,
		),
	}
}

// commurzServiceClient implements CommurzServiceClient.
type commurzServiceClient struct {
	listUsers              *connect_go.Client[v1.ListUsersRequest, v1.ListUsersResponse]
	findUserByID           *connect_go.Client[v1.FindByIDRequest, v1.User]
	listAppProducts        *connect_go.Client[v1.ListAppProductsRequest, v1.ListAppProductsResponse]
	listBackofficeProducts *connect_go.Client[v1.ListBackofficeProductsRequest, v1.ListBackofficeProductsResponse]
	createProduct          *connect_go.Client[v1.CreateProductRequest, v1.Product]
	addProductStock        *connect_go.Client[v1.ChangeProductStockRequest, v1.Product]
	reduceProductStock     *connect_go.Client[v1.ChangeProductStockRequest, v1.Product]
	addProductToCart       *connect_go.Client[v1.AddProductToCartRequest, v1.Cart]
	checkoutAll            *connect_go.Client[v1.CheckoutAllRequest, v1.Order]
}

// ListUsers calls commurz.v1.CommurzService.ListUsers.
func (c *commurzServiceClient) ListUsers(ctx context.Context, req *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error) {
	return c.listUsers.CallUnary(ctx, req)
}

// FindUserByID calls commurz.v1.CommurzService.FindUserByID.
func (c *commurzServiceClient) FindUserByID(ctx context.Context, req *connect_go.Request[v1.FindByIDRequest]) (*connect_go.Response[v1.User], error) {
	return c.findUserByID.CallUnary(ctx, req)
}

// ListAppProducts calls commurz.v1.CommurzService.ListAppProducts.
func (c *commurzServiceClient) ListAppProducts(ctx context.Context, req *connect_go.Request[v1.ListAppProductsRequest]) (*connect_go.Response[v1.ListAppProductsResponse], error) {
	return c.listAppProducts.CallUnary(ctx, req)
}

// ListBackofficeProducts calls commurz.v1.CommurzService.ListBackofficeProducts.
func (c *commurzServiceClient) ListBackofficeProducts(ctx context.Context, req *connect_go.Request[v1.ListBackofficeProductsRequest]) (*connect_go.Response[v1.ListBackofficeProductsResponse], error) {
	return c.listBackofficeProducts.CallUnary(ctx, req)
}

// CreateProduct calls commurz.v1.CommurzService.CreateProduct.
func (c *commurzServiceClient) CreateProduct(ctx context.Context, req *connect_go.Request[v1.CreateProductRequest]) (*connect_go.Response[v1.Product], error) {
	return c.createProduct.CallUnary(ctx, req)
}

// AddProductStock calls commurz.v1.CommurzService.AddProductStock.
func (c *commurzServiceClient) AddProductStock(ctx context.Context, req *connect_go.Request[v1.ChangeProductStockRequest]) (*connect_go.Response[v1.Product], error) {
	return c.addProductStock.CallUnary(ctx, req)
}

// ReduceProductStock calls commurz.v1.CommurzService.ReduceProductStock.
func (c *commurzServiceClient) ReduceProductStock(ctx context.Context, req *connect_go.Request[v1.ChangeProductStockRequest]) (*connect_go.Response[v1.Product], error) {
	return c.reduceProductStock.CallUnary(ctx, req)
}

// AddProductToCart calls commurz.v1.CommurzService.AddProductToCart.
func (c *commurzServiceClient) AddProductToCart(ctx context.Context, req *connect_go.Request[v1.AddProductToCartRequest]) (*connect_go.Response[v1.Cart], error) {
	return c.addProductToCart.CallUnary(ctx, req)
}

// CheckoutAll calls commurz.v1.CommurzService.CheckoutAll.
func (c *commurzServiceClient) CheckoutAll(ctx context.Context, req *connect_go.Request[v1.CheckoutAllRequest]) (*connect_go.Response[v1.Order], error) {
	return c.checkoutAll.CallUnary(ctx, req)
}

// CommurzServiceHandler is an implementation of the commurz.v1.CommurzService service.
type CommurzServiceHandler interface {
	// user
	ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error)
	FindUserByID(context.Context, *connect_go.Request[v1.FindByIDRequest]) (*connect_go.Response[v1.User], error)
	// app
	ListAppProducts(context.Context, *connect_go.Request[v1.ListAppProductsRequest]) (*connect_go.Response[v1.ListAppProductsResponse], error)
	// backoffice
	ListBackofficeProducts(context.Context, *connect_go.Request[v1.ListBackofficeProductsRequest]) (*connect_go.Response[v1.ListBackofficeProductsResponse], error)
	// product
	CreateProduct(context.Context, *connect_go.Request[v1.CreateProductRequest]) (*connect_go.Response[v1.Product], error)
	AddProductStock(context.Context, *connect_go.Request[v1.ChangeProductStockRequest]) (*connect_go.Response[v1.Product], error)
	ReduceProductStock(context.Context, *connect_go.Request[v1.ChangeProductStockRequest]) (*connect_go.Response[v1.Product], error)
	// cart
	AddProductToCart(context.Context, *connect_go.Request[v1.AddProductToCartRequest]) (*connect_go.Response[v1.Cart], error)
	CheckoutAll(context.Context, *connect_go.Request[v1.CheckoutAllRequest]) (*connect_go.Response[v1.Order], error)
}

// NewCommurzServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCommurzServiceHandler(svc CommurzServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	commurzServiceListUsersHandler := connect_go.NewUnaryHandler(
		CommurzServiceListUsersProcedure,
		svc.ListUsers,
		opts...,
	)
	commurzServiceFindUserByIDHandler := connect_go.NewUnaryHandler(
		CommurzServiceFindUserByIDProcedure,
		svc.FindUserByID,
		opts...,
	)
	commurzServiceListAppProductsHandler := connect_go.NewUnaryHandler(
		CommurzServiceListAppProductsProcedure,
		svc.ListAppProducts,
		opts...,
	)
	commurzServiceListBackofficeProductsHandler := connect_go.NewUnaryHandler(
		CommurzServiceListBackofficeProductsProcedure,
		svc.ListBackofficeProducts,
		opts...,
	)
	commurzServiceCreateProductHandler := connect_go.NewUnaryHandler(
		CommurzServiceCreateProductProcedure,
		svc.CreateProduct,
		opts...,
	)
	commurzServiceAddProductStockHandler := connect_go.NewUnaryHandler(
		CommurzServiceAddProductStockProcedure,
		svc.AddProductStock,
		opts...,
	)
	commurzServiceReduceProductStockHandler := connect_go.NewUnaryHandler(
		CommurzServiceReduceProductStockProcedure,
		svc.ReduceProductStock,
		opts...,
	)
	commurzServiceAddProductToCartHandler := connect_go.NewUnaryHandler(
		CommurzServiceAddProductToCartProcedure,
		svc.AddProductToCart,
		opts...,
	)
	commurzServiceCheckoutAllHandler := connect_go.NewUnaryHandler(
		CommurzServiceCheckoutAllProcedure,
		svc.CheckoutAll,
		opts...,
	)
	return "/commurz.v1.CommurzService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CommurzServiceListUsersProcedure:
			commurzServiceListUsersHandler.ServeHTTP(w, r)
		case CommurzServiceFindUserByIDProcedure:
			commurzServiceFindUserByIDHandler.ServeHTTP(w, r)
		case CommurzServiceListAppProductsProcedure:
			commurzServiceListAppProductsHandler.ServeHTTP(w, r)
		case CommurzServiceListBackofficeProductsProcedure:
			commurzServiceListBackofficeProductsHandler.ServeHTTP(w, r)
		case CommurzServiceCreateProductProcedure:
			commurzServiceCreateProductHandler.ServeHTTP(w, r)
		case CommurzServiceAddProductStockProcedure:
			commurzServiceAddProductStockHandler.ServeHTTP(w, r)
		case CommurzServiceReduceProductStockProcedure:
			commurzServiceReduceProductStockHandler.ServeHTTP(w, r)
		case CommurzServiceAddProductToCartProcedure:
			commurzServiceAddProductToCartHandler.ServeHTTP(w, r)
		case CommurzServiceCheckoutAllProcedure:
			commurzServiceCheckoutAllHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCommurzServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCommurzServiceHandler struct{}

func (UnimplementedCommurzServiceHandler) ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.ListUsers is not implemented"))
}

func (UnimplementedCommurzServiceHandler) FindUserByID(context.Context, *connect_go.Request[v1.FindByIDRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.FindUserByID is not implemented"))
}

func (UnimplementedCommurzServiceHandler) ListAppProducts(context.Context, *connect_go.Request[v1.ListAppProductsRequest]) (*connect_go.Response[v1.ListAppProductsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.ListAppProducts is not implemented"))
}

func (UnimplementedCommurzServiceHandler) ListBackofficeProducts(context.Context, *connect_go.Request[v1.ListBackofficeProductsRequest]) (*connect_go.Response[v1.ListBackofficeProductsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.ListBackofficeProducts is not implemented"))
}

func (UnimplementedCommurzServiceHandler) CreateProduct(context.Context, *connect_go.Request[v1.CreateProductRequest]) (*connect_go.Response[v1.Product], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.CreateProduct is not implemented"))
}

func (UnimplementedCommurzServiceHandler) AddProductStock(context.Context, *connect_go.Request[v1.ChangeProductStockRequest]) (*connect_go.Response[v1.Product], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.AddProductStock is not implemented"))
}

func (UnimplementedCommurzServiceHandler) ReduceProductStock(context.Context, *connect_go.Request[v1.ChangeProductStockRequest]) (*connect_go.Response[v1.Product], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.ReduceProductStock is not implemented"))
}

func (UnimplementedCommurzServiceHandler) AddProductToCart(context.Context, *connect_go.Request[v1.AddProductToCartRequest]) (*connect_go.Response[v1.Cart], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.AddProductToCart is not implemented"))
}

func (UnimplementedCommurzServiceHandler) CheckoutAll(context.Context, *connect_go.Request[v1.CheckoutAllRequest]) (*connect_go.Response[v1.Order], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commurz.v1.CommurzService.CheckoutAll is not implemented"))
}
