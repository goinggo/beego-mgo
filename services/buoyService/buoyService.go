// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package buoyService implements the service for the buoy functionality.
package buoyService

import (
	"strings"

	"github.com/goinggo/beego-mgo/models/buoyModels"
	"github.com/goinggo/beego-mgo/services"
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	"github.com/goinggo/tracelog"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//** TYPES

type (
	// buoyConfiguration contains settings for running the buoy service.
	buoyConfiguration struct {
		Database string
	}
)

//** PACKAGE VARIABLES

// Config provides buoy configuration.
var Config buoyConfiguration

//** INIT

func init() {
	// Pull in the configuration.
	if err := envconfig.Process("buoy", &Config); err != nil {
		tracelog.CompletedError(err, helper.MainGoRoutine, "Init")
	}
}

//** PUBLIC FUNCTIONS

// FindStation retrieves the specified station
func FindStation(service *services.Service, stationID string) (*buoyModels.BuoyStation, error) {
	tracelog.Startedf(service.UserID, "FindStation", "stationID[%s]", stationID)

	var buoyStation buoyModels.BuoyStation
	if err := service.DBAction(Config.Database, "buoy_stations",
		func(collection *mgo.Collection) error {
			queryMap := bson.M{"station_id": stationID}

			tracelog.Trace(service.UserID, "FindStation", "MGO : db.buoy_stations.find(%s).limit(1)", mongo.ToString(queryMap))
			return collection.Find(queryMap).One(&buoyStation)
		}); err != nil {
		if strings.Contains(err.Error(), "not found") == false {
			tracelog.CompletedError(err, service.UserID, "FindStation")
			return nil, err
		}
	}

	tracelog.Completedf(service.UserID, "FindStation", "buoyStation%+v", &buoyStation)
	return &buoyStation, nil
}

// FindRegion retrieves the stations for the specified region
func FindRegion(service *services.Service, region string) ([]buoyModels.BuoyStation, error) {
	tracelog.Startedf(service.UserID, "FindRegion", "region[%s]", region)

	var buoyStations []buoyModels.BuoyStation
	if err := service.DBAction(Config.Database, "buoy_stations",
		func(collection *mgo.Collection) error {
			queryMap := bson.M{"region": region}

			tracelog.Trace(service.UserID, "FindRegion", "Query : db.buoy_stations.find(%s)", mongo.ToString(queryMap))
			return collection.Find(queryMap).All(&buoyStations)
		}); err != nil {
		tracelog.CompletedError(err, service.UserID, "FindRegion")
		return nil, err
	}

	tracelog.Completedf(service.UserID, "FindRegion", "buoyStations%+v", buoyStations)
	return buoyStations, nil
}
