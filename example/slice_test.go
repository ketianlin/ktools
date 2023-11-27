package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

// 测试将float64切片转换为字符串切片
func TestFloat64ToS(t *testing.T) {
	fs := []float64{12.3, 14.23, 15.44}
	fs2 := ktools.Slice.Conv.Float64ToS(fs)
	for _, v := range fs2 {
		fmt.Printf("%T\t%#v\n", v, v)
	}
}

// 测试将将float64切片转换为interface切片
func TestFloat64ToI(t *testing.T) {
	fs := []float64{12.3, 14.23, 15.44}
	fs2 := ktools.Slice.Conv.Float64ToI(fs)
	for _, v := range fs2 {
		fmt.Printf("%T\t%#v\n", v, v)
	}
}

// 测试判断切片是否存在某个值
func TestSliceIsExist(t *testing.T) {
	fs := []float64{12.3, 14.23, 15.44, 14.23, 15.99, 12.3, 1.23, 12.3}
	exist := ktools.Slice.Extend.IsExist(fs, 12.3)
	fmt.Printf("%T\t%v\n", exist, exist)
}

// 测试去重切片
func TestSliceUnique(t *testing.T) {
	fs := []float64{12.3, 14.23, 15.44, 14.23, 15.99, 12.3, 1.23, 12.3}
	fs2 := ktools.Slice.Extend.Unique(fs)
	for _, v := range fs2 {
		fmt.Printf("%T\t%#v\n", v, v)
	}
}

// 测试删除切片中的某个值
func TestSliceRemove(t *testing.T) {
	fs := []float64{12.3, 14.23, 15.44}
	fs2 := ktools.Slice.Extend.Remove(fs, 14.23)
	for _, v := range fs2 {
		fmt.Printf("%T\t%#v\n", v, v)
	}
}

// 测试合并泛型数组
func TestSliceInterfaces(t *testing.T) {
	lst1 := []interface{}{"hello", "diaomao", 69, false}
	lst2 := []interface{}{12, 13.5, "a", true}
	lst := ktools.Slice.Merge.Interfaces(lst1, lst2)
	for _, v := range lst {
		fmt.Printf("%T\t%#v\n", v, v)
	}
}
