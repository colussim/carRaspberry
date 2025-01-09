package main

import (
    "flag"
    "fmt"
    "log"

    "time"
    "periph.io/x/conn/v3/gpio"
    "periph.io/x/conn/v3/gpio/gpioreg"
    "periph.io/x/host/v3"
)

func main() {
    // Initialiser periph
    if _, err := host.Init(); err != nil {
        log.Fatalf("Erreur d'initialisation de periph: %v", err)
    }

    // Configuration du GPIO
    pin := gpioreg.ByName("P1_18") // Remplacez par le nom correct de la broche que vous utilisez
    if pin == nil {
        log.Fatal("Erreur : la broche GPIO P1_18 n'est pas disponible.")
    }

    // Configurer le GPIO en sortie
    if err := pin.Out(gpio.Low); err != nil {
        log.Fatalf("Erreur lors de la configuration de la broche GPIO : %v", err)
    }

    // Lecture de la commande
    command := flag.String("command", "", "start or stop")
    flag.Parse()

    // Exécuter la commande appropriée
    switch *command {
    case "start":
        pin.Out(gpio.High) // Ouvrir le MOSFET
        fmt.Println("Amplificateur démarré.")
    case "stop":
        pin.Out(gpio.Low) // Fermer le MOSFET
        fmt.Println("Amplificateur arrêté.")
    default:
        fmt.Println("Erreur : utilisez 'start' ou 'stop'.")
    }
}
