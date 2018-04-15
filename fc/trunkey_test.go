package fc

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type User struct {
	Name    string
	Pass    string
	company string
	Age     int
}

func newUser(name, pass string, age int) *User {
	return &User{
		Name: name,
		Pass: pass,
		Age:  age,
	}
}

var DefaultData []string = []string{"name-00", "name-01", "name-02", "name-03", "name-04", "name-05", "name-06", "name-07", "name-08", "name-09", "name-10", "name-11", "name-12", "name-13", "name-14", "name-15", "name-16", "name-17", "name-18", "name-19"}

func TestArrayToSimple(t *testing.T) {
	convey.Convey("测试ArrayToSimple函数", t, func() {
		convey.Convey("测试正常执行:", func() {
			data := make([]*User, 0, 21)
			data2 := make([]User, 0, 21)
			data3 := make(map[int]*User)
			data4 := make(map[int]User)
			var arr []string
			var err error
			for i := 0; i < 20; i++ {
				istr := strconv.Itoa(i)
				if i < 10 {
					istr = "0" + istr
				}
				name := fmt.Sprintf("name-%s", istr)
				pass := fmt.Sprintf("pass-%s", istr)
				age := i*100 + 1
				n := newUser(name, pass, age)
				data = append(data, n)
				data2 = append(data2, *n)
				data3[i] = n
				data4[i] = *n
			}
			data = append(data, nil) //特意追加一个空格的nil
			arr, err = ArrayToSimple(data, "Name")
			convey.So(err, convey.ShouldBeNil)
			defaultData := append(DefaultData, "")
			convey.So(checkStringSliceEqual(arr, defaultData), convey.ShouldBeTrue)

			arr, err = ArrayToSimple(data, "gaoqilin") //测试私有属性
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(len(arr), convey.ShouldEqual, 0)

			arr, err = ArrayToSimple(data2, "Name")
			convey.So(err, convey.ShouldBeNil)
			convey.So(checkStringSliceEqual(arr, DefaultData), convey.ShouldBeTrue)
			arr, err = ArrayToSimple(data3, "Name")
			sort.Strings(arr)
			convey.So(err, convey.ShouldBeNil)
			convey.So(checkStringSliceEqual(arr, DefaultData), convey.ShouldBeTrue)
			arr, err = ArrayToSimple(data4, "Name")
			sort.Strings(arr)
			convey.So(err, convey.ShouldBeNil)
			convey.So(checkStringSliceEqual(arr, DefaultData), convey.ShouldBeTrue)
			data5 := make([]*User, 0, 24)
			data5 = append(data5, data...)
			data5 = append(data5, newUser("name-22", "pass-19", 190))
			data5 = append(data5, newUser("name-23", "pass-19", 190))
			data5 = append(data5, newUser("name-24", "pass-19", 190))
			arr, err = ArrayToSimple(data5, "Name")
			convey.So(err, convey.ShouldBeNil)
			convey.So(24, convey.ShouldEqual, len(arr))
		})
		convey.Convey("测试ArrayToSimple报错部分:", func() {
			type userInt int

			var ui userInt = 3
			arr, err := ArrayToSimple([]userInt{ui}, "Name")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(len(arr), convey.ShouldEqual, 0)
			arr, err = ArrayToSimple([]*userInt{&ui}, "Name")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(len(arr), convey.ShouldEqual, 0)

			arr, err = ArrayToSimple(1234, "Name")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(len(arr), convey.ShouldEqual, 0)

			arr, err = ArrayToSimple("abc", "Name")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(len(arr), convey.ShouldEqual, 0)

			arr, err = ArrayToSimple([]string{"abc"}, "Name")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(len(arr), convey.ShouldEqual, 0)
		})

	})
}

