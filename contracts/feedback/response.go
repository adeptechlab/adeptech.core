package feedback

import "github.com/adeptechlab/adeptech.core/contracts/common"

type AddFeedbackResponse struct {
	common.BaseErrorResponse
	Data *NewFeedback `json:"data"`
}
