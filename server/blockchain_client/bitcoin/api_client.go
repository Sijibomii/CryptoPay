package bitcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type BlockchainClient struct {
	BSUrl string
}

type BlockCountStatsResponse struct {
	Blocks int `json:"height"`
}

func New(bc_url, bs_url string) *BlockchainClient {
	return &BlockchainClient{
		BSUrl: bs_url,
	}
}

// https://blockstream.info/api/blocks/tip/height

// https://api-r.bitcoinchain.com/v1/status
func (client *BlockchainClient) Block_count_endpoint() string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/blocks/tip/height")})
	return u.String()
}

func (client *BlockchainClient) get_block_count() (int, error) {
	//fmt.Printf("block count message \n")
	// client.Block_count_endpoint()
	response, err := http.Get("https://blockstream.info/testnet/api/blocks/tip/height")
	//
	if err != nil {
		//fmt.Printf("block count message \n #############  %s", err.Error())
		return 0, err
	}
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return 0, err
	}
	//fmt.Println(string(body))
	// Parse response to integer
	value, err := strconv.Atoi(string(body))
	if err != nil {
		return 0, err
	}
	return value, nil
}

type Block struct {
	ID                string `json:"id"`
	Height            int    `json:"height"`
	Version           int    `json:"version"`
	Timestamp         int    `json:"timestamp"`
	TxCount           int    `json:"tx_count"`
	Size              int    `json:"size"`
	Weight            int    `json:"weight"`
	MerkleRoot        string `json:"merkle_root"`
	PreviousBlockHash string `json:"previousblockhash"`
	MedianTime        int    `json:"mediantime"`
	Nonce             int    `json:"nonce"`
	Bits              int    `json:"bits"`
	Difficulty        int    `json:"difficulty"`
}

func (b Block) String() string {
	return fmt.Sprintf("Block: ID=%s, Height=%d, Version=%d, Timestamp=%d, TxCount=%d, Size=%d, Weight=%d, MerkleRoot=%s, PreviousBlock=%s, MedianTime=%d, Nonce=%d, Bits=%d, Difficulty=%d",
		b.ID, b.Height, b.Version, b.Timestamp, b.TxCount, b.Size, b.Weight, b.MerkleRoot, b.PreviousBlockHash, b.MedianTime, b.Nonce, b.Bits, b.Difficulty)
}

func (client *BlockchainClient) Get_Block_endpoint(block_hash string) string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/testnet/api/block/%s", block_hash)})
	//fmt.Printf(u.String())
	return u.String()
}

func (client *BlockchainClient) get_Block(block_hash string) (*Block, error) {

	resp, err := http.Get(client.Get_Block_endpoint(block_hash))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data Block

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (client *BlockchainClient) Get_Block_Hash_with_height_endpoint(block_height int) string {
	baseURL, _ := url.Parse(client.BSUrl)
	//fmt.Printf(baseURL.String())
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/testnet/api/block-height/%s", strconv.Itoa(block_height))})
	//fmt.Printf(u.String())
	return u.String()
}

func (client *BlockchainClient) get_Block_Hash_with_height(block_height int) (string, error) {
	//fmt.Printf("Getting block with height \n")

	response, err := http.Get(client.Get_Block_Hash_with_height_endpoint(block_height))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	value := string(body)

	//fmt.Printf("value: %s \n", value)

	return value, nil
}

func (client *BlockchainClient) Get_Transactions_by_Block_hash_endpoint(block_hash string) string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/testnet/api/block/%s/txids", block_hash)})
	//fmt.Printf(u.String())
	return u.String()
}

func (client *BlockchainClient) get_Transactions_id_by_Block_hash(block_hash string) ([]string, error) {

	response, err := http.Get(client.Get_Transactions_by_Block_hash_endpoint(block_hash))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result []string
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *BlockchainClient) Get_Transaction_by_hash_endpoint(tx_hash string) string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/testnet/api/tx/%s", tx_hash)})
	//fmt.Printf(u.String())
	return u.String()
}

// GET /tx/:txid/outspends

type Transaction struct {
	TxID     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []Vin  `json:"vin"`
	Vout     []Vout `json:"vout"`
	Size     int    `json:"size"`
	Weight   int    `json:"weight"`
	Fee      int    `json:"fee"`
	Status   Status `json:"status"`
}

type Vin struct {
	TxID         string   `json:"txid"`
	Vout         int      `json:"vout"`
	ScriptSig    string   `json:"scriptsig"`
	ScriptSigAsm string   `json:"scriptsig_asm"`
	Witness      []string `json:"witness"`
	IsCoinbase   bool     `json:"is_coinbase"`
	Sequence     int      `json:"sequence"`
}

type Vout struct {
	ScriptPubKey        string `json:"scriptpubkey"`
	ScriptPubKeyAsm     string `json:"scriptpubkey_asm"`
	ScriptPubKeyType    string `json:"scriptpubkey_type"`
	ScriptPubKeyAddress string `json:"scriptpubkey_address"`
	Value               int    `json:"value"`
}

