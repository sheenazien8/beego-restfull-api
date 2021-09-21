package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func CryptoMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func CryptoDoubleMD5Hash(text string) string {
	return CryptoMD5Hash(CryptoMD5Hash(text))
}

func CryptoSHA256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
