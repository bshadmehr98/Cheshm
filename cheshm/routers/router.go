package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/bshadmehr98/cheshm/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signup", &controllers.SignupController{})
}
