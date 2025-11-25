package pkg

import (
	"math/big"
	"strings"
)

const BASE62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(input []byte) string {
	if len(input) == 0 {
		return ""
	}

	value := new(big.Int).SetBytes(input)
	var result strings.Builder

	for value.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		value.DivMod(value, big.NewInt(62), mod)
		result.WriteByte(BASE62[mod.Int64()])
	}

	for _, b := range input {
		if b == 0 {
			result.WriteByte(BASE62[0])
		} else {
			break
		}
	}

	encoded := result.String()
	runes := []rune(encoded)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Decode(input string) []byte {
	if len(input) == 0 {
		return []byte{}
	}

	value := big.NewInt(0)
	for _, c := range input {
		value.Mul(value, big.NewInt(62))
		value.Add(value, big.NewInt(int64(strings.IndexRune(BASE62, c))))
	}

	decoded := value.Bytes()
	leadingZeroes := 0
	for _, c := range input {
		if c == rune(BASE62[0]) {
			leadingZeroes++
		} else {
			break
		}
	}

	result := make([]byte, leadingZeroes+len(decoded))
	copy(result[leadingZeroes:], decoded)
	return result
}
