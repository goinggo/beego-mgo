// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	This package provides mongo connectivity support
*/
package mongo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/tracelog"
	"github.com/kelseyhightower/envconfig"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//** CONSTANTS

const (
	MASTER_SESSION    = "master"
	MONOTONIC_SESSION = "monotonic"
)

//** PACKAGE VARIABLES

var (
	singleton *mongoManager // Reference to the singleton
)

//** TYPES

type (
	// mongoConfiguration contains settings for initialization
	mongoConfiguration struct {
		Hosts    string
		Database string
		UserName string
		Password string
	}

	// mongoManager contains dial and session information
	mongoSession struct {
		mongoDBDialInfo *mgo.DialInfo
		mongoSession    *mgo.Session
	}

	// mongoManager manages a map of session
	mongoManager struct {
		sessions map[string]*mongoSession
	}

	// MongoCall defines a type of function that can be used
	// to excecute code against MongoDB
	MongoCall func(*mgo.Collection) error
)

//** PUBLIC FUNCTIONS

// Startup brings the manager to a running state
func Startup(sessionId string) (err error) {
	defer helper.CatchPanic(&err, sessionId, "Startup")

	// If the system has already been started ignore the call
	if singleton != nil {
		return err
	}

	tracelog.Started(sessionId, "Startup")

	// Pull in the configuration
	config := mongoConfiguration{}
	err = envconfig.Process("mgo", &config)
	if err != nil {
		tracelog.CompletedError(err, sessionId, "Startup")
		return err
	}

	// Create the Mongo Manager
	singleton = &mongoManager{
		sessions: map[string]*mongoSession{},
	}

	// Log the mongodb connection straps
	tracelog.Trace(sessionId, "Startup", "MongoDB : Hosts[%s]", config.Hosts)
	tracelog.Trace(sessionId, "Startup", "MongoDB : Database[%s]", config.Database)
	tracelog.Trace(sessionId, "Startup", "MongoDB : Username[%s]", config.UserName)

	hosts := strings.Split(config.Hosts, ",")

	// Create the strong and monotonic sessions
	err = CreateSession(sessionId, "strong", MASTER_SESSION, hosts, config.Database, config.UserName, config.Password)
	err = CreateSession(sessionId, "monotonic", MONOTONIC_SESSION, hosts, config.Database, config.UserName, config.Password)

	tracelog.Completed(sessionId, "Startup")
	return err
}

// Shutdown systematically brings the manager down gracefully
func Shutdown(sessionId string) (err error) {
	defer helper.CatchPanic(&err, sessionId, "Shutdown")

	tracelog.Started(sessionId, "Shutdown")

	// Close the databases
	for _, session := range singleton.sessions {
		CloseSession(sessionId, session.mongoSession)
	}

	tracelog.Completed(sessionId, "Shutdown")
	return err
}

// CreateSession creates a connection pool for use
func CreateSession(sessionId string, mode string, sessionName string, hosts []string, databaseName string, username string, password string) (err error) {
	defer helper.CatchPanic(nil, sessionId, "CreateSession")

	tracelog.Startedf(sessionId, "CreateSession", "Mode[%s] SessionName[%s] Hosts[%s] DatabaseName[%s] Username[%s]", mode, sessionName, hosts, databaseName, username)

	// Create the database object
	mongoSession := &mongoSession{
		mongoDBDialInfo: &mgo.DialInfo{
			Addrs:    hosts,
			Timeout:  60 * time.Second,
			Database: databaseName,
			Username: username,
			Password: password,
		},
	}

	// Establish the master session
	mongoSession.mongoSession, err = mgo.DialWithInfo(mongoSession.mongoDBDialInfo)
	if err != nil {
		tracelog.CompletedError(err, sessionId, "CreateSession")
		return err
	}

	switch mode {
	case "strong":
		// Reads and writes will always be made to the master server using a
		// unique connection so that reads and writes are fully consistent,
		// ordered, and observing the most up-to-date data.
		// http://godoc.org/labix.org/v2/mgo#Session.SetMode
		mongoSession.mongoSession.SetMode(mgo.Strong, true)
		break

	case "monotonic":
		// Reads may not be entirely up-to-date, but they will always see the
		// history of changes moving forward, the data read will be consistent
		// across sequential queries in the same session, and modifications made
		// within the session will be observed in following queries (read-your-writes).
		// http://godoc.org/labix.org/v2/mgo#Session.SetMode
		mongoSession.mongoSession.SetMode(mgo.Monotonic, true)
	}

	// Have the session check for errors
	// http://godoc.org/labix.org/v2/mgo#Session.SetSafe
	mongoSession.mongoSession.SetSafe(&mgo.Safe{})

	// Add the database to the map
	singleton.sessions[sessionName] = mongoSession

	tracelog.Completed(sessionId, "CreateSession")
	return err
}

