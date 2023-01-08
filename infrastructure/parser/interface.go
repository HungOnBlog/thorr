package parser

import "github.com/HungOnBlog/thorr/core/models"

type IParse interface {
	Parse(path string) (models.TestSuit, error)
}
