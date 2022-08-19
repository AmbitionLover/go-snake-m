package controller

import (
	"context"
	"go-snake/app/svc/api/v1"
)

var (
	Echo = cEcho{}
)

type cEcho struct{}

func (h *cEcho) Say(ctx context.Context, req *v1.EchoSayReq) (res *v1.EchoSayRes, err error) {
	return &v1.EchoSayRes{Content: req.Content}, nil
}
