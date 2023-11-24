package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
	"time"
)

type user struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	Score     float64 `json:"score"`
	CreatedAt int64   `json:"created_at"`
}

type member struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	Score     float64 `json:"score"`
	CreatedAt int64   `json:"created_at"`
}

var u = user{
	Id:        666,
	Name:      "吊毛",
	Status:    true,
	Score:     99.99,
	CreatedAt: time.Now().Unix(),
}

// 结构体转map
func TestStruct2Map(t *testing.T) {
	s2m := ktools.Struct.Struct2Map(u)
	fmt.Printf("%T\t%#v", s2m, s2m)
}

// 结构体转map[string]string
func TestStruct2MapString(t *testing.T) {
	s2m := ktools.Struct.Struct2MapString(u)
	fmt.Printf("%T\t%#v", s2m, s2m)
}

// 获取结构体的所有的key
func TestGetStructFields(t *testing.T) {
	s2m := ktools.Struct.GetStructFields(u)
	fmt.Printf("%T\t%#v", s2m, s2m)
}

// 获取结构体的所有的json的tag
func TestGetStructJsonTags(t *testing.T) {
	s2m := ktools.Struct.GetStructJsonTags(u)
	fmt.Printf("%T\t%#v", s2m, s2m)
}

func TestAnyToMap(t *testing.T) {
	s2m := ktools.Struct.AnyToMap(u)
	fmt.Printf("%T\t%#v", s2m, s2m)
}

// 一个结构体复制到另一个结构体
func TestCopyStruct(t *testing.T) {
	mem := new(member)
	ktools.Struct.CopyStruct(u, mem)
	fmt.Printf("%T\t%#v\n", u, u)
	fmt.Println("--------------------------------")
	fmt.Printf("%T\t%#v\n", mem, mem)
}

func TestClone(t *testing.T) {
	mem := new(member)
	ktools.Struct.Clone(u, mem)
	fmt.Printf("%T\t%#v\n", u, u)
	fmt.Println("--------------------------------")
	fmt.Printf("%T\t%#v\n", mem, mem)
}
