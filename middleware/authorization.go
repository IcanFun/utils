package middleware

import (
	"net/http"

	"redis-trade-query-server/utils/log"

	"redis-trade-query-server/model"
)

func CustomClaimsHasPermissionTo(customClaims model.CustomClaims, permission *model.Permission) bool {
	return CheckIfRolesGrantPermission(customClaims.GetUserRoles(), permission.Id)
}

func CheckIfRolesGrantPermission(roles []string, permissionId string) bool {
	for _, roleId := range roles {
		if role, ok := model.BuiltInRoles[roleId]; !ok {
			log.Debug("Bad role in system " + roleId)
			return false
		} else {
			permissions := role.Permissions
			for _, permission := range permissions {
				if permission == permissionId {
					return true
				}
			}
		}
	}

	return false
}

func GetProtocol(r *http.Request) string {
	if r.Header.Get(model.HEADER_FORWARDED_PROTO) == "https" || r.TLS != nil {
		return "https"
	} else {
		return "http"
	}
}
