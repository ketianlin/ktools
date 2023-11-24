package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

func TestToJson(t *testing.T) {
	json := ktools.Json.ToJson(u)
	fmt.Printf("%T\t%#v\n", json, json)
}

func TestToObjS(t *testing.T) {
	str := `{"id":666,"name":"吊毛","status":true,"score":99.99,"created_at":1700797295}`
	var m member
	result := ktools.Json.ToObjS(str, m)
	fmt.Printf("%T\t%#v\n", result, result)
	fmt.Printf("%T\t%#v\n", m, m)
}

func TestGetValues(t *testing.T) {
	str := `{"id":666,"name":{"first":"吊毛","last":"一个"},"status":true,"score":99.99,"created_at":1700797295}`
	result := ktools.Json.GetValues(str, "name.first")
	fmt.Printf("%T\t%#v\n", result, result[0].Str)
}
