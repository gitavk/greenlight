package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkMarshalIndent(b *testing.B) {
	w := httptest.NewRecorder()
	r := new(http.Request)
	for b.Loop() {
		healthcheckHandlerMarshalIndent(w, r)
	}
}

func BenchmarkMarshal(b *testing.B) {
	w := httptest.NewRecorder()
	r := new(http.Request)
	for b.Loop() {
		healthcheckHandlerMarshal(w, r)
	}
}
