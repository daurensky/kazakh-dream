package models

type Order struct {
	Id        int
	Status    string
	CreatedAt string
	Products  []Product
	Client    Client
}

func (o Order) StatusText() string {
	switch o.Status {
	case "PREPARING":
		return "Готовится"
	case "SENT":
		return "Отправлено курьером"
	case "DELIVERED":
		return "Доставлено"
	default:
		return "Неизвестный статус заказа"
	}
}
