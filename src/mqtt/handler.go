package mqtt

import (
	"crypto/tls"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/pixeljonas/frontipi/src/config"
	"github.com/pixeljonas/frontipi/src/homeassistant"
)

var onConnectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	cfg := config.Get()

	if token := client.Publish(fmt.Sprintf("%s/%s", cfg.RootTopic, cfg.AvailTopic), 0, false, cfg.AvailMessage); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

var defaultHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// Connect returns a default mqtt.Client
func Connect() MQTT.Client {
	cfg := config.Get()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         tls.NoClientCert,
	}

	opts := &MQTT.ClientOptions{
		ClientID:     cfg.ClientID,
		CleanSession: true,
		Username:     cfg.Username,
		Password:     cfg.Password,
		TLSConfig:    tlsConfig,
		OnConnect:    onConnectHandler,
		WillEnabled:  true,
		WillPayload:  []byte(cfg.DisconMessage),
		WillQos:      1,
		WillRetained: false,
		WillTopic:    fmt.Sprintf("%s/%s", cfg.RootTopic, cfg.AvailTopic),
	}

	url := fmt.Sprintf("ssl://%s:%s", cfg.Host, cfg.Port)
	opts.AddBroker(url)
	opts.SetDefaultPublishHandler(defaultHandler)
	client := MQTT.NewClient(opts)
	return client
}

// AutoDiscover sends a HomeAssistant Auto-Discovery message
func AutoDiscover(client MQTT.Client) {
	cfg := config.Get()
	if token := client.Publish(fmt.Sprintf("%s/%s", cfg.RootTopic, cfg.DiscoveryTopic), 1, true, homeassistant.CreateDiscoveryMessage()); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

// PublishState publishes the current state to MQTT
func PublishState(client MQTT.Client) {
	cfg := config.Get()
	if token := client.Publish(fmt.Sprintf("%s/%s", cfg.RootTopic, cfg.StateTopic), 0, false, homeassistant.GetState()); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

//ApplyCommand sends the command string to homeassistant handler
func ApplyCommand(command string) {
	homeassistant.ApplyState(command)
}
