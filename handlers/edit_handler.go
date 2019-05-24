package handlers

import (
    "net/http"

    u "github.com/Junyong-Suh/review-wiki-go/utils"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
    title, err := u.GetTitle(w, r)
    if err != nil {
        return
    }
    p, err := u.LoadPage(title)
    if err != nil {
        p = &u.Page{Title: title}
    }
    u.RenderTemplate(w, "edit", p)
}
