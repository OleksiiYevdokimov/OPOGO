package postgres

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
