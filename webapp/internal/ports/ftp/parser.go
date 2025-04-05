package ftp

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"

	"github.com/OleksiiYevdokimov/OPOGO/webapp/internal"
)

type Parser struct {
	service internal.Service
}

func NewParser(service internal.Service) Parser {
	return Parser{service: service}
}

func (p Parser) Run(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("не вдалося відкрити файл %s: %w", filePath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	reader := csv.NewReader(bytes.NewReader(data))

	categories, products, err := p.parse(reader)
	if err != nil {
		return err
	}

	// Додаємо категорії в БД
	if err = p.processCategories(ctx, categories); err != nil {
		return err
	}

	// Додаємо товари в БД
	if err = p.processProducts(ctx, products); err != nil {
		return err
	}

	return nil
}

func (p Parser) parse(r *csv.Reader) ([]internal.Category, []internal.Product, error) {
	var categories []internal.Category
	var products []internal.Product

	if err := p.parseHeader(r); err != nil {
		return nil, nil, err
	}

	for i := 1; ; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("помилка при парсингу рядка #%d: %w", i, err)
		}

		// Якщо 2 колонки → це категорія
		if len(row) != 5 {
			return nil, nil, fmt.Errorf("невідомий формат у рядку #%d", i)
		}
		tax, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			return nil, nil, fmt.Errorf("некоректний формат податку #%d, значення = %q", i, row[4])
		}

		category := internal.Category{
			Name: row[0],
			Tax:  tax,
		}
		categories = append(categories, category)
	}

	return categories, products, nil
}

var supportedHeader = []string{"ID", "Product Name", "Category", "Price", "Tax"}

func (p Parser) parseHeader(r *csv.Reader) error {
	row, err := r.Read()
	if err == io.EOF {
		return fmt.Errorf("файл пустий %w", err)
	}

	if err != nil {
		return fmt.Errorf("не вдалось прочитати хедер %w", err)
	}

	if !slices.Equal(row, supportedHeader) {
		return fmt.Errorf("отриманий хедер не відповідає стандарту %#v", row)
	}

	return nil
}

func (p Parser) processCategories(ctx context.Context, categories []internal.Category) error {
	for _, category := range categories {
		if _, err := p.service.CreateCategory(ctx, category); err != nil {
			return fmt.Errorf("не вдалося створити категорію %s: %w", category.Name, err)
		}
	}
	return nil
}

func (p Parser) processProducts(ctx context.Context, products []internal.Product) error {
	for _, product := range products {
		if _, err := p.service.CreateProduct(ctx, product); err != nil {
			return fmt.Errorf("не вдалося створити товар %s: %w", product.Name, err)
		}
	}
	return nil
}
