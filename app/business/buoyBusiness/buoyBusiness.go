// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Buoy implements the business layer for the buoy functionality
*/
package buoyBusiness

import (
	"github.com/goinggo/beego-mgo/app/services/buoyService"
	bc "github.com/goinggo/beego-mgo/controllers/baseController"
	"github.com/goinggo/tracelog"
)

//** PUBLIC FUNCTIONS

// Station handles the higher level business processing for this API Call
func Station(controller *bc.BaseController, stationId string) {
	defer bc.CatchPanic(controller, "Station")

	tracelog.STARTEDf(controller.UserId, "Station", "StationId[%s]", stationId)

	buoyStation, err := buoyService.FindStation(&controller.Service, stationId)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, controller.UserId, "Station", "StationId[%s]", stationId)
		controller.ServeError(err)
		return
	}

	controller.Data["json"] = &buoyStation
	controller.ServeJson()

	tracelog.COMPLETED(controller.UserId, "Station")
}

// Region handles the higher level business processing for this API Call
func Region(controller *bc.BaseController, region string) {
	defer bc.CatchPanic(controller, "Region")

	tracelog.STARTEDf(controller.UserId, "Region", "Region[%s]", region)

	buoyStations, err := buoyService.FindRegion(&controller.Service, region)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, controller.UserId, "Region", "Region[%s]", region)
		controller.ServeError(err)
		return
	}

	controller.Data["json"] = &buoyStations
	controller.ServeJson()

	tracelog.COMPLETED(controller.UserId, "Region")
}

// ShowRegions provides a sample of generating a view with a slice
func ShowRegions(controller *bc.BaseController, region string) {
	defer bc.CatchPanic(controller, "Region")

	tracelog.STARTEDf(controller.UserId, "Region", "Region[%s]", region)

	buoyStations, err := buoyService.FindRegion(&controller.Service, region)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, controller.UserId, "Region", "Region[%s]", region)
		controller.ServeError(err)
		return
	}

	controller.Data["stations"] = buoyStations
	controller.Layout = "buoy/layout.html"
	controller.TplNames = "buoy/content.html"
	controller.LayoutSections = map[string]string{}
	controller.LayoutSections["Stations"] = "buoy/stations.html"
}
