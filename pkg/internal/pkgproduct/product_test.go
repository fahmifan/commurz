package pkgproduct_test

import (
	"testing"
	"time"

	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/stretchr/testify/require"
)

func TestProduct(t *testing.T) {
	now := time.Now()

	t.Run("able to add & reduce stock", func(t *testing.T) {
		product := pkgproduct.CreateProduct("product 1", 1000)
		product, _ = product.AddStock(1, now)
		require.True(t, product.HaveStock(1))
		_, _, err := product.ReduceStock(1, now)
		require.NoError(t, err)
	})

	t.Run("should failed reduce insufficient stock", func(t *testing.T) {
		product := pkgproduct.CreateProduct("product 1", 1000)
		product, _ = product.AddStock(1, now)

		_, _, err := product.ReduceStock(3, now)
		require.ErrorAs(t, err, &pkgproduct.ErrInsufficientStock)
	})

}
