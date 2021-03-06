package db

import (
	"github.com/daurensky/kazakh-dream/models"
	"github.com/lib/pq"
)

func ShowProduct(productId int) (models.Product, error) {
	product := models.Product{}

	db, err := Connect()

	if err != nil {
		return product, err
	}

	row := db.QueryRow(`
		SELECT id,
		       name,
		       price,
		       photo_url,
		       composition
		FROM kazakh_dream.public.products
		WHERE id = $1
	`, productId)

	err = row.Scan(&product.Id, &product.Name, &product.Price, &product.PhotoUrl, pq.Array(&product.Composition))

	if err != nil {
		return product, err
	}

	return product, nil
}