func TestDataTrunKey(t *testing.T) {
	convey.Convey("测试DataTrunKey", t, func() {
		convey.Convey("测试DataTrunKey正常执行:", func() {
			data := make([]*User, 0, 21)
			data2 := make([]User, 0, 21)
			data3 := make(map[int]*User)
			data4 := make(map[int]User)
			var result map[string]interface{}
			var err error
			for i := 0; i < 20; i++ {
				istr := strconv.Itoa(i)
				if i < 10 {
					istr = "0" + istr
				}
				name := fmt.Sprintf("name-%s", istr)
				pass := fmt.Sprintf("pass-%s", istr)
				age := i*100 + 1
				n := newUser(name, pass, age)
				data = append(data, n)
				data2 = append(data2, *n)
				data3[i] = n
				data4[i] = *n
			}
			fieldValue := "name-19"
			convey.Convey("测试slice指针类型:", func() {
				data = append(data, nil) //特意追加一个空格的nil
				result, err = DataTrunKey(data, "Name")
				convey.So(err, convey.ShouldBeNil)
				convey.So(20, convey.ShouldEqual, len(result)) //如果为nil的情况下, DataTrunKey内部直接跳过
				val, ok := result[fieldValue]
				convey.So(ok, convey.ShouldBeTrue)
				value, ok := val.(*User) //必须为*User 类型
				convey.So(ok, convey.ShouldBeTrue)
				convey.So(fieldValue, convey.ShouldEqual, value.Name)
			})

			convey.Convey("测试slice struct类型:", func() {
				result, err = DataTrunKey(data2, "Name")
				convey.So(err, convey.ShouldBeNil)
				convey.So(20, convey.ShouldEqual, len(result)) //如果为nil的情况下, DataTrunKey内部直接跳过
				val, ok := result[fieldValue]
				convey.So(ok, convey.ShouldBeTrue)
				value2, ok := val.(User) //User 类型
				convey.So(ok, convey.ShouldBeTrue)
				convey.So(fieldValue, convey.ShouldEqual, value2.Name)
			})

			convey.Convey("测试map 指针类型:", func() {

				result, err = DataTrunKey(data3, "Name")
				convey.So(err, convey.ShouldBeNil)
				convey.So(20, convey.ShouldEqual, len(result)) //如果为nil的情况下, DataTrunKey内部直接跳过
				val, ok := result[fieldValue]
				convey.So(ok, convey.ShouldBeTrue)
				value, ok := val.(*User) //必须为*User 类型
				convey.So(ok, convey.ShouldBeTrue)
				convey.So(fieldValue, convey.ShouldEqual, value.Name)
			})

			convey.Convey("测试map struct类型:", func() {
				result, err = DataTrunKey(data4, "Name")
				convey.So(err, convey.ShouldBeNil)
				convey.So(20, convey.ShouldEqual, len(result)) //如果为nil的情况下, DataTrunKey内部直接跳过
				val, ok := result[fieldValue]
				convey.So(ok, convey.ShouldBeTrue)
				value2, ok := val.(User) //User 类型
				convey.So(ok, convey.ShouldBeTrue)
				convey.So(fieldValue, convey.ShouldEqual, value2.Name)

			})

		})
	})

}

func TestDataTrunMulti(t *testing.T) {
	convey.Convey("测试DataTrunMulti", t, func() {
		convey.Convey("测试指针情况:", func() {
			data := make([]*User, 0, 10)
			data = append(data, newUser("name-00", "pass-00", 0))
			data = append(data, newUser("name-00", "pass-01", 1))
			data = append(data, newUser("name-00", "pass-02", 2))

			data = append(data, newUser("name-01", "pass-01", 0))
			data = append(data, newUser("name-01", "pass-02", 1))
			data = append(data, newUser("name-01", "pass-03", 2))
			data = append(data, newUser("name-01", "pass-04", 3))
			arr, err := DataTrunMulti(data, "Name")
			convey.So(err, convey.ShouldBeNil)
			convey.So(3, convey.ShouldEqual, len(arr["name-00"]))
			convey.So(4, convey.ShouldEqual, len(arr["name-01"]))
		})

	})
}
func TestComputation(t *testing.T) {
	convey.Convey("测试Computation", t, func() {

		data := make([]*User, 0, 10)
		data = append(data, newUser("name-00", "pass-00", 0))
		data = append(data, newUser("name-00", "pass-01", 1))
		data = append(data, newUser("name-00", "pass-02", 2))

		data = append(data, newUser("name-01", "pass-01", 0))
		data = append(data, newUser("name-01", "pass-02", 1))
		data = append(data, newUser("name-01", "pass-03", 2))
		data = append(data, newUser("name-01", "pass-04", 3))
		vals, err := Computation(data, "Name", "Pass")
		convey.So(err, convey.ShouldBeNil)
		var v string
		var ok bool
		v, ok = vals["name-00"]
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(v, convey.ShouldEqual, "pass-02") //只有数组才固定顺序
		v, ok = vals["name-01"]
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(v, convey.ShouldEqual, "pass-04") //只有数组才固定顺序 , 值应该为最后一个
	})
}

//去重
func sliceStringUnique(slice []string) []string {
	sort.Strings(slice)
	i := 0
	var j int
	for {
		if i >= len(slice)-1 {
			break
		}
		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {
		}
		slice = append(slice[:i+1], slice[j:]...)
		i++
	}
	return slice

}

//检测slice是否相等
func checkStringSliceEqual(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
