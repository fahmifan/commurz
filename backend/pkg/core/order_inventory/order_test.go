package order_inventory_test

import (
	"testing"
	"time"

	"github.com/fahmifan/commurz/pkg/core/order_inventory"
	"github.com/fahmifan/commurz/pkg/pkgmoney"
	"github.com/fahmifan/ulids"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCart_AddItem(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cart := order_inventory.NewCart(uuid.New())
		now := time.Now()

		prod1, err := order_inventory.CreateProduct("prod1", pkgmoney.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		cart, item, err := cart.AddItem(prod1, 1)
		require.NoError(t, err)
		require.NotEmpty(t, item)

		require.Equal(t, 1, len(cart.Items))
	})

	t.Run("failed out of stock", func(t *testing.T) {
		cart := order_inventory.NewCart(uuid.New())
		now := time.Now()

		prod1, err := order_inventory.CreateProduct("prod1", pkgmoney.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		_, _, err = cart.AddItem(prod1, 11)
		require.ErrorAs(t, err, &order_inventory.ErrOutOfStock)
	})

	t.Run("failed product not found", func(t *testing.T) {
	})
}

func TestCart_RemoveItem(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cart := order_inventory.NewCart(uuid.New())
		now := time.Now()

		prod1, err := order_inventory.CreateProduct("prod1", pkgmoney.New(1_000))
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
		cart := order_inventory.NewCart(uuid.New())

		_, _, err := cart.RemoveItem(ulids.New())
		require.ErrorAs(t, err, &order_inventory.ErrNotFound)
	})

	t.Run("wrong id, should not found", func(t *testing.T) {
		cart := order_inventory.NewCart(uuid.New())
		now := time.Now()

		prod1, err := order_inventory.CreateProduct("prod1", pkgmoney.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		cart, _, err = cart.AddItem(prod1, 1)
		require.NoError(t, err)
		require.Equal(t, 1, len(cart.Items))

		_, _, err = cart.RemoveItem(ulids.New())
		require.ErrorAs(t, err, &order_inventory.ErrNotFound)
	})
}

func TestCart_CheckoutAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cart := order_inventory.NewCart(uuid.New())
		now := time.Now()

		prod1, err := order_inventory.CreateProduct("prod1", pkgmoney.New(1_000))
		require.NoError(t, err)

		prod1, _, _ = prod1.AddStock(10, now)
		require.Equal(t, int64(10), prod1.CurrentStock())

		cart, _, err = cart.AddItem(prod1, 1)
		require.NoError(t, err)
		require.Equal(t, 1, len(cart.Items))

		newOrderNumber := order_inventory.OrderNumber("order-123")
		cart, order, checkedOutStocks, err := cart.CheckoutAll(newOrderNumber, now)
		require.NoError(t, err)
		require.Equal(t, prod1.ID, checkedOutStocks[0].ProductID)
		require.Equal(t, pkgmoney.New(1_000), order.TotalPrice())
	})
}
