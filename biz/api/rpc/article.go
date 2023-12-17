package rpc

import (
	"context"
	"speedy/read/kitex_gen/speedy_read"
)

type ArticleHandler struct{}

func NewArticleHandler () *ArticleHandler {
	return &ArticleHandler{}
}

// Echo implements the SpeedyReadImpl interface.
func (s *ArticleHandler) Echo(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error) {
	return &speedy_read.Response{Message: req.Message}, nil
}

