package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/virtuality0/ecommerce-listing/internal/config"
	"github.com/virtuality0/ecommerce-listing/internal/database"
	"github.com/virtuality0/ecommerce-listing/internal/handlers"
	"github.com/virtuality0/ecommerce-listing/internal/product"
	"github.com/virtuality0/ecommerce-listing/internal/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env file found")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading the config %v", err)
	}

	ctx := context.Background()

	db, err := database.NewPool(ctx, cfg.DB)
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}

	defer db.Close()

	productRepo := product.NewRepository(db)
	h := handlers.NewHandler(productRepo)
	router := routes.NewRouter(h)

	if err := http.ListenAndServe(cfg.Http.Port, router); err != nil {
		log.Fatalf("Server stopped %v", err)
	}
}
