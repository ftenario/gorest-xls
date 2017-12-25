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
  "strconv"
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
  xlsx.SetCellValue("Sheet1", "A2", "X-Axis")
  xlsx.SetCellValue("Sheet1", "B2", "Y-Axis")
  xlsx.SetCellValue("Sheet1", "C2", "Z-Axis")

  //create a dummy data
  data := [][]uint32{
    {10, 2, 12},
    {20, 4, 55},
    {30, 10, 15},
    {40, 11, 11},
    {50, 12, 17},
    {60, 20, 17},
    {70, 35, 15},
    {80, 10, 45},
    {90, 86, 55},
    {100, 44, 60},
    {110, 55, 90},
    {120, 70, 100},
  }

  //start row ehere we put the data
  row := 3

  //loop over the data
  for i := range data {
    xlsx.SetCellValue("Sheet1", "A" + strconv.Itoa(row), data[i][0])
    xlsx.SetCellValue("Sheet1", "B" + strconv.Itoa(row), data[i][1])
    xlsx.SetCellValue("Sheet1", "C" + strconv.Itoa(row), data[i][2])
    row++
  }

  //generate a temp filename
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
