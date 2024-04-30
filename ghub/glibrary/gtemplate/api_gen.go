package gtemplate

type ApiGen struct {
	GTemplate
}

func (m *ApiGen) GenCodeStr() (string, error) {
	return "", nil
}

func (m *ApiGen) GenCode(code string) error {
	return nil
}
