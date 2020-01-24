// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api/controllers"
	"api/middleware"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "get:Get")

	ns := beego.NewNamespace("/v1",

		beego.NSBefore(middleware.Jwt),

		beego.NSNamespace("/order-promo-codes",
			beego.NSInclude(
				&controllers.OrderPromoCodeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
