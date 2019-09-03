package fc

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
		ret = fmt.Sprintf("Memory Total: %v, Free:%v, Cached:%v , FreeAll:%v , Used: %v , Cost Mem:%v  UsedPercent:%f%%\n",
			ByteFormat(float64(m.Total)),
			ByteFormat(float64(m.Free)),
			ByteFormat(float64(m.Cached)),
			ByteFormat(float64(m.Free+m.Cached)),
			ByteFormat(float64(m.Used)),
			memCost(),
			m.UsedPercent,
		)
	} else {
		ret = fmt.Sprintf("Memory Total: %v, Free:%v, Cached:%v , FreeAll:%v , Used: %v ,  UsedPercent:%f%%\n",
			ByteFormat(float64(m.Total)),
			ByteFormat(float64(m.Free)),
			ByteFormat(float64(m.Cached)),
			ByteFormat(float64(m.Free+m.Cached)),
			ByteFormat(float64(m.Used)),
			m.UsedPercent,
		)
	}
	return
}

func memCost() string {
	start := memInfo[0]
	end := memInfo[len(memInfo)-1]
	return ByteFormat(float64(end.Used - start.Used))
}
