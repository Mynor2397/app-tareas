package middlewares

import (
	"context"
	"net/http"
	"strings"

	help "github.com/Mynor2397/api-simple/src/helpers"
	mod "github.com/Mynor2397/api-simple/src/models"
)

type rol string

var (
	// Claims  contiene las reclamaciones
	Claims rol = "claims"
)

// WithAuth es el middleware para evaluar el token
func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorizationHeader, "Bearer") {
			next.ServeHTTP(w, r)
			return
		}

		var (
			uid  int64
			urol string
		)

		uid, urol, err := help.AuthUserID(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, Claims, mod.Rol{IDRol: uid, NameRol: urol})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
