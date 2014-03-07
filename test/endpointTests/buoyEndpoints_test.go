// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements tests for the buoy endpoints
*/
package endpointTests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/goinggo/tracelog"
	. "github.com/smartystreets/goconvey/convey"
)

// TestStation is a sample to run an endpoint test
func TestStation(t *testing.T) {
	r, _ := http.NewRequest("GET", "/buoy/station/42002", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	tracelog.TRACE("testing", "TestStation", "Code[%d]\n%s", w.Code, w.Body.String())

	response := struct {
		StationId string `json:"station_id"`
		Name      string `json:"name"`
		LocDesc   string `json:"location_desc"`
		Condition struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"condition"`
		Location struct {
			WindSpeed     float64 `json:"wind_speed_milehour"`
			WindDirection int     `json:"wind_direction_degnorth"`
			WindGust      float64 `json:"gust_wind_speed_milehour"`
		} `json:"location"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("The Result Should Be For Station 42002", func() {
			So(response.StationId, ShouldEqual, "42002")
		})
	})
}

// TestInvalidStation is a sample to run an endpoint test that returns
// an empty result set
func TestInvalidStation(t *testing.T) {
	r, _ := http.NewRequest("GET", "/buoy/station/000000", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	tracelog.TRACE("testing", "TestStation", "Code[%d]\n%s", w.Code, w.Body.String())

	response := struct {
		StationId string `json:"station_id"`
		Name      string `json:"name"`
		LocDesc   string `json:"location_desc"`
		Condition struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"condition"`
		Location struct {
			WindSpeed     float64 `json:"wind_speed_milehour"`
			WindDirection int     `json:"wind_direction_degnorth"`
			WindGust      float64 `json:"gust_wind_speed_milehour"`
		} `json:"location"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("The Result Should Be Empty For Station 00000", func() {
			So(response.StationId, ShouldBeBlank)
		})
	})
}

// TestInvalidStation is a sample to run an endpoint test that returns
// an empty result set
func TestMissingStation(t *testing.T) {
	r, _ := http.NewRequest("GET", "/buoy/station/420", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	tracelog.TRACE("testing", "TestStation", "Code[%d]\n%s", w.Code, w.Body.String())

	err := struct {
		Errors []string `json:"errors"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &err)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 409", func() {
			So(w.Code, ShouldEqual, 409)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("The Should Be An Error In The Result", func() {
			So(len(err.Errors), ShouldEqual, 1)
		})
	})
}
