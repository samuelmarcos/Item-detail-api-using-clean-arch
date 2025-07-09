package entity

import (
	"errors"
	"time"
)

var (
	ErrInvalidProductID          = errors.New("invalid product ID")
	ErrInvalidProductName        = errors.New("invalid product name")
	ErrInvalidProductDescription = errors.New("invalid product description")
	ErrInvalidProductBrand       = errors.New("invalid product brand")
	ErrInvalidProductModel       = errors.New("invalid product model")
	ErrInvalidProductColor       = errors.New("invalid product color")
	ErrInvalidProductCategory    = errors.New("invalid product category")
	ErrInvalidProductImages      = errors.New("invalid product images")
	ErrInvalidProductPrice       = errors.New("invalid product price")
	ErrInvalidProductStock       = errors.New("invalid product stock")
	ErrInvalidSellerName         = errors.New("invalid seller name")
	ErrInvalidSellerType         = errors.New("invalid seller type")
	ErrInvalidSellerReputation   = errors.New("invalid seller reputation")
	ErrInvalidSellerSales        = errors.New("invalid seller sales")
	ErrInvalidWarranty           = errors.New("invalid warranty")
	ErrInvalidPaymentOptions     = errors.New("invalid payment options")
	ErrInvalidSpecs              = errors.New("invalid product specs")
	ErrInvalidRelatedProducts    = errors.New("invalid related products")
	ErrInvalidRating             = errors.New("invalid rating")
	ErrInvalidReviewCount        = errors.New("invalid review count")
	ErrInvalidPurchaseOptions    = errors.New("invalid purchase options")
	ErrInvalidHighlights         = errors.New("invalid highlights")
)

// Informações do vendedor
type SellerInfo struct {
	Name       string `json:"name" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Reputation string `json:"reputation" validate:"required"`
	Sales      int    `json:"sales" validate:"required,min=0"`
	Official   bool   `json:"official" validate:"required"`
}

// Características técnicas do produto
type ProductSpecs struct {
	Attributes map[string]string `json:"attributes" validate:"required,min=1"`
}

type ProductDetail struct {
	ID              int64        `json:"id"`                             // ID interno auto-incrementado
	ProductID       string       `json:"product_id" validate:"required"` // ID externo do produto (ex: MLB123456)
	Name            string       `json:"name" validate:"required"`
	Description     string       `json:"description" validate:"required"`
	Brand           string       `json:"brand" validate:"required"`
	Model           string       `json:"model" validate:"required"`
	Color           string       `json:"color" validate:"required"`
	Category        string       `json:"category" validate:"required"`
	Images          []string     `json:"images" validate:"required,min=1"`
	Price           float64      `json:"price" validate:"required,gt=0"`
	OriginalPrice   float64      `json:"original_price" validate:"required,gt=0"`
	DiscountPercent float64      `json:"discount_percent" validate:"required,gte=0,lte=100"`
	Stock           int          `json:"stock" validate:"required,gte=0"`
	Seller          SellerInfo   `json:"seller" validate:"required"`
	Warranty        string       `json:"warranty" validate:"required"`
	PaymentOptions  []string     `json:"payment_options" validate:"required,min=1"` // Lista de opções de pagamento aceitas
	Specs           ProductSpecs `json:"specs" validate:"required"`
	RelatedProducts []string     `json:"related_products" validate:"required"`
	Rating          float64      `json:"rating" validate:"required,gte=0,lte=5"` // Avaliação média do produto
	ReviewCount     int          `json:"review_count" validate:"required,gte=0"`
	FreeShipping    bool         `json:"free_shipping" validate:"required"`
	PurchaseOptions []string     `json:"purchase_options" validate:"required,min=1"` // Opções de compra disponíveis
	Highlights      []string     `json:"highlights" validate:"required,min=1"`       // Lista de destaques do produto
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

func (p *ProductDetail) SetDiscount(discount float64) {
	newPrice := p.OriginalPrice * (1 - discount/100)
	p.DiscountPercent = discount
	p.Price = newPrice
}
