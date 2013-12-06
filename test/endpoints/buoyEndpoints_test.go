// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements tests for the buoy endpoints
*/
package testEndpoints

import (
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestStation is a sample to run an endpoint test
func TestStation(t *testing.T) {
	r, _ := http.NewRequest("GET", "/station/42002", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		So(w.Code, ShouldEqual, 200)
		So(w.Body.Len(), ShouldBeGreaterThan, 0)
	})
}

// TestRegion is a sample to run an endpoint test
func TestRegion(t *testing.T) {
	r, _ := http.NewRequest("GET", "/region/Gulf%20Of%20Mexico", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Region Endpoint\n", t, func() {
		So(w.Code, ShouldEqual, 200)
		So(w.Body.Len(), ShouldBeGreaterThan, 0)
	})
}
