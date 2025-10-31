package helpers

import (
	"crypto/md5"
	"math/big"
)

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func ShortURL(longURL string) string {
	hash := md5.Sum([]byte(longURL))
	num := new(big.Int).SetBytes(hash[:])

	var short string
	for num.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		num.DivMod(num, big.NewInt(62), mod)
		short = string(base62[mod.Int64()]) + short
	}

	return short[:7] // take first 7 chars
}
