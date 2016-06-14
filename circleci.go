package main

import (
  "io"
  "net/http"
  "log"
  //"os"
)

func main() {
  //port := os.Getenv("PORT")

  http.HandleFunc("/", base)
  http.HandleFunc("/iwasclicked", iwasclicked)

  serverStart := http.ListenAndServe(":5000", nil)
  log.Fatal(serverStart)
}


func base(response http.ResponseWriter, request *http.Request)  {
  io.WriteString(response,"<!DOCTYPE html><body><a href=\"/iwasclicked\"><button name=\"Click\">Click me</button></a></body></html>")
}


func iwasclicked(response http.ResponseWriter, request *http.Request) {
  io.WriteString(response, "<!DOCTYPE html><body><a href=\"/iwasclicked\"><button name=\"Click\">Click me</button></a><br><p>Yay new text is showing :)</p></body></html>")
}
