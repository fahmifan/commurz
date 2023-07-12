package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
	commurzpbv1 "github.com/fahmifan/commurz/protogen/commurzpb/v1"
	"github.com/fahmifan/commurz/protogen/commurzpb/v1/commurzpbv1connect"
)

var _ commurzpbv1connect.CommurzServiceHandler = &CommurzService{}

type CommurzService struct {
	*Config
}

func (service *CommurzService) CreateUser(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CreateUserRequest],
) (res *connect.Response[commurzpbv1.User], err error) {
	userRepo := pkguser.UserRepository{}

	user := pkguser.NewUser(req.Msg.Email)
	user, err = userRepo.CreateUser(ctx, service.db, user)
	if err != nil {
		return res, fmt.Errorf("[CreateUser] SaveUser: %w", err)
	}

	// use complete user fields
	user, err = userRepo.FindUserByID(ctx, service.db, user.ID)
	if err != nil {
		return res, fmt.Errorf("[CreateUser] FindUserByID: %w", err)
	}

	res = &connect.Response[commurzpbv1.User]{
		Msg: protoserde.FromUserPkg(user),
	}

	return res, nil
}

func (service *CommurzService) CreateProduct(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CreateProductRequest],
) (res *connect.Response[commurzpbv1.Product], err error) {
	productRepo := pkgproduct.ProductRepository{}

	product := pkgproduct.CreateProduct(req.Msg.Name, pkgproduct.Price(req.Msg.Price))
	product, err = productRepo.SaveProduct(ctx, service.db, product)
	if err != nil {
		return res, fmt.Errorf("[CreateProduct] CreateProduct: %w", err)
	}

	// use complete product fields
	product, err = productRepo.FindProductByID(ctx, service.db, product.ID)
	if err != nil {
		return res, fmt.Errorf("[CreateProduct] FindProductByID: %w", err)
	}

	res = &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return res, nil
}

func (service *CommurzService) AddProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ChangeProductStockRequest],
) (res *connect.Response[commurzpbv1.Product], err error) {
	var product pkgproduct.Product

	err = Transaction(ctx, service.db, func(tx *sql.Tx) error {
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
		product, stock = product.AddStock(req.Msg.GetStockQty(), time.Now())

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

func (service *CommurzService) ReduceProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.ChangeProductStockRequest],
) (*connect.Response[commurzpbv1.Product], error) {
	productRepo := pkgproduct.ProductRepository{}

	product := pkgproduct.Product{}
	productID, err := pkgutil.ParseULID(req.Msg.ProductId)
	if err != nil {
		return nil, fmt.Errorf("[ReduceProductStock] ParseULID: %w", err)
	}

	err = Transaction(ctx, service.db, func(tx *sql.Tx) error {
		product, err = productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] FindProductByID: %w", err)
		}

		var stock pkgproduct.ProductStock
		product, stock, err = product.ReduceStock(req.Msg.GetStockQty(), time.Now())
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

	product, err = productRepo.FindProductByID(ctx, service.db, productID)
	if err != nil {
		return nil, fmt.Errorf("[ReduceProductStock] FindProductByID: %w", err)
	}

	res := &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return res, nil
}
