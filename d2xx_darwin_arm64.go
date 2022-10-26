// Copyright 2022 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

//go:build cgo && !no_d2xx
// +build cgo,!no_d2xx

package d2xx

/*
#cgo LDFLAGS: -framework CoreFoundation -framework IOKit ${SRCDIR}/third_party/libftd2xx_darwin_arm64_v1.4.24.a
*/
import "C"
