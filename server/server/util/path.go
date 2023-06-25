package util

import (
	"strconv"
	"strings"

	"github.com/tyler-smith/go-bip32"
)

func NewChildKeyFromString(key *bip32.Key, path string) (*bip32.Key, error) {

	path = strings.TrimPrefix(path, "m/")

	childIdx, err := parseDerivationPath(path)

	if err != nil {
		return nil, err
	}
	return key.NewChildKey(childIdx)
}

func parseDerivationPath(path string) (uint32, error) {
	// Split the path into levels
	levels := strings.Split(path, "/")

	// Start from the master key
	childIdx := uint32(0)

	for _, level := range levels {
		var hardened bool

		// Check if the level ends with "'"
		if strings.HasSuffix(level, "'") {
			hardened = true
			// Remove the "'" from the level
			level = strings.TrimSuffix(level, "'")
		}

		// Parse the level as an index
		idx, err := strconv.ParseUint(level, 10, 32)

		if err != nil {
			return 0, err
		}

		// Calculate the child index based on the hardened flag
		if hardened {
			childIdx = childIdx + bip32.FirstHardenedChild + uint32(idx)
		} else {
			childIdx = childIdx + uint32(idx)
		}
	}

	return childIdx, nil
}
