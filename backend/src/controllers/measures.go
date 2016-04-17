package controllers

import (
  "net/http"
  "fmt"
  //"os"
  //"io"
  "model"
  "bufio"
  "utils"
  "math"
  "sort"
  "net/http/httputil"
)

type Wrapper struct {
  ir model.InputRecord

  Z float64
}

type Wrappers []Wrapper

func (ws Wrappers) Len() int {
  return len(ws)
}

func (ws Wrappers) Swap(i, j int) {
  ws[i], ws[j] = ws[j], ws[i]
}
func (ws Wrappers) Less(i, j int) bool {
  return ws[i].Z < ws[j].Z
}


func MeasureHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    fmt.Fprintf(w, `
      <h1>Upload data</h1><div>Upload a CSV file</div> 
      <form action="/measures" method="post" enctype="multipart/form-data">
        <input type="file" name="uploadfile" />
        <input type="hidden" name="token" value="{{.}}"/>
        <input type="submit" value="upload" />
      </form>
    `)
  }
  
  if r.Method == "POST" {

    r.ParseMultipartForm(32 << 20)
    dump, err := httputil.DumpRequest(r, true)
    s := string(dump[:])
    fmt.Println(s)
    file, _, err := r.FormFile("uploadfile")
    if err != nil {
      fmt.Println(err)
      return
    }


    defer file.Close()
    
    items := model.FromCSVFile(bufio.NewReader(file))

    var z_val Wrappers

    // 1) z-val
    for i, e := range items {
      z_val = append(z_val, Wrapper{ir: items[i], Z: math.Sqrt(math.Pow(e.AcX,2) + math.Pow(e.AcY,2) + math.Pow(e.AcZ,2))})
    }
    // 2) .99 quantile
    sort.Sort(z_val)

    topI := len(z_val) * 99 / 100
    fmt.Println(topI)
    topZSlice := z_val[topI:]

    //for _, e := range topZSlice {
    //  fmt.Println(e)
    //}

    var diff Wrappers
    // 3) diff
    for i := 0; i < len(topZSlice) - 1; i++ {
      diff = append(diff, topZSlice[i])
      diff[i].Z = topZSlice[i + 1].Z - topZSlice[i].Z
      //fmt.Println(diff[i])
    }
    //diff[len(topZSlice) - 1] = topZSlice[len(topZSlice)]
    // 4) pow2

    for i, e := range diff {
      diff[i].Z = math.Pow(e.Z,2)
    }

    var filtered []bool
    // 5) filter < 40

    for _, e := range(diff) {
      filtered = append(filtered, e.Z > 40)
    }



    for i, e := range diff {
      if filtered[i] {
          utils.GetConn().Write(e.ir)
      }
    }

    fmt.Fprintf(w, `
      processed %d items
    `, len(items))
    
  }
  
}