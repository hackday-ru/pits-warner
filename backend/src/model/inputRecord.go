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
  timestamp int64
  geoX float64
  geoY float64
  geoZ float64
  acX float64
  acY float64
  acZ float64
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