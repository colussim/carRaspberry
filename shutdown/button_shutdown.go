package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

const GPIO_PIN_NAME = "P1_5"

func main() {
	// Initialiser le périphérique
	if _, err := host.Init(); err != nil {
		fmt.Printf("Erreur d'initialisation du périphérique: %v\n", err)
		os.Exit(1)
	}

	// Obtenir la référence au GPIO 3
	pin := gpioreg.ByName(GPIO_PIN_NAME)

	// Configurer le GPIO comme entrée avec une résistance de tirage vers le haut
	if err := pin.In(gpio.PullUp, gpio.FallingEdge); err != nil {
		fmt.Printf("Erreur lors de la configuration du pin: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("En attente de l'appui sur le bouton...")

	// Boucle pour surveiller l'état du bouton
	for {
		// Attendre que le bouton soit pressé
		if pin.Read() == gpio.Low {
			fmt.Println("Bouton pressé, arrêt en cours...")
			// Exécuter la commande d’arrêt
			exec.Command("sudo", "halt").Run()
			break
		}
		// Pause pour éviter de saturer le CPU
		time.Sleep(100 * time.Millisecond)
	}
}
