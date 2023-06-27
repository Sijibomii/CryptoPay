package bitcoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type BlockchainClient struct {
	Url string
}

type BlockCountStatsResponse struct {
	Data struct {
		Blocks int `json:"blocks"`
	} `json:"data"`
}

func New(url string) *BlockchainClient {
	return &BlockchainClient{
		Url: url,
	}
}

func (client *BlockchainClient) Block_count_endpoint() string {
	baseURL, _ := url.Parse(client.Url)
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/bitcoin/stats")})
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
	numBlocks := stats.Data.Blocks

	return numBlocks, nil
}
