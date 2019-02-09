// random package help you to generate random strings
package random

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

// id random generator channel
var id = make(chan string, 10)

// InitRandGenerator init the random generator
func InitRandGenerator() {
	rand.Seed(int64(time.Now().Nanosecond()))
	go func() {
		h := sha1.New()
		c := []byte(time.Now().String())
		for {
			_, _ = h.Write(c)
			id <- fmt.Sprintf("%x", h.Sum(nil))
		}
	}()
}

// GetRandID read a random string random package
func GetRandID() string {
	return <-id
}
