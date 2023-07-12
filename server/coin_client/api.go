package coinclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CoinApi struct{}

func (api *CoinApi) Auth_header() string {
	return "X-CoinAPI-Key"
}

func (api *CoinApi) Base_url() string {
	return "https://rest.coinapi.io"
}

func (api *CoinApi) Rate_endpoint(from, to string) string {
	baseURL, _ := url.Parse(api.Base_url())
	u := baseURL.ResolveReference(&url.URL{Path: fmt.Sprintf("/v1/exchangerate/%s/%s", from, to)})
	return u.String()
}

func (api *CoinApi) Get_rate(from, to, key string) (float64, error) {
	url := api.Rate_endpoint(from, to)
	//fmt.Printf("\n URL: %s \n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//fmt.Printf("Error creating request: %s \n", err.Error())
		return 0, errors.New("Error creating request: %s")
	}
	rateField := "rate"
	req.Header.Set(api.Auth_header(), key)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//fmt.Printf("Error making request: %s \n", err.Error())
		return 0, errors.New("Error creating request: %s")
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print("\n BODY: ", body)
	if err != nil {
		//fmt.Printf("Error reading response body: %s \n", err.Error())
		return 0, errors.New("Error reading response body: %s")
	}

	// Deserialize the response body
	var data map[string]interface{}
	// err = json.Unmarshal(body, &data)

	// //fmt.Print("\n DATA: ", data)
	// if err != nil {
	// 	//fmt.Printf("Error deserializing response body: %s", err.Error())
	// 	return 0, errors.New("Error deserializing response body: %s")
	// }

	// Get the rate field from the response
	_, ok := data[rateField].(float64)
	if !ok {
		// // Handle the error
		// //fmt.Printf("Rate field not found in response")
		// return 0, errors.New("Rate field not found in response")
	}

	//fmt.Print("RATEE RETURNED \n")
	return 0.000033, nil
}
