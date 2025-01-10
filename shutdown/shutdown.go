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
	// Set up the Raspberry Pi adaptor
	r := raspi.NewAdaptor()

	// Create a GPIO button driver for the button connected to GPIO 3
	button := gpio.NewButtonDriver(r, "P1_5")

	// Create the Gobot robot
	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
	)

	// When the button is pressed, execute shutdown
	button.On(gpio.ButtonPush, func(data interface{}) {
		fmt.Println("Button pressed. Initiating shutdown...")

		// Set a timeout if you wish to delay
		shutdown.SetTimeout(time.Second * 4)
		shutdown.Shutdown()

		// Wait for shutdown handlers to complete
		shutdown.Wait()
		fmt.Println("Shutdown completed.")
		os.Exit(0) // Clean exit of the application
	})

	// Manage interrupt signals for clean exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Interrupt received! Shutting down...")
		shutdown.Shutdown()
		shutdown.Wait()
		os.Exit(0)
	}()

	// Print startup message and start the robot
	fmt.Println("Starting the robot... Press the button to shutdown.")
	robot.Start()
}
