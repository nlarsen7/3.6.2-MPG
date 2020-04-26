package main

import "fmt"

func divide(miles float64, gallons float64)float64{
  return miles/gallons
}
func main() {
  miles:=768.0
  gallons:=16.0
  fmt.Println(divide(miles,gallons),"is your miles to the gallon")
}