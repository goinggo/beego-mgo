// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package endpointTests implements boilerplate code for all testing.
package endpointTests

import (
	"github.com/goinggo/beego-mgo/localize"
	_ "github.com/goinggo/beego-mgo/routes" // Initalize routes
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	log "github.com/goinggo/tracelog"
)

//** CONSTANTS

const (
	// SessionID is just mocking the id for testing.
	SessionID = "testing"
)

//** INIT

// init initializes all required packages and systems
func init() {
	log.Start(log.LevelTrace)

	// Init mongo
	log.Started("main", "Initializing Mongo")
	err := mongo.Startup(helper.MainGoRoutine)
	if err != nil {
		log.CompletedError(err, helper.MainGoRoutine, "initTesting")
		return
	}

	// Load message strings
	localize.Init("en-US")
}
