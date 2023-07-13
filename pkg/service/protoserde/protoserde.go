package protoserde

import (
	"github.com/fahmifan/commurz/pkg/internal/pkgorder"
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	commurzpbv1 "github.com/fahmifan/commurz/protogen/commurzpb/v1"
	"github.com/samber/lo"
)

func FromUserPkg(user pkguser.User) *commurzpbv1.User {
	return &commurzpbv1.User{
		Id:    user.ID.String(),
		Email: user.Email,
	}
}

func FromProductPkg(product pkgproduct.Product) *commurzpbv1.Product {
	return &commurzpbv1.Product{
		Id:           product.ID.String(),
		Name:         product.Name,
		Price:        int64(product.Price),
		CurrentStock: product.CurrentStock(),
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
		ProductPrice: int64(item.ProductPrice),
		Product:      FromProductToOrderProduct(item.Product),
	}
}

func FromProductToOrderProduct(product pkgproduct.Product) *commurzpbv1.OrderProduct {
	return &commurzpbv1.OrderProduct{
		Id:           product.ID.String(),
		Name:         product.Name,
		Price:        int64(product.Price),
		CurrentStock: product.CurrentStock(),
	}
}

func FromOrderPkg(order pkgorder.Order) *commurzpbv1.Order {
	return &commurzpbv1.Order{
		Id:         order.ID.String(),
		UserId:     order.UserID.String(),
		TotalPrice: int64(order.TotalPrice()),
	}
}
