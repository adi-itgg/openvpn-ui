package integrator

import (
	"github.com/d3vilh/openvpn-ui/dto"
	"resty.dev/v3"
)

type APIIntegrator struct {
	restyClient *resty.Client
}

func NewAPIIntegrator() *APIIntegrator {
	return &APIIntegrator{
		restyClient: resty.New().
			SetBaseURL("http://openvpn:1195").
			SetHeader("Content-Type", "application/json"),
	}
}

func (a *APIIntegrator) GetStatus() (*dto.BaseResponse[dto.VPNStatusResponse], error) {
	response, err := a.restyClient.R().
		SetResult(&dto.BaseResponse[dto.VPNStatusResponse]{}).
		SetError(&dto.BaseResponse[dto.VPNStatusResponse]{}).
		Get("/api/status")
	if err != nil {
		return nil, err
	}

	return response.Result().(*dto.BaseResponse[dto.VPNStatusResponse]), nil
}

func (a *APIIntegrator) ActivateVPN(request *dto.VPNActivateRequest) (*dto.BaseResponse[any], error) {
	response, err := a.restyClient.R().
		SetBody(request).
		SetResult(&dto.BaseResponse[any]{}).
		SetError(&dto.BaseResponse[any]{}).
		Post("/api/activate")
	if err != nil {
		return nil, err
	}

	return response.Result().(*dto.BaseResponse[any]), nil
}
