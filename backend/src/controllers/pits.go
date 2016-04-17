package controllers

import (
  "net/http"
  "fmt"
  //"utils"
  //"utils"
  //"utils"
  "utils"
  "model"
  "encoding/json"
)

func PitsHandler(w http.ResponseWriter, r *http.Request) {
  lat := r.URL.Query().Get("lat")
  lng := r.URL.Query().Get("lng")
  rad := r.URL.Query().Get("radius")

  res := utils.GetConn().ReadByLocation(model.Coord{
    Lng: lng,
    Lat: lat}, rad)

  js, err := json.Marshal(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
