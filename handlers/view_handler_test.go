package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"
)

func TestViewHandler(t *testing.T) {
	// Create a request to pass to our handler
	title := "readme"
	route := fmt.Sprintf("/view/%s", title)
    req, err := http.NewRequest("GET", route, nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(ViewHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
//     expected := fmt.Sprintf(`<h1>readme</h1>
//
// <p>[<a href="/edit/readme">edit</a>]</p>
//
// <div>this is a test readme file
// </div>`)
//     if rr.Body.String() != expected {
//         t.Errorf("handler returned unexpected body: got %v want %v",
//             rr.Body.String(), expected)
//     }
}
