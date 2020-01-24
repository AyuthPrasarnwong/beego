package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["json"] = "Horeca Report API !!!"
	this.ServeJSON()
}


