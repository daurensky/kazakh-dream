package db

import (
	"github.com/daurensky/kazakh-dream/models"
	"github.com/lib/pq"
)

func GetProducts() ([]models.Product, error) {
	var products []models.Product

	db, err := Connect()

	if err != nil {
		return nil, err
	}

	productsFromDB, err := db.Query(`
		SELECT id,
		       name,
		       price,
		       photo_url,
		       composition
		FROM kazakh_dream.public.products
		ORDER BY id
	`)

	for productsFromDB.Next() {
		product := models.Product{}

		err := productsFromDB.Scan(&product.Id, &product.Name, &product.Price, &product.PhotoUrl, pq.Array(&product.Composition))

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
