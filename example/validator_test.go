package example

import (
	"fmt"
	validatorT "github.com/ketianlin/ktools/validator"
	"testing"
)

// Member 字段都必填 contactUser字段 phone email address 三选一
type Member struct {
	Name            string         `v:"required,alphaunicode"` // v是因为 validate.SetTagName("v")
	Age             uint8          `v:"required,gte=10,lte=130"`
	Phone           string         `v:"required,e164"` // e164电话的一种格式
	Email           string         `v:"required,email"`
	FavouriteColor1 string         `v:"iscolor"`
	FavouriteColor2 string         `v:"hexcolor|rgb|rgba|hsl|hsla"` // 颜色格式 16进制
	Address         *Address       `v:"required"`
	ContactUser     []*ContactUser `v:"required,gte=1,dive"`                             // dive 深入到下一层验证
	Hobby           []string       `v:"required,gte=2,dive,required,gte=1,alphaunicode"` // 深入下一层验证，后续接长度大于等于2，并且是字母+unicode
	// 深入下一层验证，keys,alpha,gte=2,lte=10,endkeys 这一段是验证map的key， endkeys后面是跟对应value的校验
	Data map[string]string `v:"required,gte=2,dive,keys,alpha,gte=2,lte=10,endkeys,required,gte=2,alphaunicode"`
}

type ContactUser struct {
	Name    string   `v:"required,alphaunicode"`
	Age     uint8    `v:"gte=10,lte=130"`
	Phone   string   `v:"required_without_all=Email Address,omitempty,e164"`  //required_without_all 如果指定字段没有值，则必填  可以为空
	Email   string   `v:"required_without_all=Phone Address,omitempty,email"` // 可以为空
	Address *Address `v:"required_without_all=Phone Email"`
}

type Address struct {
	Province string `v:"required"`
	City     string `v:"required"`
}

func TestVerifyStruct(t *testing.T) {
	addr := &Address{
		Province: "福建",
		City:     "福州",
	}
	contactUser1 := &ContactUser{
		Name:  "吊毛",
		Age:   99,
		Phone: "+8613800009999",
	}
	contactUser2 := &ContactUser{
		Name:  "吊毛2",
		Age:   98,
		Email: "diaomao@qq.com",
	}
	member := Member{
		Name:  "yuan",
		Age:   18,
		Phone: "+8613800009999",
		//Email:           "diaomao@qq.com",
		Email:           "diaomao-qq.com",
		FavouriteColor1: "#ffff",
		FavouriteColor2: "rgb(255,255,255)",
		Address:         addr,
		ContactUser:     []*ContactUser{contactUser1, contactUser2},
		Hobby:           []string{"吹牛鼻", "西海岸", "rap"},
		Data:            map[string]string{"AB": "西海岸", "CD": "rap"},
	}
	err := validatorT.Enter[*Member](&member).VerifyStruct()
	fmt.Println("err: ", err)
}

func TestVerifyMap(t *testing.T) {
	req := map[string]interface{}{
		"domains":   "ab",
		"action":    "test",
		"startTime": "2011-11-11 11:11:11",
	}
	err := validatorT.Enter[*map[string]interface{}](&req).VerifyMap("dive,keys,endkeys,required,gte=4")
	fmt.Println(err)
}
