package jsondel

import (
	"fmt"
	"strings"

	sj "github.com/guyannanfei25/go-simplejson"
	"github.com/relunctance/goutils/fc"
)

func JsonDeleteString(json string, paths []string) (string, error) {
	j, err := sj.NewJson([]byte(json))
	if err != nil {
		return "", err
	}
	j = JsonDelete(j, paths)
	return fc.JsonDecode(j), nil
}

func JsonDeleteBytes(json []byte, paths []string) ([]byte, error) {
	j, err := sj.NewJson(json)
	if err != nil {
		return nil, err
	}
	j = JsonDelete(j, paths)
	return j.MarshalJSON()
}

// 对出唯一输出
func JsonDelete(j *sj.Json, paths []string) *sj.Json {
	v := newVjson(j, paths)
	v.run()
	return v.json()
}

func newVjson(j *sj.Json, paths []string) *vjson {
	vj := &vjson{
		j: j,
	}
	paths = fc.SliceStringUnique(paths)
	vj.fs = vj.buildFs(paths)
	return vj
}

type vjson struct {
	j     *sj.Json
	fs    [][]string
	cfs   []string
	iterJ *sj.Json
}

func (v *vjson) nextFs(i int, paths []string) (ret []string) {
	if i >= len(paths) {
		return

	}
	return paths[i:]
}

func (v *vjson) isEndFname(fname string) bool {
	return v.cfs[len(v.cfs)-1] == fname
}

func (v *vjson) run() {
	for _, fs := range v.fs {
		v.cfs = fs
		v.unset(v.j, fs)
	}
}

func (v *vjson) unset(j *sj.Json, paths []string) error {
	for i, fname := range paths {
		offset := i + 1
		nextfs := v.nextFs(offset, paths)
		if v.isEndFname(fname) {
			j.Del(fname) // 删除
			break
		}
		switch fname {
		case "*":
			vmap, err := j.Map()
			if err != nil {
				return fmt.Errorf("path is set '*' map pos error ")
			}
			for key, _ := range vmap {
				v.iterJ = j.Get(key)
				v.unset(v.iterJ, nextfs)
			}

		case "#":
			vslice, err := j.JsonArray()
			if err != nil {
				return fmt.Errorf("path is set '#' slice error ")
			}
			for _, nextJ := range vslice {
				v.iterJ = nextJ
				v.unset(v.iterJ, nextfs)
			}

		default:
			v.iterJ = j.Get(fname)
			v.unset(v.iterJ, nextfs)
		}

	}

	return nil
}

func (v *vjson) json() *sj.Json {
	return v.j
}

func (v *vjson) buildFs(paths []string) (ret [][]string) {
	ret = make([][]string, 0, len(paths))
	for _, path := range paths {
		path = strings.TrimSpace(path)
		if len(path) == 0 {
			fmt.Println("111")
			continue
		}
		fs := splitComma(path)
		if err := v.checkLast(fs[len(fs)-1]); err != nil {
			panic(err)
		}
		ret = append(ret, fs)

	}
	return
}

func (v *vjson) checkLast(val string) error {
	if val == "*" || val == "#" {
		return fmt.Errorf("last char can not be '%s'", v)
	}
	return nil
}

func splitComma(path string) []string {
	path = strings.TrimSpace(path)
	ret := make([]string, 0, 3)
	if strings.Index(path, "'") != -1 {
		arr := strings.Split(path, "'")
		for _, v := range arr {
			if v == "" {
				continue
			}
			if v == "." {
				continue
			}
			if v[0] == '.' || v[len(v)-1] == '.' {
				ret = append(ret, strings.Split(strings.Trim(v, "."), ".")...)
			} else {
				ret = append(ret, v)
			}
		}
	} else {
		ret = strings.Split(path, ".")
	}
	return ret
}
