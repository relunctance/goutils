package fc

import (
	"math/rand"
	"time"
)

// n 生成随机数的最大范围
func RandGenerator(n int) int {
	rand.Seed(time.Now().UnixNano())
	out := make(chan int)
	go func(x int) {
		for {
			out <- rand.Intn(x)
		}
	}(n)
	return <-out
}
