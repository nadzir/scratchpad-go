package stringutil

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func GetSHAHash(value string) []byte {
	stringToHash := value
	sha := sha1.New()
	sha.Write([]byte(stringToHash))
	hashedValue := sha.Sum(nil)
	return hashedValue
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
