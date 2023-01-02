// 本包提供了增强的context
package context

import (
	"context"
	_ "unsafe"
)

/*
map的抽象

&sync.Map和GoMap(map[interface{}]interface{})实现了本接口
*/
type Map interface {
	//获取key对应的value,ok表示是否获取成功
	Load(key interface{}) (value interface{}, ok bool)
}

// 为map[interface{}]interface{}实现Map
type GoMap map[interface{}]interface{}

// 实现Map
func (m GoMap) Load(key interface{}) (value interface{}, ok bool) {
	value, ok = m[key]
	return
}

type mapContext struct {
	context.Context
	m Map
}

func (m *mapContext) Value(key interface{}) interface{} {
	value, ok := m.m.Load(key)
	if !ok {
		return ctx_value(m.Context, key)
	}
	return value
}

//go:linkname ctx_value context.value
func ctx_value(c context.Context, key interface{}) interface{}

// WithValue返回带有m的所有key,value的context.Context
func WithValue(parent context.Context, m Map) context.Context {
	return &mapContext{Context: parent, m: m}

}
