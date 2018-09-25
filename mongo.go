// Package timelogutil contains utility functions for working with timelog.
package timelogutil

import (
	"log"
	"time"

	"github.com/tkanos/gonfig"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetMongoSession takes a configuration name (s) and returns a mongo session
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

// GetTasksSince returns all tasks since the date that is passed in
func GetTasksSince(startTime time.Time) []Task {
	result := []Task{} // The tasks

	configuration := Configuration{}

	// TODO: something with the file name to cope with dev / prod! environments
	gonfig.GetConf("config/config.development.json", &configuration)

	session := GetMongoSession(configuration)
	defer session.Close()

	c := session.DB(configuration.Database).C("events")

	err := c.Find(bson.M{"time": bson.M{"$gte": startTime}}).All(&result)

	if err != nil {
		log.Fatal(err)
	}
	return result

}
