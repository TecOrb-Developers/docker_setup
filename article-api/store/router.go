package store

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes {
    Route {
        "AddArticle",
        "POST",
        "/articles",
        controller.AddArticle,
    },
    Route {
        "GetArticle",
        "GET",
        "/articles/{id}",
        controller.GetArticle,
    },
    Route {
        "GetArticlesByTagAndDate",
        "GET",
        "/tags/{tagName}/{date}",
        controller.GetArticlesByTagAndDate,
    }}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
         Methods(route.Method).
         Path(route.Pattern).
         Name(route.Name).
         Handler(handler)
    }
    return router
}
