package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateRandomString(length int, charset string) string {
	if charset == "" {
		// without 0, o, O, 1, l, i, I, u, v
		charset = "abcdefghjkmnpqrstwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	}
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
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
