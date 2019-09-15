package test

import (
	"fmt"
	"github.com/dulumao/Guten-utils/conv"
	"github.com/thoas/bokchoy"
)

func Test(r *bokchoy.Request) error {
	// fmt.Println("Receive request", r)
	fmt.Println("Test:", r.Task.Payload)
	//
	// return errors.New("出粗了")
	return nil
}

func TestMutil(payload interface{}) {
	fmt.Println("TestMutil", conv.String(payload))
}

func TestQueue(payload interface{}) {
	fmt.Println("Receive request", payload)
}
