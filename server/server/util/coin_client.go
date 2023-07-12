package util

import (
	"errors"
	"time"

	"github.com/anthdm/hollywood/actor"
	coinclient "github.com/sijibomii/cryptopay/coin_client"
)

func GetRate(e *actor.Engine, conn *actor.PID, from, to string) (float64, error) {

	var resp = e.Request(conn, coinclient.GetRateMessage{
		From: from,
		To:   to,
	}, time.Millisecond*2000)

	res, err := resp.Result()

	//fmt.Print("rate received", res)

	//fmt.Printf("\n")

	if err != nil {
		return 0, errors.New("An error occured!")
	}

	var rate float64
	var errorString string

	if floatValue, ok := res.(float64); ok {
		//fmt.Print("converted to float")
		rate = floatValue
		return rate, nil
	} else if stringValue, ok := res.(string); ok {
		errorString = stringValue
		return 0, errors.New(errorString)
	}

	return 0, errors.New("an error occured")
}
