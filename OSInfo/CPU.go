package OSInfo

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CPU struct{}

var (
	cpuInfo cpu.InfoStat
)

func init() {
	cpuInfos, _ := cpu.Info()

	if len(cpuInfos) < 0 {
		cpuInfos = nil
	}

	cpuInfo = cpuInfos[0]
}

//cpu频率，单位：MHz
func (c *CPU) Total() float64 {
	return cpuInfo.Mhz
}

//cpu内核数量，单位：个
func (c *CPU) Cores() int32 {
	return cpuInfo.Cores
}

//cpu占用率，单位：百分比(%)
func (c *CPU) Usage() float64 {
	usages, err := cpu.Percent(time.Second, false)

	if len(usages) <= 0 || err != nil {
		return -1
	}

	return usages[0]
}
