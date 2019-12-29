package mathUtil

import (
	"math/rand"
	"time"
)

var secureRand *rand.Rand

func Float64FomInt(min, max int) float64 {
	return float64(secureRand.Int63n(int64(max-min)) + int64(min))
}

func Int(min, max int) int {
	return secureRand.Intn(max-min) + min
}

func init() {
	s1 := rand.NewSource(time.Now().UnixNano())
	secureRand = rand.New(s1)
}
