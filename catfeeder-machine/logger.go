package main

import (
    "os"
    "strings"
    "sync"
    "time"
)

type Logger struct {
    Lock        sync.Mutex
    Fpath       *os.File
    Prefix      string
}

const LogFilename string = "/tmp/cat-feeder.log"

/**
 * @brief:  Checks to see if the log file is present from a previous
 *          execution. If true, roll over log with timestamp. Then
 *          open a new log file.
 *
 * @return: nil on success, else error
 **/
func LoggerInit() error {
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
 * @brief:  Create a Logger with prefix and filepath
 *
 * @arg:    prefix - String to prepend to the log message
 *
 * @return: New logger, nil if error
 **/
func NewLogger(prefix string) *Logger {
    fpath, err := os.OpenFile(LogFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

    if err != nil {
        return nil
    }

    return &Logger{
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
func (l *Logger) Println(msg string) (int, error) {
    l.Lock.Lock()
    defer l.Lock.Unlock()

    output := []byte(time.Now().Format("2006/01/02 15:04:05") + " " + l.Prefix + " - " + msg + "\n")
    return l.Fpath.Write(output)
}
