package grouter

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/ghub/glibrary/gvalidate"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"regexp"
)

// RouterAutoBind 通过反射绑定路由
func RouterAutoBind(R interface{}, public *gin.RouterGroup, protectedRouter *gin.RouterGroup) error {
	typ := reflect.TypeOf(R)
	val := reflect.ValueOf(R)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to struct for R")
	}
	regex, _ := regexp.Compile(`^Bind(.+)Controller$`)
	routerGroupType := reflect.TypeOf((*gin.RouterGroup)(nil))
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		if regex.MatchString(method.Name) {
			methodFunc := val.Method(i)
			methodType := methodFunc.Type()
			if methodType.NumIn() == 2 &&
				methodType.In(0) == routerGroupType &&
				methodType.In(1) == routerGroupType {
				methodFunc.Call([]reflect.Value{reflect.ValueOf(public), reflect.ValueOf(protectedRouter)})
			} else {
				fmt.Printf("Method %s does not have the correct signature.\n", method.Name)
			}
		}
	}
	return nil
}

// autoBindMiddleware 自动绑定和验证请求数据到指定的结构体
func autoBindMiddleware[T any](c *gin.Context, data T) (T, error) {
	// 如果类型为gdto.EmptyReq，跳过绑定
	if _, ok := any(data).(gdto.EmptyReq); ok {
		return data, nil
	}
	// 使用ShouldBind自动绑定数据
	if err := c.ShouldBind(&data); err != nil {
		//检查错误返回值
		err = gvalidate.ParseValidateErrors(err, &data)
		return data, err
	}
	// 添加更多验证逻辑，如果T类型实现了验证接口
	if mValidator, ok := any(data).(interface{ Validate() error }); ok {
		if err := mValidator.Validate(); err != nil {
			return data, err
		}
	}
	return data, nil
}

// HandlerFunc 是一个封装了请求数据处理逻辑的Handler函数
func HandlerFunc[T any](handler func(*gin.Context, T) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		req, err := autoBindMiddleware(c, req)
		if err != nil {
			gres.Fail(c, gres.Response{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    gin.H{},
			})
			return
		}

		response, err := handler(c, req)
		if err != nil {
			if res, ok := response.(gres.Response); ok {
				gres.Fail(c, res)
				return
			}
			gres.Fail(c, gres.Response{
				Message: err.Error(),
			})
			return
		}
		switch res := response.(type) {
		case gres.Response:
			gres.Ok(c, res)
		}
	}
}
