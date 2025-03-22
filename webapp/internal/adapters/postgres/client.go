package postgres

import (
	"context"
	"fmt"

	"github.com/DenisGoldiner/webapp/internal"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

func NewClient(db *sqlx.DB) Client {
	return Client{db: db}
}

// Створення категорії
func (c Client) CreateCategory(ctx context.Context, category internal.Category) (int, error) {
	query := `INSERT INTO categories (name, tax) VALUES ($1, $2) RETURNING id`
	var id int
	err := c.db.QueryRowContext(ctx, query, category.Name, category.Tax).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("не вдалося створити категорію: %w", err)
	}
	return id, nil
}

// Отримання категорії за ID
func (c Client) GetCategory(ctx context.Context, id int) (internal.Category, error) {
	query := `SELECT id, name, tax FROM categories WHERE id = $1`
	var category internal.Category
	err := c.db.GetContext(ctx, &category, query, id)
	if err != nil {
		return internal.Category{}, fmt.Errorf("не вдалося отримати категорію: %w", err)
	}
	return category, nil
}

// Створення товару
func (c Client) CreateProduct(ctx context.Context, product internal.Product) (int, error) {
	query := `INSERT INTO products (name, category_id, price) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := c.db.QueryRowContext(ctx, query, product.Name, product.CategoryID, product.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("не вдалося створити товар: %w", err)
	}
	return id, nil
}

// Отримання товару за ID
func (c Client) GetProduct(ctx context.Context, id int) (internal.Product, error) {
	query := `SELECT id, name, category_id, price FROM products WHERE id = $1`
	var product internal.Product
	err := c.db.GetContext(ctx, &product, query, id)
	if err != nil {
		return internal.Product{}, fmt.Errorf("не вдалося отримати товар: %w", err)
	}
	return product, nil
}
