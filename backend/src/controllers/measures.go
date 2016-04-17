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
)

type Wrapper struct {
  model.InputRecord

  Z float64
}

func (ws []Wrapper) Len() int {
  return len(ws)
}

func (ws []Wrapper) Swap(i, j int) {
  ws[i], ws[j] = ws[j], ws[i]
}
func (ws []Wrapper) Less(i, j int) bool {
  return len(ws[i].Z) < len(ws[j].Z)
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
    fmt.Println(r.MultipartForm.File["uploadfile"])
    file, _, err := r.FormFile("uploadfile")
    if err != nil {
      fmt.Println(err)
      return
    }


    defer file.Close()
    
    items := model.FromCSVFile(bufio.NewReader(file))

    var z_val [len(items)]Wrapper

    // 1) z-val
    for i, e := range items {
      z_val[i].Z = math.Sqrt2(e.AcX^2 + e.AcY^2 + e.AcZ^2)
      z_val[i] = e
    }
    // 2) .99 quantile
    sort.Sort(z_val)

    topI := len(z_val) * 99 / 100
    topZSlice := z_val[topI:]

    var diff [len(topZSlice)]Wrapper
    // 3) diff
    for i := 0; i < len(topZSlice - 1); i++ {
      diff[i] = topZSlice[i + 1] - topZSlice[i]
      diff[i] = topZSlice[i]
    }
    diff(len(topZSlice - 1)) = topZSlice[len(topZSlice)]
    // 4) pow2

    for i, e := range diff {
      diff[i].Z = e.Z^2
    }

    var filtered [len(diff)]bool
    // 5) filter < 40

    for i, e := range(diff) {
      filtered[i] = e.Z > 40
    }



    for i, e := range diff {
      if filtered[i] {
        utils.GetConn().Write(e)
      }
    }

    fmt.Fprintf(w, `
      processed %d items
    `, len(items))
    
  }
  
}