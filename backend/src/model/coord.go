package model


type Coord struct {
  Lng float64
  Lat float64
}


type FindResult struct {
  Items []Coord
}