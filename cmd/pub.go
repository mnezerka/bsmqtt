package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var pubCmd = &cobra.Command{
	Use:   "pub topic [value]",
	Short: "Publish messasge to topic",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("invalid number of arguments (topic is mandatory argument)")
		}

		client := connect(fmt.Sprintf("bsmqtt-pub-%d", get_random_number()), nil)

		value := ""

		if len(args) > 1 {
			value = args[1]
		}

		client.Publish(args[0], 0, false, value)
		log.Infof("Value '%s' published to topic '%s'\n", value, args[0])

		client.Disconnect(250)
		log.Info("Disconnected")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pubCmd)
}
