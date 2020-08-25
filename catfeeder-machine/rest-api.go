package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func PrintFeedingTimes(feeding_times []FeedingTimes) {
    for _, ft := range feeding_times {
        fmt.Printf("{ ID(%s) -- Hour(%d) -- Minute(%d) } ",
            ft.ID, ft.Hour, ft.Minute)
    }
    fmt.Println()
}

/**
 * @brief:  Handle for creating a new feeding time proposal.
 **/
func CreateNewFeedTime(w http.ResponseWriter, r *http.Request) {
    var ft []FeedingTimes

    req_body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(req_body, &ft)

    mut.Lock()
    feeding_times = ft
    json.NewEncoder(w).Encode(feeding_times)
    mut.Unlock()

    fmt.Print("Recieved feeding times: ")
    PrintFeedingTimes(ft)
}

/***
 * @brief:  Handle to feed the cat.
 ***/
func CreateFeedNow(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Place holder to feed cat now, will update once TID003 is complete")
}

/**
 * @brief:  Handle for recieving a specific feeding time.
 **/
func ReturnSingleFeedingTime(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    mut.Lock()
    for _, ft := range feeding_times {
        if ft.ID == key {
            json.NewEncoder(w).Encode(ft)
            fmt.Print("Sending feeding time: ")
            feeding := []FeedingTimes{ft}
            PrintFeedingTimes(feeding)
            break
        }
    }
    mut.Unlock()
}

/**
 * @brief:  Handle for recieving all the feeding times.
 **/
func ReturnAllFeedingTimes(w http.ResponseWriter, r *http.Request) {
    mut.Lock()
    json.NewEncoder(w).Encode(feeding_times)
    fmt.Print("Sending feeding times: ")
    PrintFeedingTimes(feeding_times)
    mut.Unlock()
}

/**
 * @brief:  Handle to recieving info on REST API.
 **/
func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Cat Feeding times")
    fmt.Fprintln(w, "-----------------------")
    fmt.Fprintln(w, "Create new feeding times with '/feedingTime'")
    fmt.Fprintln(w, "Create request to feed cat now with '/feedNow'")
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
    myRouter.HandleFunc("/feedNow", CreateFeedNow).Methods("POST")
    myRouter.HandleFunc("/feedingTimes", ReturnAllFeedingTimes).Methods("GET")
    myRouter.HandleFunc("/feedingTime/{id}", ReturnSingleFeedingTime).Methods("GET")
    log.Fatal(http.ListenAndServe(":6969", myRouter))
}
