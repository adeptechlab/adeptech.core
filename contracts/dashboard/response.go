package dashboard

import "github.com/adeptechlab/adeptech.core/contracts/common"

type GetOperationDashboardResponse struct {
	common.BaseErrorResponse
	Dashboard Operations `json:"data"`
}
