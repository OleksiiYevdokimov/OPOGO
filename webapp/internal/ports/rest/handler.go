package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/DenisGoldiner/webapp/internal"
	"github.com/gorilla/mux"
)

type Handler struct {
	service internal.Service
}

func NewHandler(service internal.Service) Handler {
	return Handler{service: service}
}

// Створення категорії
func (h Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category internal.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		log.Printf("Помилка створення категорії: %v", err)
		http.Error(w, "Невірний формат запиту", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateCategory(r.Context(), category)
	if err != nil {
		log.Printf("Помилка створення категорії: %v", err)
		http.Error(w, "Помилка створення категорії", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// Отримання категорії за ID
func (h Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Помилка створення категорії: %v", err)
		http.Error(w, "Невірний формат ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetCategory(r.Context(), id)
	if err != nil {
		log.Printf("Помилка створення категорії: %v", err)
		http.Error(w, "Категорію не знайдено", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// Створення товару
func (h Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product internal.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Printf("Помилка створення категорії: %v", err)
		http.Error(w, "Невірний формат запиту", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateProduct(r.Context(), product)
	if err != nil {
		log.Printf("Помилка створення категорії: %v", err)
		http.Error(w, "Помилка створення товару", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// Отримання товару за ID
func (h Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Помилка створення категорії: %v", err)
		http.Error(w, "Невірний формат ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProduct(r.Context(), id)
	if err != nil {
		log.Printf("Помилка створення категорії: %v", err)
		http.Error(w, "Товар не знайдено", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}
