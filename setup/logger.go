package setup

import "go.uber.org/zap"

func setLogger() *zap.Logger{
	logger, _ := zap.NewProduction()

	return logger
}
