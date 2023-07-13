package payouter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type Payouter struct {
	PostgresClieint *actor.PID
	BtcClient       *actor.PID
	Network         string
}

type PreparePayoutResponse struct {
	PrivateKey  *bip32.Key
	PublicKey   *bip32.Key
	Transaction bitcoin.Transaction
	Store       models.Store
	Fee         float64
}

func NewPayouter(network string, pg, btc *actor.PID) *Payouter {
	return &Payouter{
		PostgresClieint: pg,
		BtcClient:       btc,
		Network:         network,
	}
}
func (p *Payouter) preparePayout(e *actor.Engine, dbConn *actor.PID, payout models.Payout) (PreparePayoutResponse, error) {

	store, err := dao.GetStoreById(e, dbConn, payout.Store_id)

	if err != nil {
		fmt.Printf("unable to get store by id in prepare payout")
		panic("unable to get store by id in prepare payout")
	}

	payment, err := dao.GetPaymentById(e, dbConn, payout.Payment_id)

	if err != nil {
		fmt.Printf("unable to get payment by id in prepare payout")
		panic("unable to get payment by id in prepare payout")
	}

	fee, err := bitcoin.GetFeeEstimates(e, p.BtcClient)

	if err != nil {
		fmt.Printf("unable to get fee by id in prepare payout")
		panic("unable to get fee by id in prepare payout")
	}

	if fee.OneHourFee == 0 {
		fmt.Printf("invalid gas fee")
		panic("invalid gas fee")
	}

	trans, err := bitcoin.GetRawTransaction(e, p.BtcClient, payment.Transaction_hash)

	path := store.Hd_path

	path += "/"

	path += strconv.FormatInt(payment.Created_at.Unix(), 10)
	path += "/"

	path += strconv.Itoa(payment.Created_at.Nanosecond() / 1000)

	stores_mnemonic := store.Mnemonic
	seed := bip39.NewSeed(stores_mnemonic, "")

	masterKey, _ := bip32.NewMasterKey(seed)

	childKey, _ := util.NewChildKeyFromString(masterKey, path)

	return PreparePayoutResponse{
		PrivateKey:  childKey,
		PublicKey:   childKey.PublicKey(),
		Transaction: *trans,
		Store:       *store,
		Fee:         fee.OneHourFee,
	}, nil

}

func (p *Payouter) payout(e *actor.Engine, dbConn *actor.PID, payout models.Payout) (string, error) {

	resp, err := p.preparePayout(e, dbConn, payout)

	if err != nil {
		fmt.Printf("unable to prepare payout")
		panic("unable to prepare payout")
	}

	if resp.Store.Btc_payout_user_addresses == "" {
		fmt.Printf("no payout address in payouter")
		panic("no payout address in payouter")
	}

	recipient := resp.PublicKey.String()

	// Btc_payout_user_addresses might be a comma sep value

	addresses := strings.Split(resp.Store.Btc_payout_user_addresses, ",")
	utxo_idx := 0
	for idx, output := range resp.Transaction.Vout {
		// pubkeyhash
		if output.ScriptPubKeyType == "pubkeyhash" {
			if output.ScriptPubKeyAddress == recipient {
				utxo_idx = idx
			}
		} else {
			fmt.Printf("unexpected script type in payouter %s", output.ScriptPubKeyType)
			panic("unexpected script type in payouter")
		}
	}

	utxo := resp.Transaction.Vout[utxo_idx]

	value := (utxo.Value * 100_000_000)

	tx_fee_per_byte := float64(resp.Fee) / 1000

	if float64(value) <= tx_fee_per_byte*192 {
		fmt.Printf("Insufficient funds to pay out.")
		panic("Insufficient funds to pay out.")
	}
	var inputs []bitcoin.TransactionInput
	var outputs []bitcoin.TransactionOutput
	inputs = append(inputs, bitcoin.TransactionInput{
		Transaction: resp.Transaction,
		Idx:         utxo_idx,
	})

	outputs = append(outputs, bitcoin.TransactionOutput{
		Address: addresses[0],
		Amount:  float64(value) - tx_fee_per_byte,
	})

	tx := bitcoin.NewUnsignedTransaction(inputs, outputs)

	tx.Sign(resp.PrivateKey, resp.PublicKey)
	raw_transaction := tx.Into_raw_transaction()

	hash, err := bitcoin.BroadcastRawTransaction(e, p.BtcClient, string(raw_transaction))

	if err != nil {
		fmt.Printf("unable to broadcast transaction in payouter")
		panic("unable to broadcast transaction in payouter")
	}

	return hash, nil
}

type ProcessPayoutMessage struct {
	Payout models.Payout
}

type PayoutMessage struct {
	Payout models.Payout
}

func (client *Payouter) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	case actor.Started:
		fmt.Println("payrouter started")

	case ProcessPayoutMessage:
		client.SendPayoutMessage(ctx.Engine(), client.BtcClient, l.Payout)

	case PayoutMessage:

	default:
		fmt.Println("UNKNOWN MESSAGE TO Payrouter")
	}
}

func (p *Payouter) SendPayoutMessage(e *actor.Engine, conn *actor.PID, payout models.Payout) {
	e.Send(conn, PayoutMessage{
		Payout: payout,
	})
}

func (p *Payouter) DoPayout(e *actor.Engine, conn *actor.PID, payout models.Payout) {
	hash, err := p.payout(e, conn, payout)

	if err != nil {
		fmt.Printf("unable to get payout")
		panic("unable to get payout")
	}

	payout_payload := models.PayoutPayload{}

	payout_payload = payout_payload.FromPayout(payout)

	payout_payload.Transaction_hash = hash
	payout_payload.Status = "paidout"

	payment_payload := models.PaymentPayload{}

	payment_payload.Status = "completed"

	_, err = dao.UpdatePayoutWithPayment(e, conn, payout.ID, payout_payload, payment_payload)

	if err != nil {
		panic("insufficient funds")
	}
	fmt.Printf("payload updated")
}
