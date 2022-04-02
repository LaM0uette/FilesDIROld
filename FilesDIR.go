package main

import (
	"flag"
	"fmt"
)

// Data : Struct of data for each file searched
// type Data struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"Nom"`
// 	Date string `json:"Date"`
// 	Path string `json:"Lien"`
// }

func main() {

	var DataSearched = flag.String("mode", "===", "Mode de recherche.")
	fmt.Println(DataSearched)

}
