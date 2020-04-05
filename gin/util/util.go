package util

import (
	"math/rand"
	"time"
)

//生成随机name
func RandomString(n int) string {
	var letters = []byte("asdjhajkdhajkAFJHASFLAHFKLAHasdahdlASKDHLA")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix()) //时间戳
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
