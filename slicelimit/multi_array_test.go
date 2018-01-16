package slicelimit

import (
	"strings"
	"testing"

	"github.com/relunctance/goutils/str"
)

var data [][]string = [][]string{
	[]string{"1a", "1b", "1c"},                                                 //3
	[]string{"2a", "2b", "2c", "2d", "2e", "2f", "2g"},                         //7
	[]string{"3a", "3b", "3c", "3d", "3e", "3f", "3g", "3h", "3i"},             //9
	[]string{"4a", "4b", "4c", "4d", "4e", "4f", "4g", "4h", "4i", "4j", "4k"}, //11
}

//获取数据结果
func getData(arr [][]int) (res []string) {

	for key, items := range arr {
		if items == nil {
			continue
		}
		for _, v := range items {
			res = append(res, data[key][v])
		}
	}
	return
}

func TestGetDyadicArrayByOffset(t *testing.T) {
	convey.Convey("测试二维数组的offset,pagesize检索", t, func() {
		multiDyadicArr, totalLength := BuildDyadicArray(data) //totalLength == 30

		offset := int64(0)
		pagesize := int64(10)
		err, arr := GetDyadicArrayByOffset(offset, pagesize, totalLength, multiDyadicArr)
		convey.Convey("err应该返回nil", func() {
			convey.So(nil, convey.ShouldEqual, err)
		})
		convey.Convey("arr值判断1:", func() {
			bl := str.CheckStringSliceEqual(strings.Fields("1a 1b 1c 2a 2b 2c 2d 2e 2f 2g"), getData(arr))
			convey.So(bl, convey.ShouldBeTrue)
		})

		err, arr = GetDyadicArrayByOffset(10000, pagesize, totalLength, multiDyadicArr)
		convey.Convey("检测err,应该不为空", func() {
			convey.So(err, convey.ShouldNotBeNil)
		})
		err, arr = GetDyadicArrayByOffset(30, pagesize, totalLength, multiDyadicArr)
		convey.Convey("检测结束符号", func() {
			convey.So(err, convey.ShouldEqual, EOF)
		})

		err, arr = GetDyadicArrayByOffset(5, 10, totalLength, multiDyadicArr)
		convey.Convey("arr值判断2:", func() {
			bl := str.CheckStringSliceEqual(strings.Fields("2c 2d 2e 2f 2g 3c 3d 3e 3f 3g"), getData(arr))
			convey.So(bl, convey.ShouldBeTrue)
			convey.So(err, convey.ShouldBeNil)
		})
		err, arr = GetDyadicArrayByOffset(3, 5, totalLength, multiDyadicArr)
		convey.Convey("检测当start为边界时 , 是否报错问题:", func() {
			bl := str.CheckStringSliceEqual(strings.Fields("2a 2b 2c 2d 2e"), getData(arr))
			convey.So(bl, convey.ShouldBeTrue)
			convey.So(err, convey.ShouldBeNil)

		})
		err, arr = GetDyadicArrayByOffset(0, 5000000, totalLength, multiDyadicArr)
		convey.Convey("当pagesize过大时检测", func() {
			bl := str.CheckStringSliceEqual(strings.Fields("1a 1b 1c 2a 2b 2c 2d 2e 2f 2g 3a 3b 3c 3d 3e 3f 3g 3h 3i 4a 4b 4c 4d 4e 4f 4g 4h 4i 4j 4k"), getData(arr))
			convey.So(bl, convey.ShouldBeTrue)
			convey.So(err, convey.ShouldBeNil)

		})
		err, arr = GetDyadicArrayByOffset(13, 5000000, totalLength, multiDyadicArr)
		convey.Convey("检测", func() {
			bl := str.CheckStringSliceEqual(strings.Fields("3d 3e 3f 3g 3h 3i 4d 4e 4f 4g 4h 4i 4j 4k"), getData(arr))
			convey.So(bl, convey.ShouldBeTrue)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestBuildDyadicArray(t *testing.T) {
	convey.Convey("测试多维数组的搜索", t, func() {
		arr := [][]string{}
		mutliarr, num := BuildDyadicArray(arr)
		convey.Convey("长度应该为0", func() {
			convey.So(0, convey.ShouldEqual, len(mutliarr))
		})
		convey.Convey("结果应该为0", func() {
			convey.So(num, convey.ShouldEqual, 0)
		})

		mutliarr, num = BuildDyadicArray(data)

		convey.Convey("结果应该相等:数组长度", func() {
			convey.So(len(mutliarr), convey.ShouldEqual, len(data))
		})
		convey.Convey("结果应该相等:总长", func() {
			convey.So(30, convey.ShouldEqual, num)
		})
	})

}
