package bip39

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	pbkdf2Rounds = 2048
	pbkdf2Bytes  = 64
)

var (
	englishWordlist []string
	frenchWordlist  []string
	englishWordmap  map[string]uint16
	frenchWordmap   map[string]uint16
)

type MnemonicType int

const (
	Words12 MnemonicType = iota
	Words24
)

func (m MnemonicType) entropyLength() int {
	switch m {
	case Words12:
		return 128
	case Words24:
		return 256
	default:
		return 0
	}
}

func (m MnemonicType) checksumLength() int {
	switch m {
	case Words12:
		return 4
	case Words24:
		return 8
	default:
		return 0
	}
}

func (m MnemonicType) wordCount() int {
	switch m {
	case Words12:
		return 12
	case Words24:
		return 24
	default:
		return 0
	}
}

type Language int

const (
	English Language = iota
	French
)

func (l Language) getWordlist() []string {
	switch l {
	case English:
		return englishWordlist
	case French:
		return frenchWordlist
	default:
		return nil
	}
}

func (l Language) getWordmap() map[string]uint16 {
	switch l {
	case English:
		return englishWordmap
	case French:
		return frenchWordmap
	default:
		return nil
	}
}

// func main() {
// 	err := loadWordlists()
// 	if err != nil {
// 		// Handle error
// 		return
// 	}

// 	// Rest of the code...
// }

func loadWordlists() error {
	englishBytes, err := ioutil.ReadFile("wordlist/english.txt")
	if err != nil {
		return err
	}
	englishWordlist = splitWords(string(englishBytes))

	frenchBytes, err := ioutil.ReadFile("wordlist/french.txt")
	if err != nil {
		return err
	}
	frenchWordlist = splitWords(string(frenchBytes))

	englishWordmap = createWordmap(englishWordlist)
	frenchWordmap = createWordmap(frenchWordlist)

	return nil
}

func splitWords(input string) []string {
	return strings.Fields(input)
}

func createWordmap(wordlist []string) map[string]uint16 {
	wordmap := make(map[string]uint16)
	for i, word := range wordlist {
		wordmap[word] = uint16(i)
	}
	return wordmap
}

func generateEntropy(length int) ([]byte, error) {
	entropy := make([]byte, length)
	_, err := rand.Read(entropy)
	if err != nil {
		return nil, err
	}
	return entropy, nil
}

func generateSeed(entropy []byte, passphrase string) ([]byte, error) {
	salt := []byte("mnemonic" + passphrase)
	seed := pbkdf2.Key(entropy, salt, pbkdf2Rounds, pbkdf2Bytes, sha256.New)
	return seed, nil
}

func getWordlist(language Language) []string {
	switch language {
	case English:
		return englishWordlist
	case French:
		return frenchWordlist
	default:
		return nil
	}
}

func getWordmap(language Language) map[string]uint16 {
	switch language {
	case English:
		return englishWordmap
	case French:
		return frenchWordmap
	default:
		return nil
	}
}

func generateMnemonic(language Language, wordCount int) ([]string, error) {
	wordlist := getWordlist(language)
	if wordlist == nil {
		return nil, errors.New("unknown language")
	}

	wordmap := getWordmap(language)
	if wordmap == nil {
		return nil, errors.New("unknown language")
	}

	mnemonic := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordlist))))
		if err != nil {
			return nil, err
		}
		mnemonic[i] = wordlist[index.Int64()]
	}

	return mnemonic, nil
}

////

// Seed represents a BIP-39 seed.
type Seed struct {
	bytes []byte
}

// Mnemonic represents a BIP-39 mnemonic.
type Mnemonic struct {
	phrase  string
	lang    Language
	entropy []byte
	seed    Seed
}

// NewSeed creates a new BIP-39 seed from the given entropy and password.
func NewSeed(entropy []byte, password string) Seed {
	salt := []byte("mnemonic" + password)
	seedBytes := pbkdf2.Key(entropy, salt, 2048, 64, sha512.New)
	return Seed{bytes: seedBytes}
}

// AsBytes returns the seed as a byte slice.
func (s Seed) AsBytes() []byte {
	return s.bytes
}

func NewMnemonic(typ MnemonicType, lang Language, password string) (Mnemonic, error) {
	entropyLength := typ.entropyLength() / 8

	entropy := make([]byte, entropyLength)

	_, err := rand.Read(entropy)

	if err != nil {
		return Mnemonic{}, err
	}

	err = loadWordlists()

	if err != nil {
		fmt.Println("########## eorldlist error", err.Error())
	}

	wordlist := lang.getWordlist()

	fmt.Println("worldlist \n", wordlist)

	checksumLength := typ.checksumLength()

	fmt.Println("checksumLength\n", checksumLength)

	checksum := calculateChecksum(entropy, checksumLength)
	fmt.Println("checksum \n", checksum)

	withChecksum := append(entropy, checksum.bits...)
	withChecksumBits := bytesToBits(withChecksum)

	withChecksumBits.Truncate(11 * typ.wordCount())

	var phrase []string
	indexBits := newBitVec()

	for _, bit := range withChecksumBits.bits {
		indexBits.Push(bit != 0)

		if indexBits.Len() == 11 {
			index, err := indexBits.ToInt()
			if err != nil {
				return Mnemonic{}, err
			}

			word := wordlist[index]
			phrase = append(phrase, word)

			indexBits.Reset()
		}
	}

	seed := NewSeed([]byte(strings.Join(phrase, " ")), password)

	mnemonic := Mnemonic{
		phrase:  strings.Join(phrase, " "),
		lang:    lang,
		entropy: entropy,
		seed:    seed,
	}

	return mnemonic, nil
}

