package main


import (
  "testing"
  "net/http"
  "net/http/httptest"
  "log"
  "strings"
  "fmt"
)


func TestAppIO (t *testing.T) {
  handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something failed", http.StatusInternalServerError)
	}

  req, err := http.NewRequest("GET", "http://google.com", nil)
  	if err != nil {
  		log.Fatal(err)
  	}

  w := httptest.NewRecorder()
	handler(w, req)

string := w.Body.String()
subString := "Yay new text is showing :)"
fmt.Printf("%s",  string)
if !strings.Contains(string, subString) {
  t.Fail()
}

}
