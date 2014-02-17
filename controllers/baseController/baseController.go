// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements boilerplate code for all controllers
*/
package baseController

import (
	"runtime"

	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/goinggo/beego-mgo/services"
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
	this.UserId = this.GetString("userId")
	if this.UserId == "" {
		this.UserId = this.GetString(":userId")
	}
	if this.UserId == "" {
		this.UserId = "Unknown"
	}

	err := this.Service.Prepare()
	if err != nil {
		tracelog.ERRORf(err, this.UserId, "BaseController.Prepare", this.Ctx.Request.URL.Path)
		this.ServeError(err)
		return
	}

	tracelog.TRACE(this.UserId, "BaseController.Prepare", "UserId[%s] Path[%s]", this.UserId, this.Ctx.Request.URL.Path)
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

//** VALIDATION

func (this *BaseController) ParseAndValidate(params interface{}) bool {
	err := this.ParseForm(params)
	if err != nil {
		this.ServeError(err)
		return false
	}

	valid := validation.Validation{}
	ok, err := valid.Valid(params)
	if err != nil {
		this.ServeError(err)
		return false
	}

	if ok == false {
		this.ValidationResponse(valid.Errors)
		return false
	}

	return true
}

//** EXCEPTIONS

// ServeError prepares and serves an error exception
func (this *BaseController) ServeError(err error) {
	this.Data["json"] = struct {
		Error string
	}{err.Error()}
	this.Ctx.Output.SetStatus(400)
	this.ServeJson()
}

// ValidationResponse prepares and serves a validation exception
func (this *BaseController) ValidationResponse(validationErrors []*validation.ValidationError) {
	this.Ctx.Output.SetStatus(409)

	response := make([]string, len(validationErrors))
	for index, validationError := range validationErrors {
		response[index] = fmt.Sprintf("%s: %s", validationError.Field, validationError.String())
	}

	this.Data["json"] = response
	this.ServeJson()
}

//** CATCHING PANICS

// CatchPanic is used to catch any Panic and log exceptions. Returns a 500 as the response
func (this *BaseController) CatchPanic(functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		tracelog.WARN(this.Service.UserId, functionName, "PANIC Defered [%v] : Stack Trace : %v", r, string(buf))

		this.ServeError(fmt.Errorf("%v", r))
	}
}

//** AJAX SUPPORT

// AjaxResponse returns a standard ajax response
func (this *BaseController) AjaxResponse(resultCode int, resultString string, data interface{}) {
	response := struct {
		Result       int
		ResultString string
		ResultObject interface{}
	}{
		Result:       resultCode,
		ResultString: resultString,
		ResultObject: data,
	}

	this.Data["json"] = response
	this.ServeJson()
}
