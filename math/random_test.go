package math

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestRandom(t *testing.T) {
	convey.Convey("测试随机数:", t, func() {
		convey.So(Random(), convey.ShouldNotEqual, Random())
	})
}
