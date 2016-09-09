package main

import (
	// "database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	// _ "github.com/go-sql-driver/mysql" //_ "github.com/lib/pq" (For PostgreSQL)
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/pjebs/restgate"
	"net/http"
)

func main() { //On Google App Engine you don't use main() use init()
	app := negroni.New()

	//These middleware is common to all routes
	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())
	app.UseHandler(NewRoute())
	http.Handle("/", context.ClearHandler(app))
	app.Run(":8080") //On Google App Engine, you don't use this
}

func NewRoute() *mux.Router {

	//Create subrouters
	restRouter := mux.NewRouter()
	restRouter.HandleFunc("/api", Handler1()) //Rest API Endpoint handler -> Use your own

	// rest2Router := mux.NewRouter()
	// rest2Router.HandleFunc("/api2", Handler2()) //A second Rest API Endpoint handler -> Use your own

	//Create negroni instance to handle different middlewares for different api routes
	negRest := negroni.New()
	negRest.Use(restgate.New("X-Auth-Key", "X-Auth-Secret", restgate.Static, restgate.Config{Context: C, Key: []string{"12345"}, Secret: []string{"secret"}}))
	negRest.UseHandler(restRouter)

	//Create main router
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.HandleFunc("/", MainHandler()) //Main Handler -> Use your own
	mainRouter.Handle("/api", negRest)        //This endpoint is protected by RestGate via hardcoded KEYs
	// mainRouter.Handle("/api2", negRest2) //This endpoint is protected by RestGate via KEYs stored in a database

	return mainRouter

}

//Optional Context - If not required, remove 'Context: C' or alternatively pass nil (see above)
//NB: Endpoint handler can determine the key used to authenticate via: context.Get(r, 0).(string)
func C(r *http.Request, authenticatedKey string) {
	context.Set(r, 0, authenticatedKey) // Read http://www.gorillatoolkit.org/pkg/context about setting arbitary context key
}

//Endpoint Handlers
func Handler1() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/api -> Handler1 - protected by RestGate (Static Mode)\n")
	}
}

func MainHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/ -> MainHandler - not protected by RestGate\n")
	}
}
