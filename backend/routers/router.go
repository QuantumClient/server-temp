package routers

import (
	"backend/controllers"
	"backend/handlers"
	"backend/util"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func GetRouter() *mux.Router {
	var r = mux.NewRouter()
	r.StrictSlash(true)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.ErrorResponse(w, r, util.Err404.Error())
	})

	var api = r.PathPrefix("/api/v1").Subrouter()

	var capes = api.PathPrefix("/capes").Subrouter()

	capes.HandleFunc("/", handlers.GetOldCapes).Methods("GET")
	capes.HandleFunc("/all", handlers.GetCapesFull).Methods("GET")
	capes.HandleFunc("/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", handlers.GetSingleCape).Methods("GET")
	capes.HandleFunc("/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/enabled", handlers.SetEnabled).Methods("POST")
	capes.HandleFunc("/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/type", handlers.SetType).Methods("POST")

	var projects = api.PathPrefix("/projects").Subrouter()

	projects.HandleFunc("/", handlers.GetProjects).Methods("GET")
	projects.HandleFunc("/{project}", handlers.UpdateProjectVersion).Methods("PUT")
	projects.HandleFunc("/{project}", handlers.GetProject).Methods("GET")
	projects.HandleFunc("/{project}/link", handlers.UpdateProjectLink).Methods("PUT")

	var auth = api.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/login", handlers.Login).Methods("PUT")
	auth.HandleFunc("/register", handlers.Signup).Methods("PUT")
	auth.HandleFunc("/refresh", handlers.Refresh).Methods("POST")

	auth.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/api/v1/users", http.StatusTemporaryRedirect)
	}).Methods("GET")
	auth.HandleFunc("/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", handlers.GetPerms).Methods("GET")
	auth.HandleFunc("/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", handlers.CanRunLeg).Methods("PUT")
	auth.HandleFunc("/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/admin", handlers.SetAdmin).Methods("POST")
	auth.HandleFunc("/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/hwid", handlers.ResetHWID).Methods("POST")
	auth.HandleFunc("/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/access", handlers.SetAccess).Methods("POST")
	auth.HandleFunc("/verify", handlers.Verify).Methods("GET")
	auth.HandleFunc("/token", handlers.CheckToken).Methods("GET")
	auth.HandleFunc("/me", handlers.Me).Methods("GET")

	var users = api.PathPrefix("/users").Subrouter()
	users.Use(basicAuth)

	users.HandleFunc("/", handlers.GetAllAccounts).Methods("GET")
	users.HandleFunc("/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/link", handlers.LinkMCAccount).Methods("POST")
	users.HandleFunc("/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/capes", handlers.GetUsersCapes).Methods("GET")
	users.HandleFunc("/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/key", handlers.GetUserKey).Methods("GET")

	var mc = api.PathPrefix("/mc").Subrouter()

	mc.HandleFunc("/online", handlers.GetOnline).Methods("GET")
	mc.HandleFunc("/online", handlers.AddToOnline).Methods("PUT")
	mc.HandleFunc("/online", handlers.RemoveFromOnline).Methods("POST", "DELETE")
	mc.HandleFunc("/capes", handlers.GetGapes).Methods("GET")

	return r
}

func basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		co, err := r.Cookie("auth._token.local")
		if err == nil {
			r.Header.Set("Authorization", strings.Replace(co.Value, "%20", " ", 1))
		}
		token, err := controllers.GetToken(r)
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
