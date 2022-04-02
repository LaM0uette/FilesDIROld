package main

import (
	"Test/build"
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

	schMode := flag.String("m", "%", "Mode de recherche.")
	schFile := flag.String("f", "", "Non de fichier.")
	schExt := flag.String("e", ".*", "Extension de fichier.")
	flag.Parse()

	build.DrawStart()

	fmt.Println(*schMode, *schFile, *schExt)

}
