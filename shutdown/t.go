package main

import (
	"fmt"
	"os"
	"time"

	shutdown "github.com/klauspost/shutdown2"
)

func main() {
	// Set up the Raspberry Pi adaptor

	fmt.Println("Button pressed. Initiating shutdown...")

	// Set a timeout if you wish to delay
	shutdown.SetTimeout(time.Second * 4)
	shutdown.Shutdown()

	// Wait for shutdown handlers to complete
	shutdown.Wait()
	fmt.Println("Shutdown completed.")
	os.Exit(0) // Clean exit of the application

}
