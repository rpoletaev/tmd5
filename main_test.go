package main

import (
	"bytes"
	"testing"

	"net/http/httptest"

	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

var r = getEngine()

func TestMD5NegativeID(t *testing.T) {
	println("TestMD5NegativeID")
	data := `{"id": -1, "text": "test text"}`
	recorder := performRequest(r, "POST", "/md5", data)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Status code should be %d, was %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestEmptyText(t *testing.T) {
	println("TestEmptyText")
	data := `{"id": 1, "text": ""}`
	recorder := performRequest(r, "POST", "/md5", data)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Status code should be %d, was %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestGreaterThen100LettersText(t *testing.T) {
	println("TestGreaterThen100LettersText")
	data := `{"id": 1, "text": "012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"}`
	recorder := performRequest(r, "POST", "/md5", data)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Status code should be %d, was %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestNormalValues(t *testing.T) {
	println("TestNormalValues")
	data := `{"id": 10, "text": "normal text"}`
	recorder := performRequest(r, "POST", "/md5", data)

	if recorder.Code != http.StatusOK {
		t.Errorf("Status code should be %d, was %d", http.StatusBadRequest, recorder.Code)
	}

	testValue := `"md577814623edc932b7ef7965d2f6c4eb47"`
	if strings.TrimSpace(recorder.Body.String()) != testValue {
		t.Errorf("Wrong body must be a %s, was %s", "md577814623edc932b7ef7965d2f6c4eb47", recorder.Body.String())
	}
}
func performRequest(eng *gin.Engine, method string, path string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	eng.ServeHTTP(recorder, req)
	return recorder
}
