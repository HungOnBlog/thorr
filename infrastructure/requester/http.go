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

func (h *HttpRequester) requestBuilder(
	test models.Test,
	globalVariables *map[string]interface{},
) *req.Request {
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
		r.SetHeaders(utils.MapStringInterfaceToMapStringString(
			h.parseMapStringInterfaceWithGlobalVariables(
				test.Request.Header,
				globalVariables,
			),
		))
	}

	if test.Request.Query != nil {
		r.SetQueryParams(utils.MapStringInterfaceToMapStringString(test.Request.Query))
	}

	if test.Request.Body != nil {
		r.SetBody(test.Request.Body)
	}

	return r
}

func (h *HttpRequester) parseMapStringInterfaceWithGlobalVariables(
	m map[string]interface{},
	globalVariables *map[string]interface{},
) map[string]interface{} {
	rawMap := m
	for k, v := range rawMap {
		if utils.ContainerSubString(utils.InterfaceToString(v), "$") {
			key := utils.InterfaceToString(v)[1:]
			rawMap[k] = (*globalVariables)[key]
		}
	}

	return rawMap
}

func (h *HttpRequester) DoRequest(
	test models.Test,
	globalVariables *map[string]interface{},
) (models.Result, error) {
	r := h.requestBuilder(test, globalVariables)
	resp := r.Do()

	status := resp.GetStatusCode()

	headers := utils.HttpHeaderToMapStringString(resp.Header)

	body := resp.Body
	returnBody, err := utils.ReadCloserToInterface(body)
	if err != nil {
		return models.Result{}, err
	}

	result := models.Result{
		Status:  status,
		Headers: headers,
		Body:    returnBody,
	}

	return result, nil
}
