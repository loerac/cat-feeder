package main

import (
    "fmt"
    "sync"
    "time"

	"gobot.io/x/gobot/platforms/raspi"
)

type FeedingTimes struct {
    ID      string  `json:"id"`
    Hour    int     `json:"hour"`
    Minute  int     `json:"minute"`
}

var has_been_fed bool = false
var feeding_times []FeedingTimes
var mut sync.Mutex

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
    raspi_adapter := raspi.NewAdaptor()
    if raspi_adapter == nil {
        fmt.Println("Error setting new adaptor for raspberry pi")
    }

    /* Run REST API */
    go HandleRequests()

    /* Monitor the feeding time */
    go func() {
        for {
            if TimeToFeedCat() && !has_been_fed {
                fmt.Println("Time to feed cats")
                RotateMotor()
                has_been_fed = true
            } else if !TimeToFeedCat() {
                fmt.Println("Not time to feed cats")
                has_been_fed = false
            }
            time.Sleep(10 * time.Second)
        }
    }()

    time.Sleep(10 * time.Minute)
}
