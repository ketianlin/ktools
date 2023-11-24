package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

var m = map[string]interface{}{
	"title":  "吊毛",
	"age":    99,
	"status": true,
}

func TestGetKeys(t *testing.T) {
	keys := ktools.Map.Extend.GetKeys(m)
	fmt.Println(keys)
}

func TestGetValues2(t *testing.T) {
	keys := ktools.Map.Extend.GetValues(m)
	fmt.Println(keys)
}

func TestFilter(t *testing.T) {
	mm := ktools.Map.Extend.Filter(m, func(key string, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	fmt.Printf("%T\t%#v", mm, mm)
}

func TestMapValue(t *testing.T) {
	v := ktools.Map.Extend.MapValue(m, func(key string, value interface{}) interface{} {
		return value
	})
	fmt.Println(v)
}

func TestForEach(t *testing.T) {
	ktools.Map.Extend.ForEach(m, func(key string, value interface{}) {
		fmt.Println(key, value)
	})
}

// 判断map里面是否有这个key
func TestMapIsExist(t *testing.T) {
	exist := ktools.Map.Extend.IsExist(m, "title")
	fmt.Println(exist)
}

// 判断map里面是这个key的val是否为空
func TestMapIsEmptyV(t *testing.T) {
	exist := ktools.Map.Extend.IsEmptyV(m, "title")
	fmt.Println(exist)
}
