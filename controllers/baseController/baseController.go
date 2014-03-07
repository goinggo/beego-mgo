// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements boilerplate code for all controllers
*/
package baseController

import (
	"reflect"
	"runtime"

	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/goinggo/beego-mgo/localize"
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

// ParseAndValidate will run the params through the validation framework and then
// response with the specified localized or provided message
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
		// Build a map of the error messages for each field
		messages2 := map[string]string{}
		val := reflect.ValueOf(params).Elem()
		for i := 0; i < val.NumField(); i++ {
			// Look for an error tag in the field
			typeField := val.Type().Field(i)
			tag := typeField.Tag
			tagValue := tag.Get("error")

			// Was there an error tag
			if tagValue != "" {
				messages2[typeField.Name] = tagValue
			}
		}

		// Build the error response
		errors := []string{}
		for _, err := range valid.Errors {
			// Match an error from the validation framework errors
			// to a field name we have a mapping for
			message, ok := messages2[err.Field]
			if ok == true {
				// Use a localized message if one exists
				errors = append(errors, localize.T(message))
				continue
			}

			// No match, so use the message as is
			errors = append(errors, err.Message)
		}

		this.ServeValidationErrors(errors)
		return false
	}

	return true
}

//** EXCEPTIONS

// ServeError prepares and serves an error exception
func (this *BaseController) ServeError(err error) {
	this.Data["json"] = struct {
		Error string `json:"error"`
	}{err.Error()}
	this.Ctx.Output.SetStatus(500)
	this.ServeJson()
}

// ServeValidationErrors prepares and serves a validation exception
func (this *BaseController) ServeValidationErrors(errors []string) {
	this.Data["json"] = struct {
		Errors []string `json:"errors"`
	}{errors}
	this.Ctx.Output.SetStatus(409)
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
