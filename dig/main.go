package main

import (
	"fmt"
	"time"

	"github.com/dexterorion/dependency-injection-benchmark/shared/benchmark"
	"github.com/dexterorion/dependency-injection-benchmark/shared/services"
	"github.com/dexterorion/dependency-injection-benchmark/shared/storage"
	"go.uber.org/dig"
)

func main() {
	var container *dig.Container

	benchmark.LogTimeAndMem("Starting container", func() {
		container = dig.New()
		// container.Provide(services.NewFakeService) -> if we want to use in Invoke, cannot be interface
		container.Provide(services.NewFakeServiceStruct)
		container.Provide(storage.NewFakeInMem)
		container.Provide(storage.NewFakeLazy)
	})

	ticker := time.NewTicker(time.Millisecond * 500)
	counter := 0

	for {
		<-ticker.C

		benchmark.LogTimeAndMem("Invoking", func() {
			if err := container.Invoke(func(s *services.FakeServiceStruct) {
				fmt.Println(s)
			}); err != nil {
				fmt.Println(err)
			}
		})

		counter++
		if counter == 20 {
			return
		}
	}
}
