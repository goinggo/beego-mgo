// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements boilerplate code for all controllers
*/
package controllerbase

import (
	"runtime"

	"github.com/astaxie/beego"
	"github.com/goinggo/beego-mgo/app/services"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	"github.com/goinggo/tracelog"
)

//** TYPES

type (
	// BaseController composes all required types and behavior
	BaseController struct {
		beego.Controller
		services.Service
	}
)

//** INTERCEPT FUNCTIONS

// Prepare is called prior to the controller method
func (this *BaseController) Prepare() {
	this.UserId = "unknown" // TODO: Deal With This Later
	tracelog.TRACE(this.UserId, "Before", "UserId[%s] Path[%s]", this.UserId, this.Ctx.Request.URL.Path)

	var err error
	this.MongoSession, err = mongo.CopyMonotonicSession(this.UserId)
	if err != nil {
		tracelog.ERRORf(err, this.UserId, "Before", this.Ctx.Request.URL.Path)
		this.Abort("500")
	}
}

// Finish is called once the controller method completes
func (this *BaseController) Finish() {
	defer func() {
		if this.MongoSession != nil {
			mongo.CloseSession(this.UserId, this.MongoSession)
			this.MongoSession = nil
		}
	}()

	tracelog.COMPLETEDf(this.UserId, "Finish", this.Ctx.Request.URL.Path)
}

//** CATCHING PANICS

// CatchPanic is used to catch any Panic and log exceptions. Returns a 500 as the response
func CatchPanic(controller *BaseController, functionName string) {
	if r := recover(); r != nil {
		if r != "500" {
			buf := make([]byte, 10000)
			runtime.Stack(buf, false)

			tracelog.WARN(controller.Service.UserId, functionName, "PANIC Defered [%v] : Stack Trace : %v", r, string(buf))

			controller.Abort("500")
		}
	}
}
