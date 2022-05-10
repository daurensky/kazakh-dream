package db

import (
	"github.com/daurensky/kazakh-dream/models"
	"github.com/lib/pq"
)

func StoreProduct(product models.Product) error {
	db, err := Connect()

	if err != nil {
		return err
	}

	_, err = db.Exec(
		"INSERT INTO kazakh_dream.public.products (price, photo_url, composition, name) VALUES ($1, $2, $3, $4)",
		product.Price,
		product.PhotoUrl,
		pq.Array(product.Composition),
		product.Name,
	)

	if err != nil {
		return err
	}

	return nil
}
