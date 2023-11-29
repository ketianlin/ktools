package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"reflect"
	"testing"
	"time"
)

func TestHttpGet(t *testing.T) {
	url := "https://movie.douban.com/j/search_tags?type=movie&source=index"
	headerEntity := map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	}
	res, err := ktools.Http.Get(url, headerEntity)
	if err != nil {
		fmt.Println("http.Get error: ", err.Error())
		return
	}
	fmt.Println(string(res))
}

func TestHttpPost(t *testing.T) {
	url := "http://localhost:8080/ktools/post"
	headerEntity := map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	}
	res, err := ktools.Http.Post(url, map[string]interface{}{
		"title":         "吊毛",
		"score":         99,
		"status":        true,
		"money":         9999.99,
		"register_time": time.Now(),
	}, headerEntity)
	if err != nil {
		fmt.Println("http.Get error: ", err.Error())
		return
	}
	fmt.Println(string(res))
}

func TestHttpFormPost(t *testing.T) {
	url := "http://localhost:8080/submit"
	headerEntity := map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	}
	res, err := ktools.Http.FormPost(url, map[string]interface{}{
		"name":          "吊毛",
		"email":         "diaomao@qq.com",
		"status":        true,
		"id":            666,
		"register_time": time.Now().Format(time.RFC3339),
	}, headerEntity)
	if err != nil {
		fmt.Println("http.Get error: ", err.Error())
		return
	}
	fmt.Println(string(res))
}

func checkType2[T any](data T) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Struct:
		fmt.Println("This is a struct")
	case reflect.Map:
		fmt.Println("This is a map")
	default:
		fmt.Println("Unknown type")
	}
}

func checkType(data interface{}) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Struct:
		fmt.Println("This is a struct")
	case reflect.Map:
		fmt.Println("This is a map")
	default:
		fmt.Println("Unknown type")
	}
}
