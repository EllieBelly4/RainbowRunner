package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config RRConfig

type LoggingOptions struct {
	LogMoves             bool            `mapstructure:"log_moves"`
	LogReceivedMoves     bool            `mapstructure:"log_received_moves"`
	LogGenericSent       bool            `mapstructure:"log_generic_sent"`
	LogSmallAs           bool            `mapstructure:"log_small_as"`
	LogHashes            bool            `mapstructure:"log_hashes"`
	LogGCObjectSerialise bool            `mapstructure:"log_gc_object_serialise"`
	LogRandomEquipment   bool            `mapstructure:"log_random_equipment"`
	LogFilterMessages    bool            `mapstructure:"log_filter_messages"`
	LogSentMessageTypes  map[string]bool `mapstructure:"log_sent_message_types"`
	LogFileName          string          `mapstructure:"log_file_name"`
	LogTruncate          bool            `mapstructure:"log_truncate"`
	LogEMessages         bool            `mapstructure:"log_e_messages"`
	LogIDs               bool            `mapstructure:"log_ids"`
}

type RRConfig struct {
	SendMovementMessages     bool           `mapstructure:"send_movement_messages"`
	Logging                  LoggingOptions `mapstructure:"logging"`
	ReinitialiseZonesOnEnter bool           `mapstructure:"reinitialise_zones_on_enter"`
}

func Load() {
	viper.SetConfigName("config")                      // name of config file (without extension)
	viper.SetConfigType("yaml")                        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/rainbowrunner/")         // path to look for the config file in
	viper.AddConfigPath("$HOME/.config/rainbowrunner") // call multiple times to add many search paths
	viper.AddConfigPath(".")                           // optionally look for config in the working directory
	err := viper.ReadInConfig()                        // Find and read the config file
	if err != nil {                                    // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
