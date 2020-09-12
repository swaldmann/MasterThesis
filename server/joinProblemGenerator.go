package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// GenerateTreeQueryGraph Generates a tree-shaped query graph
func GenerateTreeQueryGraph(degree uint, size uint) {

	neighborEntry := func(i uint, size uint) string {

		i64 := uint64(i + 1)
		logI := log2_64(i64)
		level := uint(logI)
		numberOfPredecessorNodes := uint(1<<(level+1) - 1)
		columnIndex := i - 1<<(level) + 1
		fmt.Println("====")
		fmt.Println(size)
		fmt.Println("---")
		fmt.Println(i)
		fmt.Println(level)
		fmt.Println(numberOfPredecessorNodes)
		fmt.Println(columnIndex)

		numberOfNeighborsOnLevel := min(1<<(level+1), size-numberOfPredecessorNodes)
		fmt.Println(numberOfNeighborsOnLevel)

		numberOfNeighbors := min(degree, numberOfNeighborsOnLevel-columnIndex*2)
		//numberOfNeighbors := min(degree, uint(log2_64(uint64(i))))
		fmt.Println(numberOfNeighbors)

		neighbors := make([]string, numberOfNeighbors)
		for j := range neighbors {
			//numberOfNeighborsInRow := size - degree *
			//if numberOfPredecessorNodes + degree *
			//column :=

			lowestNeighborIndex := i*degree + 1     //uint(1<<(uintLogI+1) - 1)
			neighborIndexOffset := uint(j) % degree //uint(log2_64(uint64(uint(i) + 1)))
			neighborIndex := lowestNeighborIndex + neighborIndexOffset
			neighbors[j] = strconv.FormatUint(uint64(neighborIndex), 10)
		}
		return strings.Join(neighbors, ",")
	}

	neighbors := func(degree uint, size uint) map[uint]string {
		dict := map[uint]string{}
		for i := uint(0); float64(i) < math.Floor(float64(size)/2); i++ {
			dict[i] = neighborEntry(i, size)
		}
		return dict
	}

	relations := func(size uint) []JSONRelation {
		array := make([]JSONRelation, size)
		for i := uint(0); i < size; i++ {
			array[i] = JSONRelation{
				RelationCardinality: rand.Float64() * 10000,
				RelationName:        "<unknown>",
				RelationPID:         0,
				RelationRID:         i,
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

	data := JSONJoinProblem{
		ProblemID:                0,
		ProblemNeighbors:         problemNeighbors,
		ProblemNumberOfRelations: size,
		ProblemRelations:         relations(size),
		ProblemSelectivities:     selectivities(degree, size),
	}

	sizeString := strconv.FormatUint(uint64(size), 10)

	file, marshallErr := json.MarshalIndent(data, "", " ")
	if marshallErr != nil {
		panic("Can't marshall query graph with shape tree and size " + sizeString)
	}
	writeErr := ioutil.WriteFile("joinproblems/tree"+"_"+".json", file, 0644)
	if writeErr != nil {
		panic("Can't write query graph with shape tree and size " + sizeString)
	}
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func ceil(a float64) uint {
	return uint(math.Ceil(a))
}
