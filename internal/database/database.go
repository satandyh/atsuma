package database

import (
	"database/sql"

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

	// Create the Task table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		command TEXT
	);`)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("module", "database").
			Msg("Error creating tasks table")
	}

	// Create the Trigger table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS triggers (
		id INTEGER PRIMARY KEY,
		task_id INTEGER,
		time TEXT
	);`)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("module", "database").
			Msg("Error creating triggers table")
	}

}
