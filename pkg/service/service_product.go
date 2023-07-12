package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fahmifan/commurz/pkg/internal/pkgproduct"
	"github.com/fahmifan/ulids"
)

type CreateProductRequest struct {
	Name  string           `json:"name"`
	Price pkgproduct.Price `json:"price"`
}

func (service *Service) CreateProduct(
	ctx context.Context,
	req CreateProductRequest,
) (product pkgproduct.Product, err error) {
	productRepo := pkgproduct.ProductRepository{}

	product = pkgproduct.CreateProduct(req.Name, req.Price)
	product, err = productRepo.SaveProduct(ctx, service.db, product)
	if err != nil {
		return pkgproduct.Product{}, fmt.Errorf("[CreateProduct] CreateProduct: %w", err)
	}

	// use complete product fields
	product, err = productRepo.FindProductByID(ctx, service.db, product.ID)
	if err != nil {
		return pkgproduct.Product{}, fmt.Errorf("[CreateProduct] FindProductByID: %w", err)
	}

	return product, nil
}

type AddProductStockRequest struct {
	ProductID ulids.ULID `json:"product_id"`
	Quantity  int64      `json:"quantity"`
}

func (service *Service) AddProductStock(ctx context.Context, req AddProductStockRequest) (product pkgproduct.Product, err error) {
	err = Transaction(ctx, service.db, func(tx *sql.Tx) error {
		productRepo := pkgproduct.ProductRepository{}

		product, err = productRepo.FindProductByID(ctx, tx, req.ProductID)
		if err != nil {
			return fmt.Errorf("[AddProductStock] FindProductByID: %w", err)
		}

		var stock pkgproduct.ProductStock
		product, stock = product.AddStock(req.Quantity, time.Now())

		_, err = productRepo.SaveProductStock(ctx, tx, stock)
		if err != nil {
			return fmt.Errorf("[AddProductStock] SaveProductStock: %w", err)
		}

		return nil
	})

	if err != nil {
		return pkgproduct.Product{}, fmt.Errorf("[AddProductStock] Transaction: %w", err)
	}

	return product, nil
}

type ReduceProductStockRequest struct {
	ProductID ulids.ULID `json:"product_id"`
	Quantity  int64      `json:"quantity"`
}

func (service *Service) ReduceProductStock(ctx context.Context, req ReduceProductStockRequest) (product pkgproduct.Product, err error) {
	productRepo := pkgproduct.ProductRepository{}

	err = Transaction(ctx, service.db, func(tx *sql.Tx) error {
		product, err = productRepo.FindProductByID(ctx, tx, req.ProductID)
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] FindProductByID: %w", err)
		}

		var stock pkgproduct.ProductStock
		product, stock, err = product.ReduceStock(req.Quantity, time.Now())
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] ReduceStock: %w", err)
		}

		_, err = productRepo.SaveProductStock(ctx, tx, stock)
		if err != nil {
			return fmt.Errorf("[ReduceProductStock] SaveProductStock: %w", err)
		}
		return nil
	})

	if err != nil {
		return pkgproduct.Product{}, fmt.Errorf("[AddProductStock] Transaction: %w", err)
	}

	product, err = productRepo.FindProductByID(ctx, service.db, req.ProductID)
	if err != nil {
		return pkgproduct.Product{}, fmt.Errorf("[ReduceProductStock] FindProductByID: %w", err)
	}

	return product, nil
}
