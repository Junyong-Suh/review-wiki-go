package handlers

import (
    "net/http"
    "fmt"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
