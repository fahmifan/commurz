package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	"github.com/fahmifan/commurz/pkg/logs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
)

func (service *Service) UpdateProductStock(
	ctx context.Context,
	req *connect.Request[commurzpbv1.UpdateProductStockRequest],
) (res *connect.Response[commurzpbv1.Empty], err error) {
	productRepo := pkgproduct.ProductReader{}
	productWriter := pkgproduct.ProductWriter{}
	product := pkgproduct.Product{}

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		productID, err := pkgutil.ParseULID(req.Msg.GetProductId())
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
		return nil, connect.NewError(connect.CodeInternal, ErrInternal)
	}

	res = &connect.Response[commurzpbv1.Empty]{}

	return
}
