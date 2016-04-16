package model

import (
  "github.com/satori/go.uuid"
)

type InputRecord struct {
  uuid.UUID Uid,
  timestamp int64,
  geoX float64,
  geoY float64,
  geoZ float64,
  acX float64,
  acY float64,
  acZ float64
}