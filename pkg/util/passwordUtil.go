package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

func GenerateRandomString(length int, charset string) string {
	if charset == "" {
		charset = "abcdefghjkmnpqrstwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
func GenerateMD5Password(password string, salt string) string {
	hash := md5.New()
	hash.Write([]byte(password + salt))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func VerifyPassword(password string, salt string, hash string) bool {
	return GenerateMD5Password(password, salt) == hash
}
