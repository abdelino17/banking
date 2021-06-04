package main

import (
	"github.com/abdelino17/banking/app"
	"github.com/abdelino17/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
