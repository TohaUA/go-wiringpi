Forked from [eternal-flame-AD/go-wiringpi](https://github.com/eternal-flame-AD/go-wiringpi)

# go-wiringpi

[![GoDoc](https://godoc.org/github.com/TohaUA/go-wiringpi?status.svg)](https://godoc.org/github.com/TohaUA/go-wiringpi)

[WiringPi](http://wiringpi.com/) bindings for golang.

Tweaked for supporting Orange Pi 5 with [wiringOP](https://github.com/orangepi-xunlong/wiringOP)

## Features & TODOs

- [X] Init functions
- [X] Core GPIO Operations
- [X] Hardware PWM
- [ ] I2C
- [ ] SPI
- [ ] Software PWM

## Example

```golang

import "github.com/TohaUA/go-wiringpi"

func main() {
    gpio, err := wiringpi.Setup(wiringpi.WiringPiSetup)
    if err != nil {
        panic(err)
    }
    gpio.PinMode(1, wiringpi.In)
    if gpio.DigitalRead(1) == wiringpi.High {
        fmt.Println("WiringPi port 1 is at high")
    } else {
        fmt.Println("WiringPi port 1 is at low")
    }
}

```