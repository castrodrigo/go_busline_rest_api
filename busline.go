package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/twinj/uuid"
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

var busLines []BusLine

// GET

func GetBusLines(writer http.ResponseWriter, request *http.Request) {
    json.NewEncoder(writer).Encode(busLines)
}

func GetBusLineById(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    for _, item := range busLines {
        if item.ID == params["id"] {
            json.NewEncoder(writer).Encode(item)
            return
        }
    }
    // Return empty object
    json.NewEncoder(writer).Encode(&BusLine{})
}

// POST

func AddBusLine(writer http.ResponseWriter, request *http.Request) {
    var busLine BusLine
    busLine.ID = uuid.Formatter(uuid.NewV4(), uuid.FormatCanonicalCurly)
    _ = json.NewDecoder(request.Body).Decode(&busLine)
    busLines = append(busLines, busLine)
    json.NewEncoder(writer).Encode(busLines)
}

// PUT

func EditBusLine(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    for index, item := range busLines {
        if item.ID == params["id"] {
            _ = json.NewDecoder(request.Body).Decode(&item)
            busLines = append(busLines[:index], item)
            break
        }
    }
    json.NewEncoder(writer).Encode(busLines)
}

// DELETE

func DeleteBusLine(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    for index, item := range busLines {
        if item.ID == params["id"] {
            busLines = append(busLines[:index], busLines[index+1:]...)
            break
        }
    }
    json.NewEncoder(writer).Encode(busLines)
}

// MAIN

func main() {
    busLines = append(busLines, BusLine {
            ID: uuid.Formatter(uuid.NewV4(), uuid.FormatCanonicalCurly),
            Name: "Manoel Torres",
            Number: "107",
            Route: &Route {
                    From: "Bingen Bus Terminal",
                    To: "Historic Center Bus Terminal" }})

    router := mux.NewRouter()
    router.HandleFunc("/busline/", GetBusLines).Methods("GET")
    router.HandleFunc("/busline/{id}", GetBusLineById).Methods("GET")
    router.HandleFunc("/busline/", AddBusLine).Methods("POST")
    router.HandleFunc("/busline/{id}", EditBusLine).Methods("PUT")
    router.HandleFunc("/busline/{id}", DeleteBusLine).Methods("DELETE")

    // RUNNING PORT

    log.Fatal(http.ListenAndServe(":9009", router))
}
