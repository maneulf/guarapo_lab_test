package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	apserver "github.com/maneulf/guarapo_lab_test/internal/server"
)

func TestLogin(t *testing.T) {
	jsonData := []byte(`{"username": "ana"}`)
	os.Setenv("PERSISTENCE_TYPE", "inmemory")
	//os.Setenv("PERSISTENCE_TYPE", "sqlite")
	s := apserver.New()
	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonData))

	w := httptest.NewRecorder()
	s.Eng.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
