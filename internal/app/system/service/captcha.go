package service

import "context"

type ICaptcha interface {
	GetVerifyImgString(ctx context.Context) (id, b64s, answer string, err error)
	VerifyString(id, answer string) bool
}
