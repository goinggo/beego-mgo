// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package helper : catch.go implements boilerplate code for the web service.
package helper

import (
	"fmt"
	"runtime"
)

// CatchPanic is used to catch any Panic and log exceptions to Stdout. It will also write the stack trace
func CatchPanic(err *error, sessionID string, functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}
