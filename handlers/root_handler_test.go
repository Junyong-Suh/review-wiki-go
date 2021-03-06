package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"
)

func TestRootHandler(t *testing.T) {
	// Create a request to pass to our handler
	someString := "some_random_1w1strins"
	rootRoute := fmt.Sprintf("/%s", someString)
    req, err := http.NewRequest("GET", rootRoute, nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(RootHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := fmt.Sprintf("Hi there, I love %s!", someString)
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
