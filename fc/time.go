package fc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/metakeule/fmtdate"
)

// 返回微秒 , unix
func Microtime() (float64, int64) {
	t := time.Now()
	return float64(t.Nanosecond()) / 1000 / 1000 / 1000, t.Unix()
}

func FormatDateString(unix string) (string, error) {
	if unix == "" {
		return "", fmt.Errorf("unix string should not be ''")
	}
	v, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		return "", err
	}
	return FormatDateInt64(v)
}

func FormatDateInt64(v int64) (string, error) {
	if v <= 0 {
		return "", fmt.Errorf("unix should not be '0'")
	}
	t := time.Unix(v+8*3600, 0)
	return fmtdate.Format(fmtdate.DefaultDateTimeFormat, t), nil
}

/*
示例:
	t, _ = Strtotime("-1 day", time.Now())
	t, _ = Strtotime("-1       day", time.Now())
	t, _ = Strtotime(" -1  month ", time.Now())
	t, _ = Strtotime(" -1  year ", time.Now())
*/
func Strtotime(format string, t time.Time) (T time.Time, err error) {
	format = strings.TrimSpace(format)
	arr := strings.Fields(format)
	var year, month, day, num int
	num, err = strconv.Atoi(arr[0])
	if err != nil {
		return
	}

	switch arr[1] {
	case "day":
		day = num
	case "year":
		year = num
	case "month":
		month = num
	default:
		err = errors.New("format parse faild , you can use: day|month|year")
		return
	}
	T = t.AddDate(year, month, day)
	return
}

//示例:
//" 2018-03-01 "
//"2018-03-01 14:57:51"
func Fstrtotime(dateStr string) (T time.Time, err error) {
    dateStr = strings.TrimSpace(dateStr)
    fieldarr := strings.Fields(dateStr)
    if len(fieldarr) > 2 {
        err = fmt.Errorf("format dateStr is error")
        return
    }
    if len(fieldarr) == 1 {
        return fmtdate.ParseDate(dateStr)
    }

    if len(fieldarr) == 2 {
        return parseTimeWithLocation(dateStr)
    }
    return
}

// 修复时区问题
func parseTimeWithLocation(s string) (time.Time, error) {
    loc, err := time.LoadLocation("Asia/Shanghai")
    if err != nil {
        return time.Time{}, err
    }
    return time.ParseInLocation("2006-01-02 15:04:05", s, loc)
}


//包含end 天
func BuildTimeInterval(start, end time.Time) []string {
	startUnix := start.Unix()
	endUnix := end.Unix()
	length := (endUnix - startUnix) / 86400 //返回的是整数
	ret := make([]string, 0, length+1)      //只分配一次内存
	for i := startUnix; i < endUnix; i += 86400 {
		tmpday := fmtdate.Format(fmtdate.DefaultDateFormat, time.Unix(i, 0))
		ret = append(ret, tmpday)
	}
	return ret
}

//相减获取对应的time.Time
func GetBeforeDayTimer(day int64) time.Time {
	if day > 0 {
		day *= -1
	}
	return GetDayTimer(day)
}

//获取对应的天数
func GetDayTimer(day int64) time.Time {
	tunix := time.Now().Unix() + day*86400
	res := fmtdate.Format(fmtdate.DefaultDateFormat, time.Unix(tunix, 0))
	t, err := fmtdate.ParseDate(res)
	if err != nil {
		panic(err)
	}
	return t
}

//返回当天 00:00:00 对应的time.Time
func GetTodayTimer() time.Time {
	t, err := fmtdate.ParseDate(fmtdate.Format(fmtdate.DefaultDateFormat, time.Now()))
	if err != nil {
		panic(err)
	}
	return t
}

//获取昨天的Time
func GetYesterdayTimer() time.Time {
	return GetBeforeDayTimer(-1)
}
