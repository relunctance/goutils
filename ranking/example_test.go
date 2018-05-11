package ranking

import (
	"fmt"
	"sort"
)

var _ ItemWeigher = new(User) // 检测接口是否继承

type User struct {
	Name         string
	Group        string
	EnglishScore int
	ChineseScore int
}

func (u *User) GroupKey() string {
	return u.Group
}

// expect max score user
func (u *User) WeigthCalc(item ItemWeigher) bool {
	itemUser := item.(*User)
	uScore := u.ChineseScore + u.EnglishScore
	itemScore := itemUser.ChineseScore + itemUser.EnglishScore
	return uScore > itemScore
}

func ExampleWeigth() {
	data := []ItemWeigher{
		&User{"xiaoming1", "class1", 80, 61}, // 141
		&User{"xiaoming2", "class1", 33, 51}, // 84
		&User{"xiaoming3", "class1", 49, 73}, // 122
		&User{"xiaoming4", "class1", 99, 87}, // 186

		&User{"xiaoming5", "class2", 44, 38}, // 82
		&User{"xiaoming6", "class2", 56, 12}, // 68
		&User{"xiaoming7", "class2", 88, 75}, // 163
		&User{"xiaoming8", "class2", 45, 67}, // 112

		&User{"xiaoming9", "class3", 80, 100},  // 180
		&User{"xiaoming10", "class3", 100, 33}, // 133

		&User{"xiaoming11", "class4", 22, 17}, // 39
		&User{"xiaoming12", "class4", 99, 31}, // 130
	}
	values := make([]string, 0, 4)
	Weigth(data, func(item ItemWeigher) {
		u := item.(*User)
		values = append(values, fmt.Sprintf("name: %s , score: %d", u.Name, u.ChineseScore+u.EnglishScore))
	})
	sort.Strings(values) //排序
	for _, val := range values {
		fmt.Println(val)
	}
	// Output:
	// name: xiaoming12 , score: 130
	// name: xiaoming4 , score: 186
	// name: xiaoming7 , score: 163
	// name: xiaoming9 , score: 180
}
