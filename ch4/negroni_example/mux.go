package main
import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)


r := mux.NewRouter()

r.HandleFunc("/user/{user}", func (w httpResponseWritter, req *http.Request)  {
	user : mux.Vars(req)["user"]
	fmt.Fprint(w, "hi %s\n", user)
}).Methods("Get")