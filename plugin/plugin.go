//  Copyright (c) 2020, pole-group. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package plugin

import (
	"github.com/pole-group/pole/common"
)

type Plugin interface {
	Name() string

	Init(ctx *common.ContextPole)

	Run()

	Destroy()
}

type TransportPlugin interface {
	Plugin
}

type StoragePlugin interface {
	Plugin
}
