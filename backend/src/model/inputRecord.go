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
  GeoX float64
  GeoY float64
  GeoZ float64
  AcX float64
  AcY float64
  AcZ float64
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