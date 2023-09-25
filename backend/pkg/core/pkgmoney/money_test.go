package pkgmoney_test

import (
	"testing"

	"github.com/fahmifan/commurz/pkg/core/pkgmoney"
	"github.com/stretchr/testify/require"
)

func TestPrice(t *testing.T) {
	pkgmoney.Divider = 100

	price := pkgmoney.New(123456)

	require.Equal(t, price.IDR(), int64(1234))
	require.Equal(t, price.IDRCent(), int64(56))
	require.Equal(t, price.Value(), int64(123456))
	require.Equal(t, price.String(), "1.234,56")
}

func TestPrice_Arithmatic(t *testing.T) {
	pkgmoney.Divider = 100

	price := pkgmoney.New(10_000)

	price = price.Times(3)
	require.Equal(t, "30,0", price.String())

	price2 := pkgmoney.New(500)
	price = price.Add(price2)

	require.Equal(t, "30,500", price.String())

	price3 := pkgmoney.New(16)
	price = price.Sub(price3)

	require.Equal(t, "30,484", price.String())
}
