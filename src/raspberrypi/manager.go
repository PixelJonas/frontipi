package raspberrypi

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	"github.com/pixeljonas/frontipi/src/config"
)

//GetDisplayStatus returns a boolean representing the TouchScreen Display status
// false = Display is offline
// true = Display is online
func GetDisplayStatus() bool {
	cfg := config.Get()

	data, err := ioutil.ReadFile(cfg.DisplayStatusFile)
	if err != nil {
		panic(err)
	}
	var state string

	if runtime.GOOS == "windows" {
		state = strings.TrimRight(string(data), "\r\n")
	} else {
		state = strings.TrimRight(string(data), "\n")
	}

	if state == cfg.DisplayStatusOn {
		return true
	}
	return false
}

//ToggleDisplay turns the Display on the Touchscreen on the RaspberryPi on or off
func ToggleDisplay(status bool) {
	cfg := config.Get()
	fmt.Println("setting display to", status)
	if status {
		ioutil.WriteFile(cfg.DisplayStatusFile, []byte(cfg.DisplayStatusOn), 0664)
	} else {
		ioutil.WriteFile(cfg.DisplayStatusFile, []byte(cfg.DisplayStatusOff), 0664)
	}
}
