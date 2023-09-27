package order_inventory_query

import (
	"github.com/fahmifan/commurz/pkg/core/order_inventory"
	commurzv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/samber/lo"
)

func ListFromProductPkg(products []order_inventory.Product) []*commurzv1.Product {
	return lo.Map(products, func(product order_inventory.Product, _ int) *commurzv1.Product {
		return FromProductPkg(product)
	})
}

func FromProductPkg(product order_inventory.Product) *commurzv1.Product {
	return &commurzv1.Product{
		Id:           product.ID.String(),
		Name:         product.Name,
		Price:        product.Price.IDR(),
		CurrentStock: product.CurrentStock(),
		TextPriceIdr: product.Price.String(),
		Version:      product.Version,
	}
}

func FromCartPkg(cart order_inventory.Cart) *commurzv1.Cart {
	return &commurzv1.Cart{
		Id:    cart.ID.String(),
		Items: FromCartItemsPkg(cart.Items),
	}
}

func FromCartItemsPkg(items []order_inventory.CartItem) []*commurzv1.CartItem {
	return lo.Map(items, func(item order_inventory.CartItem, _ int) *commurzv1.CartItem {
		return FromCartItemPkg(item)
	})
}

func FromCartItemPkg(item order_inventory.CartItem) *commurzv1.CartItem {
	return &commurzv1.CartItem{
		Id:           item.ID.String(),
		ProductId:    item.ProductID.String(),
		Quantity:     item.Quantity,
		CartId:       item.CartID.String(),
		ProductPrice: item.ProductPrice.IDR(),
		Product:      FromProductToOrderProduct(item.Product),
	}
}

func FromProductToOrderProduct(product order_inventory.Product) *commurzv1.OrderProduct {
	return &commurzv1.OrderProduct{
		Id:           product.ID.String(),
		Name:         product.Name,
		Price:        product.Price.IDR(),
		CurrentStock: product.CurrentStock(),
	}
}

func FromOrderPkg(order order_inventory.Order) *commurzv1.Order {
	return &commurzv1.Order{
		Id:         order.ID.String(),
		UserId:     order.UserID.String(),
		TotalPrice: order.TotalPrice().IDR(),
	}
}
