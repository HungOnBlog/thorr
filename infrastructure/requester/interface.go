package requester

import "github.com/HungOnBlog/thorr/core/models"

type IRequester interface {
	DoRequest(test models.Test) (models.Result, error)
}
