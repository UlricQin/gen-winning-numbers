package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	this.Layout = "layout/default.html"
	this.TplNames = "index.html"
}
