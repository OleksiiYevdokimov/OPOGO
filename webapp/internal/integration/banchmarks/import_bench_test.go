package banchmarks

import (
	"context"
	"log"
	"testing"

	"github.com/OleksiiYevdokimov/OPOGO/webapp/internal"
	"github.com/OleksiiYevdokimov/OPOGO/webapp/internal/adapters/postgres"
	"github.com/OleksiiYevdokimov/OPOGO/webapp/internal/ports/ftp"
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
	productsService := internal.NewService(productsClient)
	productsParser := ftp.NewParser(productsService)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err = productsParser.Run(ctx, `D:\GO\src\github.com\OleksiiYevdokimov\OPOGO\webapp\internal\integration\data\products.csv`); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func newDB() (*sqlx.DB, error) {
	dsn := "postgres://postgres:12345678@localhost:5432/postgres?sslmode=disable&search_path=shop"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
