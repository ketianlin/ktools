package httpT

import (
	"encoding/json"
	structT "github.com/ketianlin/ktools/struct"
	"github.com/valyala/fasthttp"
)

func (e *Enter[T]) Get(url string, headerEntity ...map[string]string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完须要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	// 默认是application/x-www-form-urlencoded
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	if len(headerEntity) > 0 {
		for k, v := range headerEntity[0] {
			req.Header.Add(k, v)
		}
	}
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}
	//fmt.Println("result:\r\n", string(resp.Body()))
	return resp.Body(), nil
}

// Post 这里的contenttype只支持application/json, reqEntity这个可以是任意类型,reqEntity如果是结构体最好有设置json的tag
func (e *Enter[T]) Post(url string, reqEntity T, headerEntity ...map[string]string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完须要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	// 默认是application/x-www-form-urlencoded
	if len(headerEntity) > 0 {
		for k, v := range headerEntity[0] {
			req.Header.Add(k, v)
		}
	}
	requestBody, _ := json.Marshal(reqEntity)
	req.SetBody(requestBody)
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}
	//fmt.Println("result:\r\n", string(resp.Body()))
	return resp.Body(), nil
}

// FormPost 这里的contenttype只支持application/x-www-form-urlencoded, reqEntity这个可以是任意类型,reqEntity如果是结构体最好有设置json的tag
func (e *Enter[T]) FormPost(url string, reqEntity T, headerEntity ...map[string]string) ([]byte, error) {
	// 创建一个 HTTP 请求
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// 设置请求 URL 和方法
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/x-www-form-urlencoded")

	if len(headerEntity) > 0 {
		for k, v := range headerEntity[0] {
			req.Header.Add(k, v)
		}
	}

	m := structT.Enter[any]{}.AnyToMap(reqEntity)
	// 设置请求表单数据
	args := req.PostArgs()
	for k, v := range m {
		args.Set(k, v)
	}

	// 发送请求并获取响应
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	// 输出响应结果
	//fmt.Println(resp.StatusCode())
	//fmt.Println(string(resp.Body()))

	return resp.Body(), nil
}
