package db

import "github.com/daurensky/kazakh-dream/models"

func UpdateOrder(order models.Order) error {
	db, err := Connect()

	if err != nil {
		return err
	}

	_, err = db.Exec(
		"UPDATE kazakh_dream.public.orders SET status = $1, created_at = $2, client_id = $3 WHERE id = $4",
		order.Status,
		order.CreatedAt,
		order.Client.TelegramId,
		order.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
