package gohooks

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// Filter 接口 - 对输入的参数进行过滤，返回值将作为下一个 Filter 的第一个入参
type Filter func(interface{}, ...interface{}) (interface{}, error)

// Action 接口
type Action func(...interface{})

type hooks struct {
	filters map[string]map[int][]Filter
	actions map[string]map[int][]Action
}

// 默认排序
const DefaultPriority = 10

// 移除所有 Filter 和 Action 时用到的排序值
const AllPriority = 0

// 添加 Filter
// priority 建议使用大于 0 的值
func (h *hooks) AddFilter(tag string, filter Filter, priority int) bool {

	if h.filters == nil {
		h.filters = make(map[string]map[int][]Filter)
	}

	if _, ok := h.filters[tag]; !ok {
		h.filters[tag] = make(map[int][]Filter)
	}

	if _, ok := h.filters[tag][priority]; !ok {
		h.filters[tag][priority] = []Filter{}
	}

	h.filters[tag][priority] = append(h.filters[tag][priority], filter)

	return true
}

// 移除指定 tag 中指针相同并且排序相同的 Filter
// priority 为 0 则忽略 priority 并移除 tag 下所有匹配的 Filter
func (h *hooks) RemoveFilter(tag string, filter Filter, priority int) bool {
	if h.filters == nil {
		return false
	}

	if _, ok := h.filters[tag]; !ok {
		return false
	}

	if _, ok := h.filters[tag][priority]; !ok && priority != AllPriority {
		return false
	}

	filterPtr := fmt.Sprintf("%v", filter)

	ok := false

	if priority == AllPriority {

		for p, filters := range h.filters[tag] {

			for i, f := range filters {

				fPtr := fmt.Sprintf("%v", f)

				if filterPtr == fPtr {
					h.filters[tag][p] = append(h.filters[tag][p][:i], h.filters[tag][p][i+1:]...)
					ok = true
				}

			}

		}

	} else {
		for i, f := range h.filters[tag][priority] {

			fPtr := fmt.Sprintf("%v", f)

			if filterPtr == fPtr {
				h.filters[tag][priority] = append(h.filters[tag][priority][:i], h.filters[tag][priority][i+1:]...)
				ok = true
			}
		}
	}

	return ok
}

// 移除指定 tag 中 指定 priority 的 Filter
// priority 为 0 则忽略 priority 并移除 tag 下所有匹配的 Filter
func (h *hooks) RemoveAllFilter(tag string, priority int) bool {

	if h.filters == nil {
		return false
	}

	if _, ok := h.filters[tag]; !ok {
		return false
	}

	if _, ok := h.filters[tag][priority]; !ok && priority != AllPriority {
		return false
	}

	if priority == AllPriority {
		for p, _ := range h.filters[tag] {
			h.filters[tag][p] = nil
		}
	} else {
		h.filters[tag][priority] = nil
	}

	return true
}

// 判断 hooks 中是否已经存在指定的 Filter，如果存在则返回其 priority 列表
func (h *hooks) HasFilter(tag string, filter Filter) ([]int, bool) {

	var priories []int

	if h.filters == nil {
		return priories, false
	}

	if _, ok := h.filters[tag]; !ok {
		return priories, false
	}

	filterPtr := fmt.Sprintf("%v", filter)

	for p, filters := range h.filters[tag] {

		for _, f := range filters {

			fPtr := fmt.Sprintf("%v", f)

			if filterPtr == fPtr {
				priories = append(priories, p)
			}

		}

	}

	return priories, len(priories) > 0
}

// 执行 Filter
func (h *hooks) ApplyFilter(tag string, data interface{}, params ...interface{}) (interface{}, error) {
	if _, ok := h.filters[tag]; !ok {
		return nil, errors.New("")
	}

	keys := make([]int, 0, len(h.filters[tag]))

	for key, _ := range h.filters[tag] {
		keys = append(keys, key)
	}

	sort.Ints(keys[:])

	var e error

	for _, key := range keys {
		filters := h.filters[tag][key]

		if filters == nil {
			continue
		}

		for _, filter := range filters {
			data, e = filter(data, params...)

			if e != nil {
				return data, e
			}
		}
	}

	return data, nil
}

