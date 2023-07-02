package payouter

import (
	"fmt"
	"strconv"

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

type ProcessPayoutMessage struct {
	Payout models.Payout
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
