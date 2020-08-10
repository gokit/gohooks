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

	hooks.AddFilter("increase", func(data interface{}) (interface{}, error) {

		switch data := data.(type) {
		case int:
			return data + 1, nil
		}

		return data, nil
	})

	hooks.AddFilter("increase", func(data interface{}) (interface{}, error) {

		switch data := data.(type) {
		case int:
			return data + 1, nil
		}

		return data, nil
	})

	hooks.RemoveAllFilter("increase", 10)

	data, e :=hooks.ApplyFilter("increase", 0)

	fmt.Println(data, e)
}
