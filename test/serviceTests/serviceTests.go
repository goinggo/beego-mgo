package serviceTests

import (
	"github.com/goinggo/beego-mgo/localize"
	"github.com/goinggo/beego-mgo/services"
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	"github.com/goinggo/tracelog"
)

//** CONSTANTS

const (
	SESSION_ID = "testing"
)

//** TYPES

type (
	// testController contains state and behavior for testing
	testController struct {
		services.Service
	}
)

//** INIT

// init initializes all required packages and systems
func init() {
	tracelog.Start(tracelog.LEVEL_TRACE)

	// Init mongo
	tracelog.STARTED("main", "Initializing Mongo")
	err := mongo.Startup(helper.MAIN_GO_ROUTINE)
	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "initTesting")
		return
	}

	// Load message strings
	localize.Init("en-US")
}

//** INTERCEPT FUNCTIONS

func Prepare() *services.Service {
	service := &services.Service{}

	// TODO: Add Test User To Environment
	service.UserId = "testing"

	err := service.Prepare()
	if err != nil {
		tracelog.ERROR(err, service.UserId, "Prepare")
		return nil
	}

	tracelog.TRACE(service.UserId, "Before", "UserId[%s]", service.UserId)
	return service
}

func Finish(service *services.Service) {
	service.Finish()

	tracelog.COMPLETED(service.UserId, "Finish")
}
