// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements the controller layer for buoy
*/
package controllers

import (
	"github.com/goinggo/beego-mgo/app/business/buoy"
	cb "github.com/goinggo/beego-mgo/controllers/base"
)

//** TYPES

type BuoyController struct {
	cb.BaseController
}

//** WEB FUNCTIONS

// Stations returns the specified station
func (this *BuoyController) Station() {
	buoyBusiness.Station(&this.BaseController, this.Ctx.Input.Param(":stationId"))
}

// Stations returns the specified region
func (this *BuoyController) Region() {
	buoyBusiness.Region(&this.BaseController, this.Ctx.Input.Param(":region"))
}
