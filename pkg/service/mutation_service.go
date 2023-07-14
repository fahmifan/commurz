package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkgorder"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/service/protoserde"
	"github.com/fahmifan/ulids"
)

func (service *Service) CheckoutAll(
	ctx context.Context,
	req *connect.Request[commurzpbv1.CheckoutAllRequest],
) (*connect.Response[commurzpbv1.Order], error) {
	userID, err := pkgutil.ParseULID(req.Msg.UserId)
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] parse userID: %w", err)
	}

	cartReader := pkgorder.CartReader{}
	cartWriter := pkgorder.CartWriter{}
	orderReader := pkgorder.OrderReader{}
	orderWriter := pkgorder.OrderWriter{}
	productWriter := pkgproduct.ProductWriter{}

	orderNumber := pkgorder.OrderNumber(ulids.New().String())
	now := time.Now()
	order := pkgorder.Order{}
	var checkOutStocks []pkgproduct.ProductStock

	err = Transaction(ctx, service.DB, func(tx *sql.Tx) error {
		cart, err := cartReader.FindCartByUserID(ctx, tx, userID)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] FindCartByUserID: %w", err)
		}

		cart, order, checkOutStocks, err = cart.CheckoutAll(orderNumber, now)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] CheckoutAll: %w", err)
		}

		_, err = orderWriter.CreateOrder(ctx, tx, order)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] CreateOrder: %w", err)
		}

		_, err = orderWriter.BulkSaveOrderItems(ctx, tx, order.Items)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] BulkSaveOrderItems: %w", err)
		}

		// bump product version while saving the reduced product stocks to avoid race condition.
		// product stocks should have rolledback if the product version is failed to bumped.
		{
			_, err = productWriter.BulkSaveProductStocks(ctx, tx, checkOutStocks)
			if err != nil {
				return fmt.Errorf("[CheckoutAll] BulkSaveProductStocks: %w", err)
			}
			_, err = productWriter.BulkBumpProductVersion(ctx, tx, order.Products())
			if err != nil {
				return fmt.Errorf("[CheckoutAll] BulkBumpProductVersion: %w", err)
			}
		}

		err = cartWriter.DeleteCart(ctx, tx, cart)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] DeleteCart: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] Transaction: %w", err)
	}

	order, err = orderReader.FindOrderByID(ctx, service.DB, order.ID)
	if err != nil {
		return nil, fmt.Errorf("[CheckoutAll] FindOrderByID: %w", err)
	}

	res := &connect.Response[commurzpbv1.Order]{
		Msg: protoserde.FromOrderPkg(order),
	}

	return res, nil
}
