package OSInfo

import (
	"github.com/shirou/gopsutil/disk"
)

type Disk struct{}

var (
	hd *disk.UsageStat
)

func init() {
	hd, _ = disk.Usage("/")
}

//硬盘容量，单位：字节（byte）
func (d *Disk) Total() uint64 {
	return hd.Total
}

//硬盘占用率，单位：百分比（%）
func (d *Disk) Usage() float64 {
	return hd.UsedPercent
}
