package order_inventory_test

import (
	"testing"
	"time"

	"github.com/fahmifan/commurz/pkg/core/order_inventory"
	"github.com/fahmifan/commurz/pkg/pkgmoney"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	now := time.Now()
	price := pkgmoney.New(100)

	t.Run("able to add & reduce stock", func(t *testing.T) {
		product, err := order_inventory.CreateProduct("product 1", price)
		require.NoError(t, err)

		product, _, _ = product.AddStock(1, now)
		require.True(t, product.HaveStock(1))
		_, _, err = product.ReduceStock(1, now)
		require.NoError(t, err)
	})

	t.Run("should failed reduce insufficient stock", func(t *testing.T) {
		product, err := order_inventory.CreateProduct("product 1", price)
		require.NoError(t, err)

		product, _, _ = product.AddStock(1, now)

		_, _, err = product.ReduceStock(3, now)
		require.ErrorAs(t, err, &order_inventory.ErrInsufficientStock)
	})
}
