package controllers

import (
  "net/http"
  //"fmt"
  //"utils"
  //"utils"
  //"utils"
  "utils"
  "model"
  "encoding/json"
  "strconv"
  //"fmt"
  "math"

  //"fmt"
)

func PitsHandler(w http.ResponseWriter, r *http.Request) {
  lat := r.URL.Query().Get("lat")
  lng := r.URL.Query().Get("lng")
  radius := r.URL.Query().Get("radius")

  lt, _ := strconv.ParseFloat(lat, 64)
  ln, _ := strconv.ParseFloat(lng, 64)
  rf, _ := strconv.ParseFloat(radius, 64)

  res := utils.GetConn().ReadByLocation(model.Coord {
      Lng: lt,
      Lat: ln,
    },
    rf,
  )


  js,err := json.Marshal(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
  }


func toRad(a float64) float64 {
  return a*math.Pi/180

}

func dist(c0, c1 model.Coord) float64 {
  var R = 6371000.0; // metres
  var φ1 = toRad(c0.Lat);
  var φ2 = toRad(c1.Lat);
  var Δφ = toRad(c1.Lat-c0.Lat);
  var Δλ = toRad(c1.Lng-c0.Lng);

  var a = math.Sin(Δφ/2) * math.Sin(Δφ/2) +
  math.Cos(φ1) * math.Cos(φ2) *
  math.Sin(Δλ/2) * math.Sin(Δλ/2);
  var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a));

  return R * c;
}
