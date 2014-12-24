// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package serviceTests implements boilerplate code for all testing
package serviceTests

import (
	"github.com/goinggo/beego-mgo/localize"
	"github.com/goinggo/beego-mgo/services"
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	log "github.com/goinggo/tracelog"
)

//** CONSTANTS

const (
	// SessionID is just mocking the id for testing.
	SessionID = "testing"
)

//** TYPES

type (
	// testController contains state and behavior for testing
	testController struct {
		services.Service
	}
)

//** INIT

// init initializes all required packages and systems
func init() {
	log.Start(log.LEVEL_TRACE)

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

//** INTERCEPT FUNCTIONS

// Prepare is called before controllers are called.
func Prepare() *services.Service {
	var service services.Service

	// TODO: Add Test User To Environment
	service.UserID = "testing"

	err := service.Prepare()
	if err != nil {
		log.Error(err, service.UserID, "Prepare")
		return nil
	}

	log.Trace(service.UserID, "Before", "UserID[%s]", service.UserID)
	return &service
}

// Finish is called after controllers are called.
func Finish(service *services.Service) {
	service.Finish()

	log.Completed(service.UserID, "Finish")
}