type Status struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight int    `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int    `json:"block_time"`
}

// https://blockstream.info/testnet/api/

func (client *BlockchainClient) get_Transaction_By_Hash(tx_hash string) (*Transaction, error) {

	resp, err := http.Get(client.Get_Transaction_by_hash_endpoint(tx_hash))

	if resp.StatusCode == 404 {
		return &Transaction{
			TxID: "0",
		}, nil
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Printf("\n transaction request body: %s \n", string(body))

	if err != nil {
		return nil, err
	}

	var data Transaction

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GET /blocks/tip/hash hash of the last block

func (client *BlockchainClient) Get_Hash_For_Last_Block_endpoint() string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/blocks/tip/hash")})
	//fmt.Printf(u.String())
	return u.String()
}

func (client *BlockchainClient) get_Hash_For_Last_Block() (string, error) {
	response, err := http.Get(client.Get_Hash_For_Last_Block_endpoint())
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	value := string(body)

	return value, nil
}

// GET /block/:hash/txid/:index

// GET /tx/:txid/status

func (client *BlockchainClient) BroadcastTransaction_endpoint() string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/testnet/api/tx")})
	//fmt.Printf(u.String())
	return u.String()
}

func (client *BlockchainClient) broadcastTransaction(rawTx string) (string, error) {
	url := client.BroadcastTransaction_endpoint()
	// Create a JSON payload containing the raw transaction
	payload := []byte(fmt.Sprintf(`{"tx": "%s"}`, rawTx))

	// Send a POST request to the Esplora API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to broadcast transaction: %s", resp.Status)
	}

	// should return hash

	// Print the response body
	//fmt.Println(string(body))

	return string(body), nil
}

type FeeEstimates struct {
	OneHourFee        float64 `json:"1"`
	ThreeHourFee      float64 `json:"3"`
	SixHourFee        float64 `json:"6"`
	TwelveHourFee     float64 `json:"12"`
	TwentyFourHourFee float64 `json:"24"`
}

func (client *BlockchainClient) GetFeeEstimates_endpoint() string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/testnet/api/fee-estimates")})
	//fmt.Printf(u.String())
	return u.String()
}

func (client *BlockchainClient) getFeeEstimates() (FeeEstimates, error) {
	url := client.GetFeeEstimates_endpoint()

	f := FeeEstimates{}

	//fmt.Printf("\n fee url: %s \n", url)
	resp, err := http.Get(url)
	//fmt.Print("\n resp: ", resp)
	if err != nil {
		return f, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Print("\n body: ", body)
	if err != nil {
		return f, err
	}

	if resp.StatusCode != http.StatusOK {
		return f, fmt.Errorf("failed to get fee estimates: %s", resp.Status)
	}

	var feeEstimates FeeEstimates

	err = json.Unmarshal(body, &feeEstimates)
	//fmt.Print("\n feeEstimates: ", feeEstimates)
	if err != nil {
		return f, err
	}

	return feeEstimates, nil
}

type MempoolEntry struct {
	TxID  string `json:"txid"`
	Fee   int    `json:"fee"`
	VSize int    `json:"vsize"`
	Value int    `json:"value"`
}

func (client *BlockchainClient) GetRawMempool_endpoint() string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/testnet/api/mempool/recent")})
	//fmt.Printf(u.String())
	return u.String()
}

func (client *BlockchainClient) GetRawMempool() ([]MempoolEntry, error) {
	url := client.GetRawMempool_endpoint()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get raw mempool: %s", resp.Status)
	}

	var mempool []MempoolEntry
	err = json.Unmarshal(body, &mempool)
	if err != nil {
		return nil, err
	}

	return mempool, nil
}

func (client *BlockchainClient) get_Block_with_height(block_height int) (*Block, error) {
	hash, err := client.get_Block_Hash_with_height(block_height)

	if err != nil {
		//fmt.Printf("error getting hash")
	}

	block, err := client.get_Block(hash)

	if err != nil {
		//fmt.Printf("error getting block %s", err.Error())
	}

	return block, nil
}

func (client *BlockchainClient) get_all_transactions_by_block_height(block_height int) ([]Transaction, error) {

	hash, err := client.get_Block_Hash_with_height(block_height)

	if err != nil {
		//fmt.Printf("error getting hash")
	}

	txids, err := client.get_Transactions_id_by_Block_hash(hash)

	result := make([]Transaction, 0)

	for _, txid := range txids {
		trans, err := client.get_Transaction_By_Hash(txid)

		if err != nil {
			//fmt.Printf("error getting trans %s \n", txid)
		}
		//fmt.Printf("transaction request complete: ", trans)
		result = append(result, *trans)

		time.Sleep(20 * time.Microsecond)
	}

	//fmt.Printf(" \n transactions: ", result)

	return result, nil
}
