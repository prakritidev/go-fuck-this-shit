package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
  fmt.Println("Welcome To Files in GOLANG!!!")
  content := "DUCK DUCK GO!!!"


  file , err := os.Create("./test.txt")

  if err != nil {
    panic(err)
  }

  length, err := io.WriteString(file, content)


  if err != nil {
    panic(err)
  }

  fmt.Println("Length is: ", length)

  defer file.Close()
  readFile("./test.txt")
}


func readFile(filename string)  {
  dataByte, err := ioutil.ReadFile(filename)
 
  if err != nil {
    panic(err)
  }

  fmt.Println("DataBype is: \n", dataByte)
}