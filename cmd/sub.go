package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

var on_message MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Println("------- msg -------")
	fmt.Printf("topic: %s\n", msg.Topic())
	fmt.Printf("value: %s\n", msg.Payload())
}

func subscribe(topics []string) {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	client := connect(fmt.Sprintf("bsmqtt-sub-%d", get_random_number()), on_message)

	for i := 0; i < len(topics); i++ {
		client.Subscribe(topics[i], 0, nil)
		log.Infof("Subscribed to topic '%s'\n", topics[i])
	}

	<-c
}

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subscribe to topic(s)",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("no topics specified")
		}

		subscribe(args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
