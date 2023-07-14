package cli

import (
	"context"
	"database/sql"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/internal/pkgutil"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/fahmifan/commurz/pkg/service"
	_ "modernc.org/sqlite"
)

func Run(args ...string) error {
	dsnURI := "file::memory:?mode=memory&cache=shared&journal_mode=wal&_fk=1"
	db, err := sql.Open("sqlite", dsnURI)
	if err != nil {
		return err
	}

	if err := migrate(db); err != nil {
		return err
	}

	ctx := context.Background()
	svc := service.NewService(&service.Config{
		DB: db,
	})

	user, err := svc.CreateUser(ctx, &connect.Request[commurzpbv1.CreateUserRequest]{
		Msg: &commurzpbv1.CreateUserRequest{
			Email: "john@doe.com",
		},
	})
	if err != nil {
		return err
	}
	pkgutil.PrintlnDebug("user >>>", pkgutil.PrettyJSON(user))

	product, err := svc.CreateProduct(ctx, &connect.Request[commurzpbv1.CreateProductRequest]{
		Msg: &commurzpbv1.CreateProductRequest{
			Name:  "Kentang",
			Price: 10,
		},
	})
	if err != nil {
		return err
	}
	pkgutil.PrintlnDebug("product >>>", pkgutil.PrettyJSON(product))

	product, err = svc.AddProductStock(ctx, &connect.Request[commurzpbv1.ChangeProductStockRequest]{
		Msg: &commurzpbv1.ChangeProductStockRequest{
			ProductId:     product.Msg.GetId(),
			StockQuantity: 4,
		},
	})
	if err != nil {
		return err
	}
	pkgutil.PrintlnDebug("product >>>", pkgutil.PrettyJSON(product))

	product, err = svc.ReduceProductStock(ctx, &connect.Request[commurzpbv1.ChangeProductStockRequest]{
		Msg: &commurzpbv1.ChangeProductStockRequest{
			ProductId:     product.Msg.GetId(),
			StockQuantity: 1,
		},
	})
	if err != nil {
		return err
	}
	pkgutil.PrintlnDebug("product >>>", pkgutil.PrettyJSON(product))

	cart, err := svc.AddItemToCart(ctx, &connect.Request[commurzpbv1.AddItemToCartRequest]{
		Msg: &commurzpbv1.AddItemToCartRequest{
			ProductId: product.Msg.GetId(),
			UserId:    user.Msg.GetId(),
			Quantity:  2,
		},
	})
	if err != nil {
		return err
	}
	pkgutil.PrintlnDebug("cart >>>", pkgutil.PrettyJSON(cart))

	order, err := svc.CheckoutAll(ctx, &connect.Request[commurzpbv1.CheckoutAllRequest]{
		Msg: &commurzpbv1.CheckoutAllRequest{
			UserId: user.Msg.GetId(),
		},
	})
	if err != nil {
		return err
	}
	pkgutil.PrintlnDebug("order >>>", pkgutil.PrettyJSON(order))

	return nil
}

func migrate(db *sql.DB) error {
	filename := "sqlcdef/schema.sql"
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(buf)); err != nil {
		return err
	}

	return nil
}
