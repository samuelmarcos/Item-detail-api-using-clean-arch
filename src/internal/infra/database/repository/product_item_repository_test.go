package repository

import (
	"context"
	"database/sql"
	"desafio_mercado_livre/src/internal/entity"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

const testDBFile = "test_product.db"

func setupTestDB(t *testing.T) (*sql.DB, func()) {
	// Remove o arquivo do banco de dados de teste se ele existir
	os.Remove(testDBFile)

	// Abre conexão com o banco de dados de teste
	db, err := sql.Open("sqlite", testDBFile)
	require.NoError(t, err)

	// Cria a tabela de produtos
	_, err = db.ExecContext(
		context.Background(),
		`DROP TABLE IF EXISTS product_details;
		 CREATE TABLE product_details (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_id TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			description TEXT,
			brand TEXT,
			model TEXT,
			color TEXT,
			category TEXT,
			images TEXT, -- JSON array of strings
			price REAL,
			original_price REAL,
			discount_percent REAL,
			stock INTEGER,
			seller_info TEXT, -- JSON object
			warranty TEXT,
			payment_options TEXT, -- JSON array of strings
			specs TEXT, -- JSON object
			related_products TEXT, -- JSON array of strings
			rating REAL,
			review_count INTEGER,
			free_shipping BOOLEAN,
			purchase_options TEXT, -- JSON array of strings
			highlights TEXT, -- JSON array of strings
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
	)
	require.NoError(t, err)

	// Retorna a conexão e uma função de cleanup
	return db, func() {
		db.Close()
		os.Remove(testDBFile)
	}
}

func createTestProduct() *entity.ProductDetail {
	now := time.Now()
	return &entity.ProductDetail{
		ProductID:       "MLB123456",
		Name:            "Test Product",
		Description:     "Test Description",
		Brand:           "Test Brand",
		Model:           "Test Model",
		Color:           "Test Color",
		Category:        "Test Category",
		Images:          []string{"http://example.com/image1.jpg"},
		Price:           100.00,
		OriginalPrice:   120.00,
		DiscountPercent: 16.67,
		Stock:           10,
		Seller: entity.SellerInfo{
			Name:       "Test Seller",
			Type:       "Official",
			Reputation: "Good",
			Sales:      100,
			Official:   true,
		},
		Warranty:       "1 year",
		PaymentOptions: []string{"Credit Card", "Boleto"},
		Specs: entity.ProductSpecs{
			Attributes: map[string]string{
				"Size":  "Medium",
				"Color": "Blue",
			},
		},
		RelatedProducts: []string{"MLB789012"},
		Rating:          4.5,
		ReviewCount:     100,
		FreeShipping:    true,
		PurchaseOptions: []string{"New"},
		Highlights:      []string{"Best Seller"},
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func TestSaveProductDetail(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewProductItemRepository(db)
	ctx := context.Background()

	t.Run("should save product successfully", func(t *testing.T) {
		product := createTestProduct()

		err := repo.SaveProductDetail(ctx, product)
		require.NoError(t, err)

		savedProduct, err := repo.GetProductDetail(ctx, product.ProductID)
		require.NoError(t, err)
		assert.Equal(t, product.ProductID, savedProduct.ProductID)
		assert.Equal(t, product.Name, savedProduct.Name)
		assert.Equal(t, product.Price, savedProduct.Price)
		assert.Equal(t, product.Stock, savedProduct.Stock)
		assert.Equal(t, product.Seller.Name, savedProduct.Seller.Name)
	})

	t.Run("should update existing product", func(t *testing.T) {
		product := createTestProduct()
		product.Name = "Updated Name"
		product.Price = 150.00

		err := repo.SaveProductDetail(ctx, product)
		require.NoError(t, err)
		savedProduct, err := repo.GetProductDetail(ctx, product.ProductID)
		require.NoError(t, err)
		assert.Equal(t, "Updated Name", savedProduct.Name)
		assert.Equal(t, 150.00, savedProduct.Price)
	})
}

func TestGetProductDetail(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewProductItemRepository(db)
	ctx := context.Background()
	product := createTestProduct()
	err := repo.SaveProductDetail(ctx, product)
	require.NoError(t, err)

	t.Run("should get existing product", func(t *testing.T) {
		savedProduct, err := repo.GetProductDetail(ctx, product.ProductID)
		require.NoError(t, err)
		assert.Equal(t, product.ProductID, savedProduct.ProductID)
		assert.Equal(t, product.Name, savedProduct.Name)
		assert.Equal(t, product.Price, savedProduct.Price)
		assert.Equal(t, product.Stock, savedProduct.Stock)
		assert.Equal(t, product.Seller.Name, savedProduct.Seller.Name)
		assert.Equal(t, product.Specs.Attributes["Size"], savedProduct.Specs.Attributes["Size"])
	})

	t.Run("should return error for non-existent product", func(t *testing.T) {
		_, err := repo.GetProductDetail(ctx, "NONEXISTENT")
		assert.Error(t, err)
	})
}

func TestGetAllProductDetails(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewProductItemRepository(db)
	ctx := context.Background()

	product1 := createTestProduct()
	product1.ProductID = "MLB123456"
	product1.Name = "Product 1"

	product2 := createTestProduct()
	product2.ProductID = "MLB789012"
	product2.Name = "Product 2"

	err := repo.SaveProductDetail(ctx, product1)
	require.NoError(t, err)
	err = repo.SaveProductDetail(ctx, product2)
	require.NoError(t, err)

	t.Run("should get all products", func(t *testing.T) {
		products, err := repo.GetAllProductDetails(ctx)
		require.NoError(t, err)
		assert.Len(t, products, 2)
		assert.Equal(t, "MLB789012", products[0].ProductID)
		assert.Equal(t, "MLB123456", products[1].ProductID)
	})

	t.Run("should return empty slice when no products exist", func(t *testing.T) {
		_, err := db.ExecContext(ctx, "DELETE FROM product_details")
		require.NoError(t, err)

		products, err := repo.GetAllProductDetails(ctx)
		require.NoError(t, err)
		assert.Empty(t, products)
	})
}
