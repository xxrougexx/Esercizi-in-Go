package main
import "fmt"
//import "math"
import "strings"
import "strconv"

//Inserita una equazione dice se è una funzione(fatto) e se è iniettiva,suriettiva o biettiva
func main() {
  var f string
  var i int
  d := []string{""}

  //d2[0] = "0"
  //var i int
  fmt.Println("Caratteri per operazioni: +-*/^(variabili con x e y)")
  fmt.Print("Inserisci equazione: ")
  fmt.Scan(&f)
  d = strings.Split(f, "")
  fmt.Printf("%q", d)

  i, _ = strconv.Atoi(d[2])
  if i%2  == 0{
    fmt.Println("L'equazione non è una funzione")
  }
/*  if strings.Contains(f, "x") == true {
    for i = 0;i < 10;i++ {
      if strings.Contains(f, strconv.Itoa(i)){
        f = strings.Replace(f, strconv.Itoa(i), "0", -1)
        fmt.Print(f)
        fmt.Println("Perfect")
      }
    }
  }*/
}
