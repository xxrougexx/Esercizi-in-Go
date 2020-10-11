package main
import (
  "fmt"
  "strings"
  "strconv"
)
//import "math"

//Inserita una equazione dice se è una funzione(fatto) e se è iniettiva,suriettiva o biettiva
func main() {
  var equaz,potenza string
  var i, j, checkFunc int
  d := []string{}

  fmt.Println("Caratteri per operazioni: +-*/^(variabili con x e y)")
  fmt.Println("Se si vuole elevare, immettere la potenza tra parentesi")
  fmt.Print("Inserisci equazione: ")
  fmt.Scan(&equaz)
  d = strings.Split(equaz, "")

  fmt.Printf("%q\n", d)
  for i = 0; i < len(d); i++ {
    p := &d[i]
    if *p == "^" {
      j = i + 2
      /*Il seguente for è in questa modalità
      perchè dovevo fare l'aggiornamento prima del check*/
      for {
        p = &d[j]
        if *p == ")" {
          break
        }
        potenza = potenza + *p
        j++
      }
    }
  }
  checkFunc, _ = strconv.Atoi(potenza)
  fmt.Println(potenza)
  if checkFunc%2  == 0 {
    fmt.Println("L'equazione non è una funzione")
  }
}
