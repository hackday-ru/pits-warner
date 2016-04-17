package controllers

import (
  "net/http"
  "fmt"
  //"os"
  //"io"
  "model"
  "bufio"
  "utils"
)

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
    
    for _, e := range items {
      utils.GetConn().Write(e)
    }

    fmt.Fprintf(w, `
      processed %d items
    `, len(items))
    
  }
  
}