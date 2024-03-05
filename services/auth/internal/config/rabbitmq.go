package config

import (
	"fmt"

	"github.com/adityasuryadi/go-shop/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func NewRabbitMqChannell(viper *viper.Viper, log *logger.Logger) (*amqp.Channel, error) {
	connAddr := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		viper.GetString("rabbitmq.username"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.host"),
		viper.GetInt("rabbitmq.port"))
	conn, err := amqp.Dial(connAddr)
	if err != nil {
		log.Error("failed to connect rabbitmq", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error("failed to create channel", err)
	}
	return ch, nil
}
