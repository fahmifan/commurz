package protoserde

import (
	"github.com/fahmifan/commurz/pkg/internal/pkgorder"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/samber/lo"
)

func ListFromUserPkg(users []pkguser.User) []*commurzpbv1.User {
	return lo.Map(users, func(user pkguser.User, _ int) *commurzpbv1.User {
		return FromUserPkg(user)
	})
}

func FromUserPkg(user pkguser.User) *commurzpbv1.User {
	return &commurzpbv1.User{
		Id:    user.ID.String(),
		Email: user.Email,
		Role:  string(user.Role),
	}
}

func ListFromProductPkg(products []pkgproduct.Product) []*commurzpbv1.Product {
	return lo.Map(products, func(product pkgproduct.Product, _ int) *commurzpbv1.Product {
		return FromProductPkg(product)
	})
}

func FromProductPkg(product pkgproduct.Product) *commurzpbv1.Product {
	return &commurzpbv1.Product{
		Id:           product.ID.String(),
		Name:         product.Name,
		Price:        product.Price.Value(),
		CurrentStock: product.CurrentStock(),
		TextPriceIdr: product.Price.String(),
	}
}

func FromCartPkg(cart pkgorder.Cart) *commurzpbv1.Cart {
	return &commurzpbv1.Cart{
		Id:    cart.ID.String(),
		Items: FromCartItemsPkg(cart.Items),
	}
}

func FromCartItemsPkg(items []pkgorder.CartItem) []*commurzpbv1.CartItem {
	return lo.Map(items, func(item pkgorder.CartItem, _ int) *commurzpbv1.CartItem {
		return FromCartItemPkg(item)
	})
}

func FromCartItemPkg(item pkgorder.CartItem) *commurzpbv1.CartItem {
	return &commurzpbv1.CartItem{
		Id:           item.ID.String(),
		ProductId:    item.ProductID.String(),
		Quantity:     item.Quantity,
		CartId:       item.CartID.String(),
		ProductPrice: item.ProductPrice.Value(),
		Product:      FromProductToOrderProduct(item.Product),
	}
}

func FromProductToOrderProduct(product pkgproduct.Product) *commurzpbv1.OrderProduct {
	return &commurzpbv1.OrderProduct{
		Id:           product.ID.String(),
		Name:         product.Name,
		Price:        product.Price.Value(),
		CurrentStock: product.CurrentStock(),
	}
}

func FromOrderPkg(order pkgorder.Order) *commurzpbv1.Order {
	return &commurzpbv1.Order{
		Id:         order.ID.String(),
		UserId:     order.UserID.String(),
		TotalPrice: order.TotalPrice().IDR(),
	}
}
