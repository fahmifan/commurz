package pkgmoney

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Currency string

const (
	IDR Currency = "IDR"
)

// Divider is the default divider of Money.
var Divider int64 = 100_000

// Money is in  IDR 1/divider.
//
// So, a Money(1_000_000) with divider(1000) is equal to 1_000 in IDR
type Money struct {
	value int64
}

// New create a new Money with default Divider of 100_000.
func New(value int64) Money {
	money := Money{
		value: value,
	}

	return money
}

func (money Money) Value() int64 {
	return money.value
}

// IDR return the integer value.
func (money Money) IDR() int64 {
	return money.value / Divider
}

// IDRCent return the decimal value of IDR.
func (money Money) IDRCent() int64 {
	return money.value % Divider
}

func (money Money) Times(times int64) Money {
	money.value *= times
	return money
}

func (money Money) Add(money2 Money) Money {
	money.value += money2.value
	return money
}

func (money Money) Sub(money2 Money) Money {
	money.value -= money2.value
	return money
}

func (money Money) String() string {
	printer := message.NewPrinter(language.Indonesian)
	return fmt.Sprintf("%s,%d", printer.Sprint(money.IDR()), money.IDRCent())
}
