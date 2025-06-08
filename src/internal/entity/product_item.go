package entity

import (
	"errors"
	"time"
)

// Erros de validação
var (
	ErrInvalidProductID    = errors.New("invalid product ID")
	ErrInvalidProductName  = errors.New("invalid product name")
	ErrInvalidProductPrice = errors.New("invalid product price")
	ErrInvalidProductStock = errors.New("invalid product stock")
	ErrInvalidSellerName   = errors.New("invalid seller name")
)

// Informações do vendedor
type SellerInfo struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Reputation string `json:"reputation"`
	Sales      int    `json:"sales"`
	Official   bool   `json:"official"`
}

// Características técnicas do produto
type ProductSpecs struct {
	Attributes map[string]string `json:"attributes"`
}

type ProductDetail struct {
	ID              int64        `json:"id"`         // ID interno auto-incrementado
	ProductID       string       `json:"product_id"` // ID externo do produto (ex: MLB123456)
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Brand           string       `json:"brand"`
	Model           string       `json:"model"`
	Color           string       `json:"color"`
	Category        string       `json:"category"`
	Images          []string     `json:"images"`
	Price           float64      `json:"price"`
	OriginalPrice   float64      `json:"original_price"`
	DiscountPercent float64      `json:"discount_percent"`
	Stock           int          `json:"stock"`
	Seller          SellerInfo   `json:"seller"`
	Warranty        string       `json:"warranty"`
	PaymentOptions  []string     `json:"payment_options"` //Lista de opções de pagamento aceitas (ex: cartão de crédito, boleto, parcelamento).
	Specs           ProductSpecs `json:"specs"`
	RelatedProducts []string     `json:"related_products"`
	Rating          float64      `json:"rating"` // Avaliação média do produto (ex: 4.8).
	ReviewCount     int          `json:"review_count"`
	FreeShipping    bool         `json:"free_shipping"`    //FreeShipping
	PurchaseOptions []string     `json:"purchase_options"` //Opções de compra disponíveis (ex: novo, usado).
	Highlights      []string     `json:"highlights"`       //Lista de destaques do produto (ex: mais vendido, promoção).
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}
