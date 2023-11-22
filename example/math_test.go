package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

func TestMath(t *testing.T) {
	random := ktools.Math.Random
	rI := random.IntFromRange(0, 10)
	rF := random.Float64FromRange(0, 10)
	fmt.Println("随机范围整数:", rI)
	fmt.Println("随机范围浮点数:", rF)
	fmt.Println("------------------------------------------------------------")
	b := random.Bool()
	fmt.Println("随机布尔值", b)
	fmt.Println("------------------------------------------------------------")
	condition := ktools.Math.Condition
	condition.IfThenElse(true, func() {
		fmt.Println("true")
	}, func() {
		fmt.Println("false")
	})
	condition.IfThen(true, func() {
		fmt.Println("true")
	})
	ifNil := condition.DefaultIfNil(nil, 1)
	fmt.Println("ifNil:", ifNil)
	fmt.Println("------------------------------------------------------------")
	ifNil = condition.DefaultIfNil(2, 1)
	fmt.Println("ifNotNil:", ifNil)
	fmt.Println("------------------------------------------------------------")
	nonNil := condition.FirstNonNil(nil, 3)
	fmt.Println("nonNil:", nonNil)
}
