package platform

import "github.com/adeptechlab/adeptech.core/contracts/common"

type StatusResponse struct {
	common.BaseErrorResponse
	Data Status `json:"data"`
}
