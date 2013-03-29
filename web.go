package main

import (
  "fmt"
  "time"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/schema"
  "labix.org/v2/mgo"
  "encoding/json"
  "os"
)

var recordsCollection *mgo.Collection

type Record struct {
  IP string `schema:"ip" json:"ip"`
  Service string `schema:"service" json:"service"`
  Path string `schema:"path" json:"path"`
  Params string `schema:"params" json:"params"`
  Time time.Time `schema:"time" json:"time"`
}

func track(w http.ResponseWriter, r *http.Request) {
  // parse form data
  r.ParseMultipartForm(1000)

  // protect endpoint with a password
  if (r.FormValue("password") != os.Getenv("PASSWORD")) {
    w.WriteHeader(401)
    return
  }

  // get Record
  decoder := schema.NewDecoder()
  record := new(Record)
  decoder.Decode(record, r.Form)
  record.Time, _ = time.Parse(time.RFC3339, r.FormValue("time"))

  // save record
  recordsCollection.Insert(record)

  // send record back
  str, _ := json.Marshal(record)
  w.WriteHeader(201)
  fmt.Fprintf(w, string(str))
}

func main() {
  mongoUrl := os.Getenv("MONGOHQ_URL")
  if (mongoUrl == "") {
    mongoUrl = "mongodb://localhost/strasbourg-stats"
  }

  mongoSession, _ := mgo.Dial(mongoUrl)

  mongoSession.SetMode(mgo.Monotonic, true)
  recordsCollection = mongoSession.DB("").C("records")

  r := mux.NewRouter()
  r.HandleFunc("/track", track).Methods("POST")
  http.Handle("/", r)
  http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
