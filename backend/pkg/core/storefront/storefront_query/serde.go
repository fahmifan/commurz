package storefront_query

import (
	"github.com/fahmifan/commurz/pkg/core/storefront"
	commurzv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/samber/lo"
)

func ListFromProductListingsPkg(productListings []storefront.ProductListing) []*commurzv1.ProductListing {
	return lo.Map(productListings, func(product storefront.ProductListing, _ int) *commurzv1.ProductListing {
		return FromProductListingPkg(product)
	})
}

func FromProductListingPkg(product storefront.ProductListing) *commurzv1.ProductListing {
	return &commurzv1.ProductListing{
		Id:           product.ID.String(),
		Name:         product.Name,
		Price:        product.Price.IDR(),
		LatestStock:  product.LatestStock,
		TextPriceIdr: product.Price.String(),
	}
}
