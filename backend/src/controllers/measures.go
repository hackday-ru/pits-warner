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
  Ir model.InputRecord

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

func SaveMe(items []model.InputRecord) {

  var z_val Wrappers

  // 1) z-val
  for i, e := range items {
    //w := I()
    //w.Ir = items[i]
    //w.Z =  math.Sqrt(math.Pow(e.AcX,2) + math.Pow(e.AcY,2) + math.Pow(e.AcZ,2))

    z_val = append(z_val, Wrapper{
      Ir:items[i],
      Z: math.Sqrt(math.Pow(e.AcX,2) + math.Pow(e.AcY,2) + math.Pow(e.AcZ,2)) })
    //fmt.Println(z_val[i])
  }

  //fmt.Println("GHFRIOGHJERWN")


  // 2) .99 quantile
  sort.Sort(z_val)

  topI := len(z_val) * 99 / 100
  //fmt.Println(topI)
  lol := z_val[topI:]
  topZSlice := lol

  var diff Wrappers
  // 3) diff
  for i := 0; i < len(topZSlice) - 1; i++ {
    diff = append(diff, Wrapper{
      Ir: topZSlice[i].Ir,
      Z: topZSlice[i + 1].Z - topZSlice[i].Z})

    //diff = append(diff, topZSlice[i])
    //diff[i].ir = topZSlice[i]
    //diff[i].Z = topZSlice[i + 1].Z - topZSlice[i].Z
  }
  diff = append(diff, Wrapper{
    Ir: topZSlice[len(topZSlice) - 1].Ir,
    Z: topZSlice[len(topZSlice) - 1].Z})
  // 4) pow2

  //fmt.Println(diff)

  for i, e := range diff {
    diff[i].Z = math.Pow(e.Z,2)
  }

  var filtered []bool
  // 5) filter < 40

  for _, e := range(diff) {
    filtered = append(filtered, e.Z > 40)
  }

  for i, _ := range diff {
    if filtered[i] {
      fmt.Println(diff[i])
      c := utils.GetConn()

      //    fmt.Println("WTF")
      //fmt.Println("WTF1")
      c.Write(diff[i].Ir)
      //fmt.Println("WTF2")
    }
  }

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

    go SaveMe(items)

    fmt.Fprintf(w, `
      processed %d items
    `, len(items))
    
  }

}


func MeasureHandlerText(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    fmt.Fprintf(w, `
      <h1>Upload data</h1>
      I can understand only POST
    `)
    return

  }
  //buf := new(bytes.Buffer)
  //buf.ReadFrom(r.Body)

  model.FromCSVFile(r.Body)

  //s := buf.String()
  //fmt.Println(s)
}