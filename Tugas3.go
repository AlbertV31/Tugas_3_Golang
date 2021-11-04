package main

import "fmt"

func main()  {
  var buah = [] string {"Apel","Jeruk","Anggur","Melon"}
  var i = 0
  buah = append(buah,"Pepaya")
  fmt.Println("Jumlah Element :", len(buah))
  fmt.Println("Isi Element :", buah)
  for {
    fmt.Println("Element ke - :",i, buah[i])
    i++
    if i == len(buah){
      break
      }
  }
}
