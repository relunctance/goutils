package fc

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type UserDemo struct {
	Name string
	pass string
}

func (u *UserDemo) GetName() string {
	return u.Name
}
func (u *UserDemo) getpass() string {
	return u.pass
}

func TestGetPtrValueByName(t *testing.T) {
	user := &UserDemo{Name: "hello", pass: "abc"}
	v, err := ValueFromPtr("Name", user)
	if err != nil {
		t.Fatalf("should be nil , err:%#v\n", err)
	}
	switch val := v.(type) {
	case string:
		if val != "hello" {
			t.Fatalf("should be equal 'hello'")
		}
	default:
		t.Fatalf("should be string \n")
	}

	_, err = ValueFromPtr("Name2", user)
	if err == nil {
		t.Fatalf("should be not nil")
	}

	_, err = ValueFromPtr("pass", user)
	if err == nil {
		t.Fatalf("should be not nil")
	}

}
func TestIsPtr(t *testing.T) {
	convey.Convey("检测IsPtr()函数", t, func() {
		convey.So(IsPtr(1), convey.ShouldBeFalse)
		convey.So(IsPtr("abc"), convey.ShouldBeFalse)
		convey.So(IsPtr([]int{1, 2, 3}), convey.ShouldBeFalse)
		convey.So(IsPtr([2]int{1, 2}), convey.ShouldBeFalse)
		convey.So(IsPtr(map[string]string{"username": "hello"}), convey.ShouldBeFalse)
		convey.So(IsPtr(UserDemo{Name: "hello"}), convey.ShouldBeFalse)
		convey.So(IsPtr(&UserDemo{Name: "hello"}), convey.ShouldBeTrue)
	})
}

//
func TestMethodExists(t *testing.T) {
	convey.Convey("测试检测方法是否存在MethodExists():", t, func() {
		user1 := &UserDemo{Name: "hello", pass: "abc"}
		convey.So(MethodExists("GetName", user1), convey.ShouldBeTrue)
		convey.So(MethodExists("getpass", user1), convey.ShouldBeFalse)
		convey.So(MethodExists("Getname", user1), convey.ShouldBeFalse)
		convey.So(MethodExists("a", user1), convey.ShouldBeFalse)
		user2 := UserDemo{Name: "hello", pass: "abc"}
		convey.So(func() {
			MethodExists("a", user2)
		}, convey.ShouldPanic)
	})
}

func TestFieldExists(t *testing.T) {
	convey.Convey("测试检测方法是否存在FieldExists():", t, func() {
		user1 := &UserDemo{Name: "hello", pass: "abc"}
		convey.So(FieldExists("Name", user1), convey.ShouldBeTrue)
		convey.So(FieldExists("pass", user1), convey.ShouldBeTrue)
		convey.So(FieldExists("a", user1), convey.ShouldBeFalse)
		user2 := UserDemo{Name: "hello", pass: "abc"}
		convey.So(func() {
			FieldExists("a", user2)
		}, convey.ShouldPanic)
	})
}
