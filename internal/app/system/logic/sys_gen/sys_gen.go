package sys_gen

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/enum"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	instance service.IGenService
	once     sync.Once
)

func New() service.IGenService {
	once.Do(func() {
		logMode := logger.Info
		if !ghub.Cfg.Mode.Develop {
			logMode = logger.Error
		}
		dsn := "root:112233@tcp(localhost:3306)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logMode),
		})
		if err != nil {
			panic("failed to connect database")
		}
		instance = &Service{
			db:       db,
			dateBase: "gin",
		}
	})
	return instance
}

type Service struct {
	db       *gorm.DB
	dateBase string
}

func (m *Service) TableList(ctx context.Context, req dto.TableInfoSearchReq) ([]dto.TableInfoRes, int64, error) {
	var tables []model.TableInfo
	var count int64
	q := m.db.WithContext(ctx).Table("tables").Where("table_schema = ?", m.dateBase)
	if req.TableName != "" {
		q = q.Where("TABLE_NAME=?", req.TableName)
	}
	if req.TableComment != "" {
		q = q.Where("TABLE_COMMENT=?", req.TableComment)
	}
	if req.Time != nil && len(req.Time) == 2 {
		q = q.Where("CREATE_TIME BETWEEN ? AND ?", req.Time[0], req.Time[1])
	}
	err := q.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = q.Offset(req.GetOffset()).Limit(req.GetLimit()).Find(&tables).Error
	if err != nil {
		return nil, 0, err
	}
	list := make([]dto.TableInfoRes, len(tables))
	for i, table := range tables {
		list[i] = dto.TableInfoRes{
			TableName:    table.TableName,
			TableComment: table.TableComment,
			CreateTime:   table.CreateTime,
			UpdateTime:   table.UpdateTime,
		}
	}
	return list, count, nil
}

func (m *Service) TableDetails(ctx context.Context, req dto.TableInfoSearchReq) (*pkg.BaseInfo, error) {
	var table model.TableInfo
	err := m.db.WithContext(ctx).Table("tables").
		Where("table_schema = ?", m.dateBase).
		Where("TABLE_NAME=?", req.TableName).First(&table).Error
	if err != nil {
		return nil, err
	}
	t := &pkg.BaseInfo{
		SortField:    "id",
		TableName:    table.TableName,
		TableComment: table.TableComment,
		Author:       constants.Author,
		SortWay:      constants.ASC,
		ModelName:    gstr.SnakeToPascal(gstr.TrimPrefix(table.TableName, "_")),
		Config:       make([]string, 0),
		Covers:       make([]string, 0),
	}
	return t, nil
}

