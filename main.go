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

func main() {

  vampire := Vampires{
    Vampire{Name: "Vlad Tepes", Country: "Romania", Died: "December 1476"},
    Vampire{Name: "Elizabeth Bathory", Country: "Hungary", Died: "August 1614"},
  }

  json.NewEncoder(os.Stdout).Encode(vampire)

}
