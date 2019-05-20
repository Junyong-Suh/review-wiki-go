# Review Wiki

Wiki service in Go from [https://golang.org/doc/articles/wiki/](https://golang.org/doc/articles/wiki/), aims to be review service.

## Requirements
* Go 1.12.5 or higher

## Install
```
$ git clone git@github.com:Junyong-Suh/review-wiki-go.git
$ go mod tidy
```

## Run
```
$ go build .
$ go run .
2019/05/20 23:25:04 review-wiki-go is running on port 8080
```

## Endpoints

### GET /{any_string}/

Prints out "Hi there, I love {any_string}!"

```
$curl localhost:8080
Hi there, I love !

$curl localhost:8080/something
Hi there, I love something!
```

### GET /view/{title}/

Try open up from <span style="color:brown">*browser*</span>.

Renders an article with the title and contents. 302 if none exists. Reads from a file on local named {title}.txt

### GET /edit/{title}/

Try open up from <span style="color:brown">*browser*</span>.

Renders an article in edit mode. Create if none exists.

### GET /save/{title}/

Try open up from <span style="color:brown">*browser*</span>.

Saves an article. 
