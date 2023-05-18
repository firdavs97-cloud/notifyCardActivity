package handlers

import "context"

const (
	Path = "/api/v1"
)

type Handler struct {
	ctx context.Context
}

func Init(ctx context.Context) Handler {
	h := Handler{ctx: ctx}
	h.initCardApi(Path)

	return h
}
