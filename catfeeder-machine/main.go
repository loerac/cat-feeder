package main

import (
    "fmt"
    "log"
    "os"
    "sync"
    "time"
)

type FeedingTimes struct {
    ID      string  `json:"id"`
    Hour    int     `json:"hour"`
    Minute  int     `json:"minute"`
}

var has_been_fed bool = false
var feeding_times []FeedingTimes
var mut sync.Mutex
var logger *log.Logger

/**
 * @brief:  Check if it's time to feed the cat
 *
 * @return: True if it's time to feed the cat,
 *          Else, false
 **/
func TimeToFeedCat() bool {
    mut.Lock()
    defer mut.Unlock()
    ts := time.Now()

    for i := range feeding_times {
        if  feeding_times[i].Hour == ts.Hour() &&
            feeding_times[i].Minute == ts.Minute() {
            return true
        }
    }

    return false
}

func main() {
    /* Open log file */
    /* TODO: perform log rollover */
    f, err := os.OpenFile("/tmp/cat-feeder.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }
    defer f.Close()
    logger := log.New(f, "cat-feeder - ", log.LstdFlags|log.Lmsgprefix)

    /* Run REST API */
    go HandleRequests()

    /* Monitor the feeding time */
    go func() {
        for {
            if TimeToFeedCat() && !has_been_fed {
                logger.Println("Time to feed cats")
                has_been_fed = true
            } else if !TimeToFeedCat() {
                logger.Println("Not time to feed cats")
                has_been_fed = false
            }
            time.Sleep(10 * time.Second)
        }
    }()

    time.Sleep(10 * time.Minute)
}
