package protoserde

import (
	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	commurzpbv1 "github.com/fahmifan/commurz/protogen/commurzpb/v1"
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
