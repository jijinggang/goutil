package goutil

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func MD5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}

func RandInt64() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63()
}

func NewMap(args ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	var k string
	for i, v := range args {
		if i%2 == 0 {
			k = v.(string)
		} else {
			m[k] = v
		}
	}
	return m
}
