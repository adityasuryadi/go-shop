package messaging

import (
	"github.com/adityasuryadi/go-shop/pkg/logger"
	"github.com/adityasuryadi/go-shop/services/user/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type UserProducer struct {
	Producer[*model.CreateUserEvent]
}

func NewUserProducer(ch *amqp.Channel, pubConfig *ProducerConfig, log *logger.Logger) *UserProducer {
	return &UserProducer{
		Producer: Producer[*model.CreateUserEvent]{
			channell: ch,
			pCfg:     pubConfig,
			Log:      log,
		},
	}
}
