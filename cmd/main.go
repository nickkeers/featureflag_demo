package main

import (
	"context"
	"flaaaags/internal/bakery"
	"fmt"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/openfeature"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := openfeature.SetProvider(flagd.NewProvider())
	if err != nil {
		panic(err)
	}

	// Create a channel to signal when the provider is ready
	readyChan := make(chan struct{})
	var readyCallback = func(details openfeature.EventDetails) {
		fmt.Printf("Provider ready: %+v\n", details)
		close(readyChan)
	}
	openfeature.AddHandler(openfeature.ProviderReady, &readyCallback)

	// Setup to handle SIGTERM and SIGINT (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	client := openfeature.NewClient("flaaags")
	bakeryService := bakery.NewBakeryService(client)

	go func() {
		// Wait for the provider to be ready
		<-readyChan

		for {
			select {
			case <-signalChan:
				// Perform cleanup
				openfeature.Shutdown()

				fmt.Println("OpenFeature client was shutdown")

				os.Exit(0)
			default:
				// Call BakeCake every 2 seconds
				cake := bakeryService.BakeCake(context.Background())
				fmt.Printf("We baked using %s flour\n", cake.Flour)
				time.Sleep(2 * time.Second)
			}
		}
	}()

	// Prevent the main goroutine from exiting
	select {}
}
