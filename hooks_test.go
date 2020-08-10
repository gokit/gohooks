// Copyright 2019 The niqingyang Authors. All rights reserved.
// Use of self source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://acme.top
// Author: niqingyang	niqy@qq.com
// TIME: 2019/10/2 21:31

package gohooks

import (
	"testing"
)

var filter = func(data interface{}) (interface{}, error) {

	switch data := data.(type) {
	case int:
		return data + 1, nil
	case string:
		return data + "," + data, nil
	case []string:
		return append(data, data[len(data)-1]), nil
	}

	return data, nil
}

var actionNum = 0

var action = func(data interface{}) {
	switch data := data.(type) {
	case int:
		actionNum += data
	}
}

func TestFilter(t *testing.T) {
	hooks := New()

	hooks.AddFilter("test", filter, 1)
	hooks.AddFilter("test", filter, 2)
	hooks.AddFilter("test", filter, 3)

	data, _ := hooks.ApplyFilter("test", 0)

	if data != 3 {
		t.Error(`hooks.ApplyFilter("test", 0) != 3`, data)
	}
}

func TestHasFilter(t *testing.T) {
	hooks := New()

	hooks.AddFilter("test", filter, 1)
	hooks.AddFilter("test", filter, 2)
	hooks.AddFilter("test", filter, 3)
	hooks.AddFilter("test", filter, 4)
	hooks.AddFilter("test", filter, 4)

	priorites, ok := hooks.HasFilter("test", filter)

	if len(priorites) != 4 || ok == false {
		t.Error(`hooks.HasFilter("test", filter) != [1,2,3,4]`, priorites)
	}
}

func TestRemoveFilter(t *testing.T) {
	hooks := New()

	hooks.AddFilter("test", filter, 1)
	hooks.AddFilter("test", filter, 2)
	hooks.AddFilter("test", filter, 3)

	data, e := hooks.ApplyFilter("test", 0)

	if data != 3 {
		t.Error(`hooks.ApplyFilter("test", 0) != 3`, e, data)
	}

	{
		hooks.RemoveFilter("test", filter, 1)

		data, _ := hooks.ApplyFilter("test", 0)

		if data != 2 {
			t.Error(`hooks.ApplyFilter("test", 0) != 2`, data)
		}
	}

}

func TestRemoveAllFilter(t *testing.T) {
	hooks := New()

	hooks.AddFilter("test", filter, 1)
	hooks.AddFilter("test", filter, 2)
	hooks.AddFilter("test", filter, 3)

	data, e := hooks.ApplyFilter("test", 0)

	if data != 3 {
		t.Error(`hooks.ApplyFilter("test", 0) != 3`, e)
	}

	{
		hooks.RemoveAllFilter("test", AllPriority)

		data, _ := hooks.ApplyFilter("test", 0)

		if data != 0 {
			t.Error(`hooks.ApplyFilter("test", 0) != 0`, data)
		}
	}

}

func TestAction(t *testing.T) {
	hooks := New()

	// 重置
	actionNum = 0

	hooks.AddAction("test", action, 1)
	hooks.AddAction("test", action, 2)
	hooks.AddAction("test", action, 3)

	hooks.DoAction("test", 1)

	if actionNum != 3 {
		t.Error(`actionNum != 3`, actionNum)
	}
}

func TestHasAction(t *testing.T) {
	hooks := New()

	hooks.AddAction("test", action, 1)
	hooks.AddAction("test", action, 2)
	hooks.AddAction("test", action, 3)
	hooks.AddAction("test", action, 4)
	hooks.AddAction("test", action, 4)

	priorites, ok := hooks.HasAction("test", action)

	if len(priorites) != 4 || ok == false {
		t.Error(`hooks.HasAction("test", filter) != [1,2,3,4]`, priorites)
	}
}

func TestRemoveAction(t *testing.T) {
	hooks := New()

	// 重置
	actionNum = 0

	hooks.AddAction("test", action, 1)
	hooks.AddAction("test", action, 2)
	hooks.AddAction("test", action, 3)

	hooks.DoAction("test", 1)

	if actionNum != 3 {
		t.Error(`actionNum != 3`, actionNum)
	}

	// 重置
	actionNum = 0

	{
		hooks.RemoveAction("test", action, 1)

		hooks.DoAction("test", 1)

		if actionNum != 2 {
			t.Error(`actionNum != 2`, actionNum)
		}
	}

}

func TestRemoveAllAction(t *testing.T) {
	hooks := New()

	// 重置
	actionNum = 0

	hooks.AddAction("test", action, 1)
	hooks.AddAction("test", action, 2)
	hooks.AddAction("test", action, 3)

	hooks.DoAction("test", 1)

	if actionNum != 3 {
		t.Error(`actionNum != 3`, actionNum)
	}

	// 重置
	actionNum = 0

	{
		hooks.RemoveAllAction("test", AllPriority)

		hooks.DoAction("test", 1)

		if actionNum != 0 {
			t.Error(`actionNum != 0`, actionNum)
		}
	}

}