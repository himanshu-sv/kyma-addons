package logger

import (
	"go.uber.org/zap"
	"log"
)

var Logger *zap.SugaredLogger

func Initialize() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal(err)
	}
	Logger = logger.Sugar()
}
