package message

import (
	"github.com/thoas/bokchoy"
)

type HandlerFunc func(v interface{})

type Message struct {
	Name       string
	HandleFunc interface{}
	Options    []bokchoy.Option
}

func Direct(name string, handleFunc interface{}, options ...bokchoy.Option) *Message {
	return &Message{
		Name:       name,
		HandleFunc: handleFunc,
		Options:    options,
	}
}
