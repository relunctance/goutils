package fc

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestMinInt(t *testing.T) {
	if num := MinInt(3, 4); num != 3 {
		t.Fatalf("should be === 3\n")
	}
	if num := MinInt(18, 3); num != 3 {
		t.Fatalf("should be === 3\n")
	}
	if num := MinInt64(7, 8); num != int64(7) {
		t.Fatalf("should be == 7 \n")
	}
	if num := MinInt64(18, 3); num != 3 {
		t.Fatalf("should be == 3")
	}
}
func TestRound(t *testing.T) {
	convey.Convey("测试math:", t, func() {
		//   example 1: round(1241757, -3)
		//   returns 1: 1242000
		//   example 2: round(3.6)
		//   returns 2: 4
		//   example 3: round(2.835, 2)
		//   returns 3: 2.84
		//   example 4: round(1.1749999999999, 2)
		//   returns 4: 1.17
		//   example 5: round(58551.799999999996, 2)
		//   returns 5: 58551.8
		fmt.Println(Round(1, 0))
		fmt.Println(Round(124, 3))
		fmt.Println(Round(3.6, 0))
		fmt.Println(Round(2.835, 2))
		fmt.Println(Round(1.1749999999999, 2))
		fmt.Println(Round(58551.799999999996, 3))
	})
}
