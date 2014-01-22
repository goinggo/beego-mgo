// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements the this layer for buoy
*/
package controllers

import (
	bc "github.com/goinggo/beego-mgo/controllers/baseController"
	"github.com/goinggo/beego-mgo/services/buoyService"
	"github.com/goinggo/tracelog"
)

//** TYPES

type BuoyController struct {
	bc.BaseController
}

//** WEB FUNCTIONS

// Stations returns the specified station
// http://localhost:9003/station/42002
func (this *BuoyController) Station() {
	defer bc.CatchPanic(&this.BaseController, "Station")

	stationId := this.GetString(":stationId")

	tracelog.STARTEDf(this.UserId, "Station", "StationId[%s]", stationId)

	buoyStation, err := buoyService.FindStation(&this.Service, stationId)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, this.UserId, "Station", "StationId[%s]", stationId)
		this.ServeError(err)
		return
	}

	this.Data["json"] = &buoyStation
	this.ServeJson()

	tracelog.COMPLETED(this.UserId, "Station")
}

// Stations returns the specified region
// http://localhost:9003/region/Gulf%20Of%20Mexico
func (this *BuoyController) Region() {
	defer bc.CatchPanic(&this.BaseController, "Region")

	region := this.GetString(":region")

	tracelog.STARTEDf(this.UserId, "Region", "Region[%s]", region)

	buoyStations, err := buoyService.FindRegion(&this.Service, region)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, this.UserId, "Region", "Region[%s]", region)
		this.ServeError(err)
		return
	}

	this.Data["json"] = &buoyStations
	this.ServeJson()

	tracelog.COMPLETED(this.UserId, "Region")
}

// ShowRegions shows a view of the stations for the region
// http://localhost:9003/region-show/Gulf%20Of%20Mexico
func (this *BuoyController) ShowRegions() {
	defer bc.CatchPanic(&this.BaseController, "ShowRegions")

	region := this.GetString(":region")

	tracelog.STARTEDf(this.UserId, "Region", "Region[%s]", region)

	buoyStations, err := buoyService.FindRegion(&this.Service, region)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, this.UserId, "Region", "Region[%s]", region)
		this.ServeError(err)
		return
	}

	this.Data["stations"] = buoyStations
	this.Layout = "buoy/layout.html"
	this.TplNames = "buoy/content.html"
	this.LayoutSections = map[string]string{}
	this.LayoutSections["Stations"] = "buoy/stations.html"
}
