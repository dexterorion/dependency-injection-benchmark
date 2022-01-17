package shared

import (
	"fmt"
	"runtime"
)

func ShowMemInfo() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("HeapAlloc = %v mb", toMb(m.HeapAlloc))
	fmt.Printf("\tHeapInuse = %v mb", toMb(m.HeapInuse))
	fmt.Printf("\tTotalAlloc = %v mb", toMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v mb", toMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func toMb(b uint64) uint64 {
	return b / 1024 / 1024
}
