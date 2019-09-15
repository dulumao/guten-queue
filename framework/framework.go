package framework

import (
	"guten-queue/framework/core/env"
	"guten-queue/framework/core/queue"
)

func Initialize() {
	env.New()
	queue.New()
}
