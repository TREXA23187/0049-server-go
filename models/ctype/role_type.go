package ctype

import "encoding/json"

type Role int

const (
	RoleSuperAdmin Role = 1
	RoleAdmin      Role = 2
	RoleResearcher Role = 3
	RoleUser       Role = 4
	RoleDisabled   Role = 5
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case RoleSuperAdmin:
		str = "super_admin"
	case RoleAdmin:
		str = "admin"
	case RoleResearcher:
		str = "researcher"
	case RoleUser:
		str = "user"
	case RoleDisabled:
		str = "disabled"
	default:
		str = "other"
	}
	return str
}

func StringToRoleType(str string) Role {
	var role Role
	switch str {
	case "super admin":
		role = RoleSuperAdmin
	case "admin":
		role = RoleAdmin
	case "researcher":
		role = RoleResearcher
	case "user":
		role = RoleUser
	case "disabled":
		role = RoleDisabled
	}
	return role
}
