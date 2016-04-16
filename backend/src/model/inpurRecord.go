package model

import (
  "github.com/satori/go.uuid"
)

type InputRecord struct {
  Uid uuid.UUID
  Timestamp int64
  GeoX float64
  GeoY float64
  GeoZ float64
  AcX float64
  AcY float64
  AcZ float64
}