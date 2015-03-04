package main

import "fmt"

func printMenu() {
  fmt.Printf("1- Load Salary from nhlnumbers.\n")
  fmt.Printf("2- Load stats from nhl.\n")
  fmt.Printf("3- Get the top 10 best valuable players.\n")
}


func main() {
	fmt.Printf("\n\n\nWelcome to hockey pool evaluator!\n")

  printMenu()
  var i int
  fmt.Scanf("%d", &i)

  fmt.Printf("Your choice is %d \n\n", i)
}
