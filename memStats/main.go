package main

import (
	"fmt"
	"runtime"
)

func main() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	// memStats are up-to-date stats
	// heap stats on the other hand will be a snapshot of the last garbage collection

	fmt.Println(stats)

	fmt.Println("Total allocated heap objects:", stats.TotalAlloc) // never goes down, only goes up
	fmt.Println("Allocated heap objects:", stats.Alloc)
	fmt.Println("HeapSys:", stats.HeapSys)
	fmt.Println("HeapIdle:", stats.HeapIdle)
	fmt.Println("HeapInuse:", stats.HeapInuse)
	fmt.Println("HeapReleased:", stats.HeapReleased)
	fmt.Println("NumGC:", stats.NumGC)
	fmt.Println("PauseTotalNs:", stats.PauseTotalNs)
	fmt.Println("Virtual memory reserved from OS:", stats.Sys)             // in bytes
	fmt.Println("Virtual memory reserved from OS:", stats.Sys/(1024*1024)) // in MB
}
