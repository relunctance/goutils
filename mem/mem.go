package mem

import (
	"fmt"

	"github.com/relunctance/goutils/str"
	"github.com/shirou/gopsutil/mem"
)

type MemDefine []*mem.VirtualMemoryStat

var memInfo MemDefine = make(MemDefine, 0)

//统计mem
func MemInfo() (ret string) {
	m, _ := mem.VirtualMemory()
	memInfo = append(memInfo, m)
	if len(memInfo) > 1 {
		ret = fmt.Sprintf("Memory Total: %v, Free:%v, Cached:%v , FreeAll:%v , Used: %v , Cost Mem:%v  UsedPercent:%f%%\n", str.ByteFormat(m.Total), str.ByteFormat(m.Free), str.ByteFormat(m.Cached), str.ByteFormat(m.Free+m.Cached), str.ByteFormat(m.Used), memCost(), m.UsedPercent)
	} else {
		ret = fmt.Sprintf("Memory Total: %v, Free:%v, Cached:%v , FreeAll:%v , Used: %v ,  UsedPercent:%f%%\n", str.ByteFormat(m.Total), str.ByteFormat(m.Free), str.ByteFormat(m.Cached), str.ByteFormat(m.Free+m.Cached), str.ByteFormat(m.Used), m.UsedPercent)
	}
	return
}

func memCost() string {
	start := memInfo[0]
	end := memInfo[len(memInfo)-1]
	return str.ByteFormat(end.Used - start.Used)
}
