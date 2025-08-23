package middleware

import (
	"net/http"

	"github.com/noyandey88/go-todo-app/pkg/response"
)

const RoleKey contextKey = "role"

func AllowedRole(roles ...string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roleValue := r.Context().Value(RoleKey)
			role, ok := roleValue.(string)

			if !ok || role == "" {
				response.JsonResponse(
					w,
					http.StatusForbidden,
					false,
					"Forbidden: Role not found",
					nil,
				)
				return
			}

			allowed := false

			for _, allowedRole := range roles {
				if role == allowedRole {
					allowed = true
					break
				}
			}

			if !allowed {
				response.JsonResponse(
					w,
					http.StatusForbidden,
					false,
					"Forbidden: Insufficient permissions",
					nil,
				)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
