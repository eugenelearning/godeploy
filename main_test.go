package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleTestRoute(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/testme", nil)
	w := httptest.NewRecorder()

	handleTestRoute(w, request)

	resp := w.Body.String()

	if resp != "it works" {
		t.Errorf("expected 'it works' got %v", resp)
	}
}
