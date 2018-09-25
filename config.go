package timelogutil

import "time"

// Configuration is used for configuration properties
type Configuration struct {
	MongoDBHost  string
	Database     string
	AuthUserName string
	AuthPassword string
}

// Task describes what it is that is being worked on at any time
type Task struct {
	Time  time.Time
	Event string
}
