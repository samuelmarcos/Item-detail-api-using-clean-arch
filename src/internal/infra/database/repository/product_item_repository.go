package repository

import (
	"context"
	"database/sql"
	"desafio_mercado_livre/src/internal/entity"
	"encoding/json"
)

type ProductItemRepository struct {
	Db *sql.DB
}

func NewProductItemRepository(db *sql.DB) *ProductItemRepository {
	return &ProductItemRepository{
		Db: db,
	}
}

func (r *ProductItemRepository) SaveProductDetail(ctx context.Context, product *entity.ProductDetail) error {
	imagesJSON, err := json.Marshal(product.Images)
	if err != nil {
		return err
	}

	sellerInfoJSON, err := json.Marshal(product.Seller)
	if err != nil {
		return err
	}

	paymentOptionsJSON, err := json.Marshal(product.PaymentOptions)
	if err != nil {
		return err
	}

	specsJSON, err := json.Marshal(product.Specs)
	if err != nil {
		return err
	}

	relatedProductsJSON, err := json.Marshal(product.RelatedProducts)
	if err != nil {
		return err
	}

	purchaseOptionsJSON, err := json.Marshal(product.PurchaseOptions)
	if err != nil {
		return err
	}

	highlightsJSON, err := json.Marshal(product.Highlights)
	if err != nil {
		return err
	}

	_, err = r.Db.ExecContext(
		ctx,
		`INSERT OR REPLACE INTO product_details (
			id, name, description, brand, model, color, category,
			images, price, original_price, discount_percent, stock,
			seller_info, warranty, payment_options, specs,
			related_products, rating, review_count, free_shipping,
			purchase_options, highlights, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`,
		product.ID, product.Name, product.Description, product.Brand,
		product.Model, product.Color, product.Category, imagesJSON,
		product.Price, product.OriginalPrice, product.DiscountPercent,
		product.Stock, sellerInfoJSON, product.Warranty, paymentOptionsJSON,
		specsJSON, relatedProductsJSON, product.Rating, product.ReviewCount,
		product.FreeShipping, purchaseOptionsJSON, highlightsJSON,
	)
	return err
}

func (r *ProductItemRepository) GetProductDetail(ctx context.Context, id string) (*entity.ProductDetail, error) {
	var product entity.ProductDetail
	var imagesJSON, sellerInfoJSON, paymentOptionsJSON, specsJSON,
		relatedProductsJSON, purchaseOptionsJSON, highlightsJSON string

	err := r.Db.QueryRowContext(
		ctx,
		`SELECT id, name, description, brand, model, color, category,
			images, price, original_price, discount_percent, stock,
			seller_info, warranty, payment_options, specs,
			related_products, rating, review_count, free_shipping,
			purchase_options, highlights, created_at, updated_at
		FROM product_details WHERE id = ?`,
		id,
	).Scan(
		&product.ID, &product.Name, &product.Description, &product.Brand,
		&product.Model, &product.Color, &product.Category, &imagesJSON,
		&product.Price, &product.OriginalPrice, &product.DiscountPercent,
		&product.Stock, &sellerInfoJSON, &product.Warranty, &paymentOptionsJSON,
		&specsJSON, &relatedProductsJSON, &product.Rating, &product.ReviewCount,
		&product.FreeShipping, &purchaseOptionsJSON, &highlightsJSON,
		&product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(imagesJSON), &product.Images); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(sellerInfoJSON), &product.Seller); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(paymentOptionsJSON), &product.PaymentOptions); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(specsJSON), &product.Specs); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(relatedProductsJSON), &product.RelatedProducts); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(purchaseOptionsJSON), &product.PurchaseOptions); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(highlightsJSON), &product.Highlights); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductItemRepository) GetAllProductDetails(ctx context.Context) ([]*entity.ProductDetail, error) {
	rows, err := r.Db.QueryContext(
		ctx,
		`SELECT id, name, description, brand, model, color, category,
			images, price, original_price, discount_percent, stock,
			seller_info, warranty, payment_options, specs,
			related_products, rating, review_count, free_shipping,
			purchase_options, highlights, created_at, updated_at
		FROM product_details ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.ProductDetail
	for rows.Next() {
		var product entity.ProductDetail
		var imagesJSON, sellerInfoJSON, paymentOptionsJSON, specsJSON,
			relatedProductsJSON, purchaseOptionsJSON, highlightsJSON string

		err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.Model, &product.Color, &product.Category, &imagesJSON,
			&product.Price, &product.OriginalPrice, &product.DiscountPercent,
			&product.Stock, &sellerInfoJSON, &product.Warranty, &paymentOptionsJSON,
			&specsJSON, &relatedProductsJSON, &product.Rating, &product.ReviewCount,
			&product.FreeShipping, &purchaseOptionsJSON, &highlightsJSON,
			&product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(imagesJSON), &product.Images); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(sellerInfoJSON), &product.Seller); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(paymentOptionsJSON), &product.PaymentOptions); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(specsJSON), &product.Specs); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(relatedProductsJSON), &product.RelatedProducts); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(purchaseOptionsJSON), &product.PurchaseOptions); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(highlightsJSON), &product.Highlights); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
