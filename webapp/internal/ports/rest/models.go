package rest

type Category struct {
	ID   int     `db:"id" json:"id"`
	Name string  `db:"name" json:"name"`
	Tax  float64 `db:"tax" json:"tax"`
}

type Product struct {
	ID         int     `db:"id" json:"id"`
	Name       string  `db:"name" json:"name"`
	CategoryID int     `db:"category_id" json:"category_id"`
	Price      float64 `db:"price" json:"price"`
}
