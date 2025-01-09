package main

import (
    "flag"
    "fmt"
    "log"
    "periph.io/x/periph/conn/gpio"
    "periph.io/x/periph/conn/gpio/rpi"
    "periph.io/x/periph/devices/gpio"
    "periph.io/x/periph/host"
)

func main() {
    // Init periph
    if _, err := host.Init(); err != nil {
        log.Fatalf("Error d'init periph: %v", err)
    }

    // Configuration du GPIO
    pin := rpi.P1_18
    out := pin.Out(gpio.Low)

    command := flag.String("command", "", "start or stop")
    flag.Parse()

    switch *command {
    case "start":
        out.High() // open MOSFET
        fmt.Println("Amplificateur start.")
    case "stop":
        out.Low() // close MOSFET
        fmt.Println("Amplificator down.")
    default:
        fmt.Println("Erro use 'start' ou 'stop'.")
    }
}
