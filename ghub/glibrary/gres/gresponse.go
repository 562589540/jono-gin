package gres

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Response struct {
	Status  int    `json:"-"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Total   int64  `json:"total,omitempty"`
	Success bool   `json:"success,omitempty"`
}

func (m Response) IsEmpty() bool {
	return reflect.DeepEqual(m, Response{})
}

func HttpResponse(ctx *gin.Context, status int, resp Response) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
	}
	ctx.AbortWithStatusJSON(status, resp)
	//AbortWithStatusJSON 会终止以后的代码执行
	//Json返回后会继续执行后续代码
}

func buildStatus(resp Response, defaultStatus int) int {
	if 0 == resp.Status {
		return defaultStatus
	}
	return resp.Status
}

func Ok(ctx *gin.Context, resp Response) {
	if resp.Code == 0 {
		resp.Code = 200
	}
	HttpResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

func Error(ctx *gin.Context, resp Response) {
	if resp.Code == 0 {
		resp.Code = 400
	}
	HttpResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

func Fail(ctx *gin.Context, resp Response) {
	if resp.Code == 0 {
		resp.Code = 400
	}
	HttpResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

func ServerFail(ctx *gin.Context, resp Response) {
	HttpResponse(ctx, buildStatus(resp, http.StatusInternalServerError), resp)
}
