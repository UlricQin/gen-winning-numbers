package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ulricqin/gen-winning-numbers/g"
	"github.com/ulricqin/gen-winning-numbers/models"
	"github.com/ulricqin/goutils/filetool"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) CheckLogin() {
	if !this.IsAdmin {
		this.Redirect("/login", 302)
	}
}

func (this *AdminController) ConfigSync() {
	b, err := json.Marshal(models.Global)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("json.Marshal(models.Global). error: %s", err))
		return
	}

	_, err = filetool.WriteBytesToFile(g.ConfigJsonFile, b)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("write json to %s fail. error: %s", g.ConfigJsonFile, err))
		return
	}

	this.Ctx.WriteString("")
}

func (this *AdminController) Index() {
	this.Data["Global"] = models.Global
	this.Layout = "layout/admin.html"
	this.TplNames = "admin/index.html"
}

func (this *AdminController) ResetConfig() {
	baseDir := this.GetString("dir")
	if baseDir == "" {
		this.Ctx.WriteString("parameter dir is blank")
		return
	}

	if !filetool.IsExist(baseDir) {
		this.Ctx.WriteString(fmt.Sprintf("dir: %s is not exists", baseDir))
		return
	}

	models.SetBaseDir(baseDir)

	err := models.LoadDuoYuanDirs()
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("load duo_yuan dir of %s fail. error: %s", baseDir, err))
		return
	}

	this.Redirect(this.UrlFor("AdminController.Index"), 302)
}

func (this *AdminController) UpdateDuoYuanRongCuo() {
	val, err := this.GetInt("val")
	if err != nil {
		this.Ctx.WriteString("number format error")
		return
	}

	t := this.GetString("t")
	if t == "min" {
		models.Global.DuoYuanRongCuoMin = int(val)
	} else {
		models.Global.DuoYuanRongCuoMax = int(val)
	}

	this.Ctx.WriteString("")
}

func (this *AdminController) UpdateJiYiRongCuo() {
	val, err := this.GetInt("val")
	if err != nil {
		this.Ctx.WriteString("number format error")
		return
	}

	duoYuanDir := this.GetString("dy")

	t := this.GetString("t")
	if t == "min" {
		models.Global.DuoYuanList[duoYuanDir].RongCuoMin = int(val)
	} else {
		models.Global.DuoYuanList[duoYuanDir].RongCuoMax = int(val)
	}

	this.Ctx.WriteString("")
}

func (this *AdminController) UpdateDaDiRongCuo() {
	val, err := this.GetInt("val")
	if err != nil {
		this.Ctx.WriteString("number format error")
		return
	}

	duoYuanDir := this.GetString("dy")
	jiYiDir := this.GetString("jy")

	t := this.GetString("t")
	if t == "min" {
		models.Global.DuoYuanList[duoYuanDir].JiYiList[jiYiDir].RongCuoMin = int(val)
	} else {
		models.Global.DuoYuanList[duoYuanDir].JiYiList[jiYiDir].RongCuoMax = int(val)
	}

	this.Ctx.WriteString("")
}

func (this *AdminController) UpdateDuoYuanChecked() {
	checked, err := this.GetInt("checked")
	if err != nil {
		this.Ctx.WriteString("checked must be 0 or 1")
		return
	}
	dir := this.GetString("dir")
	models.Global.DuoYuanList[dir].Checked = int(checked)
	this.Ctx.WriteString("")
}

func (this *AdminController) UpdateJiYiChecked() {
	checked, err := this.GetInt("checked")
	if err != nil {
		this.Ctx.WriteString("checked must be 0 or 1")
		return
	}

	dy := this.GetString("dy")
	jy := this.GetString("jy")
	models.Global.DuoYuanList[dy].JiYiList[jy].Checked = int(checked)
	this.Ctx.WriteString("")
}

func (this *AdminController) UpdateDaDiChecked() {
	checked, err := this.GetInt("checked")
	if err != nil {
		this.Ctx.WriteString("checked must be 0 or 1")
		return
	}

	dy := this.GetString("dy")
	jy := this.GetString("jy")
	name := this.GetString("name")
	models.Global.DuoYuanList[dy].JiYiList[jy].DaDiList[name].Checked = int(checked)
	this.Ctx.WriteString("")
}

func (this *AdminController) UpdateSearchNum() {
	num, err := this.GetInt("num")
	if err != nil {
		this.Ctx.WriteString("num must be integer")
		return
	}

	dy := this.GetString("dy")
	jy := this.GetString("jy")
	name := this.GetString("name")
	models.Global.DuoYuanList[dy].JiYiList[jy].DaDiList[name].SearchNum = int(num)
	this.Ctx.WriteString("")
}

func (this *AdminController) LoadJiYiListOfDuoYuan() {
	dy := this.Ctx.Input.Param(":dy")
	this.Data["DuoYuan"] = models.Global.DuoYuanList[dy]
	this.TplNames = "inc/jyl.html"
}

func (this *AdminController) LoadDaDiListOfJiYi() {
	dy := this.Ctx.Input.Param(":dy")
	jy := this.Ctx.Input.Param(":jy")
	this.Data["DuoYuan"] = models.Global.DuoYuanList[dy]
	this.Data["JiYi"] = models.Global.DuoYuanList[dy].JiYiList[jy]
	this.TplNames = "inc/ddl.html"
}

// no need
func (this *AdminController) DaDiNew() {
	dy := this.GetString("dy")
	jy := this.GetString("jy")
	name := this.GetString("name")
	s_num := this.GetIntWithDefault("s_num", 0)

	ddl := models.Global.DuoYuanList[dy].JiYiList[jy].DaDiList
	if ddl == nil {
		models.Global.DuoYuanList[dy].JiYiList[jy].DaDiList = make(map[string]*models.DaDi)
	}
	models.Global.DuoYuanList[dy].JiYiList[jy].DaDiList[name] = &models.DaDi{Name: name, SearchNum: s_num}
	this.Ctx.WriteString("")
}

func (this *AdminController) DuoYuanNew() {
	min, err := this.GetInt("min")
	if err != nil {
		this.Ctx.WriteString("parameter min must be integer")
		return
	}

	var max int64
	max, err = this.GetInt("max")
	if err != nil {
		this.Ctx.WriteString("parameter max must be integer")
		return
	}

	// generate 20 ji yi and read, then reload page
	// 读取caipiao_data目录，拿到一个duo_yuan dir，然后把数据生成到这个目录下
	err = models.GenerateOneDuoYuanDir(min, max)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("Generate duo yuan dir fail. error: %s", err))
		return
	}

	this.Ctx.WriteString("")
}

func (this *AdminController) UpdateWinningNumber() {
	order := this.GetString("order")
	val := this.GetString("val")
	if order == "1" {
		models.Global.Last1 = val
	} else {
		models.Global.Last2 = val
	}
	this.Ctx.WriteString("")
}
