package casbin

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gbootstrap"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"sync"
)

var (
	instance *Service
	once     sync.Once
)

type Service struct {
	Enforcer *casbin.SyncedEnforcer // 使用SyncedEnforcer以支持多线程环境下的线程安全
}

func New() *Service {
	once.Do(func() {
		// 初始化数据库适配器，指定自定义的Casbin规则表
		adapter, err := gormadapter.NewAdapterByDBWithCustomTable(ghub.Db, &model.CasbinRule{}, "sys_casbin_rule")
		if err != nil {
			gutils.ErrorPanic(err) // 错误处理：打印错误并退出
		}

		// 使用SyncedEnforcer以支持并发访问
		enforcer, err := casbin.NewSyncedEnforcer(gbootstrap.Cfg.Casbin.ModelFile, adapter)
		if err != nil {
			gutils.ErrorPanic(err) // 错误处理：打印错误并退出
		}

		instance = &Service{
			Enforcer: enforcer,
		}

		// 刷新策略规则
		instance.RefreshCasbinRule()
	})
	return instance
}

func (m *Service) clearDatabaseTables() error {
	// 清空数据库中的Casbin规则表
	if err := ghub.Db.Exec("DELETE FROM sys_casbin_rule").Error; err != nil {
		return err
	}
	// 重置 AUTO_INCREMENT
	if err := ghub.Db.Exec("ALTER TABLE sys_casbin_rule AUTO_INCREMENT = 1").Error; err != nil {
		return err
	}
	return nil
}

func (m *Service) RefreshCasbinRule() {

	ghub.Log.Info("casbin策略刷新")

	menus, err := dal.Menu.Preload(dal.Menu.Roles).Find()
	if err != nil {
		ghub.Log.Error(err)
		return
	}

	_ = m.clearDatabaseTables() // 清空数据库中旧的策略规则
	m.Enforcer.ClearPolicy()    // 清空执行器中的所有策略

	// 遍历菜单，为每个有效的角色和API端点添加策略
	for _, menu := range menus {
		if menu.Api != "" {
			for _, role := range menu.Roles {
				if role.Status {
					subject := fmt.Sprintf("u_%d", role.ID)
					// AddPolicy线程安全的
					if _, err := m.Enforcer.AddPolicy(subject, menu.Api, "All"); err != nil {
						ghub.Log.Error(fmt.Sprintf("添加策略失败：角色ID %d，API %s:", role.ID, menu.Api), err)
					}
				}
			}
		}
	}
}
