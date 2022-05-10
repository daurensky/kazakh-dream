package db

import (
	"github.com/daurensky/kazakh-dream/models"
	"github.com/lib/pq"
)

func ShowOrder(orderId int) (models.Order, error) {
	var order models.Order

	db, err := Connect()

	if err != nil {
		return order, err
	}

	orderFromDB := db.QueryRow(`
		SELECT o.id,
		       o.status,
		       to_char(o.created_at, 'DD.MM.YYYY HH24:MI'),
		       c.telegram_id,
		       c.name,
		       c.phone,
		       c.address
		FROM kazakh_dream.public.orders o
			INNER JOIN clients c on c.telegram_id = o.client_id
		WHERE id = $1
	`, orderId)

	var orderProducts []models.Product

	err = orderFromDB.Scan(
		&order.Id,
		&order.Status,
		&order.CreatedAt,
		&order.Client.TelegramId,
		&order.Client.Name,
		&order.Client.Phone,
		&order.Client.Address,
	)

	if err != nil {
		return order, err
	}

	orderProductsFromDB, err := db.Query(`
		SELECT p.id,
			   SUM(p.price),
			   p.photo_url,
			   p.composition,
			   CASE WHEN COUNT(p.id) > 1 THEN CONCAT(p.name, ' ', COUNT(p.id), ' шт') ELSE p.name END AS name
		FROM kazakh_dream.public.order_product op
			INNER JOIN kazakh_dream.public.products p on p.id = op.product_id
		WHERE op.order_id = $1
		GROUP BY p.id
	`, order.Id)

	if err != nil {
		panic(err)
	}

	for orderProductsFromDB.Next() {
		product := models.Product{}

		err := orderProductsFromDB.Scan(
			&product.Id,
			&product.Price,
			&product.PhotoUrl,
			pq.Array(&product.Composition),
			&product.Name,
		)

		if err != nil {
			panic(err)
		}

		orderProducts = append(orderProducts, product)
	}

	order.Products = orderProducts

	return order, nil
}
