package controllers

import (
	"github.com/d3vilh/openvpn-ui/integrator"
)

type StatusController struct {
	BaseController
	ApiIntegrator *integrator.APIIntegrator
}

func (c *StatusController) Get() {
	response, err := c.ApiIntegrator.GetStatus()
	if err != nil {
		c.CustomAbort(500, err.Error())
		return
	}
	c.Data["json"] = map[string]any{
		"server":  response.Data.Server,
		"active":  response.Data.Active,
		"logs":    response.Data.Logs,
		"servers": response.Data.Servers,
	}
	c.ServeJSON()
}
