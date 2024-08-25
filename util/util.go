package util

import (
	"fmt"

	"golang.org/x/exp/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetNewBucketKey() string {
	return fmt.Sprintf("TokenBucket:%s", randStringBytesRmndr(6))
}

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
