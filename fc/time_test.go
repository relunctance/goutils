package fc

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFstrtotime(t *testing.T) {
	convey.Convey("测试Fstrtotime", t, func() {
		convey.Convey("开始测试", func() {
			str := " 2018-03-01 "
			ti, err := Fstrtotime(str)
			fmt.Println("aaaaaaa:", ti, err)

			str = "2018-3-01"
			ti, err = Fstrtotime(str)
			fmt.Println("aaaaaaa:", ti, err)

			str = "2018-03-01 14:57:51"
			ti, err = Fstrtotime(str)
			fmt.Println("aaaaaaa:", ti, err)

		})
	})
}
