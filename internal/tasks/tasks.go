package tasks

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gscheduler"
)

func init() {
	fmt.Println("-------------------注册定时任务------------------------")
	gscheduler.RegisterTask("Add", AddNum)
	gscheduler.RegisterTask("Demo", Demo)
}

func AddNum(a int, b int) error {
	fmt.Println(a + b)
	return nil
}

func Demo(a int, b string) error {
	fmt.Println(a)
	fmt.Println(b)
	return fmt.Errorf("测试错误")
}
