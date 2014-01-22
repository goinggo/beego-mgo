// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Buoy implements the service for the buoy functionality
*/
package buoyService

import (
	"github.com/goinggo/beego-mgo/models/buoyModels"
	"github.com/goinggo/beego-mgo/services"
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	"github.com/goinggo/tracelog"
	"github.com/kelseyhightower/envconfig"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//** TYPES

type (
	// buoyConfiguration contains settings for running the buoy service
	buoyConfiguration struct {
		Database string
	}
)

//** PACKAGE VARIABLES

var Config buoyConfiguration

//** INIT

func init() {
	// Pull in the configuration
	err := envconfig.Process("buoy", &Config)
	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "Init")
	}
}

//** PUBLIC FUNCTIONS

// FindStation retrieves the specified station
func FindStation(service *services.Service, stationId string) (buoyStation *buoyModels.BuoyStation, err error) {
	defer helper.CatchPanic(&err, service.UserId, "FindStation")

	tracelog.STARTED(service.UserId, "FindStation")

	queryMap := bson.M{"station_id": stationId}
	tracelog.TRACE(service.UserId, "FindStation", "Query : %s", mongo.ToString(queryMap))

	buoyStation = &buoyModels.BuoyStation{}
	err = service.DBAction(Config.Database, "buoy_stations",
		func(collection *mgo.Collection) error {
			return collection.Find(queryMap).One(buoyStation)
		})

	if err != nil {
		tracelog.COMPLETED_ERROR(err, service.UserId, "FindStation")
		return buoyStation, err
	}

	tracelog.COMPLETED(service.UserId, "FindStation")
	return buoyStation, err
}

// FindRegion retrieves the stations for the specified region
func FindRegion(service *services.Service, region string) (buoyStations []*buoyModels.BuoyStation, err error) {
	defer helper.CatchPanic(&err, service.UserId, "FindRegion")

	tracelog.STARTED(service.UserId, "FindRegion")

	queryMap := bson.M{"region": region}
	tracelog.TRACE(service.UserId, "FindRegion", "Query : %s", mongo.ToString(queryMap))

	buoyStations = []*buoyModels.BuoyStation{}
	err = service.DBAction(Config.Database, "buoy_stations",
		func(collection *mgo.Collection) error {
			return collection.Find(queryMap).All(&buoyStations)
		})

	if err != nil {
		tracelog.COMPLETED_ERROR(err, service.UserId, "FindRegion")
		return buoyStations, err
	}

	tracelog.COMPLETED(service.UserId, "FindRegion")
	return buoyStations, err
}
