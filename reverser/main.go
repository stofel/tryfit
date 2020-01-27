package main

import (
  "fmt"
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
    fmt.Fprintln(os.Stderr, "string not found")
    return
  }

  // Странно что в нет стандартной string.Reverse(str)
  rev := reverse(*str)

	fmt.Fprintln(os.Stdout, rev)
}


// string reverse
func reverse(str string) string {
  chars := []rune(str)
  for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
    chars[i], chars[j] = chars[j], chars[i]
  }
  return string(chars)
}
 
