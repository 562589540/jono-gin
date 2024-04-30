package gvalidate

import (
	"errors"
	"fmt"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// ParseValidateErrors 错误解析返回值标签信息
func ParseValidateErrors(errs error, target any) error {
	var errResult error

	var errValidation validator.ValidationErrors
	ok := errors.As(errs, &errValidation)
	if !ok {
		return errs
	}

	//通过反射获取指定元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, exists := fields.FieldByName(fieldErr.Field())
		if !exists {
			continue
		}
		//自定义错误标签
		errMessageTag := fmt.Sprintf("%s_m", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)

		if errMessage == "" {
			//默认错误标签
			errMessage = field.Tag.Get("m")
		}

		if errMessage == "" {
			errMessage = fmt.Sprintf("%s:%s Error", fieldErr.Field(), fieldErr.Tag())
		}

		errResult = gutils.AppendError(errResult, fmt.Errorf(errMessage))
	}
	return errResult
}
