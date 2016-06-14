package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWithNoText(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something failed", http.StatusInternalServerError)
	}

	req, err := http.NewRequest("GET", "http://cjcircle.herokuapp.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	string := w.Body.String()
	subString := "Yay new text is showing :)"
	fmt.Printf("%s", string)
	if strings.Contains(string, subString) {
		t.Fail()
	}
}

func TestWithText(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something failed", http.StatusInternalServerError)
	}

	req, err := http.NewRequest("GET", "http://cjcircle.herokuapp.com/iwasclicked", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	string := w.Body.String()
	subString := "Yay new text is showing :)"
	fmt.Printf("%s", string)
	if strings.Contains(string, subString) {
		t.Fail()
	}

}
