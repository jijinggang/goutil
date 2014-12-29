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
