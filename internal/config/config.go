// got from
// https://github.com/RobinUS2/indispenso/blob/master/conf.go

package config

import (
	"fmt"
	"os"

	logging "github.com/satandyh/atsuma/internal/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Conf struct {
	Nmap struct {
		Ip   string
		Port string
	}
	confFlags *pflag.FlagSet
}

// logger
var logConfig = logging.LogConfig{
	ConsoleLoggingEnabled: true,
	EncodeLogsAsJson:      true,
	FileLoggingEnabled:    false,
	Directory:             ".",
	Filename:              "log.log",
	MaxSize:               10,
	MaxBackups:            1,
	MaxAge:                1,
	LogLevel:              6,
}

var log = logging.Configure(logConfig)

func NewConfig() Conf {
	var c Conf

	// all env will look like ASM_SOMETHING
	// for embedded use ASM_LEV1.VALUE
	viper.SetEnvPrefix("asm")

	// Defaults
	viper.SetDefault("Nmap.Ip", "127.0.0.1")
	viper.SetDefault("Nmap.Port", "443")

	//Flags
	c.confFlags = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
	configFile := c.confFlags.StringP("config", "c", "", "Config file location. Supported formats {json,toml,yaml}. Default path {'$HOME/.atsuma','.','./config','/opt/atsuma'}/config.yml")
	//c.confFlags.StringP("verbose", "v", "6", "Logging verbosity level. Default 6 level (Info)")
	help := c.confFlags.BoolP("help", "h", false, "Print help message")

	//parse flags
	arg_err := c.confFlags.Parse(os.Args[1:])
	if arg_err != nil {
		log.Fatal().
			Err(arg_err).
			Str("module", "config").
			Msg("")
	}
	if *help {
		fmt.Println("Usage of atsuma:")
		c.confFlags.PrintDefaults()
		os.Exit(0)
	}

	if len(*configFile) > 2 {
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName("config.yml")    // name of config file (without extension)
		viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("/opt/atsuma")   // path to look for the config file in
		viper.AddConfigPath("$HOME/.atsuma") // call multiple times to add many search paths
		viper.AddConfigPath("./config")
		viper.AddConfigPath(".")
	}

	// bind flags from pflags
	arg_bind_err := viper.BindPFlags(c.confFlags)
	if arg_bind_err != nil {
		log.Fatal().
			Err(arg_bind_err).
			Str("module", "config").
			Msg("")
	}

	// try to get values from env
	viper.AutomaticEnv()

	// get values from config
	file_read_err := viper.ReadInConfig()
	if file_read_err != nil {
		log.Fatal().
			Err(file_read_err).
			Str("module", "config").
			Msg("")
		//os.Exit(0)
	}

	// do all above and get our values
	dec_err := viper.Unmarshal(&c)
	if dec_err != nil {
		log.Fatal().
			Err(dec_err).
			Str("module", "config").
			Msg("")
	}

	return c
}
