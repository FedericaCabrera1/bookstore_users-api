package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	inputToByteArray := []byte(input)
	hash.Write(inputToByteArray)
	return hex.EncodeToString(hash.Sum(nil))
}
