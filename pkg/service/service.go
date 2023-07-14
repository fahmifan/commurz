// Package service is the API entry point,
// it orchestrates between business logic, 3rd party services, and persistence storage.
package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkgorder"
	"github.com/fahmifan/commurz/pkg/internal/pkgprice"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/pb/commurz/v1/commurzv1connect"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
	"github.com/fahmifan/ulids"
)

type Config struct {
	DB *sql.DB
}

var _ commurzv1connect.CommurzServiceHandler = &Service{}

type Service struct {
	*Config
}

func NewService(config *Config) *Service {
	return &Service{config}
}

func (service *Service) ListUsers(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ListUsersRequest],
) (*connect.Response[commurzpbv1.ListUsersResponse], error) {
	userRepo := pkguser.UserRepository{}
	users, err := userRepo.FindAllUsers(ctx, service.DB)
	if err != nil {
		return nil, fmt.Errorf("[ListUsers] FindAllUsers: %w", err)
	}

	res := &connect.Response[commurzpbv1.ListUsersResponse]{
		Msg: &commurzpbv1.ListUsersResponse{
			Users: protoserde.ListFromUserPkg(users),
		},
	}

	return res, nil
}

func (service *Service) CreateUser(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CreateUserRequest],
) (res *connect.Response[commurzpbv1.User], err error) {
	userRepo := pkguser.UserRepository{}

	user, err := pkguser.NewUser(req.Msg.Email)
	if err != nil {
		return res, fmt.Errorf("[CreateUser] NewUser: %w", err)
	}

	user, err = userRepo.CreateUser(ctx, service.DB, user)
	if err != nil {
		return res, fmt.Errorf("[CreateUser] SaveUser: %w", err)
	}

	// use complete user fields
	user, err = userRepo.FindUserByID(ctx, service.DB, user.ID)
	if err != nil {
		return res, fmt.Errorf("[CreateUser] FindUserByID: %w", err)
	}

	res = &connect.Response[commurzpbv1.User]{
		Msg: protoserde.FromUserPkg(user),
	}

	return res, nil
}

func (service *Service) CreateProduct(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CreateProductRequest],
) (res *connect.Response[commurzpbv1.Product], err error) {
	productRepo := pkgproduct.ProductRepository{}

	product, err := pkgproduct.CreateProduct(req.Msg.Name, pkgprice.New(req.Msg.Price))
	if err != nil {
		return res, fmt.Errorf("[CreateProduct] CreateProduct: %w", err)
	}

	product, err = productRepo.SaveProduct(ctx, service.DB, product)
	if err != nil {
		return res, fmt.Errorf("[CreateProduct] CreateProduct: %w", err)
	}

	// use complete product fields
	product, err = productRepo.FindProductByID(ctx, service.DB, product.ID)
	if err != nil {
		return res, fmt.Errorf("[CreateProduct] FindProductByID: %w", err)
	}

	res = &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return res, nil
}

func (service *Service) AddProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ChangeProductStockRequest],
) (res *connect.Response[commurzpbv1.Product], err error) {
	var product pkgproduct.Product

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		productRepo := pkgproduct.ProductRepository{}

		productID, err := pkgutil.ParseULID(req.Msg.GetProductId())
		if err != nil {
			return fmt.Errorf("[AddProductStock] ParseULID: %w", err)
		}

		product, err = productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[AddProductStock] FindProductByID: %w", err)
		}

		var stock pkgproduct.ProductStock
		product, stock = product.AddStock(req.Msg.GetStockQuantity(), time.Now())

		_, err = productRepo.SaveProductStock(ctx, tx, stock)
		if err != nil {
			return fmt.Errorf("[AddProductStock] SaveProductStock: %w", err)
		}

		return nil
	})

	if err != nil {
		return res, fmt.Errorf("[AddProductStock] Transaction: %w", err)
	}

	res = &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return
}

