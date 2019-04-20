package main

import (
  "./model"

  "fmt"
  "flag"
  "log"
	"net/http"
  "encoding/json"
  "strconv"
)


var addr = flag.String("addr", ":8080", "http service address")

func main() {

  flag.Parse()
  log.Println("start service");
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
  log.Println("getEvents()");
  d, _ := json.Marshal(model.GetEvents())
  fmt.Fprintf(w, string(d))
}
func getScores(w http.ResponseWriter, r *http.Request) {
  // 指定がないと全てのデータを出力するようなものにする
  log.Println("getScores()");
  d, _ := json.Marshal(model.GetScores())
  fmt.Fprintf(w, string(d))
}
func addEvent(w http.ResponseWriter, r *http.Request) {
  // 指定がないと全てのデータを出力するようなものにする
  log.Println("addEvent()");
  bet,_ := strconv.Atoi(r.URL.Query().Get("bet"))
  term,_ := strconv.Atoi(r.URL.Query().Get("term"))
  model.AddBet(term,bet);
  fmt.Fprintf(w, "json")
}
func addDistance(w http.ResponseWriter, r *http.Request) {
  // 指定がないと全てのデータを出力するようなものにする
  log.Println("addDistance()");
  model.AddDistance();
  fmt.Fprintf(w, "json")
}
