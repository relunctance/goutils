package mapkey

import "fmt"

const (
	mapkey_default_name = "default"
)

// 用于统计不重复的值,出现的次数
type MapKey struct {
	names []string
	uniq  map[string]map[string]struct{}
	// 重复的key
	duplicate map[string]map[string]struct{}
	num       map[string]int
}

func NewMapKey() *MapKey {
	names := []string{mapkey_default_name}
	return NewMapKeyByNames(names)
}
func NewMapKeyByNames(names []string) *MapKey {
	m := &MapKey{}
	m.names = names
	m.num = make(map[string]int, len(m.names))
	m.uniq = make(map[string]map[string]struct{}, len(m.names))
	m.duplicate = make(map[string]map[string]struct{}, len(m.names))
	for _, name := range m.names {
		m.uniq[name] = make(map[string]struct{}, 5)      // 初始化
		m.duplicate[name] = make(map[string]struct{}, 5) //初始化
		m.num[name] = 0                                  // 初始化
	}
	return m
}

func (m *MapKey) MapCount(v string) {
	m.MapCountByName(v, mapkey_default_name)
}

func (m *MapKey) MapCountByName(v, name string) error {
	uniqMap, ok := m.uniq[name]
	if !ok {
		return fmt.Errorf("not exists name:[%s]", name)
	}
	dupMap := m.duplicate[name]
	if _, ok := uniqMap[v]; !ok {
		uniqMap[v] = struct{}{}
		m.num[name] += 1
	} else {
		dupMap[v] = struct{}{}
	}
	return nil
}

// 获取slice的结果
func (m *MapKey) GetUniqDataSlice() (ret []string) {
	ret = make([]string, 0, 10)
	for _, vmap := range m.uniq {
		for value, _ := range vmap {
			ret = append(ret, value)
		}
	}
	return
}

func (m *MapKey) GetUniqData() map[string]map[string]struct{} {
	return m.uniq
}

func (m *MapKey) GetUniqDataByName(name string) map[string]struct{} {
	return m.uniq[name]
}

func (m *MapKey) GetUniqDataSliceByName(name string) (ret []string) {
	ret = make([]string, 0, 10)
	for value, _ := range m.uniq[name] {
		ret = append(ret, value)
	}
	return
}

func (m *MapKey) GetDuplicate() map[string]map[string]struct{} {
	return m.duplicate
}

func (m *MapKey) GetDuplicateByName(name string) map[string]struct{} {
	return m.duplicate[name]
}

func (m *MapKey) GetNum() map[string]int {
	return m.num
}
