package main

import (
  "fmt"
  "flag"
  "log"
	"net/http"
)


var addr = flag.String("addr", ":8080", "http service address")

func main() {

  flag.Parse()
  log.Println("hogehoge");
  // POST
  http.HandleFunc("/run",addDistance)
  http.HandleFunc("/join_event",addEvent)
  // GET
  http.HandleFunc("/scores",getScores)
  http.HandleFunc("/events",getEvents)

  err := http.ListenAndServe(*addr, nil)
  if err != nil {
      log.Fatal("ListenAndServer:", err);
  }
}

func getEvents(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "json")
}
func getScores(w http.ResponseWriter, r *http.Request) {
  // 指定がないと全てのデータを出力するようなものにする
  fmt.Fprintf(w, "json")
}
func addEvent(w http.ResponseWriter, r *http.Request) {
  // 指定がないと全てのデータを出力するようなものにする
  fmt.Fprintf(w, "json")
}
func addDistance(w http.ResponseWriter, r *http.Request) {
  // 指定がないと全てのデータを出力するようなものにする
  fmt.Fprintf(w, "json")
}
