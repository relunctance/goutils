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


func TestFstime2(t *testing.T) {
    ms := map[string]int64{
        "2020-01-29 18:49:46": 1580294986,
        "2020-01-30 02:49:46": 1580323786,
        "2019-01-11 21:49:42": 1547214582,
    }
    for k, v := range ms {
        ti, _ := Fstrtotime(k)
        if ti.Unix() != v { // 不允许自动增加8小时
            t.Fatalf("不允许自动增加8小时时区,should be === %d", v)
        }
    }
}
