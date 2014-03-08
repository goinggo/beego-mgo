// Copyright 2013 Ardan Studios. All rights reserved.
// Use of baseController source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements boilerplate code for all baseControllers
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

// Prepare is called prior to the baseController method
func (baseController *BaseController) Prepare() {
	baseController.UserId = baseController.GetString("userId")
	if baseController.UserId == "" {
		baseController.UserId = baseController.GetString(":userId")
	}
	if baseController.UserId == "" {
		baseController.UserId = "Unknown"
	}

	err := baseController.Service.Prepare()
	if err != nil {
		tracelog.ERRORf(err, baseController.UserId, "BaseController.Prepare", baseController.Ctx.Request.URL.Path)
		baseController.ServeError(err)
		return
	}

	tracelog.TRACE(baseController.UserId, "BaseController.Prepare", "UserId[%s] Path[%s]", baseController.UserId, baseController.Ctx.Request.URL.Path)
}

// Finish is called once the baseController method completes
func (baseController *BaseController) Finish() {
	defer func() {
		if baseController.MongoSession != nil {
			mongo.CloseSession(baseController.UserId, baseController.MongoSession)
			baseController.MongoSession = nil
		}
	}()

	tracelog.COMPLETEDf(baseController.UserId, "Finish", baseController.Ctx.Request.URL.Path)
}

//** VALIDATION

// ParseAndValidate will run the params through the validation framework and then
// response with the specified localized or provided message
func (baseController *BaseController) ParseAndValidate(params interface{}) bool {
	err := baseController.ParseForm(params)
	if err != nil {
		baseController.ServeError(err)
		return false
	}

	valid := validation.Validation{}
	ok, err := valid.Valid(params)
	if err != nil {
		baseController.ServeError(err)
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

		baseController.ServeValidationErrors(errors)
		return false
	}

	return true
}

//** EXCEPTIONS

// ServeError prepares and serves an error exception
func (baseController *BaseController) ServeError(err error) {
	baseController.Data["json"] = struct {
		Error string `json:"error"`
	}{err.Error()}
	baseController.Ctx.Output.SetStatus(500)
	baseController.ServeJson()
}

// ServeValidationErrors prepares and serves a validation exception
func (baseController *BaseController) ServeValidationErrors(errors []string) {
	baseController.Data["json"] = struct {
		Errors []string `json:"errors"`
	}{errors}
	baseController.Ctx.Output.SetStatus(409)
	baseController.ServeJson()
}

//** CATCHING PANICS

// CatchPanic is used to catch any Panic and log exceptions. Returns a 500 as the response
func (baseController *BaseController) CatchPanic(functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		tracelog.WARN(baseController.Service.UserId, functionName, "PANIC Defered [%v] : Stack Trace : %v", r, string(buf))

		baseController.ServeError(fmt.Errorf("%v", r))
	}
}

//** AJAX SUPPORT

// AjaxResponse returns a standard ajax response
func (baseController *BaseController) AjaxResponse(resultCode int, resultString string, data interface{}) {
	response := struct {
		Result       int
		ResultString string
		ResultObject interface{}
	}{
		Result:       resultCode,
		ResultString: resultString,
		ResultObject: data,
	}

	baseController.Data["json"] = response
	baseController.ServeJson()
}
