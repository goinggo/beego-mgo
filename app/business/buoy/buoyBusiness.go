// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Buoy implements the business layer for the buoy functionality
*/
package buoyBusiness

import (
	"github.com/goinggo/beego-mgo/app/services/buoy"
	cb "github.com/goinggo/beego-mgo/controllers/base"
	"github.com/goinggo/tracelog"
)

//** PUBLIC FUNCTIONS

// Station handles the higher level business processing for this API Call
func Station(controller *cb.BaseController, stationId string) {
	defer cb.CatchPanic(controller, "Station")

	tracelog.STARTEDf(controller.UserId, "Station", "StationId[%s]", stationId)

	buoyStation, err := buoyService.FindStation(&controller.Service, stationId)
	if err != nil {
		controller.Abort("500")
		return
	}

	controller.Data["json"] = &buoyStation
	controller.ServeJson()

	tracelog.COMPLETED(controller.UserId, "Station")
}

// Region handles the higher level business processing for this API Call
func Region(controller *cb.BaseController, region string) {
	defer cb.CatchPanic(controller, "Region")

	tracelog.STARTEDf(controller.UserId, "Region", "Region[%s]", region)

	buoyStations, err := buoyService.FindRegion(&controller.Service, region)
	if err != nil {
		controller.Abort("500")
		return
	}

	controller.Data["json"] = &buoyStations
	controller.ServeJson()

	tracelog.COMPLETED(controller.UserId, "Region")
}
