package main

import (
	"fmt"
  "math/rand"
)

func main () {
    var oyun bool = true
    var secim string
    var botsecim int = rand.Intn(3)
    for oyun {
        fmt.Println("Taş : 1");
        fmt.Println("Kağıt : 2");
        fmt.Println("Makas : 3");
        fmt.Scan(&secim)
        botsecim = rand.Intn(3)
        if secim == "1" && botsecim == 1 {
         fmt.Println("berabere")
        } else if secim == "2" && botsecim == 2 {
           fmt.Println("berabere")
        } else if secim == "3" && botsecim == 3 {
            fmt.Println("berabere")
        } else if secim == "1" && botsecim == 2 {
            fmt.Println("kaybettin")
        } else if secim == "2" && botsecim == 3 {
            fmt.Println("kaybettin")
        } else if secim == "3" && botsecim == 1 { 
            fmt.Println("kaybettin")
        } else if secim == "2" && botsecim == 1 {
            fmt.Println("kazandın")
        } else if secim == "1" && botsecim == 3 {
            fmt.Println("kazandın")
        } else if secim == "3" && botsecim == 2 {
            fmt.Println("kazandın")
        }
    }
}