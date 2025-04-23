package authz

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authn"
)

func Authorize(r *http.Request) bool {
	target := fmt.Sprintf("%s:%s", r.Method, r.URL.Path)
	roles := authn.GetClientRoles(r.Context())

	for _, role := range roles {
		for _, perm := range RolePermissions[role] {
			if perm == "*" {
				return true
			}

			if target == perm {
				return true
			}

			if strings.HasPrefix(target, perm+"/") {
				return true
			}
		}
	}

	return false
}
