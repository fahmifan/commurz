// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commurz/v1/commurz.proto

package commurzv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

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
	// CommurzServiceFindUserByTokenProcedure is the fully-qualified name of the CommurzService's
	// FindUserByToken RPC.
	CommurzServiceFindUserByTokenProcedure = "/commurz.v1.CommurzService/FindUserByToken"
	// CommurzServiceListAppProductsProcedure is the fully-qualified name of the CommurzService's
	// ListAppProducts RPC.
	CommurzServiceListAppProductsProcedure = "/commurz.v1.CommurzService/ListAppProducts"
	// CommurzServiceFindCartByUserTokenProcedure is the fully-qualified name of the CommurzService's
	// FindCartByUserToken RPC.
	CommurzServiceFindCartByUserTokenProcedure = "/commurz.v1.CommurzService/FindCartByUserToken"
	// CommurzServiceFindProductByIDProcedure is the fully-qualified name of the CommurzService's
	// FindProductByID RPC.
	CommurzServiceFindProductByIDProcedure = "/commurz.v1.CommurzService/FindProductByID"
	// CommurzServiceListBackofficeProductsProcedure is the fully-qualified name of the CommurzService's
	// ListBackofficeProducts RPC.
	CommurzServiceListBackofficeProductsProcedure = "/commurz.v1.CommurzService/ListBackofficeProducts"
	// CommurzServiceCreateProductProcedure is the fully-qualified name of the CommurzService's
	// CreateProduct RPC.
	CommurzServiceCreateProductProcedure = "/commurz.v1.CommurzService/CreateProduct"
	// CommurzServiceUpdateProductStockProcedure is the fully-qualified name of the CommurzService's
	// UpdateProductStock RPC.
	CommurzServiceUpdateProductStockProcedure = "/commurz.v1.CommurzService/UpdateProductStock"
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
	ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error)
	FindUserByID(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.User], error)
	FindUserByToken(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.User], error)
	// storefront query
	ListAppProducts(context.Context, *connect.Request[v1.ListAppProductsRequest]) (*connect.Response[v1.ListAppProductsResponse], error)
	// order & inventory query
	FindCartByUserToken(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Cart], error)
	FindProductByID(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Product], error)
	ListBackofficeProducts(context.Context, *connect.Request[v1.ListBackofficeProductsRequest]) (*connect.Response[v1.ListBackofficeProductsResponse], error)
	// order & inventory command
	CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.Empty], error)
	UpdateProductStock(context.Context, *connect.Request[v1.UpdateProductStockRequest]) (*connect.Response[v1.Empty], error)
	AddProductToCart(context.Context, *connect.Request[v1.AddProductToCartRequest]) (*connect.Response[v1.Empty], error)
	CheckoutAll(context.Context, *connect.Request[v1.CheckoutAllRequest]) (*connect.Response[v1.Empty], error)
}

