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

type JsonRelation struct {
	Relation_cardinality float64
	Relation_name        string
	Relation_pid         uint
	Relation_rid         uint
}

type JsonJoinProblem struct {
	problem_id            uint
	Problem_neighbors     map[uint]string
	Problem_num_relations uint
	Problem_relations     []JsonRelation
	Problem_selectivities map[string]float64
}

func mapper(JJPs []JsonJoinProblem) []QueryGraph {
	res := make([]QueryGraph, len(JJPs))
	for idx, jjp := range JJPs {
		var QG QueryGraph

		// set relations
		relations := make([]uint, len(jjp.Problem_relations))
		for i := 0; i < len(jjp.Problem_relations); i++ {
			relations[i] = uint(jjp.Problem_relations[i].Relation_cardinality)
		}
		QG.R = relations

		// set selectivities
		QG.S = map[uint]float64{}
		for key, sel := range jjp.Problem_selectivities {
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
			QG.SetSel(uint(idxRel0), uint(idxRel1), sel)
		}

		// Set neighbors
		QG.N = map[uint][]uint{}
		for rel, neighborString := range jjp.Problem_neighbors {
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

		// store result
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

	var JJPs []JsonJoinProblem
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
