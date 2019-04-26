package fc

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSliceDiff(t *testing.T) {

	arr1 := []string{"a", "f", "b", "c", "g", "e", "i"}
	arr2 := []string{"c", "b", "d"}
	arr3 := []string{"d", "e"}
	arr := SliceStringDiff(arr1, arr2, arr3)
	fmt.Println(arr)
	if !CheckStringSliceEqual(arr, []string{"a", "f", "g", "i"}) {

		t.Fatalf("should be equal")
	}
}
func TestInIntsArray(t *testing.T) {
	//arr := []int{1, 2, 3, 4, 5, 6}
	//InStringArray(3, arr)
	convey.Convey("是否存在Int数组中", t, func() {

		convey.Convey("传入的数组必须是排序后的数组", func() {
			num := 129
			arr := []int{3, 11, 16, 18, 19, 20, 55, 58, 104, 142, 142, 142, 142, 143, 10006, 10015, 10027, 10046, 10053, 10057, 10058, 10059, 10060, 10062, 10063, 10064, 10069, 10221, 129}
			convey.So(InArrayInts(num, arr), convey.ShouldBeTrue)
		})
		convey.Convey("存在于数组中", func() {
			arr := []int{6, 7, 8, 9}
			convey.So(InArrayInts(8, arr), convey.ShouldBeTrue)
		})
		convey.Convey("传入的数组为空", func() {
			arr := make([]int, 0)
			convey.So(InArrayInts(8, arr), convey.ShouldBeFalse)
		})
		convey.Convey("不存在于数组中", func() {
			convey.So(InArrayInts(100, []int{6, 7, 8, 9}), convey.ShouldBeFalse)
		})
	})
}

func TestCheckIntSliceEqual(t *testing.T) {
	convey.Convey("判断int数组是否相等", t, func() {
		convey.Convey("开始检测", func() {
			var b bool
			b = CheckIntSliceEqual([]int{1, 2}, []int{1})
			convey.So(b, convey.ShouldBeFalse)
			b = CheckIntSliceEqual([]int{1, 2}, []int{1, 3})
			convey.So(b, convey.ShouldBeFalse)
			b = CheckIntSliceEqual([]int{1, 2, 18, 20, 88, 9}, []int{1, 2, 18, 20, 88, 9})
			convey.So(b, convey.ShouldBeTrue)
		})
	})
}

func TestCheckStringSliceEqual(t *testing.T) {
	convey.Convey("判断string数组是否相等", t, func() {
		convey.Convey("开始检测", func() {
			var b bool
			b = CheckStringSliceEqual([]string{"a", "b"}, []string{"a"})
			convey.So(b, convey.ShouldBeFalse)
			b = CheckStringSliceEqual([]string{"a", "b"}, []string{"a", "c"})
			convey.So(b, convey.ShouldBeFalse)
			b = CheckStringSliceEqual([]string{"a", "b", "c", "d", "E", "F"}, []string{"a", "b", "c", "d", "E", "F"})
			convey.So(b, convey.ShouldBeTrue)
		})
	})
}

func TestEndStringSlice(t *testing.T) {
	convey.Convey("测试slice获取最后一个值", t, func() {
		convey.Convey("测试字符串", func() {
			arr := []string{
				"a",
				"b",
				"c",
				"d",
				"e",
				"f",
			}
			convey.So("f", convey.ShouldEqual, EndStringSlice(arr))
			convey.So("", convey.ShouldEqual, EndStringSlice([]string{}))
		})

		convey.Convey("测试int数组", func() {
			arr := []int{0, 1, 2, 3, 4, 5, 6}
			convey.So(6, convey.ShouldEqual, EndIntSlice(arr))
			convey.So(-1, convey.ShouldEqual, EndIntSlice([]int{}))
		})

	})
}

func TestIssetSlice(t *testing.T) {
	convey.Convey("测试是否slice下标是否设置:", t, func() {

		convey.Convey("测试int数组:", func() {
			arr := []int{0, 1, 2, 3}
			convey.So(IssetSlice(arr, 0), convey.ShouldBeTrue)
			convey.So(IssetSlice(arr, 100), convey.ShouldBeFalse)
			arr2 := [][]int{
				{0, 1, 2, 3, 4},
				{5, 6, 7},
			}
			convey.So(IssetSlice(arr2, 0), convey.ShouldBeTrue)
			convey.So(IssetSlice(arr2, 3), convey.ShouldBeFalse)

		})

		convey.Convey("测试string数组:", func() {
			arr := []string{"a", "b", "c", "d"}
			convey.So(IssetSlice(arr, 0), convey.ShouldBeTrue)
			convey.So(IssetSlice(arr, 100), convey.ShouldBeFalse)
			arr2 := [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
			}
			convey.So(IssetSlice(arr2, 0), convey.ShouldBeTrue)
			convey.So(IssetSlice(arr2, 3), convey.ShouldBeFalse)

		})

		convey.Convey("测试panic", func() {
			type User struct {
				Name string
				Pass string
			}
			convey.So(func() {
				arr := []User{
					User{"a", "b"},
					User{"c", "d"},
				}
				IssetSlice(arr, 0)
			}, convey.ShouldPanic)
		})
	})
}
