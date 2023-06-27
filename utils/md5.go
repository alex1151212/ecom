package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 小寫
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tmpStr := h.Sum(nil)
	return hex.EncodeToString(tmpStr)
}

// 加密
func MakePasssword(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}

// 解密
func ValidPassword(plainPwd, salt string, password string) bool {
	return Md5Encode(plainPwd+salt) == password
}
