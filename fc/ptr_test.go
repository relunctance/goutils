package fc

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type User struct {
	Name string
	pass string
}

func (u *User) GetName() string {
	return u.Name
}
func (u *User) getpass() string {
	return u.pass
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

//
func TestMethodExists(t *testing.T) {
	convey.Convey("测试检测方法是否存在MethodExists():", t, func() {
		user1 := &User{Name: "hello", pass: "abc"}
		convey.So(MethodExists("GetName", user1), convey.ShouldBeTrue)
		convey.So(MethodExists("getpass", user1), convey.ShouldBeFalse)
		convey.So(MethodExists("Getname", user1), convey.ShouldBeFalse)
		convey.So(MethodExists("a", user1), convey.ShouldBeFalse)
		user2 := User{Name: "hello", pass: "abc"}
		convey.So(func() {
			MethodExists("a", user2)
		}, convey.ShouldPanic)
	})
}

func TestFieldExists(t *testing.T) {
	convey.Convey("测试检测方法是否存在FieldExists():", t, func() {
		user1 := &User{Name: "hello", pass: "abc"}
		convey.So(FieldExists("Name", user1), convey.ShouldBeTrue)
		convey.So(FieldExists("pass", user1), convey.ShouldBeTrue)
		convey.So(FieldExists("a", user1), convey.ShouldBeFalse)
		user2 := User{Name: "hello", pass: "abc"}
		convey.So(func() {
			FieldExists("a", user2)
		}, convey.ShouldPanic)
	})
}
