package main

import(
  "encoding/json"
  "os"
)

type Vampire struct {
  Name string `json:"Name"`
  Country string `json:"Country"`
  Died string `json:"Died"`
}

type Vampires []Vampire

// func returnVampires() {
//
//   vampire := Vampires{
//     Vampire{Name: "Vlad Tepes", Country: "Romania", Died: "December 1476"},
//     Vampire{Name: "Elizabeth Bathory", Country: "Hungary", Died: "August 1614"},
//   }
//
//   v := json.NewEncoder(os.Stdout).Encode(vampire)
//
//   return v
//
// }

func main() {

  // returnVampires()

  vampire := Vampires{
    Vampire{Name: "Vlad Tepes", Country: "Romania", Died: "December 1476"},
    Vampire{Name: "Elizabeth Bathory", Country: "Hungary", Died: "August 1614"},
  }

  json.NewEncoder(os.Stdout).Encode(vampire)

}
