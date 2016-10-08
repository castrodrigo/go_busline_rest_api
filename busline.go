package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type BusLine struct {
    ID        string    `json:"id,omitempty"`
    Name      string    `json:"name,omitempty"`
    Number    string    `json:"number,omitempty"`
    Route     *Route    `json:"route,omitempty"`
}

type Route struct {
    From      string `json:"from,omitempty"`
    To        string `json:"to,omitempty"`
}

var busLine []BusLine

// GET

func GetBusLines(writer http.ResponseWriter, request *http.Request) {
    json.NewEncoder(writer).Encode(busLine)
}

func GetBusLineById(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    for _, item := range busLine {
        if item.ID == params["id"] {
            json.NewEncoder(writer).Encode(item)
            return
        }
    }
}



// MAIN

func main() {
    busLine = append(busLine, BusLine {
            ID: "1",
            Name: "Manoel Torres",
            Number: "107",
            Route: &Route {
                    From: "Bingen Bus Terminal",
                    To: "Historic Center Bus Terminal" }})

    router := mux.NewRouter()
    router.HandleFunc("/busLine/", GetBusLines).Methods("GET")
    router.HandleFunc("/busLine/{id}", GetBusLineById).Methods("GET")

    // RUNNING PORT

    log.Fatal(http.ListenAndServe(":9009", router))
}
