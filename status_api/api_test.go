package status_api

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func beforeTest(method string, url string) (*httptest.ResponseRecorder, *http.Request, error) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)

	return w, req, err
}

func testStringPattern(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Pattern handler should be \"%s\", was \"%s\"", expected, actual)
	}
}

func testStatusCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Status code should be %v, was %v", expected, actual)
	}
}

func TestBadPatternReturns404(t *testing.T) {
	w, req, _ := beforeTest("GET", "/harbles")
	handler, _ := Handlers().Handler(req)
	handler.ServeHTTP(w, req)
	testStatusCode(t, http.StatusNotFound, w.Code)
}

func TestHomeUsesCorrectPattern(t *testing.T) {
	_, req, _ := beforeTest("GET", "/index.html")

	_, pattern := Handlers().Handler(req)

	// Handler pattern should be "/"
	testStringPattern(t, "/", pattern)
}

func TestHomeRedirects(t *testing.T) {	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/index.html", nil)

	handler, _ := Handlers().Handler(req)
	handler.ServeHTTP(w, req)

	// Root returns a redirect for directory indexes
	testStatusCode(t, http.StatusMovedPermanently, w.Code)
}

func TestDockerUsesCorrectPattern(t *testing.T) {
	_, req, _ := beforeTest("GET", "/docker/containers")
	_, pattern := Handlers().Handler(req)

	testStringPattern(t, "/docker/containers", pattern)
}

func TestDockerReturns200(t *testing.T) {
	w, req, _ := beforeTest("GET", "/docker/containers")

	handler, _ := Handlers().Handler(req)
	handler.ServeHTTP(w, req)

	testStatusCode(t, http.StatusOK, w.Code)	
}

func TestDockerReturnsJSONHeader(t *testing.T) {
	w, req, _ := beforeTest("GET", "/docker/containers")
	handler, _ := Handlers().Handler(req)
	handler.ServeHTTP(w, req)

	contentType := w.Header().Get("Content-Type")
	expected := "application/json"
	if contentType != expected {
		t.Errorf("Incorrect contentType, should be \"%s\", was \"%s\"", expected, contentType)
	}
}

func TestDockerReturnsCorrectJSON(t *testing.T) {
	// I am going to have to refactor extensively for dependency
	// injection in order to make this part testable, I think.
	t.Skip()
}

func TestDockerRejectsNonGET(t *testing.T) {
	w, req, _ := beforeTest("PUT", "/docker/containers")
	handler, _ := Handlers().Handler(req)
	handler.ServeHTTP(w, req)

	testStatusCode(t, http.StatusMethodNotAllowed, w.Code)
}

