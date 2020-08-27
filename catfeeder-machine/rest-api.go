package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

var rest_log *Golog

func FeedingTimeStr(feeding_times []FeedingTimes) string {
    ft_str := ""

    for _, ft := range feeding_times {
        ft_str += fmt.Sprintf(" { ID(%s) -- Hour(%d) -- Minute(%d) } ",
            ft.ID, ft.Hour, ft.Minute,
        )
    }

    return ft_str
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

    rest_log.Println("Recieved feeding times:" + FeedingTimeStr(ft))
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
            rest_log.Println("Sending feeding time:" + FeedingTimeStr([]FeedingTimes{ft}))
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
    rest_log.Println("Sending feeding times:" + FeedingTimeStr(feeding_times))
    mut.Unlock()
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
    rest_log = OpenGolog("rest-api")

    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", HomePage)
    myRouter.HandleFunc("/feedingTime", CreateNewFeedTime).Methods("POST")
    myRouter.HandleFunc("/feedingTimes", ReturnAllFeedingTimes).Methods("GET")
    myRouter.HandleFunc("/feedingTime/{id}", ReturnSingleFeedingTime).Methods("GET")

    rest_log.Println("Listening on 127.0.0.1:6969")
    rest_log.Println(http.ListenAndServe(":6969", myRouter).Error())
}
