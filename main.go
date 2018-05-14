package main

import (
	"github.com/visheyra/prometheus-nats-gateway/cmd"
	"go.uber.org/zap"
	"time"
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		return
	}
	defer l.Sync()
	logger := l.Sugar()

	logger.Debugw("Starting PNG (prometheus nats gateway)",
		"start", time.Now().String())
	cmd.Execute()
}
