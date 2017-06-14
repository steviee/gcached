package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var m *mux.Router
var rec *httptest.ResponseRecorder

func TestBucketsIndex(t *testing.T) {
	m = NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	m.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

func TestBucketCreate(t *testing.T) {
	m = NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testbucket", nil)

	m.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Home page didn't return %v", http.StatusCreated)
	}
}

func TestBucketIndex(t *testing.T) {
	m = NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testbucket", nil)

	m.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusCreated)
	}

}

func TestItemCreate(t *testing.T) {
	m = NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testbucket/testitem", nil)

	m.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

func TestItemSet(t *testing.T) {
	m = NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/testbucket/testitem", nil)

	m.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

func TestItemShow(t *testing.T) {
	m = NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testbucket/testitem", nil)

	m.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

func TestItemDelete(t *testing.T) {
	m = NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/testbucket/testitem", nil)

	m.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

func TestBucketDelete(t *testing.T) {
	m = NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/testbucket", nil)

	m.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}
