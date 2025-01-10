package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	shutdown "github.com/klauspost/shutdown2"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	// Configure a new Raspberry Pi instance
	r := raspi.NewAdaptor()

	// Configure GPIO Pin 3 (P1_5) as an input with a pull-up resistor
	button := gpio.NewButtonDriver(r, "3")

	// Creating a Gobot robot
	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
	)

	// When the button is pressed, execute shutdown
	button.On(raspi.ButtonPush, func(data interface{}) {
		fmt.Println("Button pressed, initiating shutdown...")

		// Perform shutdown, signaling shutdown listeners
		shutdown.SetTimeout(time.Second * 4)
		shutdown.Shutdown()

		// Wait for all shutdown handlers to complete
		shutdown.Wait()

		fmt.Println("Shutdown completed.")
		os.Exit(0) // Exit the program after shutdown
	})

	// Manage interrupt signals for a clean stop
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Interrupt received, initiating shutdown...")

		// Signal the shutdown
		shutdown.Shutdown()

		// Wait for all shutdown handlers to complete
		shutdown.Wait()

		os.Exit(0) // Exit the program after shutdown
	}()

	// Start the robot
	robot.Start()
}
