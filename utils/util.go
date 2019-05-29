package utils

import (
    "net/http"
    "io/ioutil"
    "html/template"
    "errors"
    "regexp"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) Save() error {
    // update Page to have proper unique filename to store as txt
    filename := "./data/" + p.Title + ".txt"

    // The octal integer literal 0600 indicates that the file should be created with read-write permissions for the current user only.
    return ioutil.WriteFile(filename, p.Body, 0600)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var templates = template.Must(template.ParseFiles("./templates/edit.html", "./templates/view.html"))

func LoadPage(title string) (*Page, error) {
    filename := "./data/" + title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func GetTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}
