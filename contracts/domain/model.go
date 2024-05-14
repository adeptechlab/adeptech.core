package domain

import "github.com/adeptechlab/adeptech.core/contracts/common"

type Domain struct {
	TenantID common.TenantId
	UserID   common.UserId
}
