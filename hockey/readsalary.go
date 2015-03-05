package main

import "fmt"

type Player struct {
	Gp, G, A, Pts, Salary int
  Avg float64
	LastName, FirstName, Team string
}

func (p Player) String() string {
	return fmt.Sprintf("%+v %+v %+v)", p.LastName, p.FirstName, p.Team)
}

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
  mCharland := Player{3, 2, 4, 6, 1, 2.0, "Charland", "Maxime", "Jaune"}

  var choice = printMenu()

  for choice != 4 {
    switch choice {
    case 1:
      fmt.Println("Loading salary")
    case 2:
      fmt.Println("Loading stats")
    case 3:
      fmt.Println("The top 10 is ...")
      fmt.Println(mCharland)
    }

    choice = printMenu()
  }
  fmt.Println("Bye!")
}
