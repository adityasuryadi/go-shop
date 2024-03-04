package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer[T any] struct {
	channell *amqp.Channel
	pCfg     *ProducerConfig
	Log      *logger.Logger
}

type ProducerConfig struct {
	Exchange    string
	QueueName   string
	RoutingKey  string
	ConsumerTag string
}

func (p *Producer[T]) SetupExchangeAndQueuePublisher() {
	fmt.Println("declare exchange")
	ch := p.channell
	err := ch.ExchangeDeclare(
		p.pCfg.Exchange, // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments)
	)
	p.Log.Error("Failed to declare an exchange", err)

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	q, err := ch.QueueDeclare(
		p.pCfg.QueueName, // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	p.Log.Error("Failed to declare a queue", err)

	err = ch.QueueBind(q.Name, p.pCfg.RoutingKey, p.pCfg.Exchange, false, nil)
	p.Log.Error("Failed to declare a queue", err)
}

func (p *Producer[T]) CloseChannel() {
	if err := p.channell.Close(); err != nil {
		p.Log.Error("Failed to close channel", err)
	}
}

func (p *Producer[T]) Publish(event T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	value, err := json.Marshal(event)
	if err != nil {
		p.Log.Error("failed to marshal event", err)
	}
	defer cancel()
	p.channell.PublishWithContext(ctx, p.pCfg.Exchange, p.pCfg.RoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        value,
	})
	log.Printf(" [x] Sent %s", value)
}
