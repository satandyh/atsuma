package main

import (
	config "github.com/satandyh/atsuma/internal/config"
	database "github.com/satandyh/atsuma/internal/database"
	logging "github.com/satandyh/atsuma/internal/logger"
)

// Global vars for logs
var logConfig = logging.LogConfig{
	ConsoleLoggingEnabled: true,
	EncodeLogsAsJson:      true,
	FileLoggingEnabled:    true,
	Directory:             "./data",
	Filename:              "atsuma.log",
	MaxSize:               10,
	MaxBackups:            7,
	MaxAge:                7,
	LogLevel:              6,
}

var logger = logging.Configure(logConfig)

func main() {

	conf := config.NewConfig()
	println(conf.Nmap.Ip)   // test only
	println(conf.Nmap.Port) // test only

	// init db
	database.InitDB(conf.DBPath)

	logger.Info().
		Str("module", "main").
		Msg("All tasks completed.")

}
