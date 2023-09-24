package order_inventory_cmd

import (
	"github.com/fahmifan/commurz/pkg/core/order_inventory"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/samber/lo"
)

func FromCartPkg(cart order_inventory.Cart) *commurzpbv1.Cart {
	return &commurzpbv1.Cart{
		Id:    cart.ID.String(),
		Items: FromCartItemsPkg(cart.Items),
	}
}

func FromCartItemsPkg(items []order_inventory.CartItem) []*commurzpbv1.CartItem {
	return lo.Map(items, func(item order_inventory.CartItem, _ int) *commurzpbv1.CartItem {
		return FromCartItemPkg(item)
	})
}

func FromCartItemPkg(item order_inventory.CartItem) *commurzpbv1.CartItem {
	return &commurzpbv1.CartItem{
		Id:           item.ID.String(),
		ProductId:    item.ProductID.String(),
		Quantity:     item.Quantity,
		CartId:       item.CartID.String(),
		ProductPrice: item.ProductPrice.IDR(),
		Product:      FromProductToOrderProduct(item.Product),
	}
}

func FromProductToOrderProduct(product order_inventory.Product) *commurzpbv1.OrderProduct {
	return &commurzpbv1.OrderProduct{
		Id:           product.ID.String(),
		Name:         product.Name,
		Price:        product.Price.IDR(),
		CurrentStock: product.CurrentStock(),
	}
}

func FromOrderPkg(order order_inventory.Order) *commurzpbv1.Order {
	return &commurzpbv1.Order{
		Id:         order.ID.String(),
		UserId:     order.UserID.String(),
		TotalPrice: order.TotalPrice().IDR(),
	}
}
