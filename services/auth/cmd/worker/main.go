package main

import (
	"github.com/adityasuryadi/go-shop/pkg/logger"
)

func main() {
	// viper := config.NewViper()
	logger := logger.NewLogger()
	logger.Info("Starting worker...")
}
