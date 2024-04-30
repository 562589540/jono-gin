package validate

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

//func RegisterCostValidator() {
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		_ = v.RegisterValidation("first_is_a", firstIsA)
//	}
//}

func RegisterCustomValidators() {
	fmt.Println("------------------注册验证器---------------------")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	valType := reflect.TypeOf(Validator{}) // 注意这里使用的是类型，而不是指针
	for i := 0; i < valType.NumMethod(); i++ {
		method := valType.Method(i)
		methodName := method.Name

		if strings.HasPrefix(methodName, "Validate") && checkValidatorSignature(method.Type) {
			ruleName := strings.ToLower(strings.TrimPrefix(methodName, "Validate"))
			// 断言函数类型，并注册
			if validateFunc, ok := method.Func.Interface().(func(Validator, validator.FieldLevel) bool); ok {
				wrapperFunc := func(fl validator.FieldLevel) bool {
					return validateFunc(Validator{}, fl) // 使用一个新的 Validator 实例调用函数
				}
				if err := v.RegisterValidation(ruleName, wrapperFunc); err != nil {
					panic(err)
				}
				fmt.Println("------------------注册验证器成功---------------------", ruleName)
			}
		}
	}
}

func checkValidatorSignature(t reflect.Type) bool {
	// 方法应有两个输入：一个是 Validator 结构体的接收者，一个是符合 validator.FieldLevel 的参数
	if t.NumIn() != 2 {
		fmt.Printf("Incorrect number of inputs: got %d, want 2\n", t.NumIn())
		return false
	}
	// 确认第二个参数实现了 validator.FieldLevel 接口
	if !t.In(1).Implements(reflect.TypeOf((*validator.FieldLevel)(nil)).Elem()) {
		fmt.Printf("Second input does not implement validator.FieldLevel: %v\n", t.In(1))
		return false
	}
	// 确认有一个布尔型返回值
	if t.NumOut() != 1 || t.Out(0).Kind() != reflect.Bool {
		fmt.Printf("Incorrect output: got %d outputs, first output is not bool\n", t.NumOut())
		return false
	}
	return true
}

type Validator struct{}
