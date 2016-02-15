// Copyright 2013 Ardan Studios. All rights reserved.
// Use of controller source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package controllers implements the controller layer for the buoy API.
package controllers

import (
	bc "github.com/goinggo/beego-mgo/controllers/baseController"
	"github.com/goinggo/beego-mgo/services/buoyService"
	log "github.com/goinggo/tracelog"
)

//** TYPES

// BuoyController manages the API for buoy related functionality.
type BuoyController struct {
	bc.BaseController
}

//** WEB FUNCTIONS

// Index is the initial view for the buoy system.
func (controller *BuoyController) Index() {
	region := "Gulf Of Mexico"
	log.Startedf(controller.UserID, "BuoyController.Index", "Region[%s]", region)

	buoyStations, err := buoyService.FindRegion(&controller.Service, region)
	if err != nil {
		log.CompletedErrorf(err, controller.UserID, "BuoyController.Index", "Region[%s]", region)
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

// RetrieveStation handles the example 2 tab.
func (controller *BuoyController) RetrieveStation() {
	var params struct {
		StationID string `form:"stationID" valid:"Required; MinSize(4)" error:"invalid_station_id"`
	}

	if controller.ParseAndValidate(&params) == false {
		return
	}

	buoyStation, err := buoyService.FindStation(&controller.Service, params.StationID)
	if err != nil {
		log.CompletedErrorf(err, controller.UserID, "BuoyController.RetrieveStation", "StationID[%s]", params.StationID)
		controller.ServeError(err)
		return
	}

	controller.Data["Station"] = buoyStation
	controller.Layout = ""
	controller.TplNames = "buoy/modal/pv_station-detail.html"
	view, _ := controller.RenderString()

	controller.AjaxResponse(0, "SUCCESS", view)
}

// RetrieveStationJSON handles the example 3 tab.
// http://localhost:9003/buoy/station/42002
func (controller *BuoyController) RetrieveStationJSON() {
	// The call to ParseForm inside of ParseAndValidate is failing. This is a BAD FIX
	params := struct {
		StationID string `form:":stationId" valid:"Required; MinSize(4)" error:"invalid_station_id"`
	}{controller.GetString(":stationId")}

	if controller.ParseAndValidate(&params) == false {
		return
	}

	buoyStation, err := buoyService.FindStation(&controller.Service, params.StationID)
	if err != nil {
		log.CompletedErrorf(err, controller.UserID, "Station", "StationID[%s]", params.StationID)
		controller.ServeError(err)
		return
	}

	controller.Data["json"] = buoyStation
	controller.ServeJSON()
}
