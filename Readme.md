# Gorest-xls

A web application demo using ReST API that returns an Excel spreadsheet
The data on the spreadsheet can come from a database or other source,
encoded in the document and sent back to the browser for download.

This is using and excel package, multiplexer and http middleware go packages.

github.com/360EntSecGroup-Skylar/excelize
github.com/go-zoo/bone
github.com/codegangsta/negroni

 
Requirements:

Install Go Language.
Go to https://golang.org/doc/ for installation instructions.

Build:
```
$ go build main.go
```

Run:
```
$ ./main
```

Query using a browser:
```
http://localhost:3000/api/v1/excel
```
