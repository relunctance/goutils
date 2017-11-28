package math

import (
	"math/rand"
	"time"
)

func Random() float64 {
	source := rand.NewSource(time.Now().Unix() + 20033332)
	r := rand.New(source)
	return r.Float64()
}