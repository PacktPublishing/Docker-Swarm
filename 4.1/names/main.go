package main

import (
	"bitbucket.org/albertogviana/go-microservice/names/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Env struct {
	db models.Datastore
}

type Health struct {
	Version string `json:"version" form:"version" binding:"required"`
	Status  string `json:"status" form:"status" binding:"required"`
}

func (env *Env) Insert(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data models.Person
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")

		response := `{"error": "Firstname and Lastname must be sent"}`

		io.WriteString(w, response)
		return
	}

	if data.Firstname == "" || data.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")

		response := `{"error": "Firstname and Lastname must be sent"}`

		io.WriteString(w, response)
		return
	}

	err = env.db.InsertName(data)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "")

}

func (env *Env) List(w http.ResponseWriter, r *http.Request) {
	result, err := env.db.GetNames()
	if err != nil {
		log.Fatal(err)
	}

	respBody, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}


	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)

}

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	healthCheck := Health{os.Getenv("APP_VERSION"), "OK"}

	response, err := json.Marshal(healthCheck)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	RunServer()
}

func RunServer() {
	db, err := models.NewDB()
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	http.HandleFunc("/person/add", env.Insert)
	http.HandleFunc("/person", env.List)
	http.HandleFunc("/health", getHealthCheck)
	log.Fatal("ListenAndServe: ", http.ListenAndServe(":8080", nil))
}
