package db

import (
	"github.com/daurensky/kazakh-dream/models"
	"github.com/lib/pq"
)

func UpdateProduct(product models.Product) error {
	db, err := Connect()

	if err != nil {
		return err
	}

	_, err = db.Exec(
		"UPDATE kazakh_dream.public.products SET price = $1, photo_url = $2, composition = $3, name = $4 WHERE id = $5",
		product.Price,
		product.PhotoUrl,
		pq.Array(product.Composition),
		product.Name,
		product.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
