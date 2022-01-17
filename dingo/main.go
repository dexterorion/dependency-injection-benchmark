package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dexterorion/dependency-injection-benchmark/dingo/deps/dic"
	"github.com/dexterorion/dependency-injection-benchmark/shared"
)

func main() {
	start := time.Now()

	container, err := dic.NewContainer()
	if err != nil {
		panic(err)
	}

	end := time.Now()
	diff := end.Sub(start).Microseconds()
	fmt.Printf("Total spent start container: %d micro seconds\n", diff)
	shared.ShowMemInfo()

	ticker := time.NewTicker(time.Millisecond * 500)
	counter := 0

	for {
		<-ticker.C

		bef := time.Now()

		// If the object does not already exist, it is created and saved in the Container.
		// If the object can not be created, it returns an error.
		dsclient := container.GetDsclient()
		if dsclient == nil {
			log.Fatal("error getting datastore")
		}
		fakeservice := container.GetFakeservice()
		if fakeservice == nil {
			log.Fatal("error getting fakeservice")
		}

		aft := time.Now()

		spent := aft.Sub(bef).Microseconds()
		fmt.Printf("Total spent: %d micro seconds\n", spent)
		shared.ShowMemInfo()

		counter++
		if counter == 20 {
			return
		}
	}
}
