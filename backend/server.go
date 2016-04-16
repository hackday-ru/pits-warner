package main

import (
	"fmt"
	"net/http"
  //"encoding/json"
  //"model"
  //"github.com/satori/go.uuid"

  "utils"
  "model"
)

var conn = new(utils.CompoundConnector)


func pointsHandler(w http.ResponseWriter, r *http.Request) {

  //h1 := model.GeoData { Lat:10, Lng:20 }
  //h2 := model.GeoData{ Lat:11.21312, Lng:20.1232 }
  //res := model.FindResult{ []model.GeoData{ h1, h2} }
  //js, err := json.Marshal(res)
  //if err != nil {
  //  http.Error(w, err.Error(), http.StatusInternalServerError)
  //  return
  //}
  //w.Header().Set("Content-Type", "application/json")
  //w.Write(js)
  fmt.Fprintf(w, "yo")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  var title = "Hello to REST serverr"
  var body = "use get / post to /points"
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, body)
}

func addMockHandler(w http.ResponseWriter, r *http.Request) {
  conn.RedisConnector.Set("sample", "val", 0)
}
func getMockHandler(w http.ResponseWriter, r *http.Request) {
  val, _ := conn.RedisConnector.Get("sample").Result()
  //rad, _ := conn.RedisConnector.geo

  fmt.Fprintf(w, "<div>%s</div>", val)
}


func main() {

  model.CSVTest()


  //conn.Init("52.58.116.75:6379","")
  //http.HandleFunc("/hollows", pointsHandler)
  //http.HandleFunc("/", indexHandler)
  //http.HandleFunc("/addMock", addMockHandler)
  //http.HandleFunc("/getMock", getMockHandler)
	//http.ListenAndServe(":8080", nil)
}
