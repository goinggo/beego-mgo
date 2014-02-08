// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Implements boilerplate code for all services
*/
package services

import (
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	"github.com/goinggo/tracelog"
	"labix.org/v2/mgo"
)

//** TYPES

type (
	// Services contains common properties
	Service struct {
		MongoSession *mgo.Session
		UserId       string
	}
)

//** PUBLIC FUNCTIONS

func (this *Service) Prepare() (err error) {
	this.MongoSession, err = mongo.CopyMonotonicSession(this.UserId)
	if err != nil {
		tracelog.ERROR(err, this.UserId, "Service.Prepare")
		return err
	}

	return err
}

func (this *Service) Finish() (err error) {
	defer helper.CatchPanic(&err, this.UserId, "Service.Finish")

	if this.MongoSession != nil {
		mongo.CloseSession(this.UserId, this.MongoSession)
		this.MongoSession = nil
	}

	return err
}

// Execute the MongoDB literal function
func (this *Service) DBAction(databaseName string, collectionName string, mongoCall mongo.MongoCall) (err error) {
	return mongo.Execute(this.UserId, this.MongoSession, databaseName, collectionName, mongoCall)
}
