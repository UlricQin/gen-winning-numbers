package controllers

import (
	"github.com/ulricqin/gen-winning-numbers/g"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	this.Data["PageTitle"] = "登录"
	this.Layout = "layout/admin.html"
	this.TplNames = "login/login.html"
}

func (this *LoginController) DoLogin() {
	name := this.GetString("name")
	if name == "" {
		this.Ctx.WriteString("name is blank")
		return
	}
	pwd := this.GetString("pwd")
	if pwd == "" {
		this.Ctx.WriteString("pwd is blank")
		return
	}

	if g.RootName != name {
		this.Ctx.WriteString("name is incorrect")
		return
	}

	if g.RootPass != pwd {
		this.Ctx.WriteString("pwd is incorrect")
		return
	}

	this.Ctx.SetCookie(g.CookieRootName, name, 2592000, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", g.CookieRootPass+"="+pwd+"; Max-Age=2592000; Path=/; httponly")

	this.Ctx.WriteString("")
}

func (this *LoginController) Logout() {
	this.Ctx.SetCookie(g.CookieRootName, g.RootName, 0, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", g.CookieRootPass+"="+g.RootPass+"; Max-Age=0; Path=/; httponly")
	this.Ctx.WriteString("logout")
}
