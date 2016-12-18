package routers

import (
	"github.com/astaxie/beego"
	"solo-ci/controllers"
	"net/http"
)

func init() {
	ns := beego.NewNamespace("/v1",
		//beego.NSRouter("/solohook/:project_id", &controllers.ProjectController{},"post:WebHook"),
		beego.NSRouter("/status",&controllers.ProjectController{},"get:Status"),
	)
	beego.AddNamespace(ns)
	beego.ErrorHandler("404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}
