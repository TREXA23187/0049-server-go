package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1
	PermissionUser        Role = 2
	PermissionVisitor     Role = 3
	PermissionDisableUser Role = 4
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "admin"
	case PermissionUser:
		str = "user"
	case PermissionVisitor:
		str = "visitor"
	case PermissionDisableUser:
		str = "disabled_user"
	default:
		str = "other"
	}
	return str
}
