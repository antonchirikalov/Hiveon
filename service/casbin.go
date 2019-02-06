package service

import (
	"github.com/casbin/casbin"
	. "github.com/ory/hydra/sdk/go/hydra/swagger"
	log "github.com/sirupsen/logrus"
	. "hiveon-api/utils"
	"net/http"
)

func Authorizer(e *casbin.Enforcer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			token := r.Context().Value("token").(OAuth2TokenIntrospection)
			if GetConfig().GetBool("security.useCasbin") {

				method := r.Method
				path := r.URL.Path

				user, _ := GetUserByEmail(token.Sub)
				log.Info(user)

				if e.Enforce(token.Sub, path, method) {
					next.ServeHTTP(w, r)
				} else {
					http.Error(w, http.StatusText(403), 403)
					return
				}
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
