package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/virtuality0/ecommerce-listing/internal/product"
)

type Handler struct {
	// registering repositories as dependencies
	productRepo *product.Repository
}

func NewHandler(productRepo *product.Repository) *Handler {
	return &Handler{
		productRepo: productRepo,
	}
}

func (h *Handler) GetProductList(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.GetProducts(r.Context(), 0, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting products : %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid id : %v", id), http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.GetProductById(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting product : %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid id : %v", id), http.StatusBadRequest)
		return
	}

	err = h.productRepo.DeleteProduct(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting product : %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req product.CreateProductRequestDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body : %v", err), http.StatusBadRequest)
		return
	}

	p := product.CreateProductRequestParam{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	err := h.productRepo.Create(r.Context(), p)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating product : %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		ID uuid.UUID `json:"id"`
	}{ID: p.ID})
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req product.ProductUpdateRequestDto

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid id : %v", id), http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body : %v", err), http.StatusBadRequest)
		return
	}

	p := product.ProductUpdateRequestParam{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	updatedProduct, err := h.productRepo.UpdateProductById(r.Context(), p)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating product : %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedProduct)
}
