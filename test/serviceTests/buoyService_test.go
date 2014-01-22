package serviceTests

import (
	"github.com/goinggo/beego-mgo/services/buoyService"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Test_Station checks the station service call is working
func Test_Station(t *testing.T) {
	service := Prepare()
	defer Finish(service)

	stationId := "42002"

	buoyStation, err := buoyService.FindStation(service, stationId)

	Convey("Subject: Test Station Service", t, func() {
		Convey("Should Be Able To Perform A Search", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Should Have Station Data", func() {
			So(buoyStation.StationId, ShouldEqual, stationId)
		})
	})
}

// Test_Region checks the region service call is working
func Test_Region(t *testing.T) {
	service := Prepare()
	defer Finish(service)

	region := "Gulf Of Mexico"

	buoyStations, err := buoyService.FindRegion(service, region)

	Convey("Subject: Test Region Service", t, func() {
		Convey("Should Be Able To Perform A Search", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Should Have Region Data", func() {
			So(len(buoyStations), ShouldBeGreaterThan, 0)
		})
	})
}
