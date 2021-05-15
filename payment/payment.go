package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"payment/queue"
	"time"

	"github.com/streadway/amqp"
)

type Order struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ProductId string    `json:"product_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at,string"`
}

type BoolGen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsuming("order_queue", connection, in)

	r := BoolRand()

	var order Order
	for payload := range in {
		json.Unmarshal(payload, &order)
		if r.Bool() {
			order.Status = "aprovado"
		} else {
			order.Status = "reprovado"
		}
		notifyPaymentProcessed(order, connection)
	}
}

func notifyPaymentProcessed(order Order, ch *amqp.Channel) {
	json, _ := json.Marshal(order)
	queue.Notify(json, "payment_ex", "", ch)
	fmt.Println(string(json))
}

func (b *BoolGen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func BoolRand() *BoolGen {
	return &BoolGen{src: rand.NewSource(time.Now().UnixNano())}
}
