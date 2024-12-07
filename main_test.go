package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealthcheckRoute(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hc", nil)
	w := httptest.NewRecorder()

	handleHealthcheckRoute(w, request)

	resp := w.Body.String()

	if resp != "it works" {
		t.Errorf("expected 'it works' got %v", resp)
	}
}
