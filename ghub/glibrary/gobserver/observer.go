package gobserver

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type Observer interface {
	Update(subject Subject)
}

// Observable 提供了观察者注册和通知的基础功能
type Observable struct {
	observers []Observer
}

// RegisterObserver 注册
func (o *Observable) RegisterObserver(observer Observer) {
	o.observers = append(o.observers, observer)
}

// RemoveObserver 移除
func (o *Observable) RemoveObserver(observer Observer) {
	// 实现观察者的移除逻辑
	for i, obs := range o.observers {
		if obs == observer {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers 通知
func (o *Observable) NotifyObservers() {
	for _, observer := range o.observers {
		observer.Update(o)
	}
}

//
//type 使用者 struct {
//	Observable // 继承Observable
//}
//
//func (m *使用者) Create() error {
//	// 这里可以包含创建菜单的逻辑
//	m.NotifyObservers() // 通知所有观察者
//	return nil
//}

// 观察者实现 接受通知方
type MenuCacheRefresher struct{}

func (mcr *MenuCacheRefresher) Update(subject Subject) {

}

//func main() {
//	menuService := new(MenuService)
//	refresher := new(MenuCacheRefresher)
//	menuService.RegisterObserver(refresher)
//
//	// 创建菜单示例
//	menuService.Create(&dto.MenuGenerateDTO{})
//}