// NewCommurzServiceClient constructs a client for the commurz.v1.CommurzService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCommurzServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CommurzServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &commurzServiceClient{
		listUsers: connect.NewClient[v1.ListUsersRequest, v1.ListUsersResponse](
			httpClient,
			baseURL+CommurzServiceListUsersProcedure,
			opts...,
		),
		findUserByID: connect.NewClient[v1.FindByIDRequest, v1.User](
			httpClient,
			baseURL+CommurzServiceFindUserByIDProcedure,
			opts...,
		),
		findUserByToken: connect.NewClient[v1.Empty, v1.User](
			httpClient,
			baseURL+CommurzServiceFindUserByTokenProcedure,
			opts...,
		),
		listAppProducts: connect.NewClient[v1.ListAppProductsRequest, v1.ListAppProductsResponse](
			httpClient,
			baseURL+CommurzServiceListAppProductsProcedure,
			opts...,
		),
		findCartByUserToken: connect.NewClient[v1.Empty, v1.Cart](
			httpClient,
			baseURL+CommurzServiceFindCartByUserTokenProcedure,
			opts...,
		),
		findProductByID: connect.NewClient[v1.FindByIDRequest, v1.Product](
			httpClient,
			baseURL+CommurzServiceFindProductByIDProcedure,
			opts...,
		),
		listBackofficeProducts: connect.NewClient[v1.ListBackofficeProductsRequest, v1.ListBackofficeProductsResponse](
			httpClient,
			baseURL+CommurzServiceListBackofficeProductsProcedure,
			opts...,
		),
		createProduct: connect.NewClient[v1.CreateProductRequest, v1.Empty](
			httpClient,
			baseURL+CommurzServiceCreateProductProcedure,
			opts...,
		),
		updateProductStock: connect.NewClient[v1.UpdateProductStockRequest, v1.Empty](
			httpClient,
			baseURL+CommurzServiceUpdateProductStockProcedure,
			opts...,
		),
		addProductToCart: connect.NewClient[v1.AddProductToCartRequest, v1.Empty](
			httpClient,
			baseURL+CommurzServiceAddProductToCartProcedure,
			opts...,
		),
		checkoutAll: connect.NewClient[v1.CheckoutAllRequest, v1.Empty](
			httpClient,
			baseURL+CommurzServiceCheckoutAllProcedure,
			opts...,
		),
	}
}

// commurzServiceClient implements CommurzServiceClient.
type commurzServiceClient struct {
	listUsers              *connect.Client[v1.ListUsersRequest, v1.ListUsersResponse]
	findUserByID           *connect.Client[v1.FindByIDRequest, v1.User]
	findUserByToken        *connect.Client[v1.Empty, v1.User]
	listAppProducts        *connect.Client[v1.ListAppProductsRequest, v1.ListAppProductsResponse]
	findCartByUserToken    *connect.Client[v1.Empty, v1.Cart]
	findProductByID        *connect.Client[v1.FindByIDRequest, v1.Product]
	listBackofficeProducts *connect.Client[v1.ListBackofficeProductsRequest, v1.ListBackofficeProductsResponse]
	createProduct          *connect.Client[v1.CreateProductRequest, v1.Empty]
	updateProductStock     *connect.Client[v1.UpdateProductStockRequest, v1.Empty]
	addProductToCart       *connect.Client[v1.AddProductToCartRequest, v1.Empty]
	checkoutAll            *connect.Client[v1.CheckoutAllRequest, v1.Empty]
}

// ListUsers calls commurz.v1.CommurzService.ListUsers.
func (c *commurzServiceClient) ListUsers(ctx context.Context, req *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	return c.listUsers.CallUnary(ctx, req)
}

// FindUserByID calls commurz.v1.CommurzService.FindUserByID.
func (c *commurzServiceClient) FindUserByID(ctx context.Context, req *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.User], error) {
	return c.findUserByID.CallUnary(ctx, req)
}

// FindUserByToken calls commurz.v1.CommurzService.FindUserByToken.
func (c *commurzServiceClient) FindUserByToken(ctx context.Context, req *connect.Request[v1.Empty]) (*connect.Response[v1.User], error) {
	return c.findUserByToken.CallUnary(ctx, req)
}

// ListAppProducts calls commurz.v1.CommurzService.ListAppProducts.
func (c *commurzServiceClient) ListAppProducts(ctx context.Context, req *connect.Request[v1.ListAppProductsRequest]) (*connect.Response[v1.ListAppProductsResponse], error) {
	return c.listAppProducts.CallUnary(ctx, req)
}

// FindCartByUserToken calls commurz.v1.CommurzService.FindCartByUserToken.
func (c *commurzServiceClient) FindCartByUserToken(ctx context.Context, req *connect.Request[v1.Empty]) (*connect.Response[v1.Cart], error) {
	return c.findCartByUserToken.CallUnary(ctx, req)
}

