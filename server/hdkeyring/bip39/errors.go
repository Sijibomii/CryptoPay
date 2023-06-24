package bip39

import (
	"fmt"
	"strconv"
)

type Error struct {
	Msg string
	Err error
}

func (e Error) Error() string {
	return e.Msg
}

func InvalidWordLength(length int) Error {
	return Error{
		Msg: fmt.Sprintf("invalid word length %d", length),
	}
}

func ParseIntError(err error) Error {
	return Error{
		Msg: "ParseIntError",
		Err: err,
	}
}

func InvalidWord(word string) Error {
	return Error{
		Msg: fmt.Sprintf("invalid word %s provided", word),
	}
}

func CustomError(msg string) Error {
	return Error{
		Msg: msg,
	}
}

func InvalidChecksum() Error {
	return Error{
		Msg: "invalid checksum",
	}
}

func main() {
	// Example usage
	err := Convert("abc")
	if err != nil {
		fmt.Println(err)
	}
}

func Convert(input string) error {
	_, err := strconv.Atoi(input)
	if err != nil {
		return ParseIntError(err)
	}
	return nil
}
