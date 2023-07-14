package bitcoin

import (
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/shopspring/decimal"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
	"github.com/sijibomii/cryptopay/core/models"
)

type Processor struct {
	Network        string
	PostgresClient *actor.PID
	Engine         *actor.Engine
	BtcClient      *actor.PID
}

type ProcessBlockMessage struct {
	BlockString string
}

type ProcessMempoolTransactionsMessage struct {
	Transactions []bitcoin.Transaction
}

type ProcessedBlockStream struct {
	Addresses []string
	Txids     map[string]string
	Outputs   map[string]bitcoin.Vout
	err       error
}

func (processor *Processor) processMempoolTransactions(pooledTransactions []bitcoin.Transaction) {
	var addresses []string
	txids := make(map[string]string)
	outputs := make(map[string]bitcoin.Vout)

	// length := len(pooledTransactions)
	//fmt.Println("\n TRANSACTION LENGTH: ", length)
	for _, transaction := range pooledTransactions {
		//fmt.Printf("TRANSACTION %s IS BEEN PROCCESSED \n", transaction.TxID)
		for _, output := range transaction.Vout {
			//fmt.Printf("OUTPUT %s from TRANS %s IS BEEN PROCCESSED \n", output.ScriptPubKey, transaction.TxID)
			if output.ScriptPubKeyAddress != "" {
				outputAddresses := output.ScriptPubKeyAddress
				addresses = append(addresses, outputAddresses)

				txids[outputAddresses] = transaction.TxID
				outputs[outputAddresses] = output
			}
		}
	}

	processedBlockStream := ProcessedBlockStream{
		Addresses: addresses,
		Txids:     txids,
		Outputs:   outputs,
	}

	payments, err := models.FindAllPaymentsByAddresses(processor.Engine, processor.PostgresClient, processedBlockStream.Addresses, "btc")

	if err != nil {
		//fmt.Printf("\n error... %s \n", err.Error())
		panic("error finding pending payments in processor")
	}

	for _, payment := range payments {
		txid := processedBlockStream.Txids[payment.Address]
		transaction := findTransaction(pooledTransactions, txid)

		if transaction == nil {
			continue
		}

		var paymentPayload models.PaymentPayload

		paymentPayload.FromPayment(payment)
		paymentPayload.Transaction_hash = transaction.TxID

		vout := processedBlockStream.Outputs[payment.Address]
		btcPaid, _ := decimal.NewFromString(fmt.Sprintf("%v", vout.Value))

		fee := payment.TotalFee

		switch payment.Status {
		case "pending":
			if btcPaid.Cmp(decimal.NewFromFloat(fee)) >= 0 {
				// paid enough
				paymentPayload.Status = "paid"
				paymentPayload.Amount_paid = btcPaid.String()
				paymentPayload.Set_paid_at()
			}
			if btcPaid.Cmp(decimal.NewFromFloat(fee)) < 0 {
				// get remaining by fee - amount_paid // remaining can still be paid before expiration
				paymentPayload.Status = "insufficient"
				paymentPayload.Amount_paid = btcPaid.String()
			}

			if payment.Expires_at.Before(time.Now().UTC()) {
				paymentPayload.Status = "expired"
			}

		default:
			//fmt.Printf("PAYMENT STATUS OF %s NOT RECOGNIZED \n", payment.Status)
		}

		models.UpdatePayment(processor.Engine, processor.PostgresClient, paymentPayload.ID, paymentPayload)

	}
}

