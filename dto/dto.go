package dto

type (
	BaseResponse[T any] struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    T      `json:"data"`
	}

	VPNStatusResponse struct {
		Active  bool     `json:"active"`
		Server  string   `json:"server"`
		Logs    string   `json:"logs"`
		Servers []string `json:"servers"`
	}

	VPNActivateRequest struct {
		Host   string `json:"host"`
		Port   string `json:"port"`
		Cookie string `json:"cookie"`
	}
)
