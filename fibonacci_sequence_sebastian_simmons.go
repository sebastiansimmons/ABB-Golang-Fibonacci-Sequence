// Author: Sebastian Simmons, hello@sebastiansimmons.com
// There were no parameters given other than to write a fibonacci sequence in GO

package main

import (
  "fmt"
  "os"
  "strconv"
)

// I am using 0 as the starting index of my sequence: f_0 = 1, f_1 = 1,...,etc
func main() {

  args_os := os.Args[1:]
  if len(args_os) < 1 {
    fmt.Println("Invalid argument. Use fibonacci n<int> to print the fibonacci sequence up to n")
  }

  n, err := strconv.Atoi(args_os[0])

  if err != nil {
    fmt.Println("Invalid argument. Use fibonacci n<int> to print the fibonacci sequence up to n")
  }

  fmt.Println(print_nth_fib(n))
}


func print_nth_fib(n int) int {
  f0 := 0
  f1 := 1

  switch n {
  case 0:
    return f0
  case 1:
    return f1
  }

  for i := 0; i < n; i++{
    temp := f1
    f1 = f0 + f1
    f0 = temp
  }
  return f1
}
