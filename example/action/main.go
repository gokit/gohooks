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

	test := func(data interface{}) {
		fmt.Println(data)
	}

	hooks.AddAction("test", test)
	hooks.AddAction("test", test)

	hooks.DoAction("test", 1)
}
