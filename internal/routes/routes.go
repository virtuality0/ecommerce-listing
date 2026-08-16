package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/virtuality0/ecommerce-listing/internal/handlers"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	// Global middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// Health check
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// define routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			getProductRoutes(r, h)
		})
	})

	return r
}

func getProductRoutes(r chi.Router, h *handlers.Handler) {
	r.Get("/", h.GetProductList)
	r.Get("/{id}", h.GetProductById)
	r.Post("/", h.CreateProduct)
	r.Patch("/{id}", h.UpdateProduct)
	r.Delete("/{id}", h.DeleteProduct)
}
