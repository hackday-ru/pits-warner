package main

import (
	"fmt"
	"net/http"
  //"encoding/json"
  //"model"
  //"github.com/satori/go.uuid"

  "utils"
  "model"
  //"github.com/gocql/gocql"
  "github.com/satori/go.uuid"
  "strconv"
  //"encoding/json"
  "log"
)

var conn = new(utils.CompoundConnector)


func pointsHandler(w http.ResponseWriter, r *http.Request) {
  //
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
  rec := model.InputRecord{
    Uid: uuid.NewV4(),
    Longitude: 0.0,
    Latitude: 0.0,
    Altitude: 0.0,
    AcX: 0.0,
    AcY: 0.0,
    AcZ: 0.0,
  }

  //conn.RedisConnector.Set("sample", "val", 0)
  conn.Write(rec)
}
func getMockHandler(w http.ResponseWriter, r *http.Request) {
  val, _ := conn.RedisConnector.Get("sample").Result()
  //rad, _ := conn.RedisConnector.geo

  fmt.Fprintf(w, "<div>%s</div>", val)
}




func addCHandler(w http.ResponseWriter, r *http.Request) {
  rec := model.InputRecord{
    Uid: uuid.NewV4(),
    Longitude: 0.0,
    Latitude: 0.0,
    Altitude: 0.0,
    AcX: 0.0,
    AcY: 0.0,
    AcZ: 0.0,
    Accuracy: 0.0,
    Bearing: 0.0,
    Speed: 0.0,
  }

  session, _ := conn.CassConnector.CreateSession()
  defer session.Close()

  if err := session.Query(
    "INSERT INTO geodata" +
    "(id, time, geoX, geoY, geoZ, acX, acY, acZ)" +
    "values (?, ?, ?, ?, ?, ?, ?, ?)",
    rec.Uid.String(),
    rec.Timestamp,
    rec.Longitude, rec.Latitude, rec.Altitude,
    rec.AcX, rec.AcY, rec.AcZ).Exec(); err != nil {
    log.Fatal(err)
  }
}

func getCHandler(w http.ResponseWriter, r *http.Request) {
  session, _ := conn.CassConnector.CreateSession()
  defer session.Close()

  var geoX float64
  var geoY float64

  var str string
  str += "["

  iter := session.Query(`SELECT geoX, geoY FROM geodata`).Iter()
  for iter.Scan(&geoX, &geoY) {
    //fmt.Println("Tweet:", geoX, geoY)
    str += "{lat: " + toString(geoX) + ",lng: " + toString(geoY) + "},"
  }

  sl := str[0: len(str) - 1]
  sl += "]"
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(sl))
}


func getJA(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(`{
    "0": {
        "lat": 59.89444,
        "lng": 30.26417
    },
    "1": {
        "lat": 59.9458321,
        "lng": 30.4765999
    },
    "3": {
        "lat": 59.8845205,
        "lng": 29.8843764
    },
    "4": {
        "lat": 60.010483,
        "lng": 30.6571437
    }
}`))
}

func toString(v float64) string {
  return strconv.FormatFloat(float64(v), 'f', 5, 64)
}


func main() {

  conn.Init("52.58.116.75:6379","52.58.116.75:9042")

  http.HandleFunc("/hollows", pointsHandler)
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/addMock", addMockHandler)
  http.HandleFunc("/getMock", getMockHandler)
  http.HandleFunc("/addCMock", addCHandler)
  http.HandleFunc("/getCMock", getCHandler)

  http.HandleFunc("/pits", getJA)

	http.ListenAndServe(":8080", nil)

}
