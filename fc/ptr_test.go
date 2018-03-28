package fc

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type User struct {
	Name string
}

func TestIsPtr(t *testing.T) {
	convey.Convey("检测IsPtr()函数", t, func() {
		convey.So(IsPtr(1), convey.ShouldBeFalse)
		convey.So(IsPtr("abc"), convey.ShouldBeFalse)
		convey.So(IsPtr([]int{1, 2, 3}), convey.ShouldBeFalse)
		convey.So(IsPtr([2]int{1, 2}), convey.ShouldBeFalse)
		convey.So(IsPtr(map[string]string{"username": "hello"}), convey.ShouldBeFalse)
		convey.So(IsPtr(User{Name: "hello"}), convey.ShouldBeFalse)
		convey.So(IsPtr(&User{Name: "hello"}), convey.ShouldBeTrue)
	})
}
