package routers

import (
	"github.com/astaxie/beego"
	"solo-ci/controllers"
	"net/http"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/solohook/:project_id", &controllers.ProjectController{}, "post:WebHook"),
		beego.NSRouter("/project",&controllers.ProjectController{},"post:Add"),
		beego.NSRouter("/project/:project_id",&controllers.ProjectController{},"delete:Delete"),
		beego.NSRouter("/project/:project_id",&controllers.ProjectController{},"put:Update"),
		beego.NSRouter("/project/:project_id",&controllers.ProjectController{},"get:Get"),
	)
	beego.AddNamespace(ns)
	beego.ErrorHandler("404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}
