package utils

import (
  "model"
  //"math"
)

type DBReader struct {
  CompoundConnector
}

func (reader DBReader) read() *model.FindResult {
  result := new(model.FindResult)
  return result
}

//func dist(c0, c1 model.Coord) float64 {
//
//
//  var R = 6371000; // metres
//  var φ1 = c0.Lat.toRadians();
//  var φ2 = c1.Lat.toRadians();
//  var Δφ = (c1.Lat-c0.Lat).toRadians();
//  var Δλ = (c1.Lng-c0.Lng).toRadians();
//
//  var a = math.Sin(Δφ/2) * math.Sin(Δφ/2) +
//  math.Cos(φ1) * math.Cos(φ2) *
//  math.Sin(Δλ/2) * math.Sin(Δλ/2);
//  var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a));
//
//  var d = R * c;
//}

func (reader DBReader) getByLocation(central, lu, lb, ru, rb model.Coord) *model.FindResult{
  result := new(model.FindResult)




  return result
}
