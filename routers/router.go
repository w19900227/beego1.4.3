// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego1.4.3/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.PostController{})
	beego.Router("/post", &controllers.PostController{})
	beego.Router("/post/:id([0-9]+", &controllers.PostController{}, "get:Show;put:Update")
	beego.Router("/post/edit/:id([0-9]+", &controllers.PostController{}, "get:Edit")
	beego.AutoRouter(&controllers.PostController{})

	// beego.RESTRouter("/post", new(controllers.PostController))
	// ns := beego.NewNamespace("/v1")
	// beego.AddNamespace(ns)
}
