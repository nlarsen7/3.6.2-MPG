//Nicholas Larsen
//4/25/2020
//takes your miles and gallons and finds the MPG
package main

import "fmt"

func divide(miles float64, gallons float64)float64{
  return miles/gallons
  //your MPG
}
func main() {
  miles:=768.0
  gallons:=16.0
  //sets values for each type
  fmt.Println(divide(miles,gallons),"is your miles to the gallon")
  //prints out the MPG
}