package serviceTests

import (
	"github.com/goinggo/beego-mgo/app/services"
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
}

//** INTERCEPT FUNCTIONS

// Prepare is called create a service object
func Prepare() *services.Service {
	service := &services.Service{}

	service.UserId = "testing" // TODO: Deal With This Later
	tracelog.TRACE(service.UserId, "Before", "UserId[%s]", service.UserId)

	var err error
	service.MongoSession, err = mongo.CopyMonotonicSession(service.UserId)
	if err != nil {
		tracelog.ERROR(err, service.UserId, "Before")
		return nil
	}

	return service
}

// Finish is called once the controller method completes
func Finish(service *services.Service) {
	defer func() {
		if service.MongoSession != nil {
			mongo.CloseSession(service.UserId, service.MongoSession)
			service.MongoSession = nil
		}
	}()

	tracelog.COMPLETED(service.UserId, "Finish")
}
