package fc

import (
	"fmt"
	"sort"
	"strings"
)

// map where condition to string
func GetWhereCondition(kws map[string]interface{}) (s string) {
	if len(kws) == 0 {
		return
	}
	keys := make([]string, 0, len(kws))
	for key, _ := range kws { //map 是无序的
		keys = append(keys, key)
	}
	sort.Strings(keys) //必须排序
	op := "and "
	for _, key := range keys {
		v := kws[key]
		switch vt := v.(type) {
		case string:
			s += fmt.Sprintf("`%s`='%s' %s", key, strings.TrimSpace(vt), op)
		case []string:
			s += fmt.Sprintf("`%s` in ('"+strings.Join(vt, "','")+"') %s", key, op)
		case []interface{}:
			w, err := sliceWhereBuild(key, op, vt)
			if err != nil {
				panic(err)
			}
			s += w
		}
	}
	s = strings.TrimRight(s, op)
	return
}

func sliceWhereBuild(field, op string, v []interface{}) (where string, err error) {
	if len(v) == 0 {
		err = fmt.Errorf("can not find val in field: [%s]", field)
		return
	}
	slice := trunInterfaceToSlice(v)
	where = fmt.Sprintf("`%s` in ('%s') %s", field, strings.Join(slice, "','"), op)
	return where, nil
}

func trunInterfaceToSlice(data []interface{}) []string {
	ret := make([]string, 0, len(data))
	for _, val := range data {
		ret = append(ret, strings.TrimSpace(fmt.Sprintf("%v", val)))
	}
	return ret
}
