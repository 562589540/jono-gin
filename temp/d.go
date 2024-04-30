package temp

//var (
//	rolesDao *RolesDao
//	once     sync.Once
//)
//
//type RolesDao struct {
//	*dao.BaseDao[model.RolesBase]
//}
//
//func NewRolesDao() *RolesDao {
//	once.Do(func() {
//		rolesDao =
//	})
//	return rolesDao
//}

//var rolesDao *RolesDao
//
//type RolesDao struct {
//	*dao.BaseDao[model.RolesBase]
//}
//
//func NewRolesDao() *RolesDao {
//	if rolesDao == nil {
//		rolesDao = &RolesDao{dao.NewBaseDaoG[model.RolesBase]()}
//	}
//	return rolesDao
//}
//
//func (m *RolesDao) GetRolesList(dto *dto.QueryRolesDTO) ([]model.RolesBase, int64, error) {
//	var list []model.RolesBase
//
//	var total int64
//
//	query := m.DB.Model(&model.RolesBase{})
//
//	query = m.BuildQuery(query, dto)
//
//	// 首先计算总数
//	if err := query.IsCount(&total).Error; err != nil {
//		return nil, total, err
//	}
//
//	// 有分页需求查分页
//	if dto.DPaginate.Page != 0 && dto.DPaginate.Limit != 0 {
//		query = query.Scopes(dao.DPaginate(dto.DPaginate))
//	}
//
//	// 查询具体数据
//	if err := query.Find(&list).Error; err != nil {
//		return nil, total, err
//	}
//
//	return list, total, nil
//}
//
//func (m *RolesDao) AddRoles(dto *dto.GenerateRolesDTO) error {
//	return m.DB.Create(&model.RolesBase{
//		NameZh:   dto.NameZh,
//		Code:   dto.Code,
//		Remark: dto.Remark,
//		Status: true,
//	}).Error
//}
//
//func (m *RolesDao) UpdateRoles(mDto *dto.GenerateRolesDTO) error {
//	var mModel model.RolesBase
//	err := m.DB.First(&mModel, mDto.ID).Error
//	if err != nil {
//		return err
//	}
//	mDto.ToModel(&mModel)
//	return m.DB.Save(&mModel).Error
//}
//
//func (m *RolesDao) DelRolesById(id uint) error {
//	return m.DB.Delete(&model.RolesBase{}, id).Error
//}
//
//func (m *RolesDao) GetRolesById(id uint) (model.RolesBase, error) {
//	var mModel model.RolesBase
//	err := m.DB.First(&mModel, id).Error
//	return mModel, err
//}
//
//func (m *RolesDao) BuildQuery(query *gorm.DB, dto *dto.QueryRolesDTO) *gorm.DB {
//	if dto.Code != "" {
//		query = query.Where("code = ?", dto.Code)
//	}
//	if dto.NameZh != "" {
//		query = query.Where("name = ?", dto.NameZh)
//	}
//	if dto.Status != "" {
//		if dto.Status == "1" {
//			query = query.Where("status = ?", 1)
//		} else {
//			query = query.Where("status = ?", 0)
//		}
//	}
//	return query
//}
