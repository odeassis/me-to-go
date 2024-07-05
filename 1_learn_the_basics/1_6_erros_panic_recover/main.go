package main

import (
	"errors"
	"fmt"
	"net/http"
)

var urls = []string {
  "http://www.google.com", //good url, 200
  "http://www.googlegoogle.com/", //bad url
  "http://www.zoogle.com", //500 example
}

func MakeRequest(url string, ch chan <- string) {
  resp, err := http.Get(url)

  if err != nil {
    fmt.Println("Error", err)
    ch <- fmt.Sprintf("err: %s", err)
    ch <- fmt.Sprintf("url: %s, status: %s", url, resp.Status)
  } else {
    ch <- fmt.Sprintf("url: %s, status: %s", url, resp.Status)
  }
}

func divide(a, b float64) (float64, error) {
  if b == 0 {
    return 0, errors.New("Cannot divide by zero")
  }

  return a / b, nil
}


func main() {
  
  result, err := divide(10, 22.32)

  if err != nil {
    fmt.Println("Error: ", err)
  } else {
    fmt.Println("Resultado: ", result)
  }

  // defer fmt.Println("Defer apos o panic")

  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recuperando do panic:", r)
    }

    fmt.Println("Defer apos o panic")
  }()
  
  fmt.Println("Antes do panic")
  panic("Um problema ocorreu!")
  fmt.Println("Nao sera executada")

  /*
  ch := make(chan string)

  for _, url := range urls {
    go MakeRequest(url, ch)
  }

  for range urls {
    fmt.Println(<-ch)
  }
  */
}
