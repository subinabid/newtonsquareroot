package functions

import "fmt"

func Newtonsqrt(n float64) (float64, int) {
  var itrprev float64 = n/2
  var residue float64 = 1
  var itr float64 = 0
  var i int = 0
  if n<0 {
    fmt.Println ("we are not dealing with imaginary numbers yet!")
  } else {
      for residue > .001 {
        i += 1
        itr = (itrprev + n/itrprev)/2
        residueraw := itr - itrprev
        residue = residueraw
        if residueraw < 0 {
          residue = residueraw *(-1)
        }
        itrprev = itr
        }
  }
  return itr, i
}
