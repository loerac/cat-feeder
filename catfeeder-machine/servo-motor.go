package main

import (
    "fmt"
    "time"

    //"gobot.io/x/gobot"
)

const OPEN_FEEDER uint8 = 45
const CLOSE_FEEDER uint8 = 0

func RotateMotor() {
	//servo := gpio.NewServoDriver(rapi_adapter, pin)

    fmt.Println("Opening feeder...")
    //servo.Move(OPEN_FEEDER)
    time.Sleep(time.Second)
    fmt.Println("Closing feeder...")
    //servo.Move(CLOSE_FEEDER)
}
