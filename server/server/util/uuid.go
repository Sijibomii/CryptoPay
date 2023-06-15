package util

import (
	"encoding/base32"

	"github.com/google/uuid"
)

type IDType byte

const (
	IDTypeNone    IDType = '7'
	IDTypeSession IDType = 's'
	IDTypeUser    IDType = 'u'
	IDTypeToken   IDType = 'k'
)

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 27
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off, and a one character alpha prefix indicating the
// type of entity or a `7` if unknown type.
func NewID(idType IDType) string {
	return string(idType) + newId()
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769").WithPadding(base32.NoPadding)

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// without the padding.
func newId() string {
	u := uuid.New()
	uuidBytes := u[:]
	return encoding.EncodeToString(uuidBytes)
}
