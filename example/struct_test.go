package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

func TestStruct2Map(t *testing.T) {
	type user struct {
		Id   int
		Name string
	}
	u := user{666, "吊毛"}
	s2m := ktools.Struct.Struct2Map(u)
	fmt.Printf("%T\t%#v", s2m, s2m)
}

func TestAnyToStruct(t *testing.T) {
	type user struct {
		Id   int
		Name string
	}
	u := user{666, "吊毛"}
	s2m := ktools.Struct.AnyToMap(u)
	fmt.Printf("%T\t%#v", s2m, s2m)
}
