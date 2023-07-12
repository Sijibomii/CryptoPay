package currency

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Crypto int

const (
	Btc Crypto = iota
	Eth
)

func (c Crypto) String() string {
	switch c {
	case Btc:
		return "btc"
	case Eth:
		return "eth"
	default:
		return ""
	}
}

func (c *Crypto) Scan(value interface{}) error {
	switch value := value.(type) {
	case []byte:
		str := string(value)
		switch str {
		case "btc":
			*c = Btc
		case "eth":
			*c = Eth
		default:
			return errors.New("invalid value for crypto")
		}
	case nil:
		return errors.New("crypto scan: value is nil")
	default:
		return fmt.Errorf("crypto scan: unsupported value type: %T", value)
	}
	return nil
}

func (c Crypto) Value() (driver.Value, error) {
	return c.String(), nil
}

func (c *Crypto) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	switch str {
	case "btc":
		*c = Btc
	case "eth":
		*c = Eth
	default:
		return errors.New("invalid value for crypto")
	}

	return nil
}

func (c Crypto) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func main() {
	// Example usage
	var crypto Crypto
	err := crypto.Scan([]byte("btc"))
	if err != nil {
		//fmt.Println("Error:", err)
		return
	}

	//fmt.Println("Crypto:", crypto)
}
