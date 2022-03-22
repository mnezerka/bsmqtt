package cmd

import (
	"math/rand"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

func on_reconnecting(client MQTT.Client, client_options *MQTT.ClientOptions) {
	log.Info("Reconnecting")
}

func on_connect(client MQTT.Client) {
	log.Info("Connected")
}

func connect(client_name string, publish_handler MQTT.MessageHandler) MQTT.Client {
	opts := MQTT.NewClientOptions()

	var mqtt_url = viper.GetString("mqtt.url")
	var mqtt_user = viper.GetString("mqtt.user")
	var mqtt_password = viper.GetString("mqtt.password")

	if len(mqtt_url) == 0 {
		log.Fatal("MQTT url not specified, use flag --mqtt-url or set env var BSMQTT_MQTT_URL")
	}

	log.Infof("Connecting to MQTT broker '%s' as client '%s'", mqtt_url, client_name)

	opts.AddBroker(mqtt_url)

	opts.SetClientID(client_name)

	if len(mqtt_user) > 0 {
		opts.SetUsername(mqtt_user)
	}

	if len(mqtt_password) > 0 {
		opts.SetPassword(mqtt_password)
	}

	if publish_handler != nil {
		opts.SetDefaultPublishHandler(publish_handler)
	}

	opts.SetAutoReconnect(true)
	opts.SetReconnectingHandler(on_reconnecting)
	opts.SetOnConnectHandler(on_connect)
	// Start the connection
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic(token.Error())
	}

	return client
}

func get_random_number() int {
	// seed to avoid getting same numbers if app is executed 1+ times
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9000

	return rand.Intn(max-min+1) + min
}
