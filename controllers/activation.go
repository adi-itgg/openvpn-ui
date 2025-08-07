package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/d3vilh/openvpn-ui/models"
	"os"
)

type ActivationController struct {
	BaseController
}

func (c *ActivationController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
	c.Data["breadcrumbs"] = &BreadCrumbs{
		Title: "Activation",
	}
}

func (c *ActivationController) Post() {
	c.TplName = "activation.html"

	settings := models.Settings{Profile: "default"}
	settings.Read("Profile")

	if err := settings.Read("OVConfigPath"); err != nil {
		logs.Error(err)
		return
	}

	fName := settings.OVConfigPath + "/cookie.txt"

	cookie := c.GetString("cookie")
	if cookie == "" {
		c.Data["Error"] = "Cookie tidak boleh kosong"
		return
	}
	err := os.WriteFile(fName, []byte(cookie), 0644)
	if err != nil {
		c.Data["Error"] = "Gagal menyimpan."
	} else {
		c.Data["Success"] = "Cookie tersimpan."
	}
}
