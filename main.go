package main

import (
	"ServerProject/modules"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/person/{id}", GetUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := os.ReadFile("data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var person []modules.User

	err = json.Unmarshal(data, &person)

	for _, user := range person {
		if user.ID == id {
			response, err := json.Marshal(user)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}

}
