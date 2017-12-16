package main

import (
  "fmt"
  "net/http"
  "github.com/360EntSecGroup-Skylar/excelize"
  "github.com/go-zoo/bone"
  "github.com/codegangsta/negroni"
  "gorest-xls/random"
  "os"
  "io/ioutil"
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

  var name = "./" + random.Generate(10)
  err := xlsx.SaveAs(name)

  //remove the filename
  defer os.Remove(name)

  if err != nil {
    fmt.Println(err)
  }

  //Read back the contents of the random named excel file
	excel, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

  w.Header().Set("Content-Disposition", "attachment; filename=test.xlsx")
	w.Header().Set("Content-Type", "application/xlsx")
	w.Write(excel)
}