func (service *Service) ReduceProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ChangeProductStockRequest],
) (*connect.Response[commurzpbv1.Product], error) {
	productRepo := pkgproduct.ProductRepository{}

	product := pkgproduct.Product{}
	productID, err := pkgutil.ParseULID(req.Msg.ProductId)
	if err != nil {
		return nil, fmt.Errorf("[ReduceProductStock] ParseULID: %w", err)
	}

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		product, err = productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] FindProductByID: %w", err)
		}

		var stock pkgproduct.ProductStock
		product, stock, err = product.ReduceStock(req.Msg.GetStockQuantity(), time.Now())
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] ReduceStock: %w", err)
		}

		_, err = productRepo.SaveProductStock(ctx, tx, stock)
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] SaveProductStock: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[AddProductStock] Transaction: %w", err)
	}

	product, err = productRepo.FindProductByID(ctx, service.DB, productID)
	if err != nil {
		return nil, fmt.Errorf("[ReduceProductStock] FindProductByID: %w", err)
	}

	res := &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return res, nil
}

func (service *Service) AddItemToCart(
	ctx context.Context,
	req *connect.Request[commurzpbv1.AddItemToCartRequest],
) (*connect.Response[commurzpbv1.Cart], error) {
	userID, err := pkgutil.ParseULID(req.Msg.GetUserId())
	if err != nil {
		return nil, fmt.Errorf("[AddItemToCart] parse userID: %w", err)
	}

	productID, err := pkgutil.ParseULID(req.Msg.ProductId)
	if err != nil {
		return nil, fmt.Errorf("[AddItemToCart] parse productID: %w", err)
	}

	cartRepo := pkgorder.CartRepository{}
	cart, err := service.getOrCreateCart(ctx, service.DB, userID)
	if err != nil {
		return nil, fmt.Errorf("[AddItemToCart] getOrCreateCart: %w", err)
	}

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		productRepo := pkgproduct.ProductRepository{}

		product, err := productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[AddItemToCart] FindProductByID: %w", err)
		}

		var cartItem pkgorder.CartItem

		cart, cartItem, err = cart.AddItem(product, req.Msg.GetQuantity())
		if err != nil {
			return fmt.Errorf("[AddItemToCart] AddItem: %w", err)
		}

		_, err = cartRepo.SaveCartItem(ctx, tx, cartItem)
		if err != nil {
			return fmt.Errorf("[AddItemToCart] SaveCartItem: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[AddItemToCart] Transaction: %w", err)
	}

	res := &connect.Response[commurzpbv1.Cart]{
		Msg: protoserde.FromCartPkg(cart),
	}

	return res, nil
}

func (service *Service) getOrCreateCart(ctx context.Context, tx sqlcs.DBTX, userID ulids.ULID) (cart pkgorder.Cart, err error) {
	cartRepo := pkgorder.CartRepository{}

	cart, err = cartRepo.FindCartByUserID(ctx, tx, userID)

	if err != nil && !isNotFoundErr(err) {
		return pkgorder.Cart{}, fmt.Errorf("[getOrCreateCart] FindCartByUserID: %w", err)
	}

	if err == nil {
		return cart, nil
	}

	cart = pkgorder.NewCart(userID)
	_, err = cartRepo.SaveCart(ctx, tx, cart)
	if err != nil && !isNotFoundErr(err) {
		return pkgorder.Cart{}, fmt.Errorf("[getOrCreateCart] SaveCart: %w", err)
	}

	// refresh cart data
	cart, err = cartRepo.FindCartByUserID(ctx, tx, userID)
	if err != nil {
		return pkgorder.Cart{}, fmt.Errorf("[getOrCreateCart] FindCartByUserID: %w", err)
	}

	return cart, nil
}

func (service *Service) CheckoutAll(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CheckoutAllRequest],
) (*connect.Response[commurzpbv1.Order], error) {
	userID, err := pkgutil.ParseULID(req.Msg.UserId)
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] parse userID: %w", err)
	}
	order := pkgorder.Order{}

	orderNumber := pkgorder.OrderNumber(ulids.New().String())
	cartRepo := pkgorder.CartRepository{}
	orderRepo := pkgorder.OrderRepository{}

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		cart, err := cartRepo.FindCartByUserID(ctx, tx, userID)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] FindCartByUserID: %w", err)
		}

		order, err = cart.CheckoutAll(orderNumber)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] CheckoutAll: %w", err)
		}

		order, err = orderRepo.CreateOrder(ctx, tx, order)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] CreateOrder: %w", err)
		}

		err = cartRepo.DeleteCart(ctx, tx, cart)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] DeleteCart: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] Transaction: %w", err)
	}

	res := &connect.Response[commurzpbv1.Order]{
		Msg: protoserde.FromOrderPkg(order),
	}

	return res, nil
}
