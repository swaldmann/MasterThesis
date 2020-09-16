package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	rainbow "github.com/fatih/color"
)

/* Algorithms */

// DPccp Generate best plan using DPccp
func DPccp(QG QueryGraph, JTC JoinTreeCreator) *Tree {
	n := uint(len(QG.R))
	BestTree := make([]*Tree, 1<<n)

	for i := uint(0); i < n; i++ {
		BestTree[1<<i] = &Tree{float64(QG.R[i]), 1 << i, nil, nil, 0, nil}
	}

	// Calculate csg-cmp pairs
	subgraphs := EnumerateCsg(QG)
	csgCmpPairs := []CsgCmpPair{}
	for _, subgraph := range subgraphs {
		subgraphCsgCmpPairs := EnumerateCmp(QG, subgraph)
		csgCmpPairs = append(csgCmpPairs, subgraphCsgCmpPairs...)
	}

	//HumanPrintCsgCmpPairArray("csg-cmp-pairs", csgCmpPairs)

	// Iterate over csg-cmp-pairs and set best trees
	for _, csgCmpPair := range csgCmpPairs {
		S1 := csgCmpPair.Subgraph1
		S2 := csgCmpPair.Subgraph2
		S := S1 | S2

		p1 := BestTree[S1]
		p2 := BestTree[S2]

		CurrTree := JTC.CreateJoinTree(p1, p2, QG)
		if BestTree[S] == nil {
			BestTree[S] = CurrTree
		} else if BestTree[S].Cost > CurrTree.Cost {
			BestTree[S] = CurrTree
		}

		CurrTree = JTC.CreateJoinTree(p2, p1, QG)
		if BestTree[S] == nil {
			BestTree[S] = CurrTree
		} else if BestTree[S].Cost > CurrTree.Cost {
			BestTree[S] = CurrTree
		}
	}
	rainbow.Green(BestTree[(1<<n)-1].ToString())
	return BestTree[(1<<n)-1]
}

// EnumerateCsg Enumerate Csg pairs
func EnumerateCsg(QG QueryGraph) []uint {
	n := uint(len(QG.R))
	𝔅 := uint(1<<n - 1)

	subgraphs := []uint{}

	for i := n - 1; i < n; i-- {
		v := uint(1 << i)
		subgraphs = append(subgraphs, v)
		recursiveSubgraphs := EnumerateCsgRec(QG, v, 𝔅)
		subgraphs = append(subgraphs, recursiveSubgraphs...)
		𝔅 = SetMinus(𝔅, v, n)
	}
	return subgraphs
}

// EnumerateCsgRec Enumerate Csg-pairs
func EnumerateCsgRec(QG QueryGraph, S uint, X uint) []uint {
	n := uint(len(QG.R))
	ℕ := ℕ(QG, S)
	N := SetMinus(ℕ, X, n)

	variableState := VariableTable{}
	variableState["S"] = IdxsOfSetBits(S)
	variableState["X"] = IdxsOfSetBits(X)
	variableState["N"] = IdxsOfSetBits(N)
	visualizeEnumerateCsgRec(QG, 0, S, X, N, variableState)

	subgraphs := []uint{}

	for _, SPrime := range PowerSet(N) {
		if SPrime == 0 {
			continue
		}
		SuSPrime := S | SPrime
		subgraphs = append(subgraphs, SuSPrime)

		variableState := VariableTable{}
		variableState["emit/S"] = IdxsOfSetBits(SuSPrime)
		visualizeEnumerateCsgRec(QG, 0, S, X, N, variableState)
	}
	for _, SPrime := range PowerSet(N) {
		if SPrime == 0 {
			continue
		}
		SuSPrime := S | SPrime
		XuN := X | N
		EnumerateCsgRec(QG, SuSPrime, XuN)
	}
	return subgraphs
}

// EnumerateCmp Enumerate complementary subgraphs
func EnumerateCmp(QG QueryGraph, S1 uint) []CsgCmpPair {
	minS1 := MinUintSetBitIndex(S1)
	𝔅minS1 := uint(1<<minS1) - 1

	X := 𝔅minS1 | S1
	n := uint(len(QG.R))
	ℕ := ℕ(QG, S1)
	N := SetMinus(ℕ, X, n)

	subgraphs := []CsgCmpPair{}
	for _, v := range IdxsOfSetBits(N) {
		pair := CsgCmpPair{Subgraph1: S1, Subgraph2: 1 << v}
		subgraphs = append(subgraphs, pair)
		𝔅i := uint(1<<v - 1)
		recursiveComplements := EnumerateCsgRec(QG, 1<<v, X|(𝔅i&N))
		for _, S2 := range recursiveComplements {
			pair := CsgCmpPair{Subgraph1: S1, Subgraph2: S2}
			subgraphs = append(subgraphs, pair)
		}
	}
	return subgraphs
}

/* Visualizations */

// visualizeDPccp Dynamic Programming connected pairs
func visualizeDPccp(QG QueryGraph, JTC JoinTreeCreator) []interface{} {
	visualize(DPccp, QG, JTC)
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

/* Helpers */

// ℕ Neighborhood of a subset S
func ℕ(QG QueryGraph, S uint) uint {
	indexes := IdxsOfSetBits(S)
	result := uint(0)
	for _, index := range indexes {
		for _, neighbor := range QG.N[index] {
			result = result | (1 << neighbor)
		}
	}
	return result
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

// HumanPrintCsgCmpPair Print csg-cmp-pair in human-readable format
func HumanPrintCsgCmpPair(pair CsgCmpPair) {
	HumanPrint("S1", pair.Subgraph1)
	HumanPrint("S2", pair.Subgraph2)
}

// HumanPrintCsgCmpPairArray Print csg-cmp-pair array in human-readable format
func HumanPrintCsgCmpPairArray(name string, pairs []CsgCmpPair) {
	fmt.Println(name)
	for _, pair := range pairs {
		HumanPrintCsgCmpPair(pair)
		rainbow.Yellow("------------")
	}
}
