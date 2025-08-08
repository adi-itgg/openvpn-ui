package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
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

func (c *ActivationController) Get() {
	c.TplName = "activation.html"
}

func (c *ActivationController) Post() {
	c.TplName = "activation.html"

	flash := web.NewFlash()

	cookie := c.GetString("cookie")
	if cookie == "" {
		flash.Error("Cookie tidak boleh kosong")
		flash.Store(&c.Controller)
		return
	}

	if len(cookie) <= 2000 && len(cookie) > 3000 {
		flash.Error("Cookie tidak sesuai")
		flash.Store(&c.Controller)
		return
	}

	settings := models.Settings{Profile: "default"}
	settings.Read("Profile")

	if err := settings.Read("OVConfigPath"); err != nil {
		logs.Error(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	fName := settings.OVConfigPath + "/config/cookie.txt"

	err := os.WriteFile(fName, []byte(cookie), 0644)
	if err != nil {
		logs.Error(err)
		flash.Error(err.Error())
	} else {
		flash.Success("Cookie berhasil disimpan")
	}
	flash.Store(&c.Controller)
}
