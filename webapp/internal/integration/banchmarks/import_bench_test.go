package banchmarks

import (
	"context"
	"log"
	"testing"

	"github.com/OleksiiYevdokimov/OPOGO/webapp/internal"
	"github.com/OleksiiYevdokimov/OPOGO/webapp/internal/adapters/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func BenchmarkProductsImport(b *testing.B) {
	dbExec, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	productsClient := postgres.NewClient(dbExec)
	productsService := internal.NewProducts(productsClient)
	productsParser := ftp.NewParser(productsService)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err = productsParser.Run(ctx, "/Users/denys/Go/src/github.com/DenisGoldiner/webapp/internal/integration/data/products.csv"); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func newDB() (sqlx.ExtContext, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/travellers?sslmode=disable"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
