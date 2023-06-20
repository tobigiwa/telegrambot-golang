package bot

import (
	tele "gopkg.in/telebot.v3"
)

type Database interface {
	// IsUser returns true if user is found in the db and false otherwise.
	IsUser(int64) bool
	Delete(int64) error
	Insert(int64, string) error
	//AllIDs returns a slice of all user id and an error
	AllIDs() ([]int64, error)
}

type Logger interface {
	LogInfo(string, string)
	LogError(error, string)
	LogFatal(error, string)
	WriteToStandarOutput(string)
}

// Application is the monolothic struct for the application
type Application struct {
	// Bot holds the Bot instance
	Bot *tele.Bot
	// Storage holds the database instance
	Storage Database
	//Logger holds the logger instance
	Logger Logger
}
