package main

func main() {
	/*QGs := GetQueryGraphs([]string{"clique"}, []uint{5})
	for _, QG := range QGs {
		csgs := EnumerateCsg(QG)
		fmt.Println(csgs)
	}*/
	//HumanPrintUIntArray("", PowerSet(uint(16)))

	/*graphType := "moerkotte"
	QG := GetQueryGraph(graphType, uint(5))
	Costfunctions := []costfunctionT{Cnlj, Chj, Csmj}
	JTC := JoinTreeCreator{false, false, Costfunctions}
	visualizeDPccp(QG, JTC)*/
	startServer()
	/*for i := uint(2); i <= 10; i++ {
		GenerateTreeQueryGraph(2, i)
	}*/
}
