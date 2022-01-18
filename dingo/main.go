package main

// constructors are forced to return an error, even if they don't error

import (
	"fmt"
	"log"
	"time"

	"github.com/dexterorion/dependency-injection-benchmark/dingo/deps/dic"
	"github.com/dexterorion/dependency-injection-benchmark/shared/benchmark"
)

func main() {
	var container *dic.Container
	var err error

	benchmark.LogTimeAndMem("Starting container", func() {
		container, err = dic.NewContainer()
		if err != nil {
			log.Fatal("error starting container: " + err.Error())
		}
	})

	ticker := time.NewTicker(time.Millisecond * 500)
	counter := 0

	for {
		<-ticker.C

		benchmark.LogTimeAndMem("Getting deps", func() {
			// If the object does not already exist, it is created and saved in the Container.
			// If the object can not be created, it returns an error.
			fakeservice := container.GetFakeservice()
			if fakeservice == nil {
				log.Fatal("error getting fakeservice")
			} else {
				fmt.Println(fakeservice)
			}
		})

		counter++
		if counter == 20 {
			return
		}
	}
}
