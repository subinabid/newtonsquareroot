package main

import "fmt"
import "math"
import  "github.com/subinabid/newtonsquareroot/functions"

func main()  {
  fmt.Println ("Intiated")
  fmt.Println ("Enter a number")
  var number float64
  fmt.Scan(&number)
  result, trials := functions.Newtonsqrt(number)
  fmt.Println ("Square root computed by Newton method is:",result, "in", trials, "iterations" )
  fmt.Println ("Square root computed by math function is:", math.Sqrt(number))
}
