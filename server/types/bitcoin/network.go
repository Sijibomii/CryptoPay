package bitcoin

import (
	"fmt"
)

// Network represents a network type.
type Network int

const (
	Mainnet Network = iota
	Test
)

func (n Network) String() string {
	switch n {
	case Mainnet:
		return "mainnet"
	case Test:
		return "test"
	default:
		return ""
	}
}

// MarshalText implements the encoding.TextMarshaler interface.
func (n Network) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (n *Network) UnmarshalText(data []byte) error {
	str := string(data)
	switch str {
	case "mainnet":
		*n = Mainnet
	case "test":
		*n = Test
	default:
		return fmt.Errorf("invalid value for bitcoin network")
	}
	return nil
}

// Scan implements the sql.Scanner interface.
func (n *Network) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return n.UnmarshalText(v)
	case string:
		return n.UnmarshalText([]byte(v))
	default:
		return fmt.Errorf("unsupported scan type for Network: %T", value)
	}
}
