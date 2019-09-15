package services

import (
	"github.com/thoas/bokchoy"
	"guten-queue/framework/core/message"
	"guten-queue/framework/src/services/test"
)

type services struct {
	Message []*message.Message
}

var Value *services

func Register() {
	Value.Message = append(Value.Message, message.Direct("tasks.message", test.Test, bokchoy.WithConcurrency(1)))
	Value.Message = append(Value.Message, message.Direct("tasks.mutil.message", test.TestMutil))
}

func init() {
	Value = new(services)

	Register()
}
