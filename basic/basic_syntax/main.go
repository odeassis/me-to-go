package main

import "fmt"

func main() {

  // Este é um comentário de uma única linha

  /*
    Este é um comentário
    de várias linhas
  */
  var x int
  fmt.Println(x)
  
  x = 10
  fmt.Println(x)

  var b string
  fmt.Println(b)

  var name string = "odeassis"
  fmt.Println(name)
  
  firstName := "francisco"
  fmt.Println(firstName)

  age := 10

  if age >= 18 {
    fmt.Println("Maior de idade")
  }

  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }


  weeakDay := "terca"

  switch weeakDay {
    case "segunda":
      fmt.Println("Dia de trabalho")
    case "sabado", "domingo":
      fmt.Println("Fim de semana")
    default:
      fmt.Println("Outro dia")
  }


  mySlice := []int{1,2,3}
  mySlice2 := make([]float64,3,5)

  mySlice2 = append(mySlice2, 2.0, 2.3)

  fmt.Println("Slice: ", mySlice)
  fmt.Println("Slice 2: ", mySlice2)
  fmt.Println("Tamanho Slice 2: ", len(mySlice2))
  fmt.Println("Capacidade Slice 2: ", cap(mySlice2))

}
