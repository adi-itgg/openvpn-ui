package controllers

import (
	"os"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/d3vilh/openvpn-ui/dto"
	"github.com/d3vilh/openvpn-ui/integrator"
	"github.com/d3vilh/openvpn-ui/models"
)

type ActivationController struct {
	BaseController
	ApiIntegrator *integrator.APIIntegrator
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

	response, err := c.ApiIntegrator.GetStatus()
	if err != nil {
		flash := web.NewFlash()
		logs.Error(err)
		flash.Error("Failed to get status: " + err.Error())
		flash.Store(&c.Controller)
		return
	} else {
		c.Data["Status"] = response.Data
	}
}

func (c *ActivationController) Post() {
	c.TplName = "activation.html"

	c.Get()

	host := c.GetString("vpn_host")
	port := c.GetString("vpn_port", "10443")
	host = strings.Split(host, " (")[0]

	flash := web.NewFlash()

	cookie := c.GetString("cookie")
	if cookie == "" {
		flash.Error("Cookie is empty")
		flash.Store(&c.Controller)
		return
	}

	if len(cookie) <= 2000 && len(cookie) > 3000 {
		flash.Error("Invalid cookie")
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
	}

	requestActivation := &dto.VPNActivateRequest{
		Host:   host,
		Port:   port,
		Cookie: cookie,
	}

	response, err := c.ApiIntegrator.ActivateVPN(requestActivation)
	if err != nil {
		logs.Error(err)
		flash.Error(err.Error())
	} else if response.Success {
		flash.Success("Successfully activated")
	} else {
		flash.Error(response.Message)
	}

	flash.Store(&c.Controller)
}
