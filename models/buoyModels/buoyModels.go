// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	BuoyModels contains the models for the buoy service
*/
package buoyModels

import (
	"fmt"

	"labix.org/v2/mgo/bson"
)

//** TYPES

type (
	// BuoyCondition contains information for an individual station
	BuoyCondition struct {
		WindSpeed     float64 `bson:"wind_speed_milehour" json:"wind_speed_milehour"`
		WindDirection int     `bson:"wind_direction_degnorth" json:"wind_direction_degnorth"`
		WindGust      float64 `bson:"gust_wind_speed_milehour" json:"gust_wind_speed_milehour"`
	}

	// BuoyLocation contains the buoys location
	BuoyLocation struct {
		Type        string    `bson:"type" json:"type"`
		Coordinates []float64 `bson:"coordinates" json:"coordinates"`
	}

	// BuoyStation contains information for an individual station
	BuoyStation struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		StationId string        `bson:"station_id" json:"station_id"`
		Name      string        `bson:"name" json:"name"`
		LocDesc   string        `bson:"location_desc" json:"location_desc"`
		Condition BuoyCondition `bson:"condition" json:"condition"`
		Location  BuoyLocation  `bson:"location" json:"location"`
	}
)

func (buoyCondition *BuoyCondition) DisplayWindSpeed() string {
	return fmt.Sprintf("%.2f", buoyCondition.WindSpeed)
}

func (buoyCondition *BuoyCondition) DisplayWindGust() string {
	return fmt.Sprintf("%.2f", buoyCondition.WindGust)
}
