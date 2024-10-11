package wiringpi

// #cgo LDFLAGS: -lwiringPi
/*
#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>

// Proxy function to call Go callback
extern void goCallbackProxy();
__attribute__((weak))
void proxyForGoCallback() {
   goCallbackProxy();
}
*/
import "C"
import "fmt"

var loaded = false

// PinMode enumerates available pin modes
type PinMode int

const (
	// Input set pin for input
	Input PinMode = C.INPUT
	// Output set pin for output
	Output PinMode = C.OUTPUT
	// PWMOutput set pin for PWM output
	PWMOutput PinMode = C.PWM_OUTPUT
	// GpioClock set pin for clock mode
	GpioClock PinMode = C.GPIO_CLOCK
)

// PullMode enumerates available pull modes
type PullMode int

const (
	// PullOff turns off pull up/down
	PullOff PullMode = C.PUD_OFF
	// PullDown pulls pin down
	PullDown PullMode = C.PUD_DOWN
	// PullUp pulls pin up
	PullUp PullMode = C.PUD_UP
)

// DigitalValue enumberates digital IO values
type DigitalValue int

const (
	// High pin is at high state
	High DigitalValue = C.HIGH
	// Low pin is at low state
	Low DigitalValue = C.LOW
)

// GPIO is the main handler for GPIO ports.
// use Setup function to get a *GPIO instead of creating directly
type GPIO struct {
	setup SetupMethod
}

// PinMode sets pin mode
func (c *GPIO) PinMode(pin int, mode PinMode) {
	C.pinMode(C.int(pin), C.int(mode))
}

// Pull sets pull up/down mode
func (c *GPIO) Pull(pin int, mode PullMode) {
	C.pullUpDnControl(C.int(pin), C.int(mode))
}

// DigitalWrite writes digital value to pin
func (c *GPIO) DigitalWrite(pin int, val DigitalValue) {
	C.digitalWrite(C.int(pin), C.int(val))
}

// PWMSetRange set PWM generator range
// defaults to 1024
func (c *GPIO) PWMSetRange(pin int, val uint) {
	C.pwmSetRange(C.int(pin), C.uint(val))
}

// PWMSetClock sets the divisor of PWM clock
func (c *GPIO) PWMSetClock(pin int, val int) {
	C.pwmSetClock(C.int(pin), C.int(val))
}

// PWMWrite writes pwn value
func (c *GPIO) PWMWrite(pin int, val int) {
	C.pwmWrite(C.int(pin), C.int(val))
}

// DigitalRead reads digital value
func (c *GPIO) DigitalRead(pin int) DigitalValue {
	return DigitalValue(C.digitalRead(C.int(pin)))
}

// goISRCallback is a global variable to store Go ISR callback
var goISRCallback func()

// goCallbackProxy is a proxy function to call Go callback
//
//export goCallbackProxy
func goCallbackProxy() {
	fmt.Printf("ISR callback\n")
	if goISRCallback != nil {
		goISRCallback()
		goISRCallback = nil
	}
}

// WiringPiISR sets up an interrupt service routine
func (c *GPIO) WiringPiISR(pin int, mode int, callback func()) {
	// Store Go callback in global variable
	goISRCallback = callback

	// Setup ISR
	res := C.wiringPiISR(C.int(pin), C.int(mode), (*[0]byte)(C.proxyForGoCallback))
	if res != 0 {
		fmt.Println("Failed to setup ISR")
	}
}

// Setup setup the GPIO interface
func Setup(method SetupMethod) (*GPIO, error) {
	if loaded {
		panic("wiring pi is already loaded")
	}
	var ret C.int
	switch method {
	case WiringPiSetup:
		ret = C.wiringPiSetup()
	case BroadcomSetup:
		ret = C.wiringPiSetupGpio()
	case PhysSetup:
		ret = C.wiringPiSetupPhys()
	case SysSetup:
		ret = C.wiringPiSetupSys()
	}
	if ret != 0 {
		return nil, RetCode{int(ret)}
	}
	loaded = true
	return &GPIO{
		setup: method,
	}, nil
}
