package main

import "fmt"

func main() {

  fmt.Println("Enter the name & age: ")
  var name string
  var age int
  fmt.Scanf("%s %d", &name ,&age)
  fmt.Printf("Your name is %s and age is %d", name ,age)
}
