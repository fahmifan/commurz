package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/auth"
	"github.com/fahmifan/commurz/pkg/core/pkgprice"
	"github.com/fahmifan/commurz/pkg/core/pkgproduct"
	"github.com/fahmifan/commurz/pkg/logs"
	"github.com/fahmifan/commurz/pkg/parseutil"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
)

func (service *Service) UpdateProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.UpdateProductStockRequest],
) (res *connect.Response[commurzpbv1.Empty], err error) {
	if err := service.can(ctx, auth.Manage, auth.Product); err != nil {
		return nil, err
	}

	productRepo := pkgproduct.ProductReader{}
	productWriter := pkgproduct.ProductWriter{}
	product := pkgproduct.Product{}

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		productID, err := parseutil.ParseULID(req.Msg.GetProductId())
		if err != nil {
			return fmt.Errorf("[AddProductStock] ParseULID: %w", err)
		}

		product, err = productRepo.FindProductByID(ctx, tx, productID)
		if err != nil {
			return fmt.Errorf("[AddProductStock] FindProductByID: %w", err)
		}

		if !product.SameVersion(req.Msg.GetVersion()) {
			return fmt.Errorf("[AddProductStock] VersionMismatch")
		}

		var stockIn, stockOut pkgproduct.ProductStock

		product, stockIn, err = product.AddStock(req.Msg.GetStockIn(), time.Now())
		if err != nil {
			return fmt.Errorf("[AddProductStock] AddStock: %w", err)
		}

		product, stockOut, err = product.ReduceStock(req.Msg.GetStockOut(), time.Now())
		if err != nil {
			return fmt.Errorf("[AddProductStock] ReduceStock: %w", err)
		}

		_, err = productWriter.BulkSaveProductStocks(ctx, tx, []pkgproduct.ProductStock{stockIn, stockOut})
		if err != nil {
			return fmt.Errorf("[AddProductStock] SaveProductStock: %w", err)
		}

		_, err = productWriter.BumpProductVersion(ctx, tx, product)
		if err != nil {
			return fmt.Errorf("[AddProductStock] BumpProductVersion: %w", err)
		}

		return nil
	})

	if err != nil {
		logs.ErrCtx(ctx, err, "UpdateProductStock: Transaction")
		return nil, ErrInternal
	}

	res = &connect.Response[commurzpbv1.Empty]{}

	return
}

func (service *Service) CreateProduct(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CreateProductRequest],
) (res *connect.Response[commurzpbv1.Product], err error) {
	productRepo := pkgproduct.ProductReader{}
	productWriter := pkgproduct.ProductWriter{}

	product, err := pkgproduct.CreateProduct(req.Msg.Name, pkgprice.New(req.Msg.Price))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	product, err = productWriter.SaveProduct(ctx, service.DB, product)
	if err != nil {
		logs.ErrCtx(ctx, err, "[CreateProduct] SaveProduct")
		return nil, ErrInternal
	}

	product, err = productRepo.FindProductByID(ctx, service.DB, product.ID)
	if err != nil {
		logs.ErrCtx(ctx, err, "[CreateProduct] FindProductByID")
		return nil, ErrInternal
	}

	res = &connect.Response[commurzpbv1.Product]{
		Msg: protoserde.FromProductPkg(product),
	}

	return res, nil
}
