package main

import (
	json2 "encoding/json"
	"github.com/daurensky/kazakh-dream/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/product", handleGetProducts)
	http.HandleFunc("/product/store", handleStoreProduct)
	http.HandleFunc("/product/update", handleUpdateProduct)

	http.HandleFunc("/order", handleGetOrders)
	http.HandleFunc("/update-order-status", handleUpdateOrderStatus)

	err := http.ListenAndServe("localhost:8080", http.DefaultServeMux)

	if err != nil {
		panic(err)
	}
}

func handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetProducts()

	if err != nil {
		panic(err)
	}

	json, err := json2.Marshal(products)

	if err != nil {
		panic(err)
	}

	_, err = w.Write(json)

	if err != nil {
		panic(err)
	}
}

func handleStoreProduct(w http.ResponseWriter, r *http.Request) {
	//name := strings.TrimSpace(r.PostFormValue("name"))
	//price := strings.TrimSpace(r.PostFormValue("price"))
	//composition := strings.TrimSpace(r.PostFormValue("composition"))
	//
	//photo, handler, err := r.FormFile("photo")

	//if err != nil {
	//	panic(err)
	//}
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {

}

func handleGetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := db.GetOrders()

	if err != nil {
		panic(err)
	}

	json, err := json2.Marshal(orders)

	if err != nil {
		panic(err)
	}

	_, err = w.Write(json)

	if err != nil {
		panic(err)
	}
}

func handleUpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderIdForm := r.PostFormValue("order_id")

	orderId, err := strconv.Atoi(orderIdForm)

	if err != nil {
		panic(err)
	}

	order, err := db.ShowOrder(orderId)

	status := r.PostFormValue("status")

	order.Status = status

	err = db.UpdateOrder(order)

	if err != nil {
		panic(err)
	}

	msg := tgbotapi.NewMessage(order.Client.TelegramId, "Новый статус вашего заказа: "+order.StatusText())

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))

	if err != nil {
		panic(err)
	}

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
