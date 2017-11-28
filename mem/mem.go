package mem

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type MemDefine []*mem.VirtualMemoryStat

var memInfo MemDefine = make(MemDefine, 0)

//统计mem
func MemInfo() (ret string) {
	m, _ := mem.VirtualMemory()
	memInfo = append(memInfo, m)
	if len(memInfo) > 1 {
		ret = fmt.Sprintf("Memory Total: %v, Free:%v, Cached:%v , FreeAll:%v , Used: %v , Cost Mem:%v  UsedPercent:%f%%\n", ByteFormat(m.Total), ByteFormat(m.Free), ByteFormat(m.Cached), ByteFormat(m.Free+m.Cached), ByteFormat(m.Used), memCost(), m.UsedPercent)
	} else {
		ret = fmt.Sprintf("Memory Total: %v, Free:%v, Cached:%v , FreeAll:%v , Used: %v ,  UsedPercent:%f%%\n", ByteFormat(m.Total), ByteFormat(m.Free), ByteFormat(m.Cached), ByteFormat(m.Free+m.Cached), ByteFormat(m.Used), m.UsedPercent)
	}
	return
}

func memCost() string {
	start := memInfo[0]
	end := memInfo[len(memInfo)-1]
	return ByteFormat(end.Used - start.Used)
}

func ByteFormat(i uint64) string {
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
	return fmt.Sprintf("%f %s", j, a[pos])
}
