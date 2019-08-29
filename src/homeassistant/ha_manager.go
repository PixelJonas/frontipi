package homeassistant

import (
	"encoding/json"
	"fmt"

	"github.com/pixeljonas/frontipi/src/config"
	"github.com/pixeljonas/frontipi/src/raspberrypi"
)

func createFullTopicName(topic string) string {
	cfg := config.Get()
	return fmt.Sprintf("%s/%s", cfg.RootTopic, topic)
}

// DiscoveryMessage represents a homeassistant MQTT Discovery Message
type DiscoveryMessage struct {
	UniqueID            string `json:"unique_id"`
	CommandTopic        string `json:"command_topic"`
	Name                string `json:"name"`
	Icon                string `json:"icon"`
	StateTopic          string `json:"state_topic"`
	StateOn             string `json:"state_on"`
	StateOff            string `json:"state_off"`
	AvailabilityTopic   string `json:"availability_topic"`
	PayloadAvailable    string `json:"payload_available"`
	PayloadNotAvailable string `json:"payload_not_available"`
}

//Device represents a homeassistant MQTT Device class
type Device struct {
	Name         string `json:"name"`
	Identifiers  string `json:"identifiers"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	SwVersion    string `json:"sw_version"`
}

//CreateDiscoveryMessage returns a homeassistant MQTT discovery message
func CreateDiscoveryMessage() string {
	cfg := config.Get()
	bytes, err := json.Marshal(DiscoveryMessage{
		"frontipi-001-switch",
		createFullTopicName(cfg.CommandTopic),
		"frontipi_display",
		"mdi:monitor",
		createFullTopicName(cfg.StateTopic),
		cfg.StateOn,
		cfg.StateOff,
		createFullTopicName(cfg.AvailTopic),
		cfg.AvailMessage,
		cfg.DisconMessage,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
	return string(bytes)
}

//GetState returns the current state
func GetState() string {
	cfg := config.Get()
	if raspberrypi.GetDisplayStatus() {
		return cfg.StateOn
	}
	return cfg.StateOff
}

//ApplyState applies the desired state to the raspberry pi
func ApplyState(command string) {
	cfg := config.Get()

	if command == cfg.DisplayStatusOnCommand {
		raspberrypi.ToggleDisplay(true)
	} else {
		raspberrypi.ToggleDisplay(false)
	}
}
