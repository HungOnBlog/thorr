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
		client: req.C(),
	}
}

func (h *HttpRequester) requestBuilder(test models.Test) *req.Request {
	c := h.client.SetBaseURL(test.Request.BaseURL)
	var r *req.Request

	switch test.Request.Method {
	case "GET":
		r = c.Get(test.Request.Path)
	case "POST":
		r = c.Post(test.Request.Path)
	case "PUT":
		r = c.Put(test.Request.Path)
	case "DELETE":
		r = c.Delete(test.Request.Path)
	default:
		r = c.Get(test.Request.Path)
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

func (h *HttpRequester) DoRequest(test models.Test) (int, map[string]interface{}, error) {
	r := h.requestBuilder(test)

	resp := r.Do()
	status := resp.GetStatusCode()
	body := resp.Body

	if status != test.Expected.Status {
		return status, nil, nil
	}

	returnBody, err := utils.ReadCloserToMapStringInterface(body)
	if err != nil {
		return status, nil, err
	}

	return status, returnBody, nil
}
