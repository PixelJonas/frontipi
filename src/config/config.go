package config

import "github.com/caarlos0/env"

// Config represents the default config parameter of the application
type Config struct {
	Username                string `env:"MQTT_USERNAME" envDefault:"frontipi"`
	Password                string `env:"MQTT_PASSWORD,required"`
	Host                    string `env:"MQTT_HOST" envDefault:"192.168.178.32"`
	Port                    string `env:"MQTT_PORT" envDefault:"4883"`
	ClientID                string `env:"MQTT_CLIENTID" envDefault:"frontipi"`
	RootTopic               string `env:"HA_TOPIC" envDefault:"homeassistant/switch/frontipi"`
	AvailTopic              string `env:"HA_AVAIL_TOPIC" envDefault:"availability"`
	AvailMessage            string `env:"HA_AVAIL_MESSAGE" envDefault:"online"`
	DisconMessage           string `env:"HA_DISCONNECT_MESSAGE" envDefault:"offline"`
	DiscoveryTopic          string `env:"HA_DISCOVERY_TOPIC" envDefault:"config"`
	CommandTopic            string `env:"HA_COMMAND_TOPIC" envDefault:"set"`
	StateTopic              string `env:"HA_STATE_TOPIC" envDefault:"state"`
	StateOn                 string `env:"HA_STATE_ON" envDefault:"ON"`
	StateOff                string `env:"HA_STATE_OFF" envDefault:"OFF"`
	StateInterval           int    `env:"HA_STATE_INTERVAL" envDefault:"5"`
	DisplayStatusFile       string `env:"PI_DISPLAY_STATUS_FILE" envDefault:"/sys/class/backlight/rpi_backlight/bl_power"`
	DisplayStatusOn         string `env:"PI_DISPLAY_STATUS_ON" envDefault:"0"`
	DisplayStatusOnCommand  string `env:"PI_DISPLAY_STATUS_ON_COMMAND" envDefault:"ON"`
	DisplayStatusOff        string `env:"PI_DISPLAY_STATUS_OFF" envDefault:"1"`
	DisplayStatusOffCommand string `env:"PI_DISPLAY_STATUS_OFF_COMMAND" envDefault:"OFF"`
}

func readConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return cfg
}

//Get returns the config object
func Get() Config {
	return readConfig()
}
