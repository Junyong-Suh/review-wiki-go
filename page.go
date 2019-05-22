package main

import (
    "io/ioutil"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    // update Page to have proper unique filename to store as txt
    filename := p.Title + ".txt"

    // The octal integer literal 0600 indicates that the file should be created with read-write permissions for the current user only.
    return ioutil.WriteFile(filename, p.Body, 0600)
}
