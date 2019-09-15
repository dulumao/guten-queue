package serv

import (
	"context"
	"encoding/json"
	"github.com/gookit/color"
	"github.com/thoas/bokchoy"
	"guten-queue/framework/core/env"
	"guten-queue/framework/core/queue"
	"guten-queue/framework/src/services"
	"os"
	"os/signal"
	"syscall"
)

func Initialize() {
	ctx := context.Background()

	engine, err := bokchoy.New(ctx, bokchoy.Config{
		Broker: bokchoy.BrokerConfig{
			Type: "redis",
			Redis: bokchoy.RedisConfig{
				Type: "client",
				Client: bokchoy.RedisClientConfig{
					Addr:     env.Value.Queue.Redis.Addr,
					Password: env.Value.Queue.Redis.Password,
					DB:       env.Value.Queue.Redis.DbNumber,
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	for _, m := range services.Value.Message {
		switch handlerFunc := m.HandleFunc.(type) {
		case func(*bokchoy.Request) error:
			engine.Queue(m.Name).HandleFunc(handlerFunc, m.Options...)
		case func(interface{}):
			engine.Queue(m.Name).HandleFunc(func(r *bokchoy.Request) error {
				var err error
				var payload []byte

				payload, err = json.Marshal(r.Task.Payload)

				if err != nil {
					return err
				}

				_, err = queue.EnqueueObject(queue.Item{
					Payload: payload,
					Name:    r.Task.Name,
				})

				return err
			}, bokchoy.WithConcurrency(1))
		}
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for c := range sigChan {
			switch c {
			case syscall.SIGINT:
				color.Info.Println("\nShutdown by Ctrl+C")
			case syscall.SIGTERM: // by kill
				color.Info.Println("\nShutdown quickly")
			case syscall.SIGQUIT:
				color.Info.Println("\nShutdown gracefully")
			}

			engine.Stop(ctx)
			queue.Close()
		}
	}()

	go queue.Run()

	engine.Run(ctx)
}
