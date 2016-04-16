package utils

import "model"

type DBReader struct {
  CompoundConnector
}

func (reader DBReader) read() model.FindResult {
  result := new(model.FindResult)
  return result
}

func (reader DBReader) getByLocation(coord model.Coord, zoom int) model.FindResult{
  result := new(model.FindResult)
  return result
}
