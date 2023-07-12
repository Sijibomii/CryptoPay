package bitcoin

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip32"
	"golang.org/x/crypto/sha3"
)

type Script struct {
	Data []byte
}

// opcodes used in Bitcoin scripting.
const (
	OP_DUP          byte = 0x76
	OP_HASH160      byte = 0xa9
	OP_PUSHBYTES_20 byte = 0x14
	OP_EQUALVERIFY  byte = 0x88
	OP_CHECKSIG     byte = 0xac
)

func (s *Script) p2pkh(to string) {
	decoded := base58.Decode(to)

	var pkh []byte
	if len(decoded) >= 21 {
		pkh = decoded[1:21]
	} else {
		panic("decoded slice is not long enough")
	}

	var script []byte
	script = append(script, OP_DUP)
	script = append(script, OP_HASH160)
	script = append(script, OP_PUSHBYTES_20)
	script = append(script, pkh...)
	script = append(script, OP_EQUALVERIFY)
	script = append(script, OP_CHECKSIG)

	s.Data = script
}

func (s *Script) script_sig(signature *btcec.Signature, publicKey *bip32.Key) {
	derSig := signature.Serialize()

	// Append 0x01 (SIGHASH_ALL) to the DER signature
	derSig = append(derSig, 0x01)

	script := make([]byte, 0)

	// Append the length of DER signature to the script
	script = append(script, byte(len(derSig)))

	// Append the DER signature to the script
	script = append(script, derSig...)

	pubKeyBytes, _ := publicKey.Serialize()

	pubKeyLen := len(pubKeyBytes)

	// Append the length of the public key to the script
	script = append(script, byte(pubKeyLen))

	// Append the public key to the script
	script = append(script, pubKeyBytes...)

	s.Data = script
}

func (s *Script) from_hex(hexStr string) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		//fmt.Printf("unable to decode hex correctly")
		panic("unable to decode hex correctly")
	}
	s.Data = bytes
}

type OutPoint struct {
	Hash  string
	Index int
}

type Input struct {
	OutPoint       OutPoint
	Script_sig     Script
	Sequence       int
	Script_witness [][]int
}

type Output struct {
	Value         int
	Script_pubkey Script
}

type UnsignedTransaction struct {
	Version   int
	Inputs    []Input
	Outputs   []Output
	Lock_time int
}

type TransactionInput struct {
	Transaction Transaction
	Idx         int
}

type TransactionOutput struct {
	Address string
	Amount  float64
}

func NewUnsignedTransaction(inp []TransactionInput, out []TransactionOutput) UnsignedTransaction {

	var inputs []Input

	var outputs []Output

	tx := UnsignedTransaction{
		Version:   1,
		Inputs:    inputs,
		Outputs:   outputs,
		Lock_time: 0,
	}

	for _, in := range inp {
		hex_script := in.Transaction.Vout[in.Idx].ScriptPubKey
		script := Script{}
		script.from_hex(hex_script)
		var witness [][]int
		input := Input{
			OutPoint: OutPoint{
				Hash:  in.Transaction.TxID,
				Index: in.Idx,
			},
			Script_sig:     script,
			Sequence:       0xFFFFFFFF,
			Script_witness: witness,
		}

		tx.Inputs = append(tx.Inputs, input)
	}

	for _, outp := range out {
		script := Script{}
		script.p2pkh(outp.Address)
		output := Output{
			Value:         int(outp.Amount),
			Script_pubkey: script,
		}
		tx.Outputs = append(tx.Outputs, output)
	}

	return tx
}

func (u *UnsignedTransaction) Sign(secretKey *bip32.Key, publicKey *bip32.Key) {

	for idx, _ := range u.Inputs {

		hash := u.signature_hash()

		// convert key to ecda
		ekey, err := keyToECDSA(secretKey)

		key, err := ecdsaToBTCEC(ekey)

		if err != nil {

		}

		_, err, sig := signWithSecp256k1(hash[:], key)
		script := Script{}
		script.script_sig(sig, publicKey)
		u.Inputs[idx].Script_sig = script
	}

}

func (u *UnsignedTransaction) Serialize(b *bytes.Buffer) {
	binary.Write(b, binary.LittleEndian, uint32(u.Version))

	serializeVarInt(b, len(u.Inputs))
	for _, input := range u.Inputs {
		reverseHash := reverseString(input.OutPoint.Hash)
		b.Write([]byte(reverseHash))
		binary.Write(b, binary.LittleEndian, uint32(input.OutPoint.Index))

		// _ := serializeVarInt(b, len(input.Script_sig.Data))
		b.Write(input.Script_sig.Data)
		binary.Write(b, binary.LittleEndian, uint32(input.Sequence))
	}

	serializeVarInt(b, len(u.Outputs))
	for _, output := range u.Outputs {
		binary.Write(b, binary.LittleEndian, uint64(output.Value))

		// _ := serializeVarInt(b, len(output.Script_pubkey.Data))
		b.Write(output.Script_pubkey.Data)
	}

	binary.Write(b, binary.LittleEndian, uint32(u.Lock_time))
}

func (u *UnsignedTransaction) signature_hash() [32]byte {
	txCopy := *u

	var serialized bytes.Buffer
	txCopy.Serialize(&serialized)
	binary.Write(&serialized, binary.LittleEndian, uint32(1)) // SIGHASH ALL

	hash := sha3.NewLegacyKeccak256()
	hash.Write(serialized.Bytes())

	var h [32]byte
	copy(h[:], hash.Sum(nil)[:32])

	return h
}

func (u *UnsignedTransaction) Into_raw_transaction() []byte {
	var buf bytes.Buffer
	u.Serialize(&buf)
	return buf.Bytes()
}
