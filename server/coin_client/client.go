package coinclient

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type CoinClient struct {
	Key string
}

type GetRateMessage struct {
	From string
	To   string
}

func (c *CoinClient) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	case GetRateMessage:
		key := c.Key
		payload, err := getRate(key, l.From, l.To)

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload)

	default:
		fmt.Println("UNKNOWN MESSAGE TO COIN CLIENT")
	}
}

func getRate(key, from, to string) (float64, error) {
	api := CoinApi{}
	return api.Get_rate(from, to, key)
}
