package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "domio_public/routes"
    "domio_public/components/config"
    "log"
)

func NewRouter() *mux.Router {
    conf := config.Config

    router := mux.NewRouter().StrictSlash(true)

    if (conf.Env == "development") {
        log.Print("Development environment, handling static files by Go...")
        router.Path("/style.css").Handler(http.FileServer(http.Dir("/usr/local/domio_client")))
        router.Path("/api_connector.js").Handler(http.FileServer(http.Dir("/usr/local/domio_client")))
        router.Path("/login_page.js").Handler(http.FileServer(http.Dir("/usr/local/domio_client")))
    }

    for _, route := range routes.RoutesList {
        var handler http.Handler
        handler = route.HandlerFunc

        router.
        Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }

    return router
}