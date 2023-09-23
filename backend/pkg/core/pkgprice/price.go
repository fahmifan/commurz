package pkgprice

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Currency string

const (
	IDR Currency = "IDR"
)

// Divider is the default divider of Price.
var Divider int64 = 100_000

// Price is in  IDR 1/divider.
//
// So, a Price(1_000_000) with divider(1000) is equal to 1_000 in IDR
type Price struct {
	value int64
}

// New create a new Price with default Divider of 100_000.
func New(value int64) Price {
	price := Price{
		value: value,
	}

	return price
}

func (price Price) Value() int64 {
	return price.value
}

// IDR return the integer value.
func (price Price) IDR() int64 {
	return price.value / Divider
}

// IDRCent return the decimal value of IDR.
func (price Price) IDRCent() int64 {
	return price.value % Divider
}

func (price Price) Times(times int64) Price {
	price.value *= times
	return price
}

func (price Price) Add(price2 Price) Price {
	price.value += price2.value
	return price
}

func (price Price) Sub(price2 Price) Price {
	price.value -= price2.value
	return price
}

func (price Price) String() string {
	printer := message.NewPrinter(language.Indonesian)
	return fmt.Sprintf("%s,%d", printer.Sprint(price.IDR()), price.IDRCent())
}
