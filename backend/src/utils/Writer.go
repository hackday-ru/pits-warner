package utils

import (
  "strconv"
)

type DBWriter struct {
  CompoundConnector
}

var MARKERS = "markers"

func toString(v float64) string {
  return strconv.FormatFloat(float64(v), 'f', 5, 64)
}

//
//func (writer DBWriter) Write(rec model.InputRecord)  {
//  acc := toString(rec.AcX) + ":" + toString(rec.AcY) + ":" + toString(rec.AcZ)
//  val := acc
//
//
//  writer.RedisConnector.Set(rec.Uid.String(), val, 0)
//  writer.RedisConnector.GeoAdd(MARKERS, &redis.GeoLocation{
//    Longitude:rec.GeoY, Latitude: rec.GeoY, Name: rec.Uid.String()})
//
//  session, _ := writer.CassConnector.CreateSession()
//  defer session.Close()
//
//  session.Query(
//    "INSERT INTO geodata" +
//    "(uuid, time, geoX, geoY, geoZ, acX, acY, acZ)" +
//    "values (?, ?, ?, ?, ?, ?, ?, ?)",
//    rec.Uid,
//    rec.Timestamp,
//    rec.GeoX, rec.GeoY, rec.GeoZ,
//    rec.AcX, rec.AcY, rec.AcZ).Exec()
//}
