package main

import (
	"github.com/stretchr/testify/require" // Імпортуємо бібліотеку для зручних асертів
	"testing"                             // Імпортуємо пакет для тестування
)

// TestFinalPrice перевіряє правильність розрахунку кінцевої ціни товару з урахуванням податку.
func TestFinalPrice(t *testing.T) {
	product := Product{
		ID:       1,                                      // ID товару
		Name:     "Laptop",                               // Назва товару
		Category: Category{Name: "Electronics", Tax: 20}, // Категорія товару з податком 20%
		Price:    1000,                                   // Базова ціна товару
	}
	expected := 1200.0                                           // Очікувана ціна з урахуванням податку (1000 + 20%)
	require.InEpsilon(t, expected, product.FinalPrice(), 0.0001) // Перевіряємо, що розрахунок коректний з допуском
}

// TestFindMostExpensiveInCategory перевіряє, чи знаходиться найдорожчий товар у заданій категорії.
func TestFindMostExpensiveInCategory(t *testing.T) {
	products := []Product{
		{ID: 1, Name: "Laptop", Category: Category{Name: "Electronics", Tax: 20}, Price: 1000},    // Товар 1
		{ID: 2, Name: "Smartphone", Category: Category{Name: "Electronics", Tax: 20}, Price: 800}, // Товар 2
	}

	result := FindMostExpensiveInCategory(products, "Electronics") // Викликаємо функцію для пошуку
	require.NotNil(t, result)                                      // Перевіряємо, що знайдено товар
	require.Equal(t, "Laptop", result.Name)                        // Перевіряємо, що знайдено саме "Laptop"
}

// TestFindMostExpensiveInCategoryEmpty перевіряє випадок, коли список товарів порожній.
func TestFindMostExpensiveInCategoryEmpty(t *testing.T) {
	products := []Product{}                                        // Порожній список товарів
	result := FindMostExpensiveInCategory(products, "Electronics") // Викликаємо функцію
	require.Nil(t, result)                                         // Очікуємо, що функція поверне nil
}

// TestFinalPrice_Table виконує тестування FinalPrice з використанням table-driven підходу.
func TestFinalPrice_Table(t *testing.T) {
	tests := map[string]struct {
		product       Product // Вхідний товар
		expectedPrice float64 // Очікувана кінцева ціна
	}{
		"Laptop":         {Product{ID: 1, Name: "Laptop", Category: Category{Name: "Electronics", Tax: 20}, Price: 1000}, 1200},
		"Phone":          {Product{ID: 2, Name: "Phone", Category: Category{Name: "Electronics", Tax: 15}, Price: 800}, 920},
		"Shirt":          {Product{ID: 3, Name: "Shirt", Category: Category{Name: "Clothing", Tax: 10}, Price: 50}, 55},
		"_first_product": {Product{ID: 4, Name: " first product", Category: Category{Name: "Clothing", Tax: 10}, Price: 50}, 55},
		"empty_product":  {Product{}, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) { // Запускаємо під-тест для кожного товару
			//require.InEpsilon(t, test.expectedPrice, test.product.FinalPrice(), 0.0001) // Перевіряємо розрахунок
			require.Equal(t, test.expectedPrice, test.product.FinalPrice())
		})
	}
}
