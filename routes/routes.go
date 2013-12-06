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
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/station/:stationId", &controllers.BuoyController{}, "get:Station")
	beego.Router("/region/:region", &controllers.BuoyController{}, "get:Region")
}
