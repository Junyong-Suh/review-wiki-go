package handlers

import (
    "net/http"

    u "github.com/Junyong-Suh/review-wiki-go/utils"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    title, err := u.GetTitle(w, r)
    if err != nil {
        return
    }
    body := r.FormValue("body")
    p := &u.Page{Title: title, Body: []byte(body)}
    err = p.Save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
