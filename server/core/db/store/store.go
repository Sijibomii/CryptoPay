package store

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
	"github.com/jinzhu/gorm"
)

type StoreClient struct {
	Conn      *gorm.DB
	ServerPID *actor.PID
}

func (c StoreClient) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("User db actor started")

	default:
		fmt.Println("UNKNOWN MESSAGE TO USER DB")
	}
}
