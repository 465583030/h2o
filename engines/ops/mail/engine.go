package mail

import "github.com/astaxie/beego"

// Controller controller
type Controller struct {
	beego.Controller
}

// GetHome home
// @router / [get]
func (p *Controller) GetHome() {
	p.Data["title"] = "mail"
	p.TplName = "home.tpl"
}
