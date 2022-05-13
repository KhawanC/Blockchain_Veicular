package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type Contador struct {
	ips []int
}

func AleatString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Encode(msg string) string {
	h := sha1.New()
	h.Write([]byte(msg))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctg := Contador{}
	fmt.Println(ctg.Contador + 1)

	msg := Encode(AleatString(20))

	fmt.Println(msg, "---", reflect.ValueOf(msg).Kind())

}