func MnemonicTypeFromWordCount(wordCount int) (MnemonicType, error) {
	switch wordCount {
	case 12:
		return MnemonicType(12), nil

	case 24:
		return MnemonicType(24), nil
	default:
		return 0, fmt.Errorf("invalid word count: %d", wordCount)
	}
}

// FromString creates a new BIP-39 mnemonic from the given phrase, language, and password.
func FromString(phrase string, lang Language, password string) (Mnemonic, error) {
	typ, err := MnemonicTypeFromWordCount(strings.Count(phrase, " ") + 1)

	if err != nil {
		return Mnemonic{}, err
	}

	checksumLength := typ.checksumLength()
	wordMap := lang.getWordmap()

	words := strings.Split(phrase, " ")
	indexes := make([]uint16, len(words))

	for i, word := range words {
		index, ok := wordMap[word]
		if !ok {
			return Mnemonic{}, errors.New("invalid word")
		}
		indexes[i] = index
	}

	withChecksumBits := newBitVec()

	for _, index := range indexes {
		for x := 10; x >= 0; x-- {
			bit := (index >> uint(x)) & 1
			withChecksumBits.Push(bit == 1)
		}
	}

	checksumBits := withChecksumBits.TruncateAndGetLastBits(checksumLength)
	checksumBits.Reverse()

	withChecksumBits.Truncate(11 * typ.wordCount())
	entropy := bitsToBytes(withChecksumBits)

	expectedChecksum := calculateChecksum(entropy, checksumLength)

	if !bitVecsEqual(checksumBits, expectedChecksum) {
		return Mnemonic{}, errors.New("invalid checksum")
	}

	seed := NewSeed([]byte(phrase), password)

	mnemonic := Mnemonic{
		phrase:  phrase,
		lang:    lang,
		entropy: entropy,
		seed:    seed,
	}

	return mnemonic, nil
}

// Phrase returns the mnemonic phrase.
func (m Mnemonic) Phrase() string {
	return m.phrase
}

// Seed returns the BIP-39 seed.
func (m Mnemonic) Seed() Seed {
	return m.seed
}

// Utility function to calculate the checksum of the entropy.
func calculateChecksum(entropy []byte, checksumLength int) *bitVec {
	hash := sha256.Sum256(entropy)
	checksumBytes := hash[:checksumLength/8]
	checksumBits := bytesToBits(checksumBytes)
	checksumBits.Truncate(checksumLength)
	return checksumBits
}

// bitVec is a simple bit vector implementation.
type bitVec struct {
	bits   []byte
	length int
}

func newBitVec() *bitVec {
	return &bitVec{}
}

func (bv *bitVec) Push(bit bool) {
	byteIndex := bv.length / 8
	bitIndex := bv.length % 8

	if byteIndex >= len(bv.bits) {
		bv.bits = append(bv.bits, 0)
	}

	if bit {
		bv.bits[byteIndex] |= (1 << bitIndex)
	}

	bv.length++
}

func (bv *bitVec) Len() int {
	return bv.length
}

func (bv *bitVec) ToInt() (int, error) {
	if bv.length > 32 {
		return 0, errors.New("bitVec length too large")
	}

	var value int
	for i := 0; i < bv.length; i++ {
		if bv.bits[i/8]&(1<<(i%8)) != 0 {
			value |= (1 << (bv.length - 1 - i))
		}
	}
	return value, nil
}

func (bv *bitVec) Reset() {
	bv.bits = bv.bits[:0]
	bv.length = 0
}

func (bv *bitVec) Truncate(length int) {

	byteLength := (length + 7) / 8
	fmt.Println("The number is:", bv.bits)
	bv.bits = bv.bits[:byteLength]
	bv.length = length
}

func (bv *bitVec) TruncateAndGetLastBits(length int) *bitVec {
	byteLength := (length + 7) / 8
	lastBits := newBitVec()
	lastBits.bits = bv.bits[:byteLength]
	lastBits.length = length
	bv.bits = bv.bits[byteLength:]
	bv.length -= length
	return lastBits
}

func (bv *bitVec) Reverse() {
	for i, j := 0, bv.length-1; i < j; i, j = i+1, j-1 {
		bv.bits[i/8], bv.bits[j/8] = bv.bits[j/8], bv.bits[i/8]
		bv.bits[i/8] = reverseByte(bv.bits[i/8])
		bv.bits[j/8] = reverseByte(bv.bits[j/8])
	}
}

func bytesToBits(bytes []byte) *bitVec {
	fmt.Println("######### bytes =", bytes)
	bits := newBitVec()
	for _, b := range bytes {
		for i := 0; i < 8; i++ {
			bit := (b & (1 << i)) != 0
			bits.Push(bit)
		}
	}
	return bits
}

func bitsToBytes(bits *bitVec) []byte {
	byteLength := (bits.length + 7) / 8
	bytes := make([]byte, byteLength)
	for i := 0; i < bits.length; i++ {
		if bits.bits[i/8]&(1<<(i%8)) != 0 {
			bytes[i/8] |= (1 << (i % 8))
		}
	}
	return bytes
}

func bitVecsEqual(bv1, bv2 *bitVec) bool {
	if bv1.length != bv2.length {
		return false
	}
	for i := 0; i < bv1.length; i++ {
		if (bv1.bits[i/8] & (1 << (i % 8))) != (bv2.bits[i/8] & (1 << (i % 8))) {
			return false
		}
	}
	return true
}

func reverseByte(b byte) byte {
	b = (b&0xF0)>>4 | (b&0x0F)<<4
	b = (b&0xCC)>>2 | (b&0x33)<<2
	b = (b&0xAA)>>1 | (b&0x55)<<1
	return b
}
