package pkgorder_test

import (
	"testing"
	"time"

	"github.com/fahmifan/commurz/pkg/internal/pkgorder"
	"github.com/fahmifan/commurz/pkg/internal/pkgprice"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/ulids"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCart_AddItem(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cart := pkgorder.NewCart(uuid.New())
		now := time.Now()

		prod1, err := pkgproduct.CreateProduct("prod1", pkgprice.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		cart, item, err := cart.AddItem(prod1, 1)
		require.NoError(t, err)
		require.NotEmpty(t, item)

		require.Equal(t, 1, len(cart.Items))
	})

	t.Run("failed out of stock", func(t *testing.T) {
		cart := pkgorder.NewCart(uuid.New())
		now := time.Now()

		prod1, err := pkgproduct.CreateProduct("prod1", pkgprice.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		_, _, err = cart.AddItem(prod1, 11)
		require.ErrorAs(t, err, &pkgorder.ErrOutOfStock)
	})

	t.Run("failed product not found", func(t *testing.T) {
	})
}

func TestCart_RemoveItem(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cart := pkgorder.NewCart(uuid.New())
		now := time.Now()

		prod1, err := pkgproduct.CreateProduct("prod1", pkgprice.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		cart, item, err := cart.AddItem(prod1, 1)
		require.NoError(t, err)
		require.Equal(t, 1, len(cart.Items))

		cart, removedItem, err := cart.RemoveItem(item.ID)
		require.NoError(t, err)
		require.Equal(t, 0, len(cart.Items))
		require.Equal(t, item.ID.String(), removedItem.ID.String())
	})

	t.Run("cart is empty, should not found", func(t *testing.T) {
		cart := pkgorder.NewCart(uuid.New())

		_, _, err := cart.RemoveItem(ulids.New())
		require.ErrorAs(t, err, &pkgorder.ErrNotFound)
	})

	t.Run("wrong id, should not found", func(t *testing.T) {
		cart := pkgorder.NewCart(uuid.New())
		now := time.Now()

		prod1, err := pkgproduct.CreateProduct("prod1", pkgprice.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		cart, _, err = cart.AddItem(prod1, 1)
		require.NoError(t, err)
		require.Equal(t, 1, len(cart.Items))

		_, _, err = cart.RemoveItem(ulids.New())
		require.ErrorAs(t, err, &pkgorder.ErrNotFound)
	})
}

func TestCart_CheckoutAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cart := pkgorder.NewCart(uuid.New())
		now := time.Now()

		prod1, err := pkgproduct.CreateProduct("prod1", pkgprice.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		cart, _, err = cart.AddItem(prod1, 1)
		require.NoError(t, err)
		require.Equal(t, 1, len(cart.Items))

		newOrderNumber := pkgorder.OrderNumber("order-123")
		cart, order, checkedOutStocks, err := cart.CheckoutAll(newOrderNumber, now)
		require.NoError(t, err)
		require.Equal(t, prod1.ID, checkedOutStocks[0].ProductID)
		require.Equal(t, pkgprice.New(1_000), order.TotalPrice())
	})
}
