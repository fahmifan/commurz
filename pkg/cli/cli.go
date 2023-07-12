package cli

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

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
	svc := service.NewService(service.NewConfig(db))

	user, err := svc.CreateUser(ctx, service.CreateUserRequest{
		Email: "jondoe@email.com",
	})
	if err != nil {
		return err
	}

	product, err := svc.CreateProduct(ctx, service.CreateProductRequest{
		Name:  "Kentang",
		Price: 10,
	})
	if err != nil {
		return err
	}

	fmt.Println("DEBUG >>> product >>>", prettyJSON(product))

	product, err = svc.AddProductStock(ctx, service.AddProductStockRequest{
		ProductID: product.ID,
		Quantity:  4,
	})
	if err != nil {
		return err
	}

	fmt.Println("DEBUG >>> product >>>", prettyJSON(product))

	product, err = svc.ReduceProductStock(ctx, service.ReduceProductStockRequest{
		ProductID: product.ID,
		Quantity:  1,
	})
	if err != nil {
		return err
	}

	fmt.Println("DEBUG >>> product >>>", prettyJSON(product))

	cart, err := svc.AddItemToCart(ctx, service.AddItemToCartRequest{
		ProductID: product.ID,
		Quantity:  2,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("DEBUG >>> cart >>>", prettyJSON(cart))

	order, err := svc.CheckoutAll(ctx, service.CheckoutAllRequest{
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("DEBUG >>> order >>>", prettyJSON(order))

	return nil
}

func migrate(db *sql.DB) error {
	filename := "schema.sql"
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(buf)); err != nil {
		return err
	}

	return nil
}

func prettyJSON(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
