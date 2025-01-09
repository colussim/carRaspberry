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
