package main

import (
"context"
"database/sql"
"encoding/csv"
"fmt"
"log"
"math/rand"
"os"
"strconv"
"time"

_ "github.com/lib/pq"
)

const (
connStr       = "postgres://postgres:12345678@localhost:5432/postgres?sslmode=disable&search_path=shop"
numCategories = 20
numProducts   = 10000
csvFilePath   = "products.csv"
)

func main() {
db, err := sql.Open("postgres", connStr)
if err != nil {
log.Fatal("Failed to connect to database:", err)
}
defer db.Close()

ctx := context.Background()

// Створюємо категорії
fmt.Println("Generating categories...")
categoryIDs, err := generateCategories(ctx, db)
if err != nil {
log.Fatal("Failed to generate categories:", err)
}

// Створюємо продукти
fmt.Println("Generating products...")
if err := generateProducts(ctx, db, categoryIDs); err != nil {
log.Fatal("Failed to generate products:", err)
}

// Експортуємо у CSV
fmt.Println("Exporting to CSV...")
if err := exportToCSV(ctx, db); err != nil {
log.Fatal("Failed to export to CSV:", err)
}

fmt.Println("Done!")
}

func generateCategories(ctx context.Context, db *sql.DB) ([]int, error) {
categoryIDs := []int{}
for i := 1; i <= numCategories; i++ {
name := fmt.Sprintf("Category_%d", i)
tax := rand.Float64() * 20 // Випадковий податок від 0 до 20
var id int
err := db.QueryRowContext(ctx, "INSERT INTO categories (name, tax) VALUES ($1, $2) RETURNING id", name, tax).Scan(&id)
if err != nil {
return nil, err
}
categoryIDs = append(categoryIDs, id)
}
return categoryIDs, nil
}

func generateProducts(ctx context.Context, db *sql.DB, categoryIDs []int) error {
rand.Seed(time.Now().UnixNano())

for i := 1; i <= numProducts; i++ {
name := fmt.Sprintf("Product_%d", i)
categoryID := categoryIDs[rand.Intn(len(categoryIDs))] // Випадкова категорія
price := rand.Float64() * 500                          // Випадкова ціна до 500

_, err := db.ExecContext(ctx, "INSERT INTO products (name, category_id, price) VALUES ($1, $2, $3)", name, categoryID, price)
if err != nil {
return err
}
}

return nil
}

func exportToCSV(ctx context.Context, db *sql.DB) error {
rows, err := db.QueryContext(ctx, `
		SELECT p.id, p.name, c.name AS category, p.price, c.tax
		FROM products p
		JOIN categories c ON p.category_id = c.id
	`)
if err != nil {
return err
}
defer rows.Close()

file, err := os.Create(csvFilePath)
if err != nil {
return err
}
defer file.Close()

writer := csv.NewWriter(file)
defer writer.Flush()

// Записуємо заголовки
writer.Write([]string{"ID", "Product Name", "Category", "Price", "Tax"})

// Записуємо дані
for rows.Next() {
var id int
var productName, category string
var price, tax float64

if err := rows.Scan(&id, &productName, &category, &price, &tax); err != nil {
return err
}

record := []string{
strconv.Itoa(id),
productName,
category,
fmt.Sprintf("%.2f", price),
fmt.Sprintf("%.2f", tax),
}
writer.Write(record)
}

return nil
}