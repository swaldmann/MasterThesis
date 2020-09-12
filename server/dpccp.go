package main

import (
	"image/color"
	"strconv"
	"strings"

	rainbow "github.com/fatih/color"
)

// visualizeDPccp Dynamic Programming connected pairs
func visualizeDPccp(QG QueryGraph) []interface{} {
	visualize(dpccp, QG)
	defer resetChanges()
	return changes
}

func visualizeEnumerateCsgRec(QG QueryGraph, i uint, S uint, X uint, N uint, emits VariableTable) {
	n := uint(len(QG.R))

	NIndexes := IdxsOfSetBits(N)
	SIndexes := IdxsOfSetBits(S)
	XIndexes := IdxsOfSetBits(X)

	nodeColors := []NodeColor{}

	// Color each node explicitly, not just changes
	for j := n - 1; int(j-1) >= -1; j-- {
		var nodeColor color.RGBA
		if contains(NIndexes, j) {
			nodeColor = greenColor
		} else if contains(SIndexes, j) {
			nodeColor = blueColor
		} else if contains(XIndexes, j) {
			nodeColor = whiteColor
		} else {
			nodeColor = grayColor
		}
		nodeConfiguration := NodeColor{NodeIndex: j, Color: nodeColor}
		nodeColors = append(nodeColors, nodeConfiguration)
	}
	changeGraphState := &GraphState{NodeColors: nodeColors}
	change := map[string]interface{}{}
	change["graphState"] = changeGraphState
	change["variables"] = emits
	changes = append(changes, change)
}

// MARK: -
// Algorithms

func dpccp(QG QueryGraph) {
	EnumerateCsg(QG)
}

// EnumerateCsg Enumerate Csg pairs
func EnumerateCsg(QG QueryGraph) [][]uint {
	n := uint(len(QG.R))
	𝔅 := uint(1<<n - 1)

	emits := [][]uint{}

	for i := n - 1; i < n; i-- {
		v := uint(1 << i)
		emits = append(emits, IdxsOfSetBits(v))
		EnumerateCsgRec(QG, v, 𝔅)
		𝔅 = SetMinus(𝔅, v, n)
		rainbow.Yellow("--------------")
	}
	return emits
}

// ℕ Neighborhood of a subset S
func ℕ(QG QueryGraph, S uint) uint {
	indexes := IdxsOfSetBits(S)
	res := uint(0)
	for _, index := range indexes {
		for _, neighbor := range QG.N[index] {
			res = res | (1 << neighbor)
		}
	}
	return res
}

// EnumerateCsgRec Enumerate Csg-pairs
func EnumerateCsgRec(QG QueryGraph, S uint, X uint) {
	rainbow.Green("EnumerateCsgRec")
	n := uint(len(QG.R))
	ℕ := ℕ(QG, S)
	N := SetMinus(ℕ, X, n)
	HumanPrint("S", S)
	HumanPrint("X", X)
	HumanPrint("N", N)
	HumanPrint("ℕ", ℕ)

	variableState := VariableTable{}
	variableState["S"] = IdxsOfSetBits(S)
	variableState["X"] = IdxsOfSetBits(X)
	variableState["N"] = IdxsOfSetBits(N)
	visualizeEnumerateCsgRec(QG, 0, S, X, N, variableState)

	HumanPrintUIntArray("S'", PowerSet(N))
	for _, SPrime := range PowerSet(N) {
		if SPrime == 0 {
			//us(emit, S, n))
			//rainbow.Red("Emit 2")
			//emit = SetMinus(emit, S, n)
			continue
		}
		SuSPrime := S | SPrime
		HumanPrint("SuSPrime1", SuSPrime)
		//emit = emit | SuSPrime

		variableState := VariableTable{}
		variableState["emit/S"] = IdxsOfSetBits(SuSPrime)
		visualizeEnumerateCsgRec(QG, 0, S, X, N, variableState)
	}
	for _, SPrime := range PowerSet(N) {
		if SPrime == 0 {
			//HumanPrint("Emit/S", SetMinus(emit, S, n))
			//rainbow.Red("Emit 3")
			//variableState := VariableTable{}
			//variableState["emit/S"] = IdxsOfSetBits(SetMinus(emit, S, n))
			//visualizeEnumerateCsgRec(QG, 0, S, X, N, variableState)
			continue
		}
		HumanPrint("SPrime", SPrime)
		HumanPrint("S", S)
		SuSPrime := S | SPrime
		HumanPrint("SuSPrime2", SuSPrime)
		XuN := X | N
		HumanPrint("XuN", XuN)
		rainbow.Red("-->")
		EnumerateCsgRec(QG, SuSPrime, XuN)
	}
}

// EnumerateCmp Enumerate Cmp-pairs
func EnumerateCmp(QG QueryGraph, S1 uint) []uint {
	X := S1 | S1
	N := SetMinus(S1, X, 4)
	emit := []uint{}
	for i := len(QG.R) - 1; i >= 0; i-- {
		v := uint(4)
		emit = append(emit, v)
		EnumerateCsgRec(QG, v, X|(S1&N))
	}
	return emit
}

// HumanPrint Prints uint variable in a human-readable format
func HumanPrint(variableName string, variable uint) {
	setBits := IdxsOfSetBits(variable)
	setBitsStringArray := make([]string, len(setBits))
	for i, value := range setBits {
		setBitsStringArray[i] = strconv.FormatUint(uint64(value), 10)
	}
	setBitsString := strings.Join(setBitsStringArray[:], ", ")
	binary := strconv.FormatUint(uint64(variable), 2)
	rainbow.Blue(variableName + ": [" + setBitsString + "] B: " + binary)
}

// HumanPrintUIntArray Print uint array in human-readable format
func HumanPrintUIntArray(variableName string, array []uint) {
	setBitsStringArray := make([]string, len(array))
	for i, value := range array {
		setBitsStringArray[i] = strconv.FormatUint(uint64(value), 2)
	}
	setBitsString := strings.Join(setBitsStringArray[:], ", ")
	rainbow.Cyan(variableName + ": [" + setBitsString + "]")
}