// transactions will be first processed in the mempool and will be marked as paid but will be eventually confirmed when the mempool becomes a block
// a payout is created for a valid (i.e confirmed) payment. bc that's when we can guarantee that the money got to us
func (processor *Processor) processBlock(block bitcoin.Block) {
	// //fmt.Printf("\n Processing block: %v \n", *&block)

	// get transactions
	// looks like the rquest returned is too large to be returned by the hollywood actor. It get's the payload but the context deadline keeps expiring before
	// it responds to the request
	// transactions, err := bitcoin.GetAllTransactionsByBlockHeight(processor.Engine, processor.BtcClient, block.Height)

	// code below is an anomaly (tempirary work around)

	//create a new client

	client := bitcoin.BlockchainClient{
		BSUrl: "https://blockstream.info/testnet/api",
	}

	transactions, err := client.Get_all_transactions_by_block_height(block.Height)

	if err != nil {
		//fmt.Printf("error... %s \n", err.Error())
		panic("error finding all transactions by block height")
	}

	var addresses []string
	txids := make(map[string]string)
	outputs := make(map[string]bitcoin.Vout)

	for _, transaction := range transactions {
		for _, output := range transaction.Vout {
			//fmt.Print("\n output: ", output)
			//fmt.Print("\n address: ", output.ScriptPubKeyAddress)
			if output.ScriptPubKeyAddress != "" {
				outputAddresses := output.ScriptPubKeyAddress

				addresses = append(addresses, outputAddresses)

				txids[outputAddresses] = transaction.TxID
				outputs[outputAddresses] = output
			}
		}
	}

	//fmt.Print("\n find all payments by addresses: ", addresses)

	payments, err := models.FindAllPaymentsByAddresses(processor.Engine, processor.PostgresClient, addresses, "btc")

	for _, payment := range payments {
		txid := txids[payment.Address]

		transaction := findTransaction(transactions, txid)
		vout := outputs[payment.Address]

		btcPaid, _ := decimal.NewFromString(fmt.Sprintf("%v", vout.Value))

		block_height_required := block.Height + payment.Confirmations_required - 1

		// insert payout. A payment session can have many payouts...

		paymentPayload := models.PaymentPayload{}

		paymentPayload = paymentPayload.FromPayment(payment)

		paymentPayload.Transaction_hash = transaction.TxID
		paymentPayload.Block_height_required = block_height_required
		paymentPayload.Set_paid_at()
		paymentPayload.Amount_paid = btcPaid.String()

		var payoutAction string

		if payment.Status == "paid" || payment.Status == "pending" {
			if btcPaid.Cmp(decimal.NewFromFloat(payment.TotalFee)) >= 0 {
				// paid enough
				paymentPayload.Status = "confirmed"
				paymentPayload.Amount_paid = btcPaid.String()
				paymentPayload.Set_paid_at()
				payoutAction = "payout"
			} else {
				payoutAction = "refund"
				paymentPayload.Status = "insufficient"
			}
		}

		// update payment
		models.UpdatePayment(processor.Engine, processor.PostgresClient, paymentPayload.ID, paymentPayload)

		models.InsertPayout(processor.Engine, processor.PostgresClient, models.PayoutPayload{
			Status:                "pending",
			Store_id:              payment.Store_id,
			Payment_id:            payment.ID,
			Type:                  "btc",
			Block_height_required: block_height_required,
			Transaction_hash:      transaction.TxID,
			Action:                payoutAction,
		})
		// transaction
		models.InsertTransaction(processor.Engine, processor.PostgresClient, models.BtcTransactionPayload{
			Hash:        transaction.TxID,
			Transaction: *transaction,
		})
	}

	// update blockchain status
	models.UpdateBtcBlockChainStatusByNetwork(processor.Engine, processor.PostgresClient, "testnet", block.Height)
}

// helper func
func findTransaction(transactions []bitcoin.Transaction, txid string) *bitcoin.Transaction {
	for _, tx := range transactions {
		if tx.TxID == txid {
			return &tx
		}
	}
	return nil
}

// recive
func (processor *Processor) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	case actor.Started:
		//fmt.Println("processor actor started")

	case ProcessBlockMessage:
		//fmt.Print("process block message received! \n")

		//fmt.Print("block string recived ", l.BlockString)
		//fmt.Println("")
		block, err := parseBlockString(l.BlockString)

		if err != nil {
			//fmt.Println("/n not valid block")
		}
		processor.processBlock(block)

	case ProcessMempoolTransactionsMessage:
		//fmt.Print("process mempool transactions message received! \n")
		processor.processMempoolTransactions(l.Transactions)

	default:
		//fmt.Println("UNKNOWN MESSAGE TO PROCESSOR CLIENT")
	}
}
