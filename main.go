// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Sample web application for beego and mgo
*/
package main

import (
	"github.com/astaxie/beego"
	_ "github.com/goinggo/beego-mgo/routes"
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	"github.com/goinggo/tracelog"
	"os"
)

func main() {
	tracelog.Start(tracelog.LEVEL_TRACE)

	// Init mongo
	tracelog.STARTED("main", "Initializing Mongo")
	err := mongo.Startup(helper.MAIN_GO_ROUTINE)
	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "initApp")
		os.Exit(1)
	}

	beego.Run()

	tracelog.STARTED(helper.MAIN_GO_ROUTINE, "Website Shutdown")
	tracelog.Stop()
}
