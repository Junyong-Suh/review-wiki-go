package handlers

import (
    "net/http"

    u "github.com/Junyong-Suh/review-wiki-go/utils"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
    title, err := u.GetTitle(w, r)
    if err != nil {
        return
    }
    p, err := u.LoadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    u.RenderTemplate(w, "view", p)
}
