package main

import (
	"guten-queue/framework"
	"guten-queue/framework/bootstrap/serv"
)

func main() {
	framework.Initialize()

	serv.Initialize()
}
