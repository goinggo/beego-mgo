// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Initializes the routes for the web service
*/
package routes

import (
	"github.com/astaxie/beego"
	"github.com/goinggo/beego-mgo/controllers"
)

func init() {
	beego.Router("/", &controllers.BuoyController{}, "get:Index")
	beego.Router("/buoy/retrievestation", &controllers.BuoyController{}, "post:RetrieveStation")
	beego.Router("/buoy/station/:stationId", &controllers.BuoyController{}, "get,post:RetrieveStationJson")
}
