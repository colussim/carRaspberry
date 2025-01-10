package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

// Replace with the correct name of the pin you are using
const GPIO_PIN_NAME = "P1_16"
const MESSAGE_START = "✅ Amplifier started..."
const MESSAGE_STOP = "✅ Amplifier stopped..."

var logger *logrus.Logger

func main() {
	// Init periph
	if _, err := host.Init(); err != nil {
		logrus.Fatalf("\n❌ Peripheral initialization error: %w", err)
		os.Exit(1)
	}

	// Configuration GPIO
	pin := gpioreg.ByName(GPIO_PIN_NAME)
	if pin == nil {
		logrus.Fatalf("❌ Error : GPIO pin %s is not available.", GPIO_PIN_NAME)
		os.Exit(1)
	}

	// Configuring the output GPIO line
	if err := pin.Out(gpio.Low); err != nil {
		logrus.Fatalf("❌ GPIO pin configuration error : %v", err)
		os.Exit(1)
	}

	// Reading the command line
	command := flag.String("command", "", "start or stop")
	flag.Parse()

	// Run the appropriate command
	switch *command {
	case "start":
		pin.Out(gpio.High) // Open MOSFET
		logger.Infof(MESSAGE_START)
	case "stop":
		pin.Out(gpio.Low) // CloseMOSFET
		logger.Infof(MESSAGE_STOP)
	default:
		logger.Infof("❌ Erreur : Usage: amplipower -command=<start> || <stop>")
	}
}
