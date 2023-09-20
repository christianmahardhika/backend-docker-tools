package main

import (
	"errors"
	"math/rand"
	"time"

	"golang-boilerplate/internal/config"
	logger "golang-boilerplate/internal/container/infrastructure/log"

	"go.uber.org/zap"
)

func main() {
	// init config
	cfg := config.InitConfig()
	
	log := logger.Initialize(cfg)
	count := 0
	for {

		if rand.Float32() > 0.8 {
			log.Error("oops...something is wrong",
				zap.Int("count", count),
				zap.Error(errors.New("error details")))
		} else {
			log.Info("everything is fine",
				zap.Int("count", count))
		}
		count++
		time.Sleep(time.Second * 2)
	}
}
