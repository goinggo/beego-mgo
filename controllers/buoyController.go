// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements the controller layer for buoy
*/
package controllers

import (
	"github.com/goinggo/beego-mgo/app/business/buoyBusiness"
	bc "github.com/goinggo/beego-mgo/controllers/baseController"
)

//** TYPES

type BuoyController struct {
	bc.BaseController
}

//** WEB FUNCTIONS

// Stations returns the specified station
// http://localhost:9003/station/42002
func (this *BuoyController) Station() {
	buoyBusiness.Station(&this.BaseController, this.GetString(":stationId"))
}

// Stations returns the specified region
// http://localhost:9003/region/Gulf%20Of%20Mexico
func (this *BuoyController) Region() {
	buoyBusiness.Region(&this.BaseController, this.GetString(":region"))
}
