package captcha

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"strings"
	"sync"
)

var (
	instance *Service
	once     sync.Once
)

type Service struct {
	driver *base64Captcha.DriverString
	store  base64Captcha.Store
}

func New() *Service {
	once.Do(func() {
		instance = &Service{
			driver: &base64Captcha.DriverString{
				Height:          80,
				Width:           240,
				NoiseCount:      50,
				ShowLineOptions: 20,
				Length:          4,
				Source:          "abcdefghjkmnpqrstuvwxyz23456789",
				Fonts:           []string{"chromohv.ttf"},
			},
			store: base64Captcha.DefaultMemStore,
		}
	})
	return instance
}

// GetVerifyImgString 获取字母数字混合验证码
func (s *Service) GetVerifyImgString(ctx context.Context) (id, b64s, answer string, err error) {
	driver := s.driver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, s.store)
	id, b64s, answer, err = c.Generate()
	return
}

// VerifyString 验证输入的验证码是否正确
func (s *Service) VerifyString(id, answer string) bool {
	c := base64Captcha.NewCaptcha(s.driver, s.store)
	answer = strings.ToLower(answer)
	return c.Verify(id, answer, true)
}
