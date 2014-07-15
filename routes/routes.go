// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package routes initializes the routes for the web service.
package routes

import (
	"github.com/astaxie/beego"
	"github.com/goinggo/beego-mgo/controllers"
)

func init() {
	beego.Router("/", new(controllers.BuoyController), "get:Index")
	beego.Router("/buoy/retrievestation", new(controllers.BuoyController), "post:RetrieveStation")
	beego.Router("/buoy/station/:stationId", new(controllers.BuoyController), "get,post:RetrieveStationJSON")
}
