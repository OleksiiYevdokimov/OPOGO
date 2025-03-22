package internal

import (
	"context"
	_ "github.com/google/uuid"
)

type Storage interface {
	CreateCategory(ctx context.Context, category Category) (int, error)
	GetCategory(ctx context.Context, id int) (Category, error)
	CreateProduct(ctx context.Context, product Product) (int, error)
	GetProduct(ctx context.Context, id int) (Product, error)
}

// Category представляє категорію товару
type Category struct {
	ID   int     `db:"id"`
	Name string  `db:"name"`
	Tax  float64 `db:"tax"`
}

// Product представляє товар
type Product struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	CategoryID int     `db:"category_id"`
	Price      float64 `db:"price"`
}
