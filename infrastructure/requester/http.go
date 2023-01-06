package requester

import (
	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/utils"
	"github.com/imroc/req/v3"
)

type HttpRequester struct {
	client *req.Client
}

func NewHttpRequester() *HttpRequester {
	return &HttpRequester{
		client: req.C().DevMode(),
	}
}

func (h *HttpRequester) requestBuilder(test models.Test) *req.Request {
	c := h.client.SetBaseURL(test.Request.BaseURL)
	var r *req.Request

	switch test.Request.Method {
	case "GET":
		r = c.Get()
	case "POST":
		r = c.Post()
	case "PUT":
		r = c.Put()
	case "DELETE":
		r = c.Delete()
	default:
		r = c.Get()
	}

	if test.Request.Header != nil {
		r.SetHeaders(utils.MapStringInterfaceToMapStringString(test.Request.Header))
	}

	if test.Request.Query != nil {
		r.SetQueryParams(utils.MapStringInterfaceToMapStringString(test.Request.Query))
	}

	if test.Request.Body != nil {
		r.SetBody(test.Request.Body)
	}

	return r
}

func (h *HttpRequester) DoRequest(test models.Test) (map[string]interface{}, error) {
	var resp map[string]interface{}
	r := h.requestBuilder(test)

	err := r.Do().Into(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
