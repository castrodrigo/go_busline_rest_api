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
