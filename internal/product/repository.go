package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

var ErrProductNotFound = errors.New("Product not found")

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, param CreateProductRequestParam) error {
	_, err := r.db.Exec(ctx, queryCreateProduct, param.ID, param.Name, param.Description, param.Price, param.Stock)

	if err != nil {
		return fmt.Errorf("Error create product %w", err)
	}

	return nil
}

func (r *Repository) GetProducts(ctx context.Context, page, limit int) (ProductListResponse, error) {
	rows, err := r.db.Query(ctx, queryGetProductList, page, limit)
	if err != nil {
		return ProductListResponse{}, fmt.Errorf("Error getting product list %w", err)
	}

	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product

		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.ModifiedAt)

		if err != nil {
			return ProductListResponse{}, fmt.Errorf("Error getting products %w", err)
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return ProductListResponse{}, fmt.Errorf("Problem iterating products %w", err)
	}

	productCount := 0
	if err := r.db.QueryRow(ctx, queryGetProductCount, page, limit).Scan(&productCount); err != nil {
		return ProductListResponse{}, fmt.Errorf("Problem getting product count %w", err)
	}

	return ProductListResponse{Products: products, Count: productCount}, nil
}

func (r *Repository) GetProductById(ctx context.Context, id uuid.UUID) (*Product, error) {
	var p Product
	err := r.db.QueryRow(ctx, queryGetProductById, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.ModifiedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("Error getting product %w", err)
	}

	return &p, nil
}

func (r *Repository) UpdateProductById(ctx context.Context, p ProductUpdateRequestParam) (*Product, error) {
	var result Product
	err := r.db.QueryRow(ctx, queryUpdateProduct, p.Name, p.Description, p.Price, p.Stock, p.ID).Scan(
		&result.ID, &result.Name, &result.Description, &result.Price, &p.Stock, &result.CreatedAt, &result.ModifiedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}

		return nil, fmt.Errorf("Error updating product %w", err)
	}

	return &result, nil
}

func (r *Repository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx, queryDeleteProduct, id)

	if err != nil {
		return fmt.Errorf("Error deleting product %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	return nil
}
