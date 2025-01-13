package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

const GPIO_PIN_NAME = "P1_5"

func main() {
	logger := logrus.New()

	// Initialize device
	if _, err := host.Init(); err != nil {
		logger.Fatalf("\n❌Device initialization error: %v\n", err)
		os.Exit(1)
	}

	// OGet GPIO 3 reference
	pin := gpioreg.ByName(GPIO_PIN_NAME)

	// Configure GPIO as input with pull-up resistor
	if err := pin.In(gpio.PullUp, gpio.FallingEdge); err != nil {
		logger.Fatalf("\n❌Pin configuration error: %v\n", err)
		os.Exit(1)
	}

	logger.Infof("Waiting for the button to be pressed...")

	// Loop to monitor button status
	for {
		// Wait for the button to be pressed
		if pin.Read() == gpio.Low {
			logger.Infof("✅ Button pressed, stop in progress...")
			// Execute stop command
			exec.Command("sudo", "halt").Run()
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
}
