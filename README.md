# Beego Mgo Example

Copyright 2013 Ardan Studios. All rights reserved.  
Use of this source code is governed by a BSD-style license that can be found in the LICENSE handle.

This application provides a sample to use the beego web framework and the Go MongoDB driver mgo. This program connects to a public MongoDB at MongoLab. A single collection is available for testing. The configuration can be found in the app.conf file.

The project also includes several shell scripts to make building and running the web application easier.

Ardan Studios  
12973 SW 112 ST, Suite 153  
Miami, FL 33186  
bill@ardanstudios.com

### Installation

	-- Get, build and install the code
	go get github.com/goinggo/beego-mgo
	
	-- Run the code
	cd $GOPATH/src/github.com/goinggo/beego-mgo/zscripts
	./runbuild.sh
	
	-- Test Web Service API's
	
	This will return a single station from Mongo
	http://localhost:9000/station/42002
	
	This will return a collection of stations for the region
	http://localhost:9000/region/Gulf%20Of%20Mexico

### Notes About Architecture

I have been asked why I have organized the code in this way?

For me the controller should do nothing more than call into the business layer. The business layer contains the business logic for processing the request.

The models folder contains the data structures for the individual services. Each service places their models in a separate folder.

The services folder contain the raw service calls that the business layer would use to implement higher level functionality.

The controller methods just exist to receive the request and process the request through the business layer.

The more that can be abstracted into the base controller and base service the better. This way, adding a new functionality is simple and you don't need to worry about forgetting to do something important. Authentication always comes to mind.

The utilities folder is just that, support for the web application, mostly used by the services. You have exception handling support, extended logging support and the mongo support.

The abstraction layer for executing MongoDB queries and commands help hide the boilerplate code away into the base service and mongo utility code.
