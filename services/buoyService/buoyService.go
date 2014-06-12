// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Buoy implements the service for the buoy functionality
*/
package buoyService

import (
	"strings"

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
		tracelog.CompletedError(err, helper.MAIN_GO_ROUTINE, "Init")
	}
}

//** PUBLIC FUNCTIONS

// FindStation retrieves the specified station
func FindStation(service *services.Service, stationId string) (buoyStation *buoyModels.BuoyStation, err error) {
	defer helper.CatchPanic(&err, service.UserId, "FindStation")

	tracelog.Started(service.UserId, "FindStation")

	queryMap := bson.M{"station_id": stationId}

	buoyStation = &buoyModels.BuoyStation{}
	err = service.DBAction(Config.Database, "buoy_stations",
		func(collection *mgo.Collection) error {
			tracelog.Trace(service.UserId, "FindStation", "Query : %s", mongo.ToString(queryMap))
			return collection.Find(queryMap).One(buoyStation)
		})

	if err != nil {
		if strings.Contains(err.Error(), "not found") == false {
			tracelog.CompletedError(err, service.UserId, "FindStation")
			return buoyStation, err
		}

		err = nil
	}

	tracelog.Completed(service.UserId, "FindStation")
	return buoyStation, err
}

// FindRegion retrieves the stations for the specified region
func FindRegion(service *services.Service, region string) (buoyStations []*buoyModels.BuoyStation, err error) {
	defer helper.CatchPanic(&err, service.UserId, "FindRegion")

	tracelog.Started(service.UserId, "FindRegion")

	queryMap := bson.M{"region": region}

	buoyStations = []*buoyModels.BuoyStation{}
	err = service.DBAction(Config.Database, "buoy_stations",
		func(collection *mgo.Collection) error {
			tracelog.Trace(service.UserId, "FindRegion", "Query : %s", mongo.ToString(queryMap))
			return collection.Find(queryMap).All(&buoyStations)
		})

	if err != nil {
		tracelog.CompletedError(err, service.UserId, "FindRegion")
		return buoyStations, err
	}

	tracelog.Completed(service.UserId, "FindRegion")
	return buoyStations, err
}
