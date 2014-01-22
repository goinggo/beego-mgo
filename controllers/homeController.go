// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements the controller later for home
*/
package controllers

import (
	bc "github.com/goinggo/beego-mgo/controllers/baseController"
)

//** TYPES

type HomeController struct {
	bc.BaseController
}

//** WEB FUNCTIONS

func (this *HomeController) Get() {
	this.Data["Email"] = "bill@ardanstudios.com"
	this.TplNames = "index.html"
}
