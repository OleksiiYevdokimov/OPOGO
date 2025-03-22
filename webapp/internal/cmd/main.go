package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/DenisGoldiner/webapp/internal"
	"github.com/DenisGoldiner/webapp/internal/adapters/postgres"
	"github.com/DenisGoldiner/webapp/internal/ports/rest"
)

func main() {
	app := newApplication()
	app.start()
}

type application struct {
	server *http.Server
}

func newApplication() application {
	db, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	server := newServer(db)

	return application{
		server: server,
	}
}

func newDB() (*sqlx.DB, error) {
	dsn := "postgres://postgres:12345678@localhost:5432/postgres?sslmode=disable&search_path=shop"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newServer(db *sqlx.DB) *http.Server {
	client := postgres.NewClient(db)
	service := internal.NewService(client)
	handler := rest.NewHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/categories", handler.CreateCategory).Methods("POST")
	r.HandleFunc("/api/v1/categories/{id:[0-9]+}", handler.GetCategory).Methods("GET")

	r.HandleFunc("/api/v1/products", handler.CreateProduct).Methods("POST")
	r.HandleFunc("/api/v1/products/{id:[0-9]+}", handler.GetProduct).Methods("GET")

	return &http.Server{
		Addr:    "localhost:8081",
		Handler: r,
	}
}

func (app application) start() {
	log.Println("Сервер запущено на http://localhost:8081")
	if err := app.server.ListenAndServe(); err != nil {
		log.Fatal("Помилка запуску сервера:", err)
	}
}
