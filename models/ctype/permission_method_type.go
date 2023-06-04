package ctype

import "encoding/json"

type PermissionMethod int

const (
	ALL    PermissionMethod = 1
	GET    PermissionMethod = 2
	POST   PermissionMethod = 3
	PUT    PermissionMethod = 4
	DELETE PermissionMethod = 5
)

func (p PermissionMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p PermissionMethod) String() string {
	var str string
	switch p {
	case ALL:
		str = "ALL"
	case GET:
		str = "GET"
	case POST:
		str = "POST"
	case PUT:
		str = "PUT"
	case DELETE:
		str = "DELETE"
	default:
		str = "other"
	}
	return str
}

func StringToPermissionMethod(str string) PermissionMethod {
	var method PermissionMethod
	switch str {
	case "ALL":
		method = ALL
	case "GET":
		method = GET
	case "POST":
		method = POST
	case "PUT":
		method = PUT
	case "DELETE":
		method = DELETE
	}
	return method
}
