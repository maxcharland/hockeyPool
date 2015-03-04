package main

import "fmt"

func printMenu() int {
  fmt.Println("1- Load Salary from nhlnumbers.")
  fmt.Println("2- Load stats from nhl.")
  fmt.Println("3- Get the top 10 best valuable players.")
  fmt.Println("4- Quit.")

  var i int
  fmt.Scanf("%d", &i)
  fmt.Printf("Your choice is %d \n\n", i)

  return i
}

func main() {
	fmt.Println("Welcome to hockey pool evaluator!")

  var choice = printMenu()

  for choice != 4 {
    switch choice {
    case 1:
      fmt.Println("Loading salary")
    case 2:
      fmt.Println("Loading stats")
    case 3:
      fmt.Println("The top 10 is ...")
    }

    choice = printMenu()
  }
  fmt.Println("Bye!")
}
