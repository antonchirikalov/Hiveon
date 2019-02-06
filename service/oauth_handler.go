package service

import (
	"context"
	"encoding/json"
	"github.com/go-errors/errors"
	. "github.com/ory/hydra/sdk/go/hydra/swagger"
	log "github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
	. "hiveon-api/utils"
	"net/http"
	"strings"
)

type oauthHandler struct {
	muxHandler    http.Handler
	casbinHandler http.Handler
}

func OAuthHandler(muxHandler http.Handler, casbinHandler http.Handler) http.Handler {
	return oauthHandler{muxHandler: muxHandler, casbinHandler: casbinHandler}
}

func (h oauthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	log.Info("OAuthHandler + " + path)

	//swagger calls and public URI
	publicURL := GetConfig().GetString("security.publicURL")
	if strings.HasPrefix(path, "/swaggerui") || strings.HasPrefix(path, publicURL) {
		//we don'n need casbin handler
		h.muxHandler.ServeHTTP(w, req)
		return
	}

	securedURI := GetConfig().GetString("security.securedURI")
	if strings.HasPrefix(path, securedURI) {
		token, err := checkToken(req)
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}
		ctx := context.WithValue(req.Context(), "token", token)
		h.casbinHandler.ServeHTTP(w, req.WithContext(ctx))
		return
	}

	h.muxHandler.ServeHTTP(w, req)
}

func checkToken(req *http.Request) (OAuth2TokenIntrospection, error) {
	reqToken := req.Header.Get("Authorization")
	var introToken OAuth2TokenIntrospection

	if len(reqToken) == 0 {
		return introToken, errors.New("Authorization token missed")
	}

	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return introToken, errors.New("Token is wrong")
	}

	token := strings.TrimSpace(splitToken[1])
	introspectUrl := GetConfig().GetString("oauth_introspect_url")

	resty.SetDebug(true)
	res, err := resty.R().SetFormData(map[string]string{"token": token}).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").Post(introspectUrl)

	if err != nil {
		log.Error(err)
		return introToken, err
	}


	err = json.Unmarshal(res.Body(), &introToken)

	if err != nil {
		log.Error(err)
		return introToken, err
	}

	if introToken.Active != true {
		return introToken, errors.New("Token is not active")
	}

	return introToken, nil
}
