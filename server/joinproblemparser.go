package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
Usage:
Call GetQueryGraphs(shapes, sizes)
Returns a slice of QueryGraph objects.
shapes is a slice of strings describing query graph shapes.
Currently shapes must be a subset of {chain, cycle, star, clique}
sizes must be a subset of {2,...,10}
*/

// JSONRelation Represents a relation in JSON
type JSONRelation struct {
	RelationCardinality float64 `json:"relationCardinality"`
	RelationName        string  `json:"relationName"`
	RelationPID         uint    `json:"relationPID"`
	RelationRID         uint    `json:"relationRID"`
}

// JSONJoinProblem Represents a join problem in JSON
type JSONJoinProblem struct {
	ProblemID                uint               `json:"problemID"`
	ProblemNeighbors         map[uint]string    `json:"problemNeighbors"`
	ProblemNumberOfRelations uint               `json:"problemNumberOfRelations"`
	ProblemRelations         []JSONRelation     `json:"problemRelations"`
	ProblemSelectivities     map[string]float64 `json:"problemSelectivities"`
}

func mapper(JJPs []JSONJoinProblem) []QueryGraph {
	res := make([]QueryGraph, len(JJPs))
	for idx, jjp := range JJPs {
		var QG QueryGraph

		// Set relations
		relations := make([]uint, len(jjp.ProblemRelations))
		for i := 0; i < len(jjp.ProblemRelations); i++ {
			relations[i] = uint(jjp.ProblemRelations[i].RelationCardinality)
		}
		QG.R = relations

		// Set selectivities
		QG.S = map[uint]float64{}
		for key, sel := range jjp.ProblemSelectivities {
			var idxs []string = strings.Split(key, ",")
			const base = 10
			const bitsize = 64

			idxRel0, err := strconv.ParseUint(idxs[0], base, bitsize)
			if err != nil {
				panic("strconv.ParseUint(idxs[0], base, bitsize) failed")
			}

			idxRel1, err := strconv.ParseUint(idxs[1], base, bitsize)
			if err != nil {
				panic("strconv.ParseUint(idxs[1], base, bitsize) failed")
			}
			QG.SetSelectivity(uint(idxRel0), uint(idxRel1), sel)
		}

		// Set neighbors
		QG.N = map[uint][]uint{}
		for rel, neighborString := range jjp.ProblemNeighbors {
			var idxs []string = strings.Split(neighborString, ",")
			const base = 10
			const bitsize = 64
			var resultIdxs = []uint{}

			for _, idx := range idxs {
				idxRel, err := strconv.ParseUint(idx, base, bitsize)
				if err != nil {
					panic("strconv.ParseUint(idx, base, bitsize) failed")
				}
				resultIdxs = append(resultIdxs, uint(idxRel))
			}
			QG.N[rel] = resultIdxs
		}

		// Store result
		res[idx] = QG
	}
	return res
}

func getSpecificQueryGraphs(shape string, size uint) []QueryGraph {
	filename := "joinproblems/" + shape + "_" + fmt.Sprint(size) + ".json"
	file, err := os.Open(filename)
	if err != nil {
		panic("Cannot open " + filename)
	}
	defer file.Close() // execute this command at the end of current function

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Error reading content of " + filename)
	}

	var JJPs []JSONJoinProblem
	json.Unmarshal(content, &JJPs)
	return mapper(JJPs)
}

// GetQueryGraphs Returns query graphs.
func GetQueryGraphs(shapes []string, sizes []uint) []QueryGraph {
	var res []QueryGraph
	for _, shape := range shapes {
		for _, size := range sizes {
			subres := getSpecificQueryGraphs(shape, size)
			res = append(res, subres...)
		}
	}
	return res
}

func getQueryGraphJSON(shape string, size uint) []JSONJoinProblem {
	filename := "joinproblems/" + shape + "_" + fmt.Sprint(size) + ".json"
	file, err := os.Open(filename)
	if err != nil {
		panic("Cannot open " + filename)
	}
	defer file.Close() // execute this command at the end of current function

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Error reading content of " + filename)
	}

	var JJPs []JSONJoinProblem
	json.Unmarshal(content, &JJPs)
	return JJPs
}
