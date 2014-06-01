package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/gen-winning-numbers/g"
	"github.com/ulricqin/goutils/paginator"
	"strconv"
)

type Checker interface {
	CheckLogin()
}

type BaseController struct {
	beego.Controller
	IsAdmin bool
}

func (this *BaseController) Prepare() {
	this.Data["RootName"] = g.RootName
	this.AssignIsAdmin()
	if app, ok := this.AppController.(Checker); ok {
		app.CheckLogin()
	}
}

func (this *BaseController) AssignIsAdmin() {
	name := this.Ctx.GetCookie(g.CookieRootName)
	pwd := this.Ctx.GetCookie(g.CookieRootPass)
	if name == "" || pwd == "" {
		this.IsAdmin = false
		return
	}

	if name != g.RootName || pwd != g.RootPass {
		this.IsAdmin = false
	}

	this.IsAdmin = true
	this.Data["IsAdmin"] = this.IsAdmin
}

func (this *BaseController) SetPaginator(per int, nums int64) *paginator.Paginator {
	p := paginator.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *BaseController) GetIntWithDefault(paramKey string, defaultVal int) int {
	valStr := this.GetString(paramKey)
	var val int
	if valStr == "" {
		val = defaultVal
	} else {
		var err error
		val, err = strconv.Atoi(valStr)
		if err != nil {
			val = defaultVal
		}
	}
	return val
}
