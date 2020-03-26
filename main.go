package main

import (
  "fmt"
)

// p initial
// r interest rate
// n compounding rate per time period
// t num of time periods
// func getInterest(p float64, r float64, n, t int) float64 {
//   result := p * math.Pow((1 + r/12), 12 * 1)
//   return math.Round(result * 100) / 100
// }

type World struct {
  players []*Player
}

func (w *World) addPlayer(p *Player) {
  w.players = append(w.players, p)
}

func (w *World) printCurrentState() {
  fmt.Println("Current state:")

  for _, p := range(w.players) {
    p.print()
  }
}

type Player struct {
  name string
  funds float64
  job *Job
}

func (p *Player) print() {
  fmt.Printf("Player: %+v with Job: %+v\n", *p, *p.job)
}

func (p *Player) addJob(j *Job) {
  p.job = j
}

type Job struct {
  name string
  salary float64
}

func main() {
  fmt.Println("Let the game begin!")

  w := World{}
  p := Player{name: "John", funds: 1000.00}
  w.addPlayer(&p)

  j := Job{name: "teacher", salary: 50000.00}
  p.addJob(&j)

  w.printCurrentState()
}
