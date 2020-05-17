package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/yun-cn/GolangWeb/utils"
)

/**
 * 管理员控制器
 */
type AdminController struct {
	Ctx iris.Context

	Session *sessions.Session
}

const (
	ADMINTABLENAME = "admin"
	ADMIN		   = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (ac *AdminController) GetSingout() mvc.Result {
	ac.Session.Delete(ADMIN)
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SIGNOUT),
		},
	}
}

func (ac *AdminController) GetCount() mvc.Result {
	count
}