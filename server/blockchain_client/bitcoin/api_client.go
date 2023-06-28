package bitcoin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type BlockchainClient struct {
	BCUrl string
	BSUrl string
}

type BlockCountStatsResponse struct {
	Blocks int `json:"height"`
}

func New(bc_url, bs_url string) *BlockchainClient {
	return &BlockchainClient{
		BCUrl: bc_url,
		BSUrl: bs_url,
	}
}

// https://api-r.bitcoinchain.com/v1/status
func (client *BlockchainClient) Block_count_endpoint() string {
	baseURL, _ := url.Parse(client.BCUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/status")})
	return u.String()
}

func (client *BlockchainClient) get_block_count() (int, error) {
	req, err := http.NewRequest("GET", client.Block_count_endpoint(), nil)

	if err != nil {
		fmt.Printf("Error creating request: %s \n", err.Error())
		return 0, errors.New("Error creating request: %s")
	}

	cl := http.Client{}
	resp, err := cl.Do(req)

	if err != nil {
		fmt.Printf("Error making request: %s \n", err.Error())
		return 0, errors.New("Error creating request: %s")
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading response body: %s \n", err.Error())
		return 0, errors.New("Error reading response body: %s")
	}

	// Parse the JSON response
	var stats BlockCountStatsResponse
	err = json.Unmarshal(body, &stats)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 0, errors.New("Error reading response body: %s")
	}

	// Extract the number of blocks
	numBlocks := stats.Blocks

	return numBlocks, nil
}

type Block struct {
	Hash       string   `json:"hash"`
	Height     int      `json:"height"`
	PrevBlock  string   `json:"prev_block"`
	NextBlock  string   `json:"next_block"`
	MerkleRoot string   `json:"mrkl_root"`
	Tx         []string `json:"tx"`
	TxCount    int      `json:"tx_count"`
	Reward     string   `json:"reward"`
	Fee        string   `json:"fee"`
	VoutSum    string   `json:"vout_sum"`
	IsMain     bool     `json:"is_main"`
	Version    string   `json:"version"`
	Difficulty string   `json:"difficulty"`
	Size       int      `json:"size"`
	Bits       string   `json:"bits"`
	Nonce      string   `json:"nonce"`
	Time       int      `json:"time"`
}

// https://api-r.bitcoinchain.com/v1/block/0000000000000000000fa2283ae08d7e4b4da949c19a40b68c4d0e1267647250/withTx
func (client *BlockchainClient) Get_Block_endpoint(block_hash string) string {
	baseURL, _ := url.Parse(client.BCUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/blocks/%s/withTx", block_hash)})
	return u.String()
}

func (client *BlockchainClient) Get_Block(block_hash string) (*Block, error) {

	resp, err := http.Get(client.Get_Block_endpoint(block_hash))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data []Block

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		return &data[0], nil
	}

	return nil, errors.New("empty response")
}

// https://api-r.bitcoinchain.com/v1/block/{blockHeight}/withTx
func (client *BlockchainClient) Get_Block_with_height_endpoint(block_height int) string {
	baseURL, _ := url.Parse(client.BCUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/blocks/%s/withTx", strconv.Itoa(block_height))})
	return u.String()
}

func (client *BlockchainClient) Get_Block_with_height(block_height int) (*Block, error) {
	resp, err := http.Get(client.Get_Block_with_height_endpoint(block_height))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data []Block

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		return &data[0], nil
	}

	return nil, errors.New("empty response")
}

func (client *BlockchainClient) Get_Transaction_by_hash_endpoint(tx_hash string) string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/tx/%s/", tx_hash)})
	return u.String()
}

type Transaction struct {
	TxID     string  `json:"txid"`
	Version  int     `json:"version"`
	LockTime int     `json:"locktime"`
	Vin      []Vin   `json:"vin"`
	Vout     []Vout  `json:"vout"`
	Size     int     `json:"size"`
	Weight   int     `json:"weight"`
	Fee      float64 `json:"fee"`
	Status   Status  `json:"status"`
}

type Vin struct {
	TxID       string  `json:"txid"`
	Vout       int     `json:"vout"`
	Prevout    Prevout `json:"prevout"`
	ScriptSig  string  `json:"scriptsig"`
	IsCoinbase bool    `json:"is_coinbase"`
	Sequence   uint32  `json:"sequence"`
}

type Prevout struct {
	ScriptPubKey        string  `json:"scriptpubkey"`
	ScriptPubKeyAsm     string  `json:"scriptpubkey_asm"`
	ScriptPubKeyType    string  `json:"scriptpubkey_type"`
	ScriptPubKeyAddress string  `json:"scriptpubkey_address"`
	Value               float64 `json:"value"`
}

type Vout struct {
	ScriptPubKey        string  `json:"scriptpubkey"`
	ScriptPubKeyAsm     string  `json:"scriptpubkey_asm"`
	ScriptPubKeyType    string  `json:"scriptpubkey_type"`
	ScriptPubKeyAddress string  `json:"scriptpubkey_address"`
	Value               float64 `json:"value"`
}

type Status struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight int    `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int    `json:"block_time"`
}

func (client *BlockchainClient) Get_Transaction_By_Hash_height(tx_hash string) (*Transaction, error) {

	resp, err := http.Get(client.Get_Transaction_by_hash_endpoint(tx_hash))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

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

func (client *BlockchainClient) BroadcastTransaction_endpoint() string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/tx/")})
	return u.String()
}

func (client *BlockchainClient) BroadcastTransaction(rawTx string) error {
	url := client.BroadcastTransaction_endpoint()
	// Create a JSON payload containing the raw transaction
	payload := []byte(fmt.Sprintf(`{"tx": "%s"}`, rawTx))

	// Send a POST request to the Esplora API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to broadcast transaction: %s", resp.Status)
	}

	// Print the response body
	fmt.Println(string(body))

	return nil
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
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/fee-estimates/")})
	return u.String()
}

func (client *BlockchainClient) GetFeeEstimates() (*FeeEstimates, error) {
	url := client.GetFeeEstimates_endpoint()

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
		return nil, fmt.Errorf("failed to get fee estimates: %s", resp.Status)
	}

	var feeEstimates FeeEstimates

	err = json.Unmarshal(body, &feeEstimates)
	if err != nil {
		return nil, err
	}

	return &feeEstimates, nil
}

type MempoolEntry struct {
	TxID  string `json:"txid"`
	Fee   int    `json:"fee"`
	VSize int    `json:"vsize"`
	Value int    `json:"value"`
}

func (client *BlockchainClient) GetRawMempool_endpoint() string {
	baseURL, _ := url.Parse(client.BSUrl)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/mempool/recent/")})
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
