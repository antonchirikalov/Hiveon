// HAVEON-API
//
// GO based backend API
//
//     Schemes: http
//     Host: 95.216.199.4:8099
//     Version: 0.0.1
//     BasePath: /api
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer:
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"hiveon-api/api"
	"hiveon-api/service"
	"hiveon-api/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	httpPort := utils.GetConfig().GetString("httpport")

	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "02-01-2006 15:04:05", PrettyPrint: true})

	mux := http.NewServeMux()
	mux.Handle("/api/", api.MakeServiceHandlers())

	//Add swagger support
	fs := http.FileServer(http.Dir("./swaggerui"))
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	errs := make(chan error, 2)
	go func() {
		log.WithFields(log.Fields{"transport:": "http", "port": httpPort}).Info("HIVEON.API has started")
		// casbin
		e := casbin.NewEnforcer("casbin/authz_model.conf", "casbin/authz_policy.csv")
		casbinHandler := service.Authorizer(e)(mux)
		errs <- http.ListenAndServe(":" + httpPort, handlers.LoggingHandler(os.Stdout, service.OAuthHandler(mux, casbinHandler)))
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Info("terminated", <-errs)
}
