package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    sum := 0
    for j := 1; j <= 100; j++ { 
    for i := 1; i <= 1000000; i++ {
        sum += i
    }
    }
    fmt.Fprint(w, "Hello, world!")
}