// CopyMasterSession makes a copy of the master session for client use
func CopyMasterSession(sessionId string) (*mgo.Session, error) {
	return CopySession(sessionId, MASTER_SESSION)
}

// CopyMonotonicSession makes a copy of the monotonic session for client use
func CopyMonotonicSession(sessionId string) (*mgo.Session, error) {
	return CopySession(sessionId, MONOTONIC_SESSION)
}

// CopySession makes a copy of the specified session for client use
func CopySession(sessionId string, useSession string) (mongoSession *mgo.Session, err error) {
	defer helper.CatchPanic(nil, sessionId, "CopySession")

	tracelog.Startedf(sessionId, "CopySession", "UseSession[%s]", useSession)

	// Find the session object
	session := singleton.sessions[useSession]

	if session == nil {
		err = fmt.Errorf("Unable To Locate Session %s", useSession)
		tracelog.CompletedError(err, sessionId, "CopySession")
		return mongoSession, err
	}

	// Copy the master session
	mongoSession = session.mongoSession.Copy()

	tracelog.Completed(sessionId, "CopySession")
	return mongoSession, err
}

// CloneMasterSession makes a clone of the master session for client use
func CloneMasterSession(sessionId string) (*mgo.Session, error) {
	return CloneSession(sessionId, MASTER_SESSION)
}

// CloneMonotonicSession makes a clone of the monotinic session for client use
func CloneMonotonicSession(sessionId string) (*mgo.Session, error) {
	return CloneSession(sessionId, MONOTONIC_SESSION)
}

// CopySession makes a clone of the specified session for client use
func CloneSession(sessionId string, useSession string) (mongoSession *mgo.Session, err error) {
	defer helper.CatchPanic(nil, sessionId, "CopySession")

	tracelog.Startedf(sessionId, "CloneSession", "UseSession[%s]", useSession)

	// Find the session object
	session := singleton.sessions[useSession]

	if session == nil {
		err = fmt.Errorf("Unable To Locate Session %s", useSession)
		tracelog.CompletedError(err, sessionId, "CloneSession")
		return mongoSession, err
	}

	// Clone the master session
	mongoSession = session.mongoSession.Clone()

	tracelog.Completed(sessionId, "CloneSession")
	return mongoSession, err
}

// CloseSession puts the connection back into the pool
func CloseSession(sessionId string, mongoSession *mgo.Session) {
	defer helper.CatchPanic(nil, sessionId, "CloseSession")

	tracelog.Started(sessionId, "CloseSession")
	mongoSession.Close()
	tracelog.Completed(sessionId, "CloseSession")
}

// GetCollection returns a reference to a collection for the specified database and collection name
func GetCollection(mongoSession *mgo.Session, useDatabase string, useCollection string) (*mgo.Collection, error) {
	return mongoSession.DB(useDatabase).C(useCollection), nil
}

// CollectionExists returns true if the collection name exists in the specified database
func CollectionExists(sessionId string, mongoSession *mgo.Session, useDatabase string, useCollection string) bool {
	database := mongoSession.DB(useDatabase)
	collections, err := database.CollectionNames()

	if err != nil {
		return false
	}

	for _, collection := range collections {
		if collection == useCollection {
			return true
		}
	}

	return false
}

// ToString converts the quer map to a string
func ToString(queryMap bson.M) string {
	json, err := json.Marshal(queryMap)
	if err != nil {
		return ""
	}
	return string(json)
}

// Execute the MongoDB literal function
func Execute(sessionId string, mongoSession *mgo.Session, databaseName string, collectionName string, mongoCall MongoCall) (err error) {
	tracelog.Startedf(sessionId, "Execute", "Database[%s] Collection[%s]", databaseName, collectionName)

	// Capture the specified collection
	collection, err := GetCollection(mongoSession, databaseName, collectionName)
	if err != nil {

		tracelog.CompletedError(err, sessionId, "Execute")
		return err
	}

	// Execute the mongo call
	err = mongoCall(collection)
	if err != nil {
		tracelog.CompletedError(err, sessionId, "Execute")
		return err
	}

	tracelog.Completed(sessionId, "Execute")

	return err
}
