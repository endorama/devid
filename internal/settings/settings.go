package settings

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Init initialize settings and default values.
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("DEVID")

	setXDGBaseDirs()
	setDefaults()
	// readConfigFile()
	setConstants()
	expandEnvs()
}

func setXDGBaseDirs() {
	if name, ok := os.LookupEnv("XDG_DATA_HOME"); !ok || name == "" {
		os.Setenv("XDG_DATA_HOME", os.ExpandEnv("$HOME/.local/share"))
	}
}

func setDefaults() {
	viper.SetDefault("personas_location", "$XDG_DATA_HOME/devid/personas")
}

func readConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.ExpandEnv("$XDG_CONFIG_HOME/devid"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("no config file found")
		} else {
			log.Fatalf(fmt.Errorf("cannot read config file: %w", err).Error())
		}
	}
}

// setConstants allow setting configuration values that MUST NOT be available in the config file.
func setConstants() {
	viper.Set("shell_loader_filename", "load.sh")
	viper.Set("shell_runner_filename", "run.sh")
}

// expandEnvs perform environment variable environment on specific configuration values.
func expandEnvs() {
	viper.Set("personas_location",
		os.ExpandEnv(viper.GetString("personas_location")))
}
