package gutils

func InitCasbin() {

	//// 初始化Casbin的GORM适配器
	//adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "", "casbin_rule")
	//if err != nil {
	//	panic("failed to initialize adapter")
	//}
	//
	//// 初始化Casbin enforcer
	//enforcer, err := casbin.NewEnforcer("path/to/rbac_model.conf", adapter)
	//if err != nil {
	//	panic("failed to create enforcer")
	//}
	//
	//// 从数据库加载策略
	//err = enforcer.LoadPolicy()
	//if err != nil {
	//	panic("failed to load policy from DB")
	//}
}
