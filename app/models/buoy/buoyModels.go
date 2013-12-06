// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	BuoyModels contains the models for the buoy service
*/
package buoyModel

import (
	"labix.org/v2/mgo/bson"
)

//** TYPES

type (
	// BuoyCondition contains information for an individual station
	BuoyCondition struct {
		WindSpeed     float64 `bson:"wind_speed_milehour"`
		WindDirection int     `bson:"wind_direction_degnorth"`
		WindGust      float64 `bson:"gust_wind_speed_milehour"`
	}

	// BuoyLocation contains the buoys location
	BuoyLocation struct {
		Type        string    `bson:"type"`
		Coordinates []float64 `bson:"coordinates"`
	}

	// BuoyStation contains information for an individual station
	BuoyStation struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		StationId string        `bson:"station_id"`
		Name      string        `bson:"name"`
		LocDesc   string        `bson:"location_desc"`
		Condition BuoyCondition `bson:"condition"`
		Location  BuoyLocation  `bson:"location"`
	}
)
