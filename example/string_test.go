package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

func TestStringExtend(t *testing.T) {
	extend := ktools.String.Extend
	s := "你好，吊毛！"
	//从索引2开始，截取长度为3
	sub := extend.Sub(s, 2, 3)
	fmt.Println("Substring:", sub)

	//从索引2开始，截取到最后
	sub = extend.Sub(s, 2)
	fmt.Println("Substring:", sub)

	//从索引-1开始，截取到最后
	sub = extend.Sub(s, -1)
	fmt.Println("Substring:", sub)

	//截取最后3个字符
	sub = extend.Sub(s, -3)
	fmt.Println("Substring:", sub)

	fmt.Println("len:", len(s))
	fmt.Println("extend.Len:", extend.Len(s))
}

func TestStringConv(t *testing.T) {
	type user struct {
		Id   int
		Name string
	}

	val := []interface{}{
		1,
		true,
		nil,
		"github",
		map[string]string{},
		[]string{},
		&user{},
	}

	for _, v := range val {
		cts := ktools.String.Conv.ConvToString(v)
		fmt.Printf("%T\t%v\n", cts, cts)
	}
}

func TestStringVerify(t *testing.T) {
	verify := ktools.String.Verify
	res := verify.IsEmail("abc@diaomao.com")
	fmt.Println("IsEmail:", res)

	res = verify.IsCNMobile("13211115555")
	fmt.Println("IsCNMobile:", res)

	res = verify.IsChinese("吊毛")
	fmt.Println("IsChinese:", res)

	res = verify.IsIDCard("350111197012121713")
	fmt.Println("IsIDCard:", res)

	res = verify.IsEnglish("hello")
	fmt.Println("IsEnglish:", res)

	res = verify.IsNumber("666")
	fmt.Println("IsNumber:", res)

	res = verify.IsLowerCase("hello")
	fmt.Println("IsLowerCase:", res)

	res = verify.IsUpperCase("HELLO")
	fmt.Println("IsUpperCase:", res)
}
