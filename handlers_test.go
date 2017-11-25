package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/olenedr/esamarathon/db"
	"github.com/olenedr/esamarathon/handlers"
	"github.com/stretchr/testify/assert"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func initDb(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		checkError(err, t)
	}
	if err := db.Connect(); err != nil {
		checkError(err, t)
	}
}

func TestIndexHandler(t *testing.T) {
	body := testHandler(t, "/", handlers.Index, http.StatusOK)
	contains := "Welcome to European Speedrunner Assembly!"
	assert.Contains(t, body, contains, "Response body differs")
}

func TestNewsHandler(t *testing.T) {
	initDb(t)
	body := testHandler(t, "/news", handlers.News, http.StatusOK)
	contains := "Latest news from ESA"
	assert.Contains(t, body, contains, "Response body differs")
}

func TestScheduleHandler(t *testing.T) {
	body := testHandler(t, "/schedule", handlers.Schedule, http.StatusOK)
	contains := "<section class=\"schedule\">"
	assert.Contains(t, body, contains, "Response body differs")
}

func TestLoginHandler(t *testing.T) {
	body := testHandler(t, "/login", handlers.HandleAuth, http.StatusOK)
	contains := "Log in with Twitch.tv"
	assert.Contains(t, body, contains, "Response body differs")
}

func TestNotFoundHandler(t *testing.T) {
	body := testHandler(t, "/randomroutethatshouldntexist", handlers.HandleNotFound, http.StatusNotFound)
	contains := "Sorry my dude, couldn't find the page you were looking for"
	assert.Contains(t, body, contains, "Response body differs")
}

func testHandler(t *testing.T, path string, h http.HandlerFunc, expectedStatus int) string {
	req, err := http.NewRequest("GET", path, nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	//Make the handler function satisfy http.Handler
	handler := http.HandlerFunc(h)
	handler.ServeHTTP(rr, req)

	//Confirm the response has the right status code
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", expectedStatus, status)
		return ""
	}

	return rr.Body.String()

}
