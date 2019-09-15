package queue

import (
	"guten-queue/framework/core/env"
	"guten-queue/framework/src/helpers/goque"
	"guten-queue/framework/src/services"
	"time"
)

var Queue *goque.Queue

func New() {
	var err error
	Queue, err = goque.OpenQueue(env.Value.Queue.Dir)

	if err != nil {
		panic(err)
	}
}

func Enqueue(v []byte) (*goque.Item, error) {
	item, err := Queue.Enqueue(v)

	if err != nil {
		panic(err)
	}

	return item, nil
}

func EnqueueString(v string) (*goque.Item, error) {
	item, err := Queue.EnqueueString(v)

	if err != nil {
		panic(err)
	}

	return item, nil
}

func EnqueueObject(v interface{}) (*goque.Item, error) {
	item, err := Queue.EnqueueObject(v)

	if err != nil {
		panic(err)
	}

	return item, nil
}

func Dequeue() (*goque.Item, error) {
	item, err := Queue.Dequeue()

	if err != nil {
		panic(err)
	}

	return item, nil
}

func IsEmpty() bool {
	if _, err := Queue.Peek(); err != nil {
		return false
	}

	return true
}

func Run() {
	for {
		if _, err := Queue.Peek(); err != nil {
			time.Sleep(time.Duration(env.Value.Queue.NoQueueSleepMinutes) * time.Minute)
			continue
		}

		var item Item

		v, _ := Queue.Dequeue()
		v.ToObject(&item)

		for _, m := range services.Value.Message {

			if m.Name == item.Name {
				handlerFunc, ok := m.HandleFunc.(func(interface{}))

				if ok {
					handlerFunc(item.Payload)
					// time.Sleep(1 * time.Minute)
				}
			}
		}

		// dump.DD2(item)
	}
}

func Close() {
	Queue.Close()
}
