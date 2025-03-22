package internal

import (
	"context"
	"fmt"
)

// Service визначає інтерфейс бізнес-логіки
type Service struct {
	db Storage
}

func NewService(db Storage) Service {
	return Service{db: db}
}

// Створення категорії
func (s Service) CreateCategory(ctx context.Context, category Category) (int, error) {
	id, err := s.db.CreateCategory(ctx, category)
	if err != nil {
		return 0, fmt.Errorf("не вдалося створити категорію: %w", err)
	}
	return id, nil
}

// Отримання категорії за ID
func (s Service) GetCategory(ctx context.Context, id int) (Category, error) {
	category, err := s.db.GetCategory(ctx, id)
	if err != nil {
		return Category{}, fmt.Errorf("не вдалося отримати категорію: %w", err)
	}
	return category, nil
}

// Створення товару
func (s Service) CreateProduct(ctx context.Context, product Product) (int, error) {
	id, err := s.db.CreateProduct(ctx, product)
	if err != nil {
		return 0, fmt.Errorf("не вдалося створити товар: %w", err)
	}
	return id, nil
}

// Отримання товару за ID
func (s Service) GetProduct(ctx context.Context, id int) (Product, error) {
	product, err := s.db.GetProduct(ctx, id)
	if err != nil {
		return Product{}, fmt.Errorf("не вдалося отримати товар: %w", err)
	}
	return product, nil
}
