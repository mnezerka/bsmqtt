package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	OUTPUT_PADDING       = 3
	LOGGER_MODULE        = "piot"
	LOGGER_FORMAT        = "[%{level:.6s}] %{message}"
	LOGGER_FORMAT_COLORS = "%{color}[%{level:.6s}] %{color:reset}%{message}"
)

var (
	config_cfg_file      string
	config_log_level     string
	config_mqtt_url      string
	config_mqtt_user     string
	config_mqtt_password string
)

// global instance of logger
var log = logging.MustGetLogger(LOGGER_MODULE)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "bsmqtt",
	Short:   "Blue Soft MQTT Client",
	Long:    ``,
	Version: appVersion,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&config_cfg_file, "config", "", "config file (default is $HOME/.bsmqtt)")

	rootCmd.PersistentFlags().StringVarP(&config_log_level, "log-level", "", "INFO", "Log level (CRITICIAL, ERROR, WARNING, NOTICE, INFO, DEBUG)")
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	rootCmd.PersistentFlags().StringVar(&config_mqtt_url, "mqtt-url", "", "MQTT broker url")
	viper.BindPFlag("mqtt.url", rootCmd.PersistentFlags().Lookup("mqtt-url"))

	rootCmd.PersistentFlags().StringVar(&config_mqtt_user, "mqtt-user", "", "User")
	viper.BindPFlag("mqtt.user", rootCmd.PersistentFlags().Lookup("mqtt-user"))

	rootCmd.PersistentFlags().StringVar(&config_mqtt_password, "mqtt-password", "", "Password")
	viper.BindPFlag("mqtt.password", rootCmd.PersistentFlags().Lookup("mqtt-password"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if config_cfg_file != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config_cfg_file)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".piot" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".bsmqtt")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("BSMQTT")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	var config_file_used string
	if err := viper.ReadInConfig(); err == nil {
		config_file_used = viper.ConfigFileUsed()
	}

	// configure logging
	var logLevelStr = viper.GetString("log.level")
	// try to convert string log level
	logLevel, err := logging.LogLevel(logLevelStr)
	if err != nil {
		fmt.Printf("Invalid logging level: \"%s\"\n", logLevelStr)
		os.Exit(1)
	}

	formatterStdErr := logging.NewBackendFormatter(
		// out, prefix flag
		logging.NewLogBackend(os.Stderr, "", 0),
		logging.MustStringFormatter(LOGGER_FORMAT_COLORS),
	)
	logging.SetBackend(formatterStdErr)
	logging.SetLevel(logLevel, LOGGER_MODULE)

	log.Debug("Logging initialized")

	if len(config_file_used) > 0 {
		log.Infof("Using config file: '%s'", config_file_used)
	}

}
