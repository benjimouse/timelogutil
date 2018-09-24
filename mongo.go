// Package timelogutil contains utility functions for working with timelog.
package timelogutil

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Takes a configuration name (s) and returns a mongo session
func GetMongoSession(conf Configuration) *mgo.Session {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{conf.MongoDBHost},
		Timeout:  60 * time.Second,
		Database: conf.Database,
		Username: conf.AuthUserName,
		Password: conf.AuthPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal(err)
	}
	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession.Copy()

}
