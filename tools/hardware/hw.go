package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	InfoCPU()
	InfoMem()
	/*for {
	}*/
}

const padding = 1

//InfoCPU Show CPU information
func InfoCPU() {
	v, _ := cpu.Info()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	fmt.Fprintln(w, "\tprocessor\tcores\tcpu MHz\tcache size(MB)\tstepping\tvendor_id\tmodel name\t")

	// almost every return value is a struct
	for i, cpu := range v {
		fmt.Fprintln(w, fmt.Sprintf("\t%d\t%v\t%v\t%v\t%v\t%v\t%v\t", i, cpu.Cores, cpu.Mhz,
			cpu.CacheSize/1024, cpu.Stepping, cpu.VendorID, cpu.ModelName))
	}
	fmt.Printf("CPU information\nnumbers %v\n", len(v))
	w.Flush()
	fmt.Println()
}

var kb uint64 = 1024
var mb = kb * 1024
var gb = mb * 1024

//InfoMem Show memory information
func InfoMem() {
	vmem, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)
	fmt.Fprintln(w, " \ttotal\tused\tfree\tshared\tcache\tbuff\tavailable\tused percent\t")
	fmt.Fprintln(w, fmt.Sprintf("memory\t%v MB\t%v MB\t%v MB\t%v MB\t%v MB\t%v MB\t%v MB\t%f%%\t",
		vmem.Total/mb, vmem.Used/mb, vmem.Free/mb, vmem.Shared/mb, vmem.Cached/mb, vmem.Buffers/mb, vmem.Available/mb, vmem.UsedPercent))
	fmt.Fprintln(w, fmt.Sprintf("swap\t%v MB\t%v MB\t%v MB\t\t\t\t\t%f%%\t",
		swap.Total/mb, swap.Used/mb, swap.Free/mb, swap.UsedPercent))
	fmt.Printf("\nMemory information\n")
	w.Flush()
	// almost every return value is a struct
	/*fmt.Printf(`
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
			swap.Total/mb, swap.Used/mb, swap.Free/mb, swap.UsedPercent)*/
}
