package main

// constructors are forced to return an error, even if they don't error
// it looks like it is used only to start application, not do return the deps

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dexterorion/dependency-injection-benchmark/shared/benchmark"
	"github.com/dexterorion/dependency-injection-benchmark/shared/services"
	"github.com/dexterorion/dependency-injection-benchmark/shared/storage"
	"go.uber.org/fx"
)

func main() {
	var app *fx.App

	benchmark.LogTimeAndMem("Starting container", func() {
		app = fx.New(
			fx.Provide(
				services.NewFakeService,
				storage.NewFakeInMem,
				// storage.NewFakeLazy, // cannot do that (two with same response structure)
			),
			fx.Invoke(func(fs services.FakeService) {
				fmt.Println(fs)
			}),
		)
	})

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	benchmark.LogTimeAndMem("Running container initialization", func() {
		if err := app.Start(startCtx); err != nil {
			log.Fatal(err)
		}
	})
}
