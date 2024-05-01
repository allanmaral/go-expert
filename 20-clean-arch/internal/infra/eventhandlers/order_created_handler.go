package eventhandlers

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/allanmaral/go-expert/20-clean-arch/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandlerAMQP struct {
	channel *amqp.Channel
}

var _ events.EventHandler = (*OrderCreatedHandlerAMQP)(nil)

func NewOrderCreatedHandlerAMQP(channel *amqp.Channel) *OrderCreatedHandlerAMQP {
	return &OrderCreatedHandlerAMQP{channel}
}

func (h *OrderCreatedHandlerAMQP) Handle(event events.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order created: %v\n", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.channel.Publish(
		"amq.direct",
		"",
		false,
		false,
		msg,
	)
}
