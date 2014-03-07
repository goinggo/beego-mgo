// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements boilerplate code for all testing
*/
package endpointTests

import (
	"github.com/goinggo/beego-mgo/localize"
	_ "github.com/goinggo/beego-mgo/routes"
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	"github.com/goinggo/tracelog"
)

//** CONSTANTS

const (
	SESSION_ID = "testing"
)

//** INIT

// init initializes all required packages and systems
func init() {
	tracelog.Start(tracelog.LEVEL_TRACE)

	// Init mongo
	tracelog.STARTED("main", "Initializing Mongo")
	err := mongo.Startup(helper.MAIN_GO_ROUTINE)
	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "initTesting")
		return
	}

	// Load message strings
	localize.Init("en-US")
}
