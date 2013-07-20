/*
 * TestApplication.cpp
 *
 * Translated into Go by Frank Davidson, ffdavidson@gmail.com
 *
 * Copyright Derek Molloy, School of Electronic Engineering, Dublin City University
 * www.derekmolloy.ie
 *
 * YouTube Channel: http://www.youtube.com/derekmolloydcu/
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED ''AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL I
 * BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE
 * GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
	"fmt"
	"github.com/davidsonff/SimpleGPIO"
	"time"
)

var (
	LEDGPIO, ButtonGPIO int = 60 /* GPIO1_28 = (1x32) + 28 = 60 */, 15 /* GPIO0_15 = (0x32) + 15 = 15 */
)

func main() {

	fmt.Println()
	fmt.Println("Testing the GPIO pins...")
	fmt.Println()

	SimpleGPIO.GPIOExport(LEDGPIO)
	SimpleGPIO.GPIOSetDirection(LEDGPIO, SimpleGPIO.OUTPUT_PIN)

	SimpleGPIO.GPIOExport(ButtonGPIO)
	SimpleGPIO.GPIOSetDirection(ButtonGPIO, SimpleGPIO.INPUT_PIN)

	for i := 0; i < 5; i++ {
		fmt.Println("Setting the LED on...")
		SimpleGPIO.GPIOSetValue(LEDGPIO, SimpleGPIO.HIGH)
		time.Sleep(1 * time.Second)
		fmt.Println("Setting the LED off...")
		SimpleGPIO.GPIOSetValue(LEDGPIO, SimpleGPIO.LOW)
		time.Sleep(1 * time.Second)
	}

	fmt.Println()
	fmt.Println("Please press the button!!!")

	var value = SimpleGPIO.LOW

	for value != SimpleGPIO.HIGH {
		value = SimpleGPIO.GPIOGetValue(ButtonGPIO)
		time.Sleep(1 * time.Millisecond)
		fmt.Print(".")
	}

	fmt.Println()
	fmt.Println("The button was just pressed!!!")
	fmt.Println()
	fmt.Println("Finished testing the GPIO pins...")
	fmt.Println()

	SimpleGPIO.GPIOUnexport(LEDGPIO)
	SimpleGPIO.GPIOUnexport(ButtonGPIO)
}
