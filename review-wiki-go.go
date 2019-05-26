package main

import (
    "log"
    "net/http"

    h "github.com/Junyong-Suh/review-wiki-go/handlers"
)

func main() {
    http.HandleFunc("/", h.RootHandler)
    http.HandleFunc("/view/", h.ViewHandler)
    http.HandleFunc("/edit/", h.EditHandler)
    http.HandleFunc("/save/", h.SaveHandler)
    log.Println("review-wiki-go is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
