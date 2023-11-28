package httpT

import (
	"encoding/json"
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

// Post 这里的contenttype只支持application/json, reqEntity这个可以是任意类型
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
