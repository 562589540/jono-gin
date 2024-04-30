package sys_gen

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub/gbootstrap"
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate"
	"github.com/562589540/jono-gin/internal/app/system/dal"
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
		if !gbootstrap.Cfg.Mode.Develop {
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

func (m *Service) TableDetails(ctx context.Context, req dto.TableInfoSearchReq) (*gtemplate.BaseInfo, error) {
	var table model.TableInfo
	err := m.db.WithContext(ctx).Table("tables").
		Where("table_schema = ?", m.dateBase).
		Where("TABLE_NAME=?", req.TableName).First(&table).Error
	if err != nil {
		return nil, err
	}
	t := &gtemplate.BaseInfo{
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

func (m *Service) GenTableFields(ctx context.Context, modeList []model.TableColumn) []*gtemplate.TableFields {
	var list []*gtemplate.TableFields
	for _, column := range modeList {
		//过滤不需要的字段
		if _, ok := gtemplate.FieldFilter[column.ColumnName]; !ok {
			list = append(list, &gtemplate.TableFields{
				Field:     column.ColumnName,
				FieldDes:  column.ColumnComment,
				MysqlType: column.ColumnType,
				GoType:    gtemplate.ConvertDBTypeToGoType(column.ColumnType),
				TsType:    gtemplate.ConvertDBTypeToTS(column.ColumnType),
				GoName:    gstr.SnakeToPascal(column.ColumnName),
				JsonName:  gstr.SnakeToCamel(column.ColumnName),
				Edit:      true,
				List:      true,
				Details:   true,
				Query:     true,
				FillUp:    true,
				QueryType: gtemplate.Equal,
				ShowType:  gtemplate.ShowInput,
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
	genInfo := &gtemplate.GenInfo{
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
			//生成代码文本
			tp := gtemplate.NewGTemplate(req.BaseInfo, req.GenInfo, req.FieldsInfo)
			indexStr, err := tp.GenerateVueIndexStr()
			if err != nil {
				return err
			}
			apiStr, err := tp.GenerateVueApiTsStr()
			if err != nil {
				return err
			}
			formStr, err := tp.GenerateVueFormStr()
			if err != nil {
				return err
			}
			hookStr, err := tp.GenerateVueHookTsxStr()
			if err != nil {
				return err
			}
			ruleStr, err := tp.GenerateVueRuleTsStr()
			if err != nil {
				return err
			}
			typesStr, err := tp.GenerateVueTypesTsStr()
			if err != nil {
				return err
			}
			goModelStr, err := tp.GenerateGoStr("model")
			if err != nil {
				return err
			}
			goDtoStr, err := tp.GenerateGoStr("dto")
			if err != nil {
				return err
			}
			goServiceStr, err := tp.GenerateGoStr("service")
			if err != nil {
				return err
			}
			goApiStr, err := tp.GenerateGoStr("api")
			if err != nil {
				return err
			}
			goRouterStr, err := tp.GenerateGoStr("router")
			if err != nil {
				return err
			}
			goLogicStr, err := tp.GenerateGoStr("logic")
			if err != nil {
				return err
			}
			var mode *model.SysGen
			if exist {
				mode = first
			} else {
				mode = &model.SysGen{}
			}
			mode.TableNamed = req.BaseInfo.TableName
			mode.TableComment = req.BaseInfo.TableComment
			mode.GoApiCode = goApiStr
			mode.GoDtoCode = goDtoStr
			mode.GoLogicCode = goLogicStr
			mode.GoModelCode = goModelStr
			mode.GoRouterCode = goRouterStr
			mode.GoServiceCode = goServiceStr
			mode.VueApiCode = apiStr
			mode.VueFormCode = formStr
			mode.VueHookCode = hookStr
			mode.VueIndexCode = indexStr
			mode.VueRuleCode = ruleStr
			mode.VueTypesCode = typesStr
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
	_, err := gen.WithContext(ctx).Where(gen.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) GenCode(ctx context.Context, id uint) error {
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
	//执行代码写入
	tp := gtemplate.NewGTemplate(baseInfo, genInfo, fieldsInfo)
	err = tp.GenCode(gtemplate.GenCodes{
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
