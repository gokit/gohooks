// Copyright 2019 The niqingyang Authors. All rights reserved.
// Use of self source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://acme.top
// Author: niqingyang	niqy@qq.com
// TIME: 2019/10/3 19:43

package main

import (
	"fmt"
	"github.com/gokit/gohooks"
)

func main() {
	hooks := gohooks.Instance()

	hooks.AddAction("test", func(params ...interface{}) {

		fmt.Println(params)

	}, gohooks.DefaultPriority)

	hooks.DoAction("test", 1, 2, 3, 4)
}