// TableInfo 获取所有字段
func (m *Service) TableInfo(ctx context.Context, req dto.TableInfoSearchReq) ([]model.TableColumn, error) {
	var columns []model.TableColumn
	err := m.db.WithContext(ctx).Table("INFORMATION_SCHEMA.COLUMNS").Select(
		"COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, COLUMN_KEY, COLUMN_DEFAULT, EXTRA, COLUMN_COMMENT",
	).Where("TABLE_SCHEMA = ? AND TABLE_NAME = ?", m.dateBase, req.TableName).
		Order("ORDINAL_POSITION"). // 添加排序依据
		Scan(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}

func (m *Service) GenTableFields(ctx context.Context, modeList []model.TableColumn) []*pkg.TableFields {
	var list []*pkg.TableFields
	for _, column := range modeList {
		//过滤不需要的字段
		if _, ok := pkg.FieldFilter[column.ColumnName]; !ok {
			var edit = true
			if column.ColumnName == "id" {
				edit = false
			}
			goName := gstr.SnakeToPascal(column.ColumnName)
			if goName == "Id" {
				goName = "ID"
			}
			list = append(list, &pkg.TableFields{
				Field:     column.ColumnName,
				FieldDes:  column.ColumnComment,
				MysqlType: column.ColumnType,
				GoType:    pkg.ConvertDBTypeToGoType(column.ColumnType),
				TsType:    pkg.ConvertDBTypeToTS(column.ColumnType),
				GoName:    goName,
				JsonName:  gstr.SnakeToCamel(column.ColumnName),
				Edit:      edit,
				List:      true,
				Details:   true,
				Query:     true,
				FillUp:    false,
				QueryType: enum.Equal,
				ShowType:  enum.ShowInput,
				Required:  column.IsNullable == "NO",
				Date:      "", //暂时还未知
			})
		}
	}
	return list
}

func (m *Service) getGinInfoByMysql(ctx context.Context, req dto.TableInfoSearchReq) (*model.GenDate, error) {
	first, err := dal.SysGen.WithContext(ctx).Where(dal.SysGen.TableNamed.Eq(req.TableName)).First()
	if err != nil {
		return nil, err
	}
	baseInfo, err := first.DeserializeBaseInfo()
	if err != nil {
		return nil, err
	}
	genInfo, err := first.DeserializeGenInfo()
	if err != nil {
		return nil, err
	}
	fieldsInfo, err := first.DeserializeFieldsInfo()
	if err != nil {
		return nil, err
	}
	return &model.GenDate{
		BaseInfo:   baseInfo,
		FieldsInfo: fieldsInfo,
		GenInfo:    genInfo,
	}, nil
}

func (m *Service) GinInfo(ctx context.Context, req dto.TableInfoSearchReq) (*model.GenDate, error) {
	byMysql, err := m.getGinInfoByMysql(ctx, req)
	if err == nil {
		return byMysql, nil
	}
	baseInfo, err := m.TableDetails(ctx, req)
	if err != nil {
		return nil, err
	}
	fieldsInfo, err := m.TableInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	tableFields := m.GenTableFields(ctx, fieldsInfo)
	genInfo := &pkg.GenInfo{
		Template:     "0",
		BusinessName: gstr.ToCamelCase(baseInfo.ModelName),
		PackPath:     "system",
		Directory:    gstr.ToCamelCase(baseInfo.ModelName),
		FunctionName: baseInfo.TableComment,
	}
	return &model.GenDate{
		BaseInfo:   baseInfo,
		FieldsInfo: tableFields,
		GenInfo:    genInfo,
	}, nil
}

func (m *Service) ImportDate(ctx context.Context, req model.GenDate) error {
	exist := false
	first, reqErr := dal.SysGen.WithContext(ctx).Where(dal.SysGen.TableNamed.Eq(req.BaseInfo.TableName)).First()
	if reqErr == nil && first.ID > 0 {
		exist = true
	}
	if req.BaseInfo != nil && req.GenInfo != nil && req.FieldsInfo != nil {
		return dal.Q.Transaction(func(tx *dal.Query) error {
			gen := m.setGenerator(req.BaseInfo, req.GenInfo, req.FieldsInfo)
			err := gen.GenCodeAllStr()
			if err != nil {
				return err
			}
			//生成代码文本
			var mode *model.SysGen
			if exist {
				mode = first
			} else {
				mode = &model.SysGen{}
			}
			mode.TableNamed = req.BaseInfo.TableName
			mode.TableComment = req.BaseInfo.TableComment
			//代码
			mode.GoApiCode = gen.Code.GoApiCode
			mode.GoDtoCode = gen.Code.GoDtoCode
			mode.GoLogicCode = gen.Code.GoLogicCode
			mode.GoModelCode = gen.Code.GoModelCode
			mode.GoRouterCode = gen.Code.GoRouterCode
			mode.GoServiceCode = gen.Code.GoServiceCode
			mode.VueApiCode = gen.Code.VueApiCode
			mode.VueFormCode = gen.Code.VueFormCode
			mode.VueHookCode = gen.Code.VueHookCode
			mode.VueIndexCode = gen.Code.VueIndexCode
			mode.VueRuleCode = gen.Code.VueRuleCode
			mode.VueTypesCode = gen.Code.VueTypesCode
			if err = mode.SerializeGenInfo(req.GenInfo); err != nil {
				return err
			}
			if err = mode.SerializeBaseInfo(req.BaseInfo); err != nil {
				return err
			}
			if err = mode.SerializeFieldsInfo(req.FieldsInfo); err != nil {
				return err
			}
			//储存代码信息
			return tx.SysGen.WithContext(ctx).Save(mode)
		})
	}
	return fmt.Errorf("导入失败")
}

func (m *Service) List(ctx context.Context, req dto.TableInfoSearchReq) ([]dto.GenListRes, int64, error) {
	gen := dal.SysGen
	q := gen.WithContext(ctx).Select(gen.ID, gen.TableNamed, gen.TableComment, gen.CreatedAt, gen.UpdatedAt)
	if req.TableName != "" {
		q = q.Where(gen.TableNamed.Eq(req.TableName))
	}
	if req.TableComment != "" {
		q = q.Where(gen.TableComment.Eq(req.TableComment))
	}
	if req.Time != nil && len(req.Time) == 2 {
		q = q.Where(gen.CreatedAt.Between(req.Time[0], req.Time[1]))
	}
	s := make([]dto.GenListRes, 0)
	total, err := q.ScanByPage(&s, req.GetOffset(), req.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return s, total, nil
}

func (m *Service) GetCodes(ctx context.Context, id uint) (*dto.GenCodeRes, error) {
	gen := dal.SysGen
	first, err := gen.WithContext(ctx).Select(gen.GoModelCode, gen.GoServiceCode, gen.GoRouterCode,
		gen.GoLogicCode, gen.GoDtoCode, gen.GoApiCode,
		gen.VueIndexCode, gen.VueTypesCode, gen.VueHookCode, gen.VueRuleCode, gen.VueApiCode, gen.VueFormCode).
		Where(gen.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return &dto.GenCodeRes{
		VueApiCode:    first.VueApiCode,
		VueRuleCode:   first.VueRuleCode,
		VueTypesCode:  first.VueTypesCode,
		VueHookCode:   first.VueHookCode,
		VueIndexCode:  first.VueIndexCode,
		VueFormCode:   first.VueFormCode,
		GoApiCode:     first.GoApiCode,
		GoDtoCode:     first.GoDtoCode,
		GoServiceCode: first.GoServiceCode,
		GoModelCode:   first.GoModelCode,
		GoRouterCode:  first.GoRouterCode,
		GoLogicCode:   first.GoLogicCode,
	}, nil
}

func (m *Service) Delete(ctx context.Context, ids []uint) error {
	gen := dal.SysGen
	info, err := gen.WithContext(ctx).Where(gen.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf(constants.DeleteError)
	}
	return nil
}

func (m *Service) GenCode(ctx context.Context, id uint) error {
	//先创建模型和dal 如果没有的话报错
	gen := dal.SysGen
	mModel, err := gen.WithContext(ctx).Where(gen.ID.Eq(id)).First()
	if err != nil {
		return err
	}
	baseInfo, err := mModel.DeserializeBaseInfo()
	if err != nil {
		return err
	}
	genInfo, err := mModel.DeserializeGenInfo()
	if err != nil {
		return err
	}
	fieldsInfo, err := mModel.DeserializeFieldsInfo()
	if err != nil {
		return err
	}
	tp := m.setGenerator(baseInfo, genInfo, fieldsInfo)
	//写入代码
	err = tp.WriteCode(pkg.GenCodes{
		GoApiCode:     mModel.GoApiCode,
		GoDtoCode:     mModel.GoDtoCode,
		GoLogicCode:   mModel.GoLogicCode,
		GoModelCode:   mModel.GoModelCode,
		GoRouterCode:  mModel.GoRouterCode,
		GoServiceCode: mModel.GoServiceCode,
		VueApiCode:    mModel.VueApiCode,
		VueFormCode:   mModel.VueFormCode,
		VueHookCode:   mModel.VueHookCode,
		VueIndexCode:  mModel.VueIndexCode,
		VueRuleCode:   mModel.VueRuleCode,
		VueTypesCode:  mModel.VueTypesCode,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) setGenerator(baseInfo *pkg.BaseInfo, genInfo *pkg.GenInfo, fieldsInfo []*pkg.TableFields) *gtemplate.Generator {
	tp := gtemplate.NewGTemplate(baseInfo, genInfo, fieldsInfo)
	gen := gtemplate.NewGenerator()
	//Vue 生成器
	gen.PushGen(gtemplate.NewGVueIndexGen(tp))
	gen.PushGen(gtemplate.NewGVueFormGen(tp))
	gen.PushGen(gtemplate.NewGVueRuleGen(tp))
	gen.PushGen(gtemplate.NewGVueHookGen(tp))
	gen.PushGen(gtemplate.NewGVueApiGen(tp))
	gen.PushGen(gtemplate.NewGVueTypesGen(tp))
	//Go 生成器
	gen.PushGen(gtemplate.NewGGoApiGen(tp))
	gen.PushGen(gtemplate.NewGGoServiceGen(tp))
	gen.PushGen(gtemplate.NewGGoDTOGen(tp))
	gen.PushGen(gtemplate.NewGGoRouterGen(tp))
	gen.PushGen(gtemplate.NewGGoLogicGen(tp))
	gen.PushGen(gtemplate.NewGGoModelGen(tp))
	return gen
}
