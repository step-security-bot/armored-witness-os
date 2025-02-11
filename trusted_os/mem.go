// Copyright 2022 The Armored Witness OS authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "unsafe"

	"github.com/usbarmory/tamago/dma"
)

const (
	// Secure Monitor
	secureStart = 0x80000000
	secureSize  = 0x0e000000 // 224MB

	// Secure Monitor DMA
	secureDMAStart = 0x8e000000
	secureDMASize  = 0x02000000 // 32MB

	// Secure Monitor Applet
	appletStart = 0x90000000
	appletSize  = 0x10000000 // 256MB
)

//go:linkname ramStart runtime.ramStart
var ramStart uint32 = secureStart

//go:linkname ramSize runtime.ramSize
var ramSize uint32 = secureSize

var appletRegion *dma.Region

func init() {
	appletRegion, _ = dma.NewRegion(appletStart, appletSize, false)
	appletRegion.Reserve(appletSize, 0)

	dma.Init(secureDMAStart, secureDMASize)
}
