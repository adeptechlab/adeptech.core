package common

import (
	"regexp"
)

type TenantId string
type UserId string
type UserEmail string
type RegistrationID string
type DomainId string

func (id *TenantId) Valid() bool {
	pattern := regexp.MustCompile(`--`)
	result := pattern.Split(string(*id), -1)
	return len(result) == 2
}

func EmptyRegistrationID() RegistrationID {
	return RegistrationID("")
}

func EmptyUserId() UserId {
	return UserId("")
}

func EmptyTenantId() TenantId {
	return TenantId("")
}

func ExisingTenantId(id string) TenantId {
	return TenantId(id)
}

func ExisingUserId(id string) UserId {
	return UserId(id)
}