// 添加 Action
// priority 建议使用大于 0 的值
func (h *hooks) AddAction(tag string, action Action, priority int) bool {

	if h.actions == nil {
		h.actions = make(map[string]map[int][]Action)
	}

	if _, ok := h.actions[tag]; !ok {
		h.actions[tag] = make(map[int][]Action)
	}

	if _, ok := h.actions[tag][priority]; !ok {
		h.actions[tag][priority] = []Action{}
	}

	h.actions[tag][priority] = append(h.actions[tag][priority], action)

	return true
}

// 移除指定 tag 中指针相同并且排序相同的 Action
// priority 为 0 则忽略 priority 并移除 tag 下所有匹配的 Action
func (h *hooks) RemoveAction(tag string, action Action, priority int) bool {
	if h.actions == nil {
		return false
	}

	if _, ok := h.actions[tag]; !ok {
		return false
	}

	if _, ok := h.actions[tag][priority]; !ok && priority != AllPriority {
		return false
	}

	actionPtr := fmt.Sprintf("%v", action)

	ok := false

	if priority == AllPriority {

		for p, actions := range h.actions[tag] {

			for i, a := range actions {

				aPtr := fmt.Sprintf("%v", a)

				if actionPtr == aPtr {
					h.actions[tag][p] = append(h.actions[tag][p][:i], h.actions[tag][p][i+1:]...)
					ok = true
				}

			}

		}

	} else {
		for i, a := range h.actions[tag][priority] {

			aPtr := fmt.Sprintf("%v", a)

			if actionPtr == aPtr {
				h.actions[tag][priority] = append(h.actions[tag][priority][:i], h.actions[tag][priority][i+1:]...)
				ok = true
			}
		}
	}

	return ok
}

// 移除指定 tag 中 指定 priority 的 Action
// priority 为 0 则忽略 priority 并移除 tag 下所有匹配的 Action
func (h *hooks) RemoveAllAction(tag string, priority int) bool {

	if h.actions == nil {
		return false
	}

	if _, ok := h.actions[tag]; !ok {
		return false
	}

	if _, ok := h.actions[tag][priority]; !ok && priority != AllPriority {
		return false
	}

	if priority == AllPriority {
		for p, _ := range h.actions[tag] {
			h.actions[tag][p] = nil
		}
	} else {
		h.actions[tag][priority] = nil
	}

	return true
}

// 判断 hooks 中是否已经存在指定的 Action，如果存在则返回其 priority 列表
func (h *hooks) HasAction(tag string, action Action) ([]int, bool) {

	var priories []int

	if h.actions == nil {
		return priories, false
	}

	if _, ok := h.actions[tag]; !ok {
		return priories, false
	}

	actionPtr := fmt.Sprintf("%v", action)

	for p, actions := range h.actions[tag] {

		for _, a := range actions {

			aPtr := fmt.Sprintf("%v", a)

			if actionPtr == aPtr {
				priories = append(priories, p)
			}

		}

	}

	return priories, len(priories) > 0
}

// 执行 Action
func (h *hooks) DoAction(tag string, params ...interface{}) {
	if _, ok := h.actions[tag]; !ok {
		return
	}

	keys := make([]int, 0, len(h.actions[tag]))

	for key, _ := range h.actions[tag] {
		keys = append(keys, key)
	}

	sort.Ints(keys[:])

	for _, key := range keys {
		actions := h.actions[tag][key]

		if actions == nil {
			continue
		}

		for _, action := range actions {
			action(params...)
		}
	}
}

var once sync.Once

var instance *hooks

// 获取 hooks 的单例
func Instance() *hooks {
	once.Do(func() {
		instance = &hooks{}
	})
	return instance
}

// 新建一个 hooks 实例
func NewInstance() *hooks {
	return &hooks{}
}
