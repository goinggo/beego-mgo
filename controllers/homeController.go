// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements the controller later for home
*/
package controllers

import (
	cb "github.com/goinggo/beego-mgo/controllers/base"
)

//** TYPES

type HomeController struct {
	cb.BaseController
}

//** WEB FUNCTIONS

func (this *HomeController) Get() {
	this.Data["Website"] = "ArdanStudios.com"
	this.Data["Email"] = "bill@ardanstudios.com"
	this.TplNames = "index.tpl"
}
