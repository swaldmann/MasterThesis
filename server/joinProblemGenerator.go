package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// GenerateTreeQueryGraph Generates a tree-shaped query graph
func GenerateTreeQueryGraph(degree uint, size uint) {

	neighborEntry := func(i uint, size uint) string {
		level := uint(log2_64(uint64(i + 1)))
		numberOfPredecessorNodes := uint(1<<(level+1) - 1)
		numberOfChildrenOnLevel := min(1<<(level+1), size-numberOfPredecessorNodes)
		columnIndex := i - 1<<(level) + 1
		numberOfChildren := min(degree, numberOfChildrenOnLevel-columnIndex*degree)

		neighbors := make([]string, numberOfChildren)
		lowestNeighborIndex := i*degree + 1
		for j := range neighbors {
			neighborIndexOffset := uint(j) % degree
			neighborIndex := lowestNeighborIndex + neighborIndexOffset
			neighbors[j] = strconv.FormatUint(uint64(neighborIndex), 10)
		}
		return strings.Join(neighbors, ",")
	}

	neighbors := func(degree uint, size uint) map[uint]string {
		dict := map[uint]string{}
		for i := uint(0); float64(i) < math.Floor(float64(size)/float64(degree)); i++ {
			dict[i] = neighborEntry(i, size)
		}
		return dict
	}

	relations := func(size uint) []JSONRelation {
		array := make([]JSONRelation, size)
		for i := uint(0); i < size; i++ {
			array[i] = JSONRelation{
				Cardinality: rand.Float64() * 10000,
				Name:        "unknown",
				ProblemID:   0,
				RelationID:  i,
			}
		}
		return array
	}

	problemNeighbors := neighbors(degree, size)

	selectivities := func(degree uint, size uint) map[string]float64 {
		result := map[string]float64{}
		for key, value := range problemNeighbors {
			neighborStrings := strings.Split(value, ",")
			for _, neighborString := range neighborStrings {
				neighborIndex, err := strconv.ParseUint(neighborString, 10, 64)
				if err != nil {
					panic("Can't convert neighbor string to uint.")
				}
				if key > uint(neighborIndex) {
					continue
				}
				keyString := strconv.FormatUint(uint64(key), 10)
				neighborIndexString := strconv.FormatUint(uint64(neighborIndex), 10)
				result[keyString+","+neighborIndexString] = rand.Float64()
			}
		}
		return result
	}

	relationsResult := relations(size)
	selectivitiesResult := selectivities(degree, size)

	// Append parents to neighbors
	for i := uint(0); i < size; i++ {
		if i == 0 {
			continue // The first node has no parent
		}
		parentIndex := (i - 1) / degree
		parentString := strconv.FormatUint(uint64(parentIndex), 10)
		if currentNeighborsString, ok := problemNeighbors[i]; ok {
			joinStrings := []string{currentNeighborsString, parentString}
			problemNeighbors[i] = strings.Join(joinStrings, ",")
		} else {
			problemNeighbors[i] = parentString
		}
	}

	data := JSONJoinProblem{
		ProblemID:         0,
		Neighbors:         problemNeighbors,
		NumberOfRelations: size,
		Relations:         relationsResult,
		Selectivities:     selectivitiesResult,
	}

	sizeString := strconv.FormatUint(uint64(size), 10)

	file, marshallErr := json.MarshalIndent(data, "", " ")
	if marshallErr != nil {
		panic("Can't marshall query graph with shape tree and size " + sizeString)
	}
	writeErr := ioutil.WriteFile("joinproblems/tree"+"_"+sizeString+".json", file, 0644)
	if writeErr != nil {
		panic("Can't write query graph with shape tree and size " + sizeString)
	}
}
