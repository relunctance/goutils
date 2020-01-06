package fc

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

// URLEncode urlencode()
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

// URLDecode urldecode()
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

/*
	"'ipinfo'.*.info.'city'":                      4,
	"'ipinfo'.*.info.city":                        4,
	"ipinfo.*.info.city":                          4,
	"ipinfo.'*'.info.city":                        4,
	"'1234.23.4.2'.ipinfo.'1.0.0.1001'.info.city": 5,
	"'ipinfo'.'1.0.0.1001'.info.name.val.'city'":  6,
	"ipinfo.'1.0.0.1001'.info.city.'a.b.c'":       5,

*/
func SplitComma(path string) []string {
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

// 字符串两个切割符号支持
// 使用示例: SplitByChar(str , "&&" , "||");
func SplitByChar(smart, spChar1, spChar2 string) (data []string) {
	arr := strings.Split(smart, spChar1)
	for i, val := range arr {
		val = strings.TrimSpace(val)
		temparr := strings.Split(val, spChar2)
		if len(temparr) == 1 {
			data = append(data, val)
			if i < len(arr)-1 {
				data = append(data, spChar1)
			}
			continue
		}
		for j, v := range temparr {
			v = strings.TrimSpace(v)
			data = append(data, v)
			if j < len(temparr)-1 {
				data = append(data, spChar2)
			}
		}
		if i < len(arr)-1 {
			data = append(data, spChar1)
		}
	}
	return

}

// unicode转码为中文
func UnicodeToString(s string) string {
	var pos int
	pos = strings.Index(s, "\\u")
	if pos == -1 {
		return s
	}

	var ret string
	ret += s[0:pos]
	if pos+6 > len(s) {
		return s
	}
	ret += unicodeChar(s[pos+2 : pos+6])
	ret += s[pos+6:]
	return UnicodeToString(ret)
}

func unicodeChar(s string) string {
	temp, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%c", temp)
}

// 字符串转码为unicode
func StringToUnicode2(s string) string {
	rs := []rune(s)
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return json
}

func Chr(c rune) string {
	return string(c)
}

func Ord(c string) []rune {
	return []rune(c)
}

func Substr(str string, begin, length int) string {
	lth := len(str)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	return string(str[begin:end])
}

func StrToInt32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("\"%s\" is not integer", str)
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0, fmt.Errorf("%d is not 32-bit integer", num)
	}
	return int32(num), nil
}

/* eg:
u := ParseStr("?username=zhangsan&passwd=123456")
u.Get("username");
*/
func ParseStr(str string) url.Values {
	if strings.Index(str, "?") == -1 {
		panic(fmt.Errorf("must exists '?' like: '?username=zhangsan&passwd=123456'"))
	}
	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}
	return u.Query()
}

//首字母大写
func Ucfirst(str string) string {
	if len(str) == 0 {
		return str
	}
	if len(str) == 1 {
		return strings.ToUpper(str)
	}
	return string(bytes.ToUpper([]byte{str[0]})) + str[1:]
}

// md5 计算
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Md5Bytes(v []byte) string {
	return fmt.Sprintf("%x", md5.Sum(v))
}

func ByteFormat(i float64) string {
	var a = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB", "UnKnown"}
	var pos int = 0
	var j float64 = float64(i)
	for {
		if i >= 1024 {
			i = i / 1024
			j = j / 1024
			pos++
		} else {
			break
		}
	}
	if pos >= len(a) { // fixed out index bug
		pos = len(a) - 1
	}
	return fmt.Sprintf("%.3f %s", j, a[pos])
}

// 统计字符串出现的次数
func SubstrCount(s, v string) int {
	arr := strings.Split(s, v)
	return len(arr) - 1

}

func IsIp(s string) bool {
	arr := strings.Split(s, ".")
	if len(arr) != 4 {
		return false
	}

	for i := 0; i < 4; i++ {
		if len(arr[i]) > 3 {
			return false
		}
		v, err := strconv.Atoi(arr[i])
		if err != nil {
			return false
		}
		if v >= 0 && v <= 255 {
			continue
		} else {
			return false
		}
	}
	return true

}

