package main

func main() {
	/*QGs := GetQueryGraphs([]string{"clique"}, []uint{5})
	for _, QG := range QGs {
		csgs := EnumerateCsg(QG)
		fmt.Println(csgs)
	}*/
	//HumanPrintUIntArray("", PowerSet(uint(16)))
	//startServer()
	GenerateTreeQueryGraph(2, 8)
}
