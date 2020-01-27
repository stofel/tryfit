package main

import (
  "fmt"
  "bytes"
  "github.com/valyala/fasthttp"
  "os/exec"
  "io/ioutil"
  "errors"
)

func main() {
  err := fasthttp.ListenAndServe(":8080", httpHandler)
	fmt.Println("Start server error:", err)
}



// Http request handler
func httpHandler(ctx *fasthttp.RequestCtx) {

  body := ctx.PostBody()
  errBody := []byte("error")
  if bytes.Equal(body, errBody) {
    ctx.Error("body is error", fasthttp.StatusBadRequest)
    return
  }
  if reversed, err := reverse(body); err != nil {
    ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
  } else {
    fmt.Fprintf(ctx, "%s", string(reversed))
  }

}


// Reverse func
func reverse(body []byte) (reversed []byte, err error) {
  cmd := exec.Command("../reverser/reverser", "-str=" + string(body))

  // все через проверочки
  stderr, err := cmd.StderrPipe();
  if err != nil {
    return
  }

  stdout, err := cmd.StdoutPipe()
  if err != nil {
    return
  }

  err = cmd.Start()
  if err != nil {
    return
  }

  errStr, err := ioutil.ReadAll(stderr)
  if err != nil {
    return
  }

  reversed, err  = ioutil.ReadAll(stdout)
  if err != nil {
    return
  }

  if !bytes.Equal([]byte(""), errStr) {
    err = errors.New(string(errStr))
    return
  }

  return
}
