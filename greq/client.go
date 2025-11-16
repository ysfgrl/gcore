package greq

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ysfgrl/gcore/gerror"
)

type HttpService struct {
	BaseUrl string
	Headers map[string]string
	Query   map[string]string
}

func (base *HttpService) Request(method string, path string, body []byte, query map[string]string) ([]byte, int, *gerror.Error) {
	req, err := http.NewRequest(method, base.BaseUrl+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, gerror.GetError(err)
	}
	for k, v := range base.Headers {
		req.Header.Set(k, v)
	}
	if query != nil {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, 0, gerror.GetError(err)
	}
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, gerror.GetError(err)
	}
	return resBytes, res.StatusCode, nil
}
