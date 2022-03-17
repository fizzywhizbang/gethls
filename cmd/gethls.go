package main

import (
	"flag"
	"fmt"

	"github.com/fizzywhizbang/gethls"
)

func main() {
	modelFlag := flag.String("m", "none", "supply model name")
	flag.Parse()
	fmt.Println("Retrieving data for broadcaster:", *modelFlag)

	data := gethls.GetBCStatus(*modelFlag)

	fmt.Println(data.HlsSource)
}
