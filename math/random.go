package math

import (
	"math/rand"
	"time"
)

func Random() float64 {
	source := rand.NewSource(time.Now().UnixNano() + 20033332)
	r := rand.New(source)
	return r.Float64()
}


// shuffle int slice
func ShuffleInt(arr []int) []int {
    for i := len(arr) - 1; i >= 0; i-- {
        p := Rand(0, i)
        arr[i], arr[p] = arr[p], arr[i]
    }
    return arr
}


// Rand between [min ~ max]
func Rand(min, max int) int {
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)
    v := max - min
    if v <= 0 {
        v = 1
    }
    return r.Intn(v) + min
}
