package route

import (
    "github.com/gorilla/mux"
    calender "github.com/heaptracetechnology/google-calender/calender"
    "log"
    "net/http"
)

//Route struct
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

//Routes array
type Routes []Route

var routes = Routes{
    Route{
        "CreateCalender",
        "POST",
        "/createCalender",
        calender.CreateCalender,
    },
    Route{
        "GetEventByID",
        "POST",
        "/getEvent",
        calender.GetEventByID,
    },
    Route{
        "EventList",
        "POST",
        "/eventList",
        calender.EventList,
    },
    Route{
        "CreateEvent",
        "POST",
        "/createEvent",
        calender.CreateEvent,
    },
}

//NewRouter route
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
