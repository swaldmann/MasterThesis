package main

import (
	"fmt"
	"strconv"
	"strings"

	rainbow "github.com/fatih/color"
)

// DPccp Generate best plan using DPccp.
func DPccp(QG QueryGraph, JTC JoinTreeCreator) *Tree {
	n := uint(len(QG.R))
	BestTree := make([]*Tree, 1<<n)

	for i := uint(0); i < n; i++ {
		BestTree[1<<i] = &Tree{float64(QG.R[i]), 1 << i, nil, nil, 0, nil}
	}

	subgraphs := EnumerateCsg(QG)
	fmt.Println(subgraphs)
	/*csgCmpPairs := []CsgCmpPair{}
	for _, subgraph := range subgraphs {
		subgraphCsgCmpPairs := EnumerateCmp(QG, subgraph)
		csgCmpPairs = append(csgCmpPairs, subgraphCsgCmpPairs...)
	}

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
	rainbow.Green(BestTree[(1<<n)-1].ToString()) // Print best tree
	return BestTree[(1<<n)-1]*/
	return BestTree[(1)]
}

// EnumerateCsg Enumerate connected subgraphs.
func EnumerateCsg(QG QueryGraph) []uint {
	if visualizationOn {
		emitObserver := ObservedRelation{Identifier: "emit", Color: orangeColor}
		observedRelations := []ObservedRelation{emitObserver}
		currentRoutine := &VisualizationRoutine{Name: "EnumerateCsg", ObservedRelations: observedRelations}
		startVisualizeRoutine(currentRoutine)
		defer popStack()
		//defer popFromRoutineStack()
	}

	n := uint(len(QG.R))
	subgraphs := []uint{}

	for i := n - 1; i < n; i-- {
		v := uint(1 << i)

		if visualizationOn {
			variableState := VariableTable{}
			variableState["emit"] = IdxsOfSetBits(v)
			addVisualizationStep(QG, variableState)
		}

		subgraphs = append(subgraphs, v)
		𝔅 := uint(1<<(i+1) - 1)
		recursiveSubgraphs := EnumerateCsgRec(QG, v, 𝔅)
		subgraphs = append(subgraphs, recursiveSubgraphs...)
	}

	if visualizationOn {
		//resultArray := IdxsOfSetBits(subgraphs)
		//description := uintArrayToString(IdxsOfSetBits(subgraphs), 2)
		description := "Test"
		result := &VisualizationRoutineResult{Description: description}
		endVisualizationRoutine(result)
	}
	return subgraphs
}

// EnumerateCsgRec Enumerate connected subgraphs recursively.
func EnumerateCsgRec(QG QueryGraph, S uint, X uint) []uint {
	n := uint(len(QG.R))
	ℕ := ℕ(QG, S)
	N := SetMinus(ℕ, X, n)

	if visualizationOn && !(N == 0 && S != 1<<(n-1)) {
		sObserver := ObservedRelation{Identifier: "S", Color: blueColor}
		xObserver := ObservedRelation{Identifier: "X", Color: grayColor}
		nObserver := ObservedRelation{Identifier: "N", Color: greenColor}
		emitObserver := ObservedRelation{Identifier: "emit/S", Color: orangeColor}
		observedRelations := []ObservedRelation{emitObserver, sObserver, xObserver, nObserver}
		currentRoutine := &VisualizationRoutine{Name: "EnumerateCsgRec", ObservedRelations: observedRelations}
		startVisualizeRoutine(currentRoutine)
		defer popStack()
		//defer popFromRoutineStack()
	}

	if visualizationOn && !(N == 0 && S != 1<<(n-1)) {
		variableState := VariableTable{}
		variableState["S"] = IdxsOfSetBits(S)
		variableState["X"] = IdxsOfSetBits(SetMinus(X, S, n))
		variableState["N"] = IdxsOfSetBits(N)
		addVisualizationStep(QG, variableState)
	}

	subgraphs := []uint{}

	for _, SPrime := range PowerSet(N) {
		if SPrime == 0 {
			continue
		}
		SuSPrime := S | SPrime
		subgraphs = append(subgraphs, SuSPrime)

		if visualizationOn {
			variableState := VariableTable{}
			variableState["emit/S"] = IdxsOfSetBits(SuSPrime)
			addVisualizationStep(QG, variableState)
		}
	}
	for _, SPrime := range PowerSet(N) {
		if SPrime == 0 {
			continue
		}
		SuSPrime := S | SPrime
		XuN := X | N
		recursiveSubgraphs := EnumerateCsgRec(QG, SuSPrime, XuN)
		subgraphs = append(subgraphs, recursiveSubgraphs...)
	}
	return subgraphs
}

// EnumerateCmp Enumerate complementary subgraphs.
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
		//HumanPrintCsgCmpPair(pair)
		//rainbow.Yellow("---------")
		𝔅i := uint(1<<v - 1)
		recursiveComplements := EnumerateCsgRec(QG, 1<<v, X|(𝔅i&N))
		for _, S2 := range recursiveComplements {
			pair := CsgCmpPair{Subgraph1: S1, Subgraph2: S2}
			///HumanPrintCsgCmpPair(pair)
			//rainbow.Yellow("---------")
			subgraphs = append(subgraphs, pair)
		}
	}
	// Begin visualization
	/*if visualizationOn {
		sObserver := ObservedRelation{Identifier: "S", Color: blueColor}
		xObserver := ObservedRelation{Identifier: "X", Color: grayColor}
		nObserver := ObservedRelation{Identifier: "N", Color: greenColor}
		emitObserver := ObservedRelation{Identifier: "emit/S", Color: orangeColor}
		observedRelations := []ObservedRelation{sObserver, xObserver, nObserver, emitObserver}
		currentRoutine = VisualizationRoutine{Name: "EnumerateCmp " + uintArrayToString(IdxsOfSetBits(S1)), Steps: steps, ObservedRelations: observedRelations}
		routines = append(routines, currentRoutine)
		defer resetSteps()
	}*/
	// End visualization

	return subgraphs
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
	n := uint(len(QG.R))
	return SetMinus(result, S, n)
}

/* Debug Helpers */

// HumanPrint Prints uint variable in a human-readable format.
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

// HumanPrintUIntArray Print uint array in human-readable format.
func HumanPrintUIntArray(variableName string, array []uint) {
	setBitsStringArray := make([]string, len(array))
	for i, value := range array {
		setBitsStringArray[i] = strconv.FormatUint(uint64(value), 2)
	}
	setBitsString := strings.Join(setBitsStringArray[:], ", ")
	rainbow.Cyan(variableName + ": [" + setBitsString + "]")
}

// HumanPrintCsgCmpPair Print csg-cmp-pair in human-readable format.
func HumanPrintCsgCmpPair(pair CsgCmpPair) {
	HumanPrint("S1", pair.Subgraph1)
	HumanPrint("S2", pair.Subgraph2)
}

// HumanPrintCsgCmpPairArray Print csg-cmp-pair array in human-readable format.
func HumanPrintCsgCmpPairArray(name string, pairs []CsgCmpPair) {
	fmt.Printf(name + ": ")
	for _, pair := range pairs {
		HumanPrintCsgCmpPair(pair)
	}
}

func uintArrayToString(array []uint, radix int) string {
	setBitsStringArray := make([]string, len(array))
	for i, value := range array {
		setBitsStringArray[i] = strconv.FormatUint(uint64(value), radix)
	}
	setBitsString := strings.Join(setBitsStringArray[:], ", ")
	return "{" + setBitsString + "}"
}

func concatStrings(array []string) string {
	return strings.Join(array[:], ",")
}
