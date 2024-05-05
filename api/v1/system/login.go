package system

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"net/http"
	"time"
)

type LoginApi struct {
	loginService      service.ILoginService
	loginLogService   service.ILoginLogService
	adminService      service.IAdminService
	tokenService      service.ITokenService
	userOnlineService service.IUserOnlineService
	captchaService    service.ICaptcha
}

func NewLoginApi(captchaService service.ICaptcha, loginService service.ILoginService, loginLogService service.ILoginLogService, adminService service.IAdminService, tokenService service.ITokenService, userOnlineService service.IUserOnlineService) *LoginApi {
	return &LoginApi{
		loginService:      loginService,
		loginLogService:   loginLogService,
		adminService:      adminService,
		tokenService:      tokenService,
		userOnlineService: userOnlineService,
		captchaService:    captchaService,
	}
}

// Captcha 做成通用api较好
func (m LoginApi) Captcha(c *gin.Context, _ gdto.EmptyReq) (any, error) {
	id, b64s, _, err := m.captchaService.GetVerifyImgString(c)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Success: true,
		Data: gdto.CaptchaRes{
			ID:   id,
			B64s: b64s,
		},
	}, nil
}

func (m LoginApi) CheckCaptcha(_ *gin.Context, req gdto.CheckCaptchaReq) (any, error) {
	return gres.Response{
		Data: gdto.CheckCaptchaRes{
			Success: m.captchaService.VerifyString(req.ID, req.Captcha),
		},
	}, nil
}

func (m LoginApi) Login(c *gin.Context, req *dto.AdminLoginReq) (any, error) {
	if !m.captchaService.VerifyString(req.CaptchaId, req.Captcha) {
		return nil, fmt.Errorf("验证码错误")
	}
	var (
		token  string
		rToken string
		eTime  time.Time
		param  = dto.LoginParam{
			UserName: req.UserName,
			Ip:       c.ClientIP(),
			Behavior: "账号登陆",
			Ua:       useragent.Parse(c.GetHeader("User-Agent")),
		}
	)
	//账号密码登陆
	mModel, err := m.loginService.Login(c, req)
	if err == nil {
		//获取token
		token, eTime, err = m.tokenService.GenerateLoginToken(c, mModel.ID, mModel.UserName)
		if err != nil {
			err = fmt.Errorf(constants.ServiceError)
		}
	}

	if err == nil {
		//获取刷新token
		rToken, err = m.tokenService.GenerateRefreshToken(c, mModel.ID, mModel.UserName)
		if err != nil {
			err = fmt.Errorf(constants.ServiceError)
		}
	}

	if err != nil {
		//记录login失败日志
		m.loginLogService.Create(c, &param)
		return nil, err
	}

	param.UserId = mModel.ID
	param.Status = true
	//记录login成功日志
	m.loginLogService.Create(c, &param)
	//更新登陆数据
	m.adminService.SetLogin(c, req.UserName, param.Ip)
	//更新在线
	m.userOnlineService.Create(c, &param)

	roles := make([]string, len(mModel.RoleSign))
	for i, role := range mModel.RoleSign {
		roles[i] = role.Code
	}
	return gres.Response{
		Success: true,
		Message: "登陆成功",
		Data: dto.AdminLoginRes{
			UserName:     mModel.UserName,
			Avatar:       mModel.Avatar,
			NickName:     mModel.NickName,
			AccessToken:  token,
			RefreshToken: rToken,
			Expires:      eTime.UnixNano() / 1e6,
			Roles:        roles,
		},
	}, nil
}

// RefreshToken 应该是单独给一个api
func (m LoginApi) RefreshToken(c *gin.Context, _ gdto.EmptyReq) (any, error) {
	token, rToken, eTime, err := m.tokenService.RefreshToken(c)
	if err != nil {
		return gres.Response{
			Status:  http.StatusUnauthorized,
			Code:    401,
			Message: "登陆状态已过期,请重新登陆",
		}, err
	}

	return gres.Response{
		Success: true,
		Message: "续签Token成功",
		Data: dto.RefreshTokenRes{
			AccessToken:  token,
			RefreshToken: rToken,
			Expires:      eTime.UnixNano() / 1e6,
		},
	}, nil
}

//// 实现DTO的验证方法，如果需要
//func (dto ExampleDTO) Validate() error {
//	if dto.ID < 1 {
//		return errors.New("ID must be greater than 0")
//	}
//	return nil
//}

//// 定义一个DTO
//type ExampleDTO struct {
//	ID   int    `form:"id" binding:"required"`
//	NameZh string `form:"name" binding:"required"`
//}
//
//type User struct {
//	ID        int
//	NameZh      string
//	Email     string
//	IsActive  bool `json:"-"`
//	ActiveInt int  `json:"isActive"`
//}
//
//// MarshalJSON 用于自定义User的JSON序列化
//func (u User) MarshalJSON() ([]byte, error) {
//	// 当IsActive为true时，设置ActiveInt为1；否则为0
//	if u.IsActive {
//		u.ActiveInt = 1
//	} else {
//		u.ActiveInt = 0
//	}
//	// 创建一个临时的类型，用来避免无限递归调用MarshalJSON
//	type Alias User
//	return json.Marshal(&struct {
//		*Alias
//	}{
//		Alias: (*Alias)(&u),
//	})
//}
//
//// UnmarshalJSON 用于自定义User的JSON反序列化
//func (u *User) UnmarshalJSON(data []byte) error {
//	// 临时接收数据的结构体
//	temp := struct {
//		ActiveInt int `json:"isActive"`
//	}{}
//
//	if err := json.Unmarshal(data, &temp); err != nil {
//		return err
//	}
//
//	// 将整数转换回布尔值
//	u.IsActive = temp.ActiveInt != 0
//	return nil
//}
//
//// Test 实际的业务处理函数
//func (m LoginApi) Test(c *gin.Context, req ExampleDTO) (any, error) {
//	//// 业务逻辑处理
//	var user User
//	if err := copier.Copy(&user, &req); err != nil {
//		fmt.Println(err)
//
//	}
//	user.IsActive = true
//	return gres.Response{
//		Data: gin.H{"message": "Success", "data": user},
//	}, nil
//}
