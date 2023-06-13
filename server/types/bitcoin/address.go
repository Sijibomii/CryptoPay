package bitcoin

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil/base58"
)

type AddressType int

const (
	P2PKH AddressType = iota
	P2SH
)

type Address string

func (addr Address) AddressType() AddressType {
	switch addr[0] {
	case '1', 'm', 'n':
		return P2PKH
	case '3', '2':
		return P2SH
	default:
		panic("invalid bitcoin address found")
	}
}

func (addr Address) Network() wire.BitcoinNet {
	switch addr[0] {
	case 'm', 'n':
		return wire.MainNet
	case '2':
		return wire.TestNet3
	default:
		panic("invalid bitcoin address found")
	}
}

func NewAddress(address string) (Address, error) {

	raw := base58.Decode(address)

	if len(raw) != 25 {
		return "", errors.New("invalid bitcoin address length")
	}
	// The checksum is a mechanism used to ensure the integrity of the bitcoin address during decoding and validation.
	checksum := chainhash.DoubleHashB(raw[:21])[:4]
	if !bytes.Equal(raw[21:], checksum) {
		return "", errors.New("invalid bitcoin address checksum")
	}

	// Only support P2PKH for now.
	switch address[0] {
	case '1', 'm', 'n':
		return Address(address), nil
	default:
		return "", errors.New("address type not supported")
	}
}

func (addr Address) String() string {
	return string(addr)
}

func (addr *Address) Scan(value interface{}) error {
	switch value := value.(type) {
	case []byte:
		*addr = Address(value)
		return nil
	case string:
		*addr = Address(value)
		return nil
	default:
		return errors.New("unsupported value type for Address")
	}
}

func (addr Address) Value() (driver.Value, error) {
	return string(addr), nil
}

type AddressString string

func (addr AddressString) String() string {
	return string(addr)
}

func (addr *AddressString) Scan(value interface{}) error {
	switch value := value.(type) {
	case []byte:
		*addr = AddressString(value)
		return nil
	case string:
		*addr = AddressString(value)
		return nil
	default:
		return errors.New("unsupported value type for AddressString")
	}
}

func (addr AddressString) Value() (driver.Value, error) {
	return string(addr), nil
}
