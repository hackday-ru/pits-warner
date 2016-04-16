package model

import (
  "github.com/satori/go.uuid"
  "encoding/csv"
  "os"
  "io"
  "bufio"
  "fmt"
)

type InputRecord struct {
  Uid uuid.UUID
  Timestamp int64
  AcX float64
  AcY float64
  AcZ float64
  Accuracy float64
  Bearing float64
  Speed float64
  Longitude float64
  Latitude float64
  Altitude float64
}

func FromCSVFile(inReader io.Reader) []InputRecord {
  r := csv.NewReader(inReader)
  for {
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    //for value := range record {
    //  if len(record) != 10 {
    //    continue
    //  }
    //
    //  fmt.Println(record[])
    //}
    fmt.Println(record[0])
    fmt.Println("---------")
  }
  return []InputRecord{}
}

func CSVTest() {
  pathh := "/Users/alexeyvelikiy/Downloads/1460834335548.csv"
  f, _ := os.Open(pathh)
  FromCSVFile(bufio.NewReader(f))
}