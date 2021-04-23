// main.go
package main

import(
//    "encoding/json"
//    "fmt"
    "log"
//    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}

func handleRequests() {
    router := mux.NewRouter()

    api := router.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("/hello", get).Methods(http.MethodGet)
    api.HandleFunc("/hello", post).Methods(http.MethodPost)
    api.HandleFunc("/hello", put).Methods(http.MethodPut)
    api.HandleFunc("/hello", delete).Methods(http.MethodDelete)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
    //TODO: Configuration

    handleRequests()
}

