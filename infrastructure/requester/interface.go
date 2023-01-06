package requester

import "github.com/HungOnBlog/thorr/core/models"

type IRequester interface {
	DoRequest(test models.Test) (map[string]interface{}, error)
}
