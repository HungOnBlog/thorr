package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpRequester struct {
	client *http.Client
}

func NewHttpRequester() *HttpRequester {
	return &HttpRequester{
		client: &http.Client{},
	}
}

func (r *HttpRequester) Request(
	method string,
	url string,
	headers map[string]interface{},
	queries map[string]interface{},
	body map[string]interface{},
) (int, []byte, error) {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return 500, nil, err
	}

	for key, value := range headers {
		valueString := value.(string)
		req.Header.Add(key, valueString)
	}

	q := req.URL.Query()
	for key, value := range queries {
		valueString := value.(string)
		q.Add(key, valueString)
	}

	req.URL.RawQuery = q.Encode()

	bodyString, err := json.Marshal(body)
	if err != nil {
		return 500, nil, err
	}
	req.Body = io.NopCloser(bytes.NewBuffer(bodyString))
	fmt.Println("Request body: ", string(bodyString))

	resp, err := r.client.Do(req)
	if err != nil {
		return 500, nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 500, nil, err
	}

	resStatus := resp.StatusCode

	return resStatus, respBody, nil
}
