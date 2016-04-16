package model

import (
  "github.com/satori/go.uuid"
  "encoding/csv"
  "os"
  "io"
  "bufio"
  "strconv"
  "fmt"
)

type InputRecord struct {
  Uid uuid.UUID
  Timestamp uint64
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

  var res []InputRecord;
  r := csv.NewReader(inReader)
  for i := 0; ;i++  {
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    if i == 0 {
      continue;
    }

    timestamp, _ := strconv.ParseUint(record[0], 10, 64)
    acX, _ := strconv.ParseFloat(record[1], 64)
    acY,_ := strconv.ParseFloat(record[2], 64)
    acZ,_ := strconv.ParseFloat(record[3], 64)
    accuracy,_ := strconv.ParseFloat(record[4], 64)
    bearing,_ := strconv.ParseFloat(record[5], 64)
    speed,_ := strconv.ParseFloat(record[6], 64)
    longitude,_ := strconv.ParseFloat(record[7], 64)
    latitude,_ := strconv.ParseFloat(record[8], 64)
    altitude,_ := strconv.ParseFloat(record[9], 64)
    
    ir := InputRecord{
      Timestamp: timestamp,
      AcX: acX,
      AcY: acY,
      AcZ: acZ,
      Accuracy: accuracy,
      Bearing: bearing,
      Speed: speed,
      Longitude: longitude,
      Latitude: latitude,
      Altitude: altitude,
    }
    res = append(res, ir)
  }
  return res
}

func CSVTest() {
  pathh := "/Users/alexeyvelikiy/Downloads/1460834335548.csv"
  f, _ := os.Open(pathh)
  items := FromCSVFile(bufio.NewReader(f))
  fmt.Println(len(items))
}