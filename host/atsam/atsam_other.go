// Copyright 2017 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// +build !arm

package atsam

const isArm = false

const isLinux = false

func isErrBusy(err error) bool {
	// This function is not used on non-linux.
	return false
}
