# Golang Hooks

参考 wordpress 的 hooks 机制实现的 golang hooks（the wordpress filter/action system in golang）

非线程安全

# Installation

```shell
$ go get -u github.com/gokit/gohooks
```

# Quick start

1. 获取 hooks 全局单例实例

```go
package instance

import "github.com/gokit/gohooks"

func main()  {
	hooks := gohooks.Instance()
	// ...
}
```

2. 获取一个新的 hooks 实例

```go
package instance

import "github.com/gokit/gohooks"

func main()  {
	hooks := gohooks.NewInstance()
	// ...
}
```

3. 添加 Filter

```go
package filter

import (
	"fmt"
	"github.com/gokit/gohooks"
)

func main()  {
	hooks := gohooks.Instance()

	hooks.AddFilter("increase", func(data interface{}, params ...interface{}) (interface{}, error) {
    
    		switch data := data.(type) {
    		case int:
    			return data + 1, nil
    		}
    
    		return data, nil
	}, gohooks.DefaultPriority)
    
	data, e :=hooks.ApplyFilter("increase", 1)
    
	fmt.Println(data, e)
}
```

4. 添加 Action

```go
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
```

5. 其他接口

```go
// 获取全局 hooks 单例
gohooks.Instance()
// 获取新的 hooks 实例
gohooks.NewInstance()
// 添加 Filter
hooks.AddFilter(tag, filter, priority)
// 移除 Filter
hooks.RemoveFilter(tag, filter, priority)
// 移除所有 Filter
hooks.RemoveAllFilter(tag, priority)
// 添加 Filter
hooks.HasFilter(tag, filter)
// 添加 Action
hooks.AddAction(tag, action, priority)
// 移除 Action
hooks.RemoveAction(tag, action, priority)
// 移除所有 Action
hooks.RemoveAllAction(tag, priority)
// 添加 Action
hooks.HasAction(tag, action)
```