// FindProductByID calls commurz.v1.CommurzService.FindProductByID.
func (c *commurzServiceClient) FindProductByID(ctx context.Context, req *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Product], error) {
	return c.findProductByID.CallUnary(ctx, req)
}

// ListBackofficeProducts calls commurz.v1.CommurzService.ListBackofficeProducts.
func (c *commurzServiceClient) ListBackofficeProducts(ctx context.Context, req *connect.Request[v1.ListBackofficeProductsRequest]) (*connect.Response[v1.ListBackofficeProductsResponse], error) {
	return c.listBackofficeProducts.CallUnary(ctx, req)
}

// CreateProduct calls commurz.v1.CommurzService.CreateProduct.
func (c *commurzServiceClient) CreateProduct(ctx context.Context, req *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.Empty], error) {
	return c.createProduct.CallUnary(ctx, req)
}

// UpdateProductStock calls commurz.v1.CommurzService.UpdateProductStock.
func (c *commurzServiceClient) UpdateProductStock(ctx context.Context, req *connect.Request[v1.UpdateProductStockRequest]) (*connect.Response[v1.Empty], error) {
	return c.updateProductStock.CallUnary(ctx, req)
}

// AddProductToCart calls commurz.v1.CommurzService.AddProductToCart.
func (c *commurzServiceClient) AddProductToCart(ctx context.Context, req *connect.Request[v1.AddProductToCartRequest]) (*connect.Response[v1.Empty], error) {
	return c.addProductToCart.CallUnary(ctx, req)
}

// CheckoutAll calls commurz.v1.CommurzService.CheckoutAll.
func (c *commurzServiceClient) CheckoutAll(ctx context.Context, req *connect.Request[v1.CheckoutAllRequest]) (*connect.Response[v1.Empty], error) {
	return c.checkoutAll.CallUnary(ctx, req)
}

// CommurzServiceHandler is an implementation of the commurz.v1.CommurzService service.
type CommurzServiceHandler interface {
	// user
	ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error)
	FindUserByID(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.User], error)
	FindUserByToken(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.User], error)
	// storefront query
	ListAppProducts(context.Context, *connect.Request[v1.ListAppProductsRequest]) (*connect.Response[v1.ListAppProductsResponse], error)
	// order & inventory query
	FindCartByUserToken(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Cart], error)
	FindProductByID(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Product], error)
	ListBackofficeProducts(context.Context, *connect.Request[v1.ListBackofficeProductsRequest]) (*connect.Response[v1.ListBackofficeProductsResponse], error)
	// order & inventory command
	CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.Empty], error)
	UpdateProductStock(context.Context, *connect.Request[v1.UpdateProductStockRequest]) (*connect.Response[v1.Empty], error)
	AddProductToCart(context.Context, *connect.Request[v1.AddProductToCartRequest]) (*connect.Response[v1.Empty], error)
	CheckoutAll(context.Context, *connect.Request[v1.CheckoutAllRequest]) (*connect.Response[v1.Empty], error)
}

// NewCommurzServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCommurzServiceHandler(svc CommurzServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	commurzServiceListUsersHandler := connect.NewUnaryHandler(
		CommurzServiceListUsersProcedure,
		svc.ListUsers,
		opts...,
	)
	commurzServiceFindUserByIDHandler := connect.NewUnaryHandler(
		CommurzServiceFindUserByIDProcedure,
		svc.FindUserByID,
		opts...,
	)
	commurzServiceFindUserByTokenHandler := connect.NewUnaryHandler(
		CommurzServiceFindUserByTokenProcedure,
		svc.FindUserByToken,
		opts...,
	)
	commurzServiceListAppProductsHandler := connect.NewUnaryHandler(
		CommurzServiceListAppProductsProcedure,
		svc.ListAppProducts,
		opts...,
	)
	commurzServiceFindCartByUserTokenHandler := connect.NewUnaryHandler(
		CommurzServiceFindCartByUserTokenProcedure,
		svc.FindCartByUserToken,
		opts...,
	)
	commurzServiceFindProductByIDHandler := connect.NewUnaryHandler(
		CommurzServiceFindProductByIDProcedure,
		svc.FindProductByID,
		opts...,
	)
	commurzServiceListBackofficeProductsHandler := connect.NewUnaryHandler(
		CommurzServiceListBackofficeProductsProcedure,
		svc.ListBackofficeProducts,
		opts...,
	)
	commurzServiceCreateProductHandler := connect.NewUnaryHandler(
		CommurzServiceCreateProductProcedure,
		svc.CreateProduct,
		opts...,
	)
	commurzServiceUpdateProductStockHandler := connect.NewUnaryHandler(
		CommurzServiceUpdateProductStockProcedure,
		svc.UpdateProductStock,
		opts...,
	)
	commurzServiceAddProductToCartHandler := connect.NewUnaryHandler(
		CommurzServiceAddProductToCartProcedure,
		svc.AddProductToCart,
		opts...,
	)
	commurzServiceCheckoutAllHandler := connect.NewUnaryHandler(
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
		case CommurzServiceFindUserByTokenProcedure:
			commurzServiceFindUserByTokenHandler.ServeHTTP(w, r)
		case CommurzServiceListAppProductsProcedure:
			commurzServiceListAppProductsHandler.ServeHTTP(w, r)
		case CommurzServiceFindCartByUserTokenProcedure:
			commurzServiceFindCartByUserTokenHandler.ServeHTTP(w, r)
		case CommurzServiceFindProductByIDProcedure:
			commurzServiceFindProductByIDHandler.ServeHTTP(w, r)
		case CommurzServiceListBackofficeProductsProcedure:
			commurzServiceListBackofficeProductsHandler.ServeHTTP(w, r)
		case CommurzServiceCreateProductProcedure:
			commurzServiceCreateProductHandler.ServeHTTP(w, r)
		case CommurzServiceUpdateProductStockProcedure:
			commurzServiceUpdateProductStockHandler.ServeHTTP(w, r)
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

func (UnimplementedCommurzServiceHandler) ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.ListUsers is not implemented"))
}

func (UnimplementedCommurzServiceHandler) FindUserByID(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.User], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.FindUserByID is not implemented"))
}

func (UnimplementedCommurzServiceHandler) FindUserByToken(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.User], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.FindUserByToken is not implemented"))
}

func (UnimplementedCommurzServiceHandler) ListAppProducts(context.Context, *connect.Request[v1.ListAppProductsRequest]) (*connect.Response[v1.ListAppProductsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.ListAppProducts is not implemented"))
}

func (UnimplementedCommurzServiceHandler) FindCartByUserToken(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Cart], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.FindCartByUserToken is not implemented"))
}

func (UnimplementedCommurzServiceHandler) FindProductByID(context.Context, *connect.Request[v1.FindByIDRequest]) (*connect.Response[v1.Product], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.FindProductByID is not implemented"))
}

func (UnimplementedCommurzServiceHandler) ListBackofficeProducts(context.Context, *connect.Request[v1.ListBackofficeProductsRequest]) (*connect.Response[v1.ListBackofficeProductsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.ListBackofficeProducts is not implemented"))
}

func (UnimplementedCommurzServiceHandler) CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.CreateProduct is not implemented"))
}

func (UnimplementedCommurzServiceHandler) UpdateProductStock(context.Context, *connect.Request[v1.UpdateProductStockRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.UpdateProductStock is not implemented"))
}

func (UnimplementedCommurzServiceHandler) AddProductToCart(context.Context, *connect.Request[v1.AddProductToCartRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.AddProductToCart is not implemented"))
}

func (UnimplementedCommurzServiceHandler) CheckoutAll(context.Context, *connect.Request[v1.CheckoutAllRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commurz.v1.CommurzService.CheckoutAll is not implemented"))
}
