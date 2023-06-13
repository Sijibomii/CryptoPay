package types

import (
	"fmt"
	"io"
)

type Client string

const (
	WebClient Client = "web"
)

func (c Client) ToString() string {
	return string(c)
}

func (c Client) String() string {
	return c.ToString()
}

func (c *Client) ToSql(w io.Writer) (err error) {
	_, err = w.Write([]byte(c.ToString()))
	return
}

func (c *Client) FromSql(value interface{}) (err error) {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("unable to convert value to string")
	}

	switch str {
	case "web":
		*c = WebClient
	default:
		return fmt.Errorf("unknown value %s for client found", str)
	}

	return
}
