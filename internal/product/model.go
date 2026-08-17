package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Price       float32   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

type ProductListResponse struct {
	Products []Product `json:"products"`
	Count    int       `json:"count"`
}

type ProductUpdateRequestDto struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float32 `json:"price,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
}

type ProductUpdateRequestParam struct {
	ID          uuid.UUID
	Name        *string
	Description *string
	Price       *float32
	Stock       *int
}

type CreateProductRequestDto struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
}

type CreateProductRequestParam struct {
	ID          uuid.UUID
	Name        string
	Description *string
	Price       float32
	Stock       int
}
