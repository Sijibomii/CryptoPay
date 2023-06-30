package bitcoin

import (
	"fmt"
	"log"
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
	Block bitcoin.Block
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

	for _, transaction := range pooledTransactions {
		for _, output := range transaction.Vout {
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

	payments, err := models.FindAllPendingPaymentsByAddresses(processor.Engine, processor.PostgresClient, processedBlockStream.Addresses, "btc")

	if err != nil {
		fmt.Printf("error...")
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

		case "insufficient":
			// cont
			paid, _ := decimal.NewFromString(paymentPayload.Amount_paid)
			if btcPaid.Cmp(decimal.NewFromFloat(fee).Sub(paid)) > 0 {
				// paid enough
				paymentPayload.Status = "paid"
				paymentPayload.Amount_paid = btcPaid.String()
				paymentPayload.Set_paid_at()
			}
			if btcPaid.Cmp(decimal.NewFromFloat(fee).Sub(paid)) < 0 {
				// get remaining by fee - amount_paid // remaining can still be paid before expiration
				paymentPayload.Status = "insufficient"
				//
				prev, _ := decimal.NewFromString(paymentPayload.Amount_paid)
				paymentPayload.Amount_paid = btcPaid.Add(prev).String()
			}

			if payment.Expires_at.Before(time.Now().UTC()) {
				paymentPayload.Status = "expired"
			}

			if payment.Expires_at.Before(time.Now().UTC()) {
				paymentPayload.Status = "expired"
			}
		}

		models.UpdatePayment(processor.Engine, processor.PostgresClient, paymentPayload.ID, paymentPayload)

	}
}

func (processor *Processor) processBlock(block bitcoin.Block) {
	log.Printf("Processing block: %v\n", *&block.Height)

	// get transactions
	transactions, err := bitcoin.GetAllTransactionsByBlockHeight(processor.Engine, processor.BtcClient)

	if err != nil {
		fmt.Printf("error...")
		panic("error finding all transactions by block height")
	}

	var addresses []string
	txids := make(map[string]string)
	outputs := make(map[string]bitcoin.Vout)

	for _, transaction := range transactions {
		for _, output := range transaction.Vout {
			if output.ScriptPubKeyAddress != "" {
				outputAddresses := output.ScriptPubKeyAddress
				addresses = append(addresses, outputAddresses)

				txids[outputAddresses] = transaction.TxID
				outputs[outputAddresses] = output
			}
		}
	}

	payments, err := models.FindAllPendingPaymentsByAddresses(processor.Engine, processor.PostgresClient, addresses, "btc")

	for _, payment := range payments {
		txid := txids[payment.Address]

		transaction := findTransaction(transactions, txid)
		vout := outputs[payment.Address]
		btcPaid, _ := decimal.NewFromString(fmt.Sprintf("%v", vout.Value))

		block_height_required := block.Height + payment.Confirmations_required - 1

		// insert payout. A payment session can have many payouts...

	}

	// filter where vout has addresses

	// store address, txids, vouts in sep []

	//
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
