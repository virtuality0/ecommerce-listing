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
	Stock       float32   `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

type ProductListResponse struct {
	Products []Product `json:"products"`
	Count    int       `json:"count"`
}

type ProductUpdateRequest struct {
	ID          uuid.UUID  `json:"id"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Price       *float32   `json:"price,omitempty"`
	Stock       *float32   `json:"stock,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ModifiedAt  *time.Time `json:"modified_at,omitempty"`
}
