package main

import "fmt"

func main() {
	QGs := GetQueryGraphs([]string{"clique"}, []uint{5})
	for _, QG := range QGs {
		csgs := EnumerateCsg(QG)
		fmt.Println(csgs)
	}
	startServer()
}
