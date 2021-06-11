package routers

import (
	"backend/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case "GET":


	
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
}

func GetRouter() *mux.Router {
	var r = mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/api/v1/capes", handlers.AddCape).Methods("PUT")
	r.HandleFunc("/api/v1/capes", handlers.GetCapes).Methods("GET")
	r.HandleFunc("/api/v1/capes/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", handlers.DeleteCape).Methods("DELETE")
	r.HandleFunc("/api/v1/capes/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", handlers.GetSingleCape).Methods("GET")
	r.HandleFunc("/api/v1/capes/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", handlers.SetType).Methods("PUT")

	r.HandleFunc("/api/v1/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/api/v1/projects/{project}", handlers.UpdateProjectVersion).Methods("PUT")
	r.HandleFunc("/api/v1/projects/{project}", handlers.GetProject).Methods("GET")
	r.HandleFunc("/api/v1/projects/{project}/link", handlers.UpdateProjectLink).Methods("Put")

	r.HandleFunc("/api/v1/auth/login", handlers.Login).Methods("PUT")
	r.HandleFunc("/api/v1/auth/register", handlers.Signup).Methods("PUT")
	r.HandleFunc("/api/v1/auth/users", handlers.GetAllAccounts).Methods("GET")
	r.HandleFunc("/api/v1/auth/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", handlers.GetPerms).Methods("GET")
	r.HandleFunc("/api/v1/auth/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", handlers.CanRun).Methods("PUT")
	r.HandleFunc("/api/v1/auth/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/admin", handlers.SetAdmin).Methods("POST")
	r.HandleFunc("/api/v1/auth/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/hwid", handlers.ResetHWID).Methods("POST")
	r.HandleFunc("/api/v1/auth/users/{uuid:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/access", handlers.SetAccess).Methods("POST")

	r.HandleFunc("/api/v1/auth/token", handlers.CheckToken).Methods("GET")
	r.HandleFunc("/api/v1/auth/me", handlers.Me).Methods("GET")


	return r
}


