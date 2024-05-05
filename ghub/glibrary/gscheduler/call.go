package gscheduler

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	funcMap = map[string]interface{}{}
	lock    sync.RWMutex
)

// RegisterTask 注册方法到函数映射表
func RegisterTask(fnName string, task interface{}) {
	lock.Lock()
	defer lock.Unlock()
	funcMap[fnName] = task
}

// UnregisterTask 从函数映射表中解绑方法
func UnregisterTask(fnName string) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := funcMap[fnName]; ok {
		delete(funcMap, fnName)
	}
}

// callFunc 根据方法名和参数调用函数
func callFunc(funcName string, paramStr string) error {
	fmt.Println("callFunc-funcName" + funcName)
	fmt.Println("callFunc-paramStr" + paramStr)
	f, ok := funcMap[funcName]
	if !ok {
		return fmt.Errorf("未定义的函数: %s", funcName)
	}
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()

	// 如果函数不需要参数，直接调用
	if funcType.NumIn() == 0 {
		result := funcValue.Call(nil)
		if len(result) == 0 || result[0].IsNil() {
			return nil // 无返回值或返回 nil error
		}
		if err, ok := result[0].Interface().(error); ok {
			return err
		}
		return nil // 函数执行成功，没有error返回
	}

	params, err := parseParams(paramStr, funcType)
	if err != nil {
		return err
	}
	result := funcValue.Call(params)
	if len(result) == 0 {
		return nil // 假定没有返回值为成功
	}
	if err, ok := result[0].Interface().(error); ok && err != nil {
		return err // 返回值为 error 类型且非 nil，表示错误
	}
	return nil // 函数执行成功
}

// parseParams 解析参数字符串，将其转换为reflect.Value切片
func parseParams(paramStr string, funcType reflect.Type) ([]reflect.Value, error) {
	params := strings.Split(paramStr, "|")
	var values []reflect.Value
	if len(params) != funcType.NumIn() {
		return nil, fmt.Errorf("参数数量不匹配，需要%d个参数", funcType.NumIn())
	}
	for i, param := range params {
		paramType := funcType.In(i)
		switch paramType.Kind() {
		case reflect.String:
			values = append(values, reflect.ValueOf(param))
		case reflect.Int:
			intParam, err := strconv.Atoi(param)
			if err != nil {
				return nil, fmt.Errorf("参数解析错误: %s", err)
			}
			values = append(values, reflect.ValueOf(intParam))
		case reflect.Float64:
			floatParam, err := strconv.ParseFloat(param, 64)
			if err != nil {
				return nil, fmt.Errorf("参数解析错误: %s", err)
			}
			values = append(values, reflect.ValueOf(floatParam))
		case reflect.Bool:
			boolParam, err := strconv.ParseBool(param)
			if err != nil {
				return nil, fmt.Errorf("参数解析错误: %s", err)
			}
			values = append(values, reflect.ValueOf(boolParam))
		default:
			return nil, fmt.Errorf("不支持的参数类型: %s", paramType.Kind())
		}
	}
	return values, nil
}

//result, err := callFunc("add", "5|3")
//if err != nil {
//fmt.Println("Error:", err)
//} else {
//fmt.Println("Result:", result.Int())
//}
//
//result, err = callFunc("concat", "hello|world")
//if err != nil {
//fmt.Println("Error:", err)
//} else {
//fmt.Println("Result:", result.String())
//}
