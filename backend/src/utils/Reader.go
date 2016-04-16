package utils

import (
  "model"
  //"math"
  "math"
)

type DBReader struct {
  CompoundConnector
}

func (reader DBReader) read() *model.FindResult {
  result := new(model.FindResult)
  return result
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

func (reader DBReader) getByLocation(central, lu, lb, ru, rb model.Coord) *model.FindResult{
  result := new(model.FindResult)

  //reader.RedisConnector.GeoRadius()



  return result
}
