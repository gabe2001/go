package main

import (
	"fmt"
	"log"
	"time"
)

var timeZones = [...]string{
	"Europe/Dublin",
	"Europe/Amsterdam",
	"Europe/Bristol", // does not exist
	"America/New_York",
	"America/Phoenix",
	"America/Denver",
	"America/Chicago",
	"America/San_Francisco", // doesn't exist either
	"America/Los_Angeles",
}

func main() {

	now := time.Now()

	for _, zone := range timeZones {
		location, err := time.LoadLocation(zone)
		if err != nil {
			log.Println("ERROR", err)
		} else {
			zonedTime := now.In(location)
			fmt.Printf("%25s: ", zone)
			fmt.Printf("%4d-%02d-%02dT%02d:%02d:%02d\n",
				zonedTime.Year(), zonedTime.Month(), zonedTime.Day(), zonedTime.Hour(), zonedTime.Minute(), zonedTime.Second())
		}
	}

}
