package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fahmifan/commurz/pkg/internal/pkgorder"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/sqlcs"
	"github.com/fahmifan/ulids"
)

type AddItemToCartRequest struct {
	ProductID ulids.ULID `json:"product_id"`
	Quantity  int64      `json:"quantity"`
	UserID    ulids.ULID `json:"user_id"`
}

func (s *Service) AddItemToCart(ctx context.Context, req AddItemToCartRequest) (cart pkgorder.Cart, err error) {
	cartRepo := pkgorder.CartRepository{}
	cart, err = s.getOrCreateCart(ctx, s.db, req.UserID)
	if err != nil {
		return pkgorder.Cart{}, fmt.Errorf("[AddItemToCart] getOrCreateCart: %w", err)
	}

	err = Transaction(s.db, func(tx *sql.Tx) error {
		productRepo := pkgproduct.ProductRepository{}

		product, err := productRepo.FindProductByID(ctx, tx, req.ProductID)
		if err != nil {
			return fmt.Errorf("[AddItemToCart] FindProductByID: %w", err)
		}

		var cartItem pkgorder.CartItem

		cart, cartItem, err = cart.AddItem(product, req.Quantity)
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
		return pkgorder.Cart{}, fmt.Errorf("[AddItemToCart] Transaction: %w", err)
	}

	return cart, nil
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

type CheckoutAllRequest struct {
	UserID ulids.ULID `json:"user_id"`
}

func (s *Service) CheckoutAll(ctx context.Context, req CheckoutAllRequest) (order pkgorder.Order, err error) {
	orderNumber := pkgorder.OrderNumber(ulids.New().String())

	err = Transaction(s.db, func(tx *sql.Tx) error {
		cartRepo := pkgorder.CartRepository{}

		cart, err := cartRepo.FindCartByUserID(ctx, tx, req.UserID)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] FindCartByUserID: %w", err)
		}

		order, err = cart.CheckoutAll(orderNumber)
		if err != nil {
			return fmt.Errorf("[CheckoutAll] CheckoutAll: %w", err)
		}

		return nil
	})
	if err != nil {
		return pkgorder.Order{}, fmt.Errorf("[CheckoutAll] Transaction: %w", err)
	}

	return order, nil
}
