package main

import (
	"context"
	"fmt"
	"github.com/thoas/bokchoy"
	"guten-queue/framework/core/env"
	"log"
	"time"
)

func main() {
	env.New()

	ctx := context.Background()

	// define the main engine which will manage queues
	engine, err := bokchoy.New(ctx, bokchoy.Config{
		Broker: bokchoy.BrokerConfig{
			Type: "redis",
			Redis: bokchoy.RedisConfig{
				Type: "client",
				Client: bokchoy.RedisClientConfig{
					// Addr: "localhost:6379",
					Addr:     env.Value.Queue.Redis.Addr,
					Password: env.Value.Queue.Redis.Password,
					DB:       env.Value.Queue.Redis.DbNumber,
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	payload := map[string]string{
		"data": "hello world " + time.Now().Local().Format("2006-01-02 15:04:05"),
	}

	task, err := engine.Queue("tasks.mutil.message").Publish(ctx, payload, bokchoy.WithMaxRetries(1), bokchoy.WithRetryIntervals([]time.Duration{
		1 * time.Second,
		2 * time.Second,
		3 * time.Second,
	}), bokchoy.WithCountdown(5*time.Second))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(task, "has been published")

	task2, err := engine.Queue("tasks.message").Publish(ctx, payload, bokchoy.WithMaxRetries(1), bokchoy.WithRetryIntervals([]time.Duration{
		1 * time.Second,
		2 * time.Second,
		3 * time.Second,
	}))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(task2, "has been published")
}
