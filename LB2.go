package main

import (
	"fmt"
)

// Category представляє категорію товару з податком
type Category struct {
	Name string  // Назва категорії
	Tax  float64 // Податок у відсотках
}

// Product представляє товар
type Product struct {
	ID       int      // Унікальний ідентифікатор товару
	Name     string   // Назва товару
	Category Category // Категорія товару
	Price    float64  // Початкова ціна без урахування податку
}

// FinalPrice обчислює кінцеву вартість товару з урахуванням податку
func (p Product) FinalPrice() float64 {
	return p.Price * (1 + p.Category.Tax/100)
}

// FindMostExpensiveInCategory знаходить найдорожчий товар у вказаній категорії
func FindMostExpensiveInCategory(products []Product, categoryName string) *Product {
	var mostExpensive *Product

	for i, product := range products {
		if product.Category.Name == categoryName {
			if mostExpensive == nil || product.FinalPrice() > mostExpensive.FinalPrice() {
				mostExpensive = &products[i]
			}
		}
	}
	return mostExpensive
}

// InitializeCategories створює список категорій
func InitializeCategories() (Category, Category) {
	return Category{Name: "Electronics", Tax: 20}, Category{Name: "Clothing", Tax: 10}
}

// InitializeProducts створює список товарів
func InitializeProducts(electronics, clothing Category) []Product {
	return []Product{
		{ID: 1, Name: "Laptop", Category: electronics, Price: 1000},
		{ID: 2, Name: "Smartphone", Category: electronics, Price: 800},
		{ID: 3, Name: "Jeans", Category: clothing, Price: 50},
		{ID: 4, Name: "Jacket", Category: clothing, Price: 120},
	}
}

// PrintMostExpensiveProduct виводить найдорожчий товар у заданій категорії
func PrintMostExpensiveProduct(products []Product, categoryName string) {
	mostExpensive := FindMostExpensiveInCategory(products, categoryName)
	if mostExpensive != nil {
		fmt.Printf("Найдорожчий товар у категорії %s: %s (Ціна з податком: %.2f)\n", categoryName, mostExpensive.Name, mostExpensive.FinalPrice())
	} else {
		fmt.Println("Товарів у цій категорії не знайдено.")
	}
}

func main() {
	electronics, clothing := InitializeCategories()
	products := InitializeProducts(electronics, clothing)
	PrintMostExpensiveProduct(products, "Electronics")
}
