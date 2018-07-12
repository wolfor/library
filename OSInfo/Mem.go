package OSInfo

import (
	"github.com/shirou/gopsutil/mem"
)

type Mem struct{}

var (
	vm *mem.VirtualMemoryStat
)

func init() {
	vm, _ = mem.VirtualMemory()
}

//内存总量，单位：字节(byte)
func (m *Mem) Total() uint64 {
	return vm.Total
}

//内存占用率，单位：百分比(%)
func (m *Mem) Usage() float64 {
	return vm.UsedPercent
}
