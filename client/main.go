package main

import (
  "fmt"
  "github.com/valyala/fasthttp"
  "flag"
  "os"
)


func main() {
  str := flag.String("str", "", "String to reverse")
  flag.Parse()
  flagset := make(map[string]bool)
  flag.Visit(func(f *flag.Flag) {flagset[f.Name]=true})
  if !flagset["str"] {
    //stderr
    fmt.Fprintln(os.Stderr, "-str text flag is mandatory ")
    return
  }


  url := "http://localhost:8080"

  req := fasthttp.AcquireRequest()
  req.SetRequestURI(url)
  req.Header.SetMethod("POST")
  req.SetBody([]byte(*str))

  resp := fasthttp.AcquireResponse()
  client := &fasthttp.Client{}
  client.Do(req, resp)

	fmt.Println("Response", resp)
}
