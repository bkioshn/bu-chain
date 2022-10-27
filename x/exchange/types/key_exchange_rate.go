package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExchangeRateKeyPrefix is the prefix to retrieve all ExchangeRate
	ExchangeRateKeyPrefix = "ExchangeRate/value/"
)

// ExchangeRateKey returns the store key to retrieve a ExchangeRate from the index fields
func ExchangeRateKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
