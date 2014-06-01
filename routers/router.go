package routers

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/gen-winning-numbers/controllers"
)

func init() {
	beego.Router("/health", &controllers.ApiController{}, "get:Health")
	beego.Router("/md5", &controllers.ApiController{}, "get:Md5")

	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/admin", &controllers.AdminController{}, "get:Index")
	beego.Router("/admin/config/reset", &controllers.AdminController{}, "get:ResetConfig")
	beego.Router("/admin/config/sync", &controllers.AdminController{}, "get:ConfigSync")
	beego.Router("/admin/config/jyl/dy/:dy", &controllers.AdminController{}, "get:LoadJiYiListOfDuoYuan")
	beego.Router("/admin/config/ddl/dy/:dy/jy/:jy", &controllers.AdminController{}, "get:LoadDaDiListOfJiYi")
	beego.Router("/admin/config/update/dyck", &controllers.AdminController{}, "get:UpdateDuoYuanChecked")
	beego.Router("/admin/config/update/dyrc", &controllers.AdminController{}, "get:UpdateDuoYuanRongCuo")
	beego.Router("/admin/config/update/jyck", &controllers.AdminController{}, "get:UpdateJiYiChecked")
	beego.Router("/admin/config/update/jyrc", &controllers.AdminController{}, "get:UpdateJiYiRongCuo")
	beego.Router("/admin/config/update/ddck", &controllers.AdminController{}, "get:UpdateDaDiChecked")
	beego.Router("/admin/config/update/ddrc", &controllers.AdminController{}, "get:UpdateDaDiRongCuo")
	beego.Router("/admin/config/update/snum", &controllers.AdminController{}, "get:UpdateSearchNum")
	beego.Router("/admin/config/update/winning-number", &controllers.AdminController{}, "get:UpdateWinningNumber")
	beego.Router("/admin/config/dd/new", &controllers.AdminController{}, "post:DaDiNew")
	beego.Router("/admin/config/dy/new", &controllers.AdminController{}, "post:DuoYuanNew")
	beego.Router("/admin/compute/model2", &controllers.ComputeController{}, "get:Model2")
}
