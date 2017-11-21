package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	InfoCPU()
	InfoMem()
	for {
	}
}

//InfoCPU Show CPU information
func InfoCPU() {
	v, _ := cpu.Info()

	// almost every return value is a struct
	//fmt.Printf("VendorID:%v, CPU Model: %v\n", v[0].VendorID, v[0].ModelName)
	for i, cpu := range v {

		fmt.Printf(`
##CPU %d
  vendor id : %v
  model name: %v
  cores     : %v
  stepping  : %v
  cpu MHz   : %v
  cache size: %v MB
`, i, cpu.VendorID, cpu.ModelName, cpu.Cores, cpu.Stepping, cpu.Mhz, cpu.CacheSize/1024)

	}

}

var kb uint64 = 1024
var mb = kb * 1024
var gb = mb * 1024

//InfoMem Show memory information
func InfoMem() {
	vmem, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()

	// almost every return value is a struct
	fmt.Printf(`
##Memory
  total : %v MB
  used  : %v MB
  free  : %v MB
  shared: %v MB
  cache : %v MB
  buff  : %v MB
  available    : %v MB
  used percent : %f%%
##Swap
  total : %v MB
  used  : %v MB
  free  : %v MB
  used percent : %f%%
`, vmem.Total/mb, vmem.Used/mb, vmem.Free/mb, vmem.Shared/mb,
		vmem.Cached/mb, vmem.Buffers/mb, vmem.Available/mb, vmem.UsedPercent,
		swap.Total/mb, swap.Used/mb, swap.Free/mb, swap.UsedPercent)
}
