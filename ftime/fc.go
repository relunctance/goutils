package ftime

import "time"

// 返回微秒 , unix
func Microtime() (millisecond float64, unix int) {
	t := time.Now()
	millisecond, unix = float64(t.Nanosecond())/1000/1000/1000, t.Unix()
	return
}
