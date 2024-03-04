package messaging

import (
	"encoding/json"
	"log"

	"github.com/adityasuryadi/go-shop/pkg/logger"
)

type UserConsumer struct {
	Log *logger.Logger
}

func NewUserConsumer(log *logger.Logger) *UserConsumer {
	return &UserConsumer{
		Log: log,
	}
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

func (c *UserConsumer) Consume(message []byte) error {
	user := new(User)
	err := json.Unmarshal(message, user)
	if err != nil {
		log.Fatal("failed unmarshal")
		return err
	}
	c.Log.Info("processDeliveries deliveryTag% v", user)
	return nil
}
