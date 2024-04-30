package gdao

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/gstruct/g"
	"gorm.io/gorm"
)

// BaseDao 提供了一个泛型版本的基础DAO接口
type BaseDao[T any] struct {
	DB *gorm.DB
}

// NewBaseDao 创建一个新的泛型DAO实例
func NewBaseDao[T any]() *BaseDao[T] {
	return &BaseDao[T]{DB: ghub.Db}
}

// Model Create 插入新记录
func (d *BaseDao[T]) Model() *gorm.DB {
	var entity T
	return d.DB.Model(&entity)
}

// Create 插入新记录
func (d *BaseDao[T]) Create(entity *T) error {
	return d.DB.Create(entity).Error
}

// DeleteById 通过ID删除记录，需要从外部传递一个实体的指针来指定类型
func (d *BaseDao[T]) DeleteById(id uint) error {
	var entity T
	return d.DB.Delete(&entity, id).Error
}

// UnscopedDeleteById 通过ID删除记录硬删除
func (d *BaseDao[T]) UnscopedDeleteById(id uint) error {
	var entity T
	return d.DB.Unscoped().Delete(&entity, id).Error
}

// UnscopedBatchDelete  批量映删除
func (d *BaseDao[T]) UnscopedBatchDelete(ids []uint) error {
	var entity T
	return d.DB.Unscoped().Delete(&entity, ids).Error
}

// DeleteAll 清空表数据
func (d *BaseDao[T]) DeleteAll() error {
	// 使用GORM的Delete方法删除所有记录
	// 传递一个空的model.LoginLog实例指针到Delete方法中
	// 要删除所有记录，不传递任何where条件
	var entity T
	return d.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity).Error
}

func (d *BaseDao[T]) UpdateValue(id uint, key string, value any) error {
	var entity T
	result := d.DB.Model(&entity).Where("id = ?", id).Update(key, value)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到要更新的数据")
	}
	return nil
}

// BatchDelete 批量删除
func (d *BaseDao[T]) BatchDelete(ids []uint) error {
	var entity T
	return d.DB.Delete(&entity, ids).Error
}

// GetById 通过ID获取记录
func (d *BaseDao[T]) GetById(id uint) (*T, error) {
	var entity T
	err := d.DB.First(&entity, id).Error
	return &entity, err
}

// CRPFind 责任链的数据查询
func (d *BaseDao[T]) CRPFind(options *g.FindOptions) ([]T, int64, error) {
	var entity T
	query := d.DB.Model(&entity)
	// 设置查询责任链
	gormCrp := ghub.GormCrp[T]()
	// 构造条件查询
	gormCrp.AddHandler(ghub.ConditionHandler())
	// 构造复杂条件查询
	gormCrp.AddHandler(ghub.ModifierHandler())
	// 构造模糊查询
	gormCrp.AddHandler(ghub.LikeHandler())
	// 构造排序
	gormCrp.AddHandler(ghub.OrderHandler())
	// 构造关联
	gormCrp.AddHandler(ghub.PreloadHandler())
	//获取查询结果
	return gormCrp.ExecuteQuery(query, options)
}
