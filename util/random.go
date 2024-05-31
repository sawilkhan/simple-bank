package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var r *rand.Rand

func init(){
	//rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

//RandomInt generates a random integer between min & max
func RandomInt(min, max int64) int64 {
	//return min + rand.Int63n(max-min+1)
	return r.Int63n(max-min+1)
}

//RandomString generates a random string 
func RandomString(n int) string{
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//Generate a random owner name
func RandomOwner() string{
	return RandomString(6)
}

//Generate a random amount of money
func RandomMoney() int64{
	return RandomInt(0, 1000)
}

//Generate a random currency
func RandomCurrency() string{
	currencies := []string{"EUR", "USD", "INR", "CAD"}
	n := len(currencies)
	return currencies[r.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}