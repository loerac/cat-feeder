package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

/**
 * @brief:  Handle for creating a new feeding time proposal.
 **/
func CreateNewFeedTime(w http.ResponseWriter, r *http.Request) {
    var ft []FeedingTimes

    req_body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(req_body, &ft)

    feeding_times = ft
    json.NewEncoder(w).Encode(feeding_times)
}

/**
 * @brief:  Handle for recieving a specific feeding time.
 **/
func ReturnSingleFeedingTime(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, ft := range feeding_times {
        if ft.ID == key {
            json.NewEncoder(w).Encode(ft)
        }
    }
}

/**
 * @brief:  Handle for recieving all the feeding times.
 **/
func ReturnAllFeedingTimes(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(feeding_times)
}

/**
 * @brief:  Handle to recieving info on REST API.
 **/
func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Cat Feeding times")
    fmt.Fprintln(w, "-----------------------")
    fmt.Fprintln(w, "Create new feeding times with '/feedingTime'")
    fmt.Fprintln(w, "Return all feeding time with '/feedingTimes'")
    fmt.Fprintln(w, "Return single feeding time with '/feedingTime/<ID>'")
}

/**
 * @brief:  Handle to monitor all the CRUD methods.
 **/
func HandleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", HomePage)
    myRouter.HandleFunc("/feedingTime", CreateNewFeedTime).Methods("POST")
    myRouter.HandleFunc("/feedingTimes", ReturnAllFeedingTimes).Methods("GET")
    myRouter.HandleFunc("/feedingTime/{id}", ReturnSingleFeedingTime).Methods("GET")
    log.Fatal(http.ListenAndServe(":6969", myRouter))
}
