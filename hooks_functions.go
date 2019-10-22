// Copyright 2019 The niqingyang Authors. All rights reserved.
// Use of self source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://acme.top
// Author: niqingyang	niqy@qq.com
// TIME: 2019/10/22 17:25

package gohooks

// 添加 Filter
// priority 建议使用大于 0 的值
func AddFilter(tag string, filter Filter, priority int) bool {
	hooks := Instance()
	return hooks.AddFilter(tag, filter, priority)
}

// 移除指定 tag 中指针相同并且排序相同的 Filter
// priority 为 0 则忽略 priority 并移除 tag 下所有匹配的 Filter
func RemoveFilter(tag string, filter Filter, priority int) bool {
	hooks := Instance()
	return hooks.RemoveFilter(tag, filter, priority)
}

// 移除指定 tag 中 指定 priority 的 Filter
// priority 为 0 则忽略 priority 并移除 tag 下所有匹配的 Filter
func RemoveAllFilter(tag string, priority int) bool {
	hooks := Instance()
	return hooks.RemoveAllFilter(tag, priority)
}

// 判断 hooks 中是否已经存在指定的 Filter，如果存在则返回其 priority 列表
func HasFilter(tag string, filter Filter) ([]int, bool) {
	hooks := Instance()
	return hooks.HasFilter(tag, filter)
}

// 执行 Filter
func ApplyFilter(tag string, data interface{}, params ...interface{}) (interface{}, error) {
	hooks := Instance()
	return hooks.ApplyFilter(tag, data, params...)
}

// 添加 Action
// priority 建议使用大于 0 的值
func AddAction(tag string, action Action, priority int) bool {
	hooks := Instance()
	return hooks.AddAction(tag, action, priority)
}

// 移除指定 tag 中指针相同并且排序相同的 Action
// priority 为 0 则忽略 priority 并移除 tag 下所有匹配的 Action
func RemoveAction(tag string, action Action, priority int) bool {
	hooks := Instance()
	return hooks.RemoveAction(tag, action, priority)
}

// 移除指定 tag 中 指定 priority 的 Action
// priority 为 0 则忽略 priority 并移除 tag 下所有匹配的 Action
func RemoveAllAction(tag string, priority int) bool {
	hooks := Instance()
	return hooks.RemoveAllAction(tag, priority)
}

// 判断 hooks 中是否已经存在指定的 Action，如果存在则返回其 priority 列表
func HasAction(tag string, action Action) ([]int, bool) {
	hooks := Instance()
	return hooks.HasAction(tag, action)
}

// 执行 Action
func DoAction(tag string, params ...interface{}) {
	hooks := Instance()
	hooks.DoAction(tag, params...)
}
