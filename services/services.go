// Copyright 2013 Ardan Studios. All rights reserved.
// Use of service source code is governed by a BSD-style
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

func (service *Service) Prepare() (err error) {
	service.MongoSession, err = mongo.CopyMonotonicSession(service.UserId)
	if err != nil {
		tracelog.Error(err, service.UserId, "Service.Prepare")
		return err
	}

	return err
}

func (service *Service) Finish() (err error) {
	defer helper.CatchPanic(&err, service.UserId, "Service.Finish")

	if service.MongoSession != nil {
		mongo.CloseSession(service.UserId, service.MongoSession)
		service.MongoSession = nil
	}

	return err
}

// Execute the MongoDB literal function
func (service *Service) DBAction(databaseName string, collectionName string, mongoCall mongo.MongoCall) (err error) {
	return mongo.Execute(service.UserId, service.MongoSession, databaseName, collectionName, mongoCall)
}
