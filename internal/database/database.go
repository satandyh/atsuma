package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	logging "github.com/satandyh/atsuma/internal/logger"
)

type DB struct {
	Name string
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

func InitDB(PathDB string) {
	db, db_err := sql.Open("sqlite3", PathDB)
	if db_err != nil {
		log.Fatal().
			Err(db_err).
			Str("module", "database").
			Msg("")
	}
	defer db.Close()

	// try to initialize
	_, init_err := db.Exec(initDB)
	if init_err != nil {
		log.Fatal().
			Err(init_err).
			Str("module", "database").
			Msg("Database initialization error")
	} else {
		log.Info().
			Str("module", "database").
			Msg("Database initialization success")
	}
}

func AddExampleData(PathDB string) {
	db, db_err := sql.Open("sqlite3", PathDB)
	if db_err != nil {
		log.Fatal().
			Err(db_err).
			Str("module", "database").
			Msg("")
	}
	defer db.Close()

	// try to insert example data
	_, add_err := db.Exec(addExampleData)
	if add_err != nil {
		log.Error().
			Err(add_err).
			Str("module", "database").
			Msg("Error cannot insert example data")
		os.Exit(0)
	} else {
		log.Info().
			Str("module", "database").
			Msg("Example data inserted")
	}
}

func DeleteExampleData(PathDB string) {
	db, db_err := sql.Open("sqlite3", PathDB)
	if db_err != nil {
		log.Fatal().
			Err(db_err).
			Str("module", "database").
			Msg("")
	}
	defer db.Close()

	// try to delete example data
	_, add_err := db.Exec(delExampleData)
	if add_err != nil {
		log.Error().
			Err(add_err).
			Str("module", "database").
			Msg("Error cannot delete example data")
		os.Exit(0)
	}
}
