package routers

import (
	"net/http"

	C "github.com/Mynor2397/api-simple/src/controllers"
	M "github.com/Mynor2397/api-simple/src/middlewares"
	"github.com/matryer/way"
)

// New retorna un router
func New() http.Handler {
	api := way.NewRouter()
	api.HandleFunc("POST", "/login", C.Login)
	api.HandleFunc("POST", "/user", C.CreateUser)

	router := way.NewRouter()
	router.Handle("*", "/api...", http.StripPrefix("/api", M.WithAuth(api)))
	return router
}
