package main

import (
  "fmt"
  "net/http"
  "github.com/360EntSecGroup-Skylar/excelize"
  "github.com/go-zoo/bone"
  "github.com/codegangsta/negroni"
)

func main() {
  mux := bone.New()
  mux.Get("/api/v1/excel", http.HandlerFunc(ExcelHandler))

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":3000")
}

func ExcelHandler(w http.ResponseWriter, r *http.Request) {
  xlsx := excelize.NewFile()

  xlsx.SetCellValue("Sheet1", "A1", "Test Data")

  err := xlsx.SaveAs("./example.xlsx")

  if err != nil {
    fmt.Println(err)
  }

}
