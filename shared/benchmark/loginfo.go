package benchmark

import (
	"fmt"
	"time"
)

func LogTimeAndMem(description string, method func()) {
	bef := time.Now()

	method()

	aft := time.Now()

	spent := aft.Sub(bef).Microseconds()
	fmt.Printf("%s - Total spent: %d micro seconds\n", description, spent)
	ShowMemInfo()
}
