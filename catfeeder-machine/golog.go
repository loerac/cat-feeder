package main

import (
    "os"
    "strings"
    "sync"
    "time"
)

type Golog struct {
    Lock        sync.Mutex
    Fpath       *os.File
    Prefix      string
}

const LogFilename string = "/tmp/cat-feeder.log"

/**
 * @brief:  Checks to see if the log file is present from a previous
 *          execution. If true, roll over log with timestamp.
 *
 * @return: nil on success, else error
 **/
func InitGolog() error {
    finfo, err := os.Stat(LogFilename)
    if err == nil {
        path := strings.Split(LogFilename, finfo.Name())[0]
        err = os.Rename(
            LogFilename,
            path + time.Now().Format("20060102T150405") + "-" + finfo.Name(),
        )
        if err != nil {
            return err
        }
    }

    return err
}

/**
 * @brief:  Opens and creates a log file. If no errors opening/creating,
 *          then create a Golog with prefix and file pointer
 *
 * @arg:    prefix - String to prepend to the log message
 *
 * @return: New logger, nil if error
 **/
func OpenGolog(prefix string) *Golog {
    fpath, err := os.OpenFile(LogFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

    if err != nil {
        return nil
    }

    return &Golog{
        Prefix: prefix,
        Fpath: fpath,
    }
}

/**
 * @brief:  Write log message to the log file. Log message contains:
 *          - Timestamp
 *          - Prefix
 *          - Message
 *
 * @arg:    msg - Message to write to log file
 *
 * @return: int - How many bytes were written
 *          error - Errors while writing
 **/
func (gl *Golog) Println(msg string) (int, error) {
    gl.Lock.Lock()
    defer gl.Lock.Unlock()

    output := []byte(time.Now().Format("2006/01/02 15:04:05") + " " + gl.Prefix + " - " + msg + "\n")
    return gl.Fpath.Write(output)
}
