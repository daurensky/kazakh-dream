package main

import (
	json2 "encoding/json"
	"github.com/daurensky/kazakh-dream/db"
	"github.com/daurensky/kazakh-dream/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/api/product", handleGetProducts)
	http.HandleFunc("/api/product/show", handleShowProduct)
	http.HandleFunc("/api/product/store", handleStoreProduct)
	http.HandleFunc("/api/product/update", handleUpdateProduct)

	http.HandleFunc("/api/order", handleGetOrders)
	http.HandleFunc("/api/update-order-status", handleUpdateOrderStatus)

	err := http.ListenAndServe(":8000", http.DefaultServeMux)

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

func handleShowProduct(w http.ResponseWriter, r *http.Request) {
	formProductId := strings.TrimSpace(r.FormValue("id"))

	productId, err := strconv.Atoi(formProductId)

	if err != nil {
		panic(err)
	}

	product, err := db.ShowProduct(productId)

	if err != nil {
		panic(err)
	}

	json, err := json2.Marshal(product)

	if err != nil {
		panic(err)
	}

	_, err = w.Write(json)

	if err != nil {
		panic(err)
	}
}

func handleStoreProduct(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.PostFormValue("name"))
	formPrice := strings.TrimSpace(r.PostFormValue("price"))
	formComposition := strings.TrimSpace(r.PostFormValue("composition"))
	photoUrl := strings.TrimSpace(r.PostFormValue("photo_url"))

	price, err := strconv.ParseFloat(formPrice, 64)

	if err != nil {
		panic(err)
	}

	composition := strings.Split(formComposition, ",")

	for i, _ := range composition {
		composition[i] = strings.TrimSpace(composition[i])
	}

	product := models.Product{
		Price:       price,
		PhotoUrl:    photoUrl,
		Composition: composition,
		Name:        name,
	}

	err = db.StoreProduct(product)

	if err != nil {
		panic(err)
	}
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	formProductId := strings.TrimSpace(r.FormValue("id"))

	productId, err := strconv.Atoi(formProductId)

	if err != nil {
		panic(err)
	}

	product, err := db.ShowProduct(productId)

	if err != nil {
		panic(err)
	}

	name := strings.TrimSpace(r.PostFormValue("name"))
	formPrice := strings.TrimSpace(r.PostFormValue("price"))
	formComposition := strings.TrimSpace(r.PostFormValue("composition"))
	photoUrl := strings.TrimSpace(r.PostFormValue("photo_url"))

	price, err := strconv.ParseFloat(formPrice, 64)

	if err != nil {
		panic(err)
	}

	composition := strings.Split(formComposition, ",")

	for i, _ := range composition {
		composition[i] = strings.TrimSpace(composition[i])
	}

	product.Price = price
	product.PhotoUrl = photoUrl
	product.Composition = composition
	product.Name = name

	err = db.UpdateProduct(product)

	if err != nil {
		panic(err)
	}
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
