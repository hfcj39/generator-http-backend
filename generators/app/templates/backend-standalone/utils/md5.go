package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str string) string {
	h := md5.New()
	_, _ = h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
