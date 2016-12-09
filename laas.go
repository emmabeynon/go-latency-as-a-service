package main

import (
    "net/http"
    "time"
)
// GET `/latency` - returns the string "OK" atfter 500ms
// GET `/latency?duration=100ms` - override the default delay duration.
// An invalid duration parameter should return a 400 along with an error message.
func latencyServer(w http.ResponseWriter, req *http.Request) {
    durationParam := req.FormValue("duration")
    if durationParam == "" {
        time.Sleep(time.Millisecond * 500)
    } else {
        duration, err := time.ParseDuration(durationParam)
        if err != nil {
            http.Error(w, "Error: Invalid duration parameter", 400)
            return
        } else {
            time.Sleep(duration)
        }
    }
    w.Write([]byte("OK"))
}

func main()  {
    http.HandleFunc("/latency", latencyServer)
    http.ListenAndServe(":8080", nil)
}
