package gtemplate

import "github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"

type GeneratorEntry struct {
	Gen    pkg.IGenCode
	Setter func(*pkg.GenCodes, string)
}

type Generator struct {
	Generators []pkg.IGenCode
	Code       pkg.GenCodes
}

func NewGenerator() *Generator {
	return &Generator{
		Generators: []pkg.IGenCode{},
	}
}

// PushGen 添加生成器
func (g *Generator) PushGen(gen pkg.IGenCode) {
	g.Generators = append(g.Generators, gen)
}

// GenCodeAllStr 生成所有代码
func (g *Generator) GenCodeAllStr() error {
	for _, generator := range g.Generators {
		str, err := generator.GenCodeStr()
		if err != nil {
			return err
		}
		generator.UpdateGeneratedCode(&g.Code, str)
	}
	return nil
}

// WriteCode 将代码写入文件
func (g *Generator) WriteCode(codes pkg.GenCodes) error {
	for _, generator := range g.Generators {
		err := generator.GenCode(codes)
		if err != nil {
			return err
		}
	}
	return nil
}