// 判断是否是域名
// 是返回true
// 不是返回false
func IsDomain(s string) bool {
	pos := strings.LastIndex(s, ".")
	if pos == -1 {
		return false
	}
	if pos == len(s)-1 {
		return false
	}
	v := s[pos+1:]
	switch v {
	case
		"ac",
		"ad",
		"ae",
		"af",
		"ag",
		"ai",
		"al",
		"am",
		"an",
		"ao",
		"aq",
		"ar",
		"arpa",
		"art",
		"as",
		"at",
		"aw",
		"ax",
		"az",
		"ba",
		"bb",
		"bd",
		"be",
		"beer",
		"bf",
		"bg",
		"bh",
		"bi",
		"biz",
		"bj",
		"bl",
		"bm",
		"bn",
		"bo",
		"bq",
		"br",
		"bs",
		"bt",
		"bv",
		"bw",
		"by",
		"bz",
		"ca",
		"cc",
		"cd",
		"cf",
		"cg",
		"ch",
		"ci",
		"ck",
		"cl",
		"club",
		"cm",
		"cn",
		"co",
		"com",
		"cr",
		"cu",
		"cv",
		"cw",
		"cx",
		"cy",
		"cz",
		"de",
		"design",
		"dj",
		"dk",
		"dm",
		"do",
		"dz",
		"ec",
		"edu",
		"ee",
		"eg",
		"eh",
		"er",
		"es",
		"et",
		"eu",
		"fi",
		"fj",
		"fk",
		"fm",
		"fo",
		"fr",
		"fun",
		"ga",
		"gb",
		"gd",
		"ge",
		"gf",
		"gg",
		"gh",
		"gi",
		"gm",
		"gn",
		"gov",
		"gp",
		"gq",
		"gr",
		"group",
		"gs",
		"gt",
		"gu",
		"gw",
		"gy",
		"hk",
		"hm",
		"hn",
		"hr",
		"ht",
		"hu",
		"id",
		"ie",
		"il",
		"im",
		"in",
		"info",
		"ink",
		"int",
		"io",
		"iq",
		"ir",
		"is",
		"it",
		"je",
		"jm",
		"jo",
		"jp",
		"ke",
		"kg",
		"kh",
		"ki",
		"kim",
		"km",
		"kn",
		"kp",
		"kr",
		"kw",
		"ky",
		"kz",
		"la",
		"lb",
		"lc",
		"li",
		"link",
		"live",
		"lk",
		"lr",
		"ls",
		"lt",
		"ltd",
		"lu",
		"luxe",
		"lv",
		"ly",
		"ma",
		"mc",
		"md",
		"me",
		"mf",
		"mg",
		"mh",
		"mil",
		"mk",
		"ml",
		"mm",
		"mn",
		"mo",
		"mobi",
		"mp",
		"mq",
		"mr",
		"ms",
		"mt",
		"mu",
		"mv",
		"mw",
		"mx",
		"my",
		"mz",
		"na",
		"name",
		"nc",
		"ne",
		"net",
		"nf",
		"ng",
		"ni",
		"nl",
		"no",
		"np",
		"nr",
		"nu",
		"nz",
		"om",
		"online",
		"org",
		"pa",
		"pe",
		"pf",
		"pg",
		"ph",
		"pk",
		"pl",
		"pm",
		"pn",
		"pr",
		"pro",
		"ps",
		"pt",
		"pub",
		"pw",
		"py",
		"qa",
		"re",
		"red",
		"ren",
		"ro",
		"rs",
		"ru",
		"rw",
		"sa",
		"sb",
		"sc",
		"sd",
		"se",
		"sg",
		"sh",
		"shop",
		"si",
		"site",
		"sj",
		"sk",
		"sl",
		"sm",
		"sn",
		"so",
		"sr",
		"ss",
		"st",
		"store",
		"su",
		"sv",
		"sx",
		"sy",
		"sz",
		"tc",
		"td",
		"tech",
		"tf",
		"tg",
		"th",
		"tj",
		"tk",
		"tl",
		"tm",
		"tn",
		"to",
		"top",
		"tp",
		"tr",
		"tt",
		"tv",
		"tw",
		"tz",
		"ua",
		"ug",
		"uk",
		"um",
		"us",
		"uy",
		"uz",
		"va",
		"vc",
		"ve",
		"vg",
		"vi",
		"vip",
		"vn",
		"vu",
		"wang",
		"wf",
		"wiki",
		"work",
		"ws",
		"xin",
		"xyz",
		"ye",
		"yt",
		"za",
		"zm",
		"zw",
		"餐厅",
		"公司",
		"集团",
		"商标",
		"网店",
		"网络",
		"网址",
		"我爱你",
		"在线",
		"中国",
		"中文网":
		return true
	}
	return false

}

// 检测MD5
func CheckMd5(v string) bool {
	return CheckHexAndLength(v, 32)
}

// 检测sha1
func CheckSha1(v string) bool {
	return CheckHexAndLength(v, 40)
}

// 检测sha256
func CheckSha256(v string) bool {
	return CheckHexAndLength(v, 64)
}
func CheckHexValid(vs string) bool {
	for _, v := range vs {
		if v >= '0' && v <= 'f' {
			continue
		}
		return false
	}
	return true
}

//
func CheckHexAndLength(v string, l int) bool {
	if len(v) != l {
		return false
	}
	if !CheckHexValid(v) {
		return false
	}
	return true
}

// 解析domain
func ParseHostFromUrl(address string) (string, error) {
	u, err := url.Parse(address)
	if err != nil { // 日志中有可能存在
		return "", err
	}
	return u.Hostname(), nil
}
