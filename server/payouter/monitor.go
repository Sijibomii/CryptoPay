package payouter

import (
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/core/models"
)

type Monitor struct {
	Payouter       *actor.PID
	Network        string
	PostgresClient *actor.PID
	PreviousBlock  int
}

func NewMonitor(network string, previosu_block int, Payouter, postgresClient *actor.PID) *Monitor {
	return &Monitor{
		Payouter:       Payouter,
		PostgresClient: postgresClient,
		PreviousBlock:  previosu_block,
		Network:        network,
	}
}

type MonitorMessage struct{}

// actor.started message
func started(e *actor.Engine, conn *actor.PID) {
	//fmt.Printf("repeated message will be sent to monitor..")
	e.SendRepeat(conn, MonitorMessage{}, time.Millisecond*10000)
}

func (m *Monitor) doMonitor(e *actor.Engine, conn *actor.PID) {
	status, err := models.FindBtcBlockChainStatusByNetwork(e, m.PostgresClient, m.Network)
	if err != nil {
		//fmt.Printf("monitor error %s", err.Error())
	}
	if status.Block_Height == m.PreviousBlock {
		//fmt.Printf("ok from monitor")
	} else {

		// why not just call the func?
		var resp = e.Request(conn, ProcessBlockMessage{
			Number:  status.Block_Height,
			Network: m.Network,
		}, time.Millisecond*500)

		_, err := resp.Result()
		if err != nil {
			//fmt.Printf("monitor error %s", err.Error())
			panic("monitor error")
		}
	}

}

type ProcessBlockMessage struct {
	Number  int
	Network string
}

func (m *Monitor) processBlock(e *actor.Engine, block_number int, network string) bool {

	// find all confirmed payouts
	payouts, err := models.FindAllConfirmedPayout(e, m.PostgresClient, block_number, network)
	if err != nil {
		//fmt.Printf("monitor error %s", err.Error())
	}

	// send processPayout for all payouts to the payrouter
	for _, payout := range payouts {
		e.Send(m.Payouter, ProcessPayoutMessage{
			Payout: payout,
		})
	}
	return true
}

func (monitor *Monitor) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	case actor.Started:
		started(ctx.Engine(), ctx.PID())

	case MonitorMessage:
		monitor.doMonitor(ctx.Engine(), ctx.PID())

	case ProcessBlockMessage:
		payload := monitor.processBlock(ctx.Engine(), l.Number, l.Network)

		ctx.Respond(payload)
	default:
		//fmt.Println("UNKNOWN MESSAGE TO Monitor CLIENT")
	}
}
