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

	err := struct {
		Error string
	}{}
	json.Unmarshal(w.Body.Bytes(), &err)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("The Should Be No Error In The Result", func() {
			So(len(err.Error), ShouldEqual, 0)
		})
	})
}
