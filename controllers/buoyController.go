// Copyright 2013 Ardan Studios. All rights reserved.
// Use of controller source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements the controller layer for buoy
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

// Index is the initial view for the buoy system
func (controller *BuoyController) Index() {
	region := "Gulf Of Mexico"
	tracelog.STARTEDf(controller.UserId, "BuoyController.Index", "Region[%s]", region)

	buoyStations, err := buoyService.FindRegion(&controller.Service, region)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, controller.UserId, "BuoyController.Index", "Region[%s]", region)
		controller.ServeError(err)
		return
	}

	controller.Data["Stations"] = buoyStations
	controller.Layout = "shared/basic-layout.html"
	controller.TplNames = "buoy/content.html"
	controller.LayoutSections = map[string]string{}
	controller.LayoutSections["PageHead"] = "buoy/page-head.html"
	controller.LayoutSections["Header"] = "shared/header.html"
	controller.LayoutSections["Modal"] = "shared/modal.html"
}

//** AJAX FUNCTIONS

// RetrieveStation handles the example 2 tab
func (controller *BuoyController) RetrieveStation() {
	params := struct {
		StationId string `form:"stationId" valid:"Required; MinSize(4)" error:"invalid_station_id"`
	}{}

	if controller.ParseAndValidate(&params) == false {
		return
	}

	buoyStation, err := buoyService.FindStation(&controller.Service, params.StationId)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, controller.UserId, "BuoyController.RetrieveStation", "StationId[%s]", params.StationId)
		controller.ServeError(err)
		return
	}

	controller.Data["Station"] = buoyStation
	controller.Layout = ""
	controller.TplNames = "buoy/modal/pv_station-detail.html"
	view, _ := controller.RenderString()

	controller.AjaxResponse(0, "SUCCESS", view)
}

// Stations handles the example 3 tab
// http://localhost:9003/buoy/station/42002
func (controller *BuoyController) RetrieveStationJson() {
	params := struct {
		StationId string `form:":stationId" valid:"Required; MinSize(4)" error:"invalid_station_id"`
	}{}

	if controller.ParseAndValidate(&params) == false {
		return
	}

	buoyStation, err := buoyService.FindStation(&controller.Service, params.StationId)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, controller.UserId, "Station", "StationId[%s]", params.StationId)
		controller.ServeError(err)
		return
	}

	controller.Data["json"] = &buoyStation
	controller.ServeJson()
}
