package products

import (
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	CategoryID  uint    `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Sku         string  `json:"sku"`
	IsActive    bool    `json:"is_active"`
	Category    struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		IsActive    bool   `json:"is_active"`
	}

	Images []Image `json:"images"`
}

type ProductResponse struct {
	Success string    `json:"success"`
	Message string    `json:"message"`
	Data    []Product `json:"product"`
	Error   string    `json:"error"`
	Meta    struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"meta"`
}

type Image struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	AltText   string    `json:"alt_text"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductDetailResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
	Error   string  `json:"error"`
}
