package enumerations

type Operation string

var (
	OperationAdd                 Operation = "add"
	OperationUpdate              Operation = "update"
	OperationSeal                Operation = "seal"
	OperationUnseal              Operation = "unseal"
	OperationDelete              Operation = "delete"
	OperationSetAsDefault        Operation = "set_as_default"
	OperationArchive             Operation = "archive"
	OperationDisable             Operation = "disable"
	OperationEnable              Operation = "enable"
	OperationInvite              Operation = "invite"
	OperationResetPassword       Operation = "reset_password"
	OperationUpdateProfile       Operation = "update_profile"
	OperationUpdateOrganization  Operation = "update_organization"
	OperationDeleteOrganization  Operation = "delete_organization"
	OperationLoginFromNewBrowser Operation = "login_from_new_browser"
	OperationLoginFromNewDevice  Operation = "login_from_new_device"
	OperationLoginFromNewIP      Operation = "login_from_new_ip"
	OperationChangeUserRole      Operation = "change_user_role"
)

func (t Operation) String() string {
	return string(t)
}
