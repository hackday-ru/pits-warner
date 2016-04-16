package model

import (
  "github.com/satori/go.uuid"
)

type InputRecord struct {
  Uid uuid.UUID
  timestamp int64
  geoX float64
  geoY float64
  geoZ float64
  acX float64
  acY float64
  acZ float64
}

func FromFile(path string) []InputRecord {
  var res []InputRecord
  //res.
  return res
}