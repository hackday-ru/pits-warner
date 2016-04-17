package utils

import (
  "model"
  //"math"
  //"math"
)

type DBReader struct {
  CompoundConnector
}

func (reader DBReader) Read() *model.FindResult {
  result := new(model.FindResult)
  return result
}


func (reader DBReader) getByLocation(central, lu, lb, ru, rb model.Coord) *model.FindResult{
  result := new(model.FindResult)

  //reader.RedisConnector.GeoRadius()



  return result
}
