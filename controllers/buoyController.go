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

// Index is the initial view for the buoy system
func (this *BuoyController) Index() {
	region := "Gulf Of Mexico"
	tracelog.STARTEDf(this.UserId, "BuoyController.Index", "Region[%s]", region)

	buoyStations, err := buoyService.FindRegion(&this.Service, region)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, this.UserId, "BuoyController.Index", "Region[%s]", region)
		this.ServeError(err)
		return
	}

	this.Data["Stations"] = buoyStations
	this.Layout = "shared/basic-layout.html"
	this.TplNames = "buoy/content.html"
	this.LayoutSections = map[string]string{}
	this.LayoutSections["PageHead"] = "buoy/page-head.html"
	this.LayoutSections["Header"] = "shared/header.html"
	this.LayoutSections["Modal"] = "shared/modal.html"
}

//** AJAX FUNCTIONS

// RetrieveStation handles the example 2 tab
func (this *BuoyController) RetrieveStation() {
	params := struct {
		StationId string `form:"stationId" valid:"Required; MinSize(4)"`
	}{}

	if this.ParseAndValidate(&params) == false {
		return
	}

	buoyStation, err := buoyService.FindStation(&this.Service, params.StationId)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, this.UserId, "BuoyController.RetrieveStation", "StationId[%s]", params.StationId)
		this.ServeError(err)
		return
	}

	this.Data["Station"] = buoyStation
	this.Layout = ""
	this.TplNames = "buoy/modal/pv_station-detail.html"
	view, _ := this.RenderString()

	this.AjaxResponse(0, "SUCCESS", view)
}

// Stations handles the example 3 tab
// http://localhost:9003/buoy/station/42002
func (this *BuoyController) RetrieveStationJson() {
	params := struct {
		StationId string `form:":stationId" valid:"Required; MinSize(4)"`
	}{}

	if this.ParseAndValidate(&params) == false {
		return
	}

	buoyStation, err := buoyService.FindStation(&this.Service, params.StationId)
	if err != nil {
		tracelog.COMPLETED_ERRORf(err, this.UserId, "Station", "StationId[%s]", params.StationId)
		this.ServeError(err)
		return
	}

	this.Data["json"] = &buoyStation
	this.ServeJson()
}
