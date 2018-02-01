package slicelimit

import (
	"testing"
	"utils"

	"github.com/smartystreets/goconvey/convey"
)

func TestCheckIsEndPanic(t *testing.T) {
	convey.Convey("检测CheckIsEnd是否报错", t, func() {
		convey.Convey("必须报错", func() {
			err, _, _ := GetBoundary(11, 5, 10)
			_, err2 := CheckIsEnd(err)
			convey.So(err2, convey.ShouldNotBeNil)
		})
	})
}

func TestOffset(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	offset := int64(10)
	pagesize := int64(5)
	l := int64(len(arr))
	err, start, end := GetBoundary(offset, pagesize, l)
	if err != EOF {
		t.Fatalf(" error must EOF")
	}
	if start != end {
		t.Fatalf("[start] should be eq [end]")
	}
	if hasNext, _ := CheckIsEnd(err); hasNext != true {
		t.Fatalf("should be is true")
	}

	err, start, end = GetBoundary(0, 3, l)
	if err != nil {
		t.Errorf("error should be not null")
	}
	if start != 0 {
		t.Fatalf("start should be eq 0")
	}
	if end != 3 {
		t.Fatalf("end should be eq 3")
	}

	err, start, end = GetBoundary(3, 5, l) //nil , 3, 8
	if err != nil {
		t.Errorf("error should be not null")
	}
	if start != 3 {
		t.Fatalf("start should be eq 3")
	}
	if end != 8 {
		t.Fatalf("end should be eq 8")
	}

	resultdata := []int{4, 5, 6, 7, 8}
	for _, val := range arr[start:end] {
		if !utils.InArrayInts(val, resultdata) {
			t.Fatalf("[%v] should be in %v\n", val, resultdata)
		}
	}
	if hasNext, _ := CheckIsEnd(err); hasNext != false {
		t.Errorf("error should be is false")
	}
	err, start, end = GetBoundary(3, 100, l) // nil , 3, 10
	if err != nil {
		t.Errorf("error should be not null")
	}
	if start != 3 {
		t.Fatalf("start should be eq 3")
	}
	if end != 10 {
		t.Fatalf("end should be eq 10")
	}

	err, start, end = GetBoundary(11, pagesize, l)
	if err == nil {
		t.Fatalf("must error")
	}
}
