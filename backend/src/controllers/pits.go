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
)

func PitsHandler(w http.ResponseWriter, r *http.Request) {
  lat := r.URL.Query().Get("lat")
  lng := r.URL.Query().Get("lng")
  rad := r.URL.Query().Get("radius")

  lt, _ := strconv.ParseFloat(lng, 64)
  ln, _ := strconv.ParseFloat(lat, 64)
  rf, _ := strconv.ParseFloat(rad, 64)

  res := utils.GetConn().ReadByLocation(model.Coord{
    Lng: lt,
    Lat: ln},rf )

  js, err := json.Marshal(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
