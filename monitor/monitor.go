package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	n,_:=host.Info()
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	//cn,_:=net.Connections("all")
	//fmt.Println("cn:",cn)
	var cores int32
	for _, sub_cpu := range c {
		modelname := sub_cpu.ModelName
		cores = cores + sub_cpu.Cores
		fmt.Printf("        CPU       : %v   %v cores \n", modelname, cores)
	}
	nv,_:=net.IOCounters(false)
	fmt.Printf("        Network: %v bytes / %v bytes\n", nv[0].BytesRecv, nv[0].BytesSent)
	fmt.Printf("        SystemBoot:%v\n", n.BootTime)
	for _, p := range cc {
		fmt.Printf("        CPU Used    : used %f%% \n", p)
	}
	fmt.Printf("        MemTotal        : %v MB  Available: %v MB Usage:%f%%\n", v.Total/1024/1024, v.Available/1024/1024, v.UsedPercent)
	fmt.Printf("        Disk        : %v GB  Free: %v GB Usage:%f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
	fmt.Printf("        OS        : %v(%v)   %v  \n", n.Platform, n.PlatformFamily, n.PlatformVersion)
	fmt.Printf("        Hostname  : %v  \n", n.Hostname)

}
