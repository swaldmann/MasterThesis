package main

import (
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"reflect"
	"runtime"
	"strings"
)

// only to get name of costfunction

// used in SubsetsSizeK for error checking

// type definition: costfunctionT is another name for "func(*Tree, *Tree, QueryGraph) float64"
type costfunctionT func(*Tree, *Tree, QueryGraph) float64

// Tree Data structure representing a join tree
type Tree struct {
	Cardinality  float64
	Relations    uint
	Left         *Tree
	Right        *Tree
	Cost         float64
	Costfunction costfunctionT
}

// CsgCmpPair Type representing a csg-cmp-pair
type CsgCmpPair struct {
	Subgraph1 uint
	Subgraph2 uint
}

// ToString string of a tree
func (T *Tree) ToString() string {
	var isBaseRel bool = (T.Relations&(T.Relations-1) == 0) // is base relation if is power of two (exactly one bit set)
	if isBaseRel {
		return fmt.Sprint(bits.TrailingZeros(T.Relations))
	}
	res := "("
	if T.Left != nil {
		res += T.Left.ToString()
	}

	res += " "
	if T.Costfunction == nil {
		res += "x"
	} else {
		res += GetJoinImplNameFromCostFunc(T.Costfunction)
	}

	res += " "
	if T.Right != nil {
		res += T.Right.ToString()
	}
	res += ")"
	return res
}

func mergeTrees(T1 *Tree, T2 *Tree, QG QueryGraph, costfunc costfunctionT) *Tree {
	var parent Tree
	parent.Cardinality = T1.Cardinality * T2.Cardinality * QG.GetUnionSelTwoSets(T1.Relations, T2.Relations)
	parent.Relations = T1.Relations | T2.Relations
	parent.Left = T1
	parent.Right = T2
	parent.Cost = costfunc(T1, T2, QG)
	parent.Costfunction = costfunc
	return &parent
}

/* Representation of a query graph:
- a slice of relations named R
- a map from uint to float for the join selectivities. Named S:
  - Let idx be a uint. Set bit i and j in idx to 1. Assign a selectivity to the S[idx]
*/

// QueryGraph Representation of a query graph
type QueryGraph struct {
	R []uint           `json:"relationCardinalities"`
	S map[uint]float64 `json:"selectivities"`
	N map[uint][]uint  `json:"neighbors"` // Neighbors
}

/*func Bitvector(i uint, neighbors uint[]) uint {
	result := uint(0)
	for range _, neighbor {
		result += 1<<neig
	}
	return result
}*/

// SetSelectivity Set a selectivity for two relations
func (QG *QueryGraph) SetSelectivity(relA uint, relB uint, sel float64) {
	QG.S[1<<relA|1<<relB] = sel
}

// GetSelectivity Get the selectivity for two relations
func (QG *QueryGraph) GetSelectivity(relA uint, relB uint) float64 {
	val, ok := QG.S[1<<relA|1<<relB]
	if ok {
		return val
	}
	return 1.0
}

// GetAllSels Get all selectivities for a given relation
func (QG *QueryGraph) GetAllSels(Srels uint, rel uint) float64 {
	var sel float64 = 1.0
	for Srels != 0 {
		idx := uint(bits.TrailingZeros(Srels)) // number of trailing zeros == position of least significant bit set
		sel *= QG.GetSelectivity(idx, rel)
		Srels &= ^(1 << idx) // "unary" ^ operator flips all bits
	}

	return sel
}

// GetUnionSelTwoSets Get the union of two selectivity sets
func (QG *QueryGraph) GetUnionSelTwoSets(Srels uint, Rrels uint) float64 {
	var sel float64 = 1.0
	for key := range QG.S {
		var matchIdxS uint = key & Srels
		var matchIdxR uint = key & Rrels
		if matchIdxS != 0 && matchIdxR != 0 {
			idx := matchIdxS | matchIdxR
			s, ok := QG.S[idx]
			if ok {
				sel *= s
			}
		}
	}
	return sel
}

// Connected Checks if two relations are connected.
func (QG *QueryGraph) Connected(R1 uint, R2 uint) bool {
	for key := range QG.S {
		if ((key & R1) != 0) && ((key & R2) != 0) {
			return true
		}
	}
	return false
}

// SubsetsSizeK Gives all subsets of size k
func (QG *QueryGraph) SubsetsSizeK(k uint) []uint {
	if k < 1 {
		fmt.Println("k cannot be smaller than 1")
		return []uint{}
	}

	// believe it or not
	var n uint = uint(len(QG.R))
	var firstsub uint = (1 << k) - 1
	var lastsub uint = firstsub << (n - k)

	var res []uint
	res = append(res, firstsub)
	var sub uint = firstsub
	for sub != lastsub {
		low := uint(bits.TrailingZeros(sub)) // index of lowest bit set
		sub += (1 << low)
		high := uint(bits.TrailingZeros(sub))
		sub |= (uint(1) << (high - low - 1)) - 1
		res = append(res, sub)
	}

	// note: len(res) == BinomialCoefficient(n, k)
	// check if correct number of values is found
	var B big.Int // Package big implements arbitrary precision arithmetic (big numbers)
	(&B).Binomial(int64(n), int64(k))
	if uint64(len(res)) != B.Uint64() {
		panic("SubsetsSizeK(" + fmt.Sprint(k) + ") did not compute the correct number of subsets.")
	}

	return res

	/*
	 * alternative way of doing it:
	 * https://graphics.stanford.edu/~seander/bithacks.html
	 * scroll down all to look at #NextBitPermutation
	 */

	/*
	 * If you think this is easy, send a detailed explanation of why the code works to the TA.
	 * Your explanation will be included in the next exercise sheet.
	 */
}

// ToString Gives description of a query graph.
func (QG *QueryGraph) ToString() string {
	res := "Relations: " + fmt.Sprint(QG.R) + "\n"
	res += "Sels:" + "\n"
	for rels, sel := range QG.S {
		res += fmt.Sprintf("%0"+fmt.Sprint(len(QG.R))+"b", rels) + " : " + fmt.Sprintf("%.5f", sel) + "\n"
	}
	res += "N: "
	res += fmt.Sprint(QG.N)
	return res
}

// Subsets Returns all subsets of S, excluding ∅ and S itself.
func Subsets(S uint) []uint {
	subsets := []uint{}
	S1 := S & (-S)
	for ok := true; ok; ok = (S1 != S && S1 != 0) {
		subsets = append(subsets, S1)
		S1 = S & (S1 - S)
	}
	return subsets
}

// SetMinus Difference between two sets
func SetMinus(S1 uint, S2 uint, length uint) uint {
	/*fmt.Println("")
	fmt.Println("Set Minus")
	fmt.Printf("%08b", S1)
	fmt.Println("")
	fmt.Printf("%08b", S2)
	fmt.Println("")
	mask := uint((1 << length) - 1)
	fmt.Printf("%08b", mask)
	fmt.Println("")
	temp := S1 & S2
	fmt.Printf("%08b", temp)
	fmt.Println("")
	temp = ^temp
	fmt.Printf("%08b", temp)
	fmt.Println("")
	temp &= mask
	fmt.Printf("%08b", temp)
	fmt.Println("")
	res := S1 & temp
	fmt.Println("Result")
	fmt.Printf("%08b", res)
	fmt.Println("")*/

	// This should be equivalent:
	mask := uint((1 << length) - 1)
	res := S1 & ^S2
	//fmt.Printf("%08b", mask)
	//fmt.Println("Result")
	//fmt.Println(res)
	//fmt.Printf("%08b", res)
	//fmt.Println("")
	return res & mask
}

// is S1 subset of S2?
func isSubset(S1 uint, S2 uint) bool {
	return ((S1 & S2) == S1)
}

// IdxsOfSetBits Get the Indices of set bits in a uint (modelling a set)
func IdxsOfSetBits(S uint) []uint {
	var res []uint
	for S != 0 {
		idx := uint(bits.TrailingZeros(S))
		S &= ^(1 << idx)
		res = append(res, idx)
	}
	return res
}

// IdxsOfUnsetBits Get the Indices of unset bits in a uint (modelling a set)
func IdxsOfUnsetBits(S uint, length uint) []uint {
	mask := uint((1 << length) - 1)
	Scomplement := (^S) & mask
	return IdxsOfSetBits(Scomplement)
}

// ValuesOfSetBits Get values of bits set
func ValuesOfSetBits(S uint) []uint {
	Res := IdxsOfSetBits(S)
	for idx, idxval := range Res {
		Res[idx] = 1 << idxval
	}
	return Res
}

// ValuesOfUnsetBits Get the values of the bits not set
func ValuesOfUnsetBits(S uint, length uint) []uint {
	Res := IdxsOfUnsetBits(S, length)
	for idx, idxval := range Res {
		Res[idx] = 1 << idxval
	}
	return Res
}

// JoinTreeCreator Struct to create join trees with
type JoinTreeCreator struct {
	RightDeepOnly bool
	LeftDeepOnly  bool
	Costfunctions []costfunctionT
}

// CreateJoinTree This implementation avoids cross products if possible
func (JTC JoinTreeCreator) CreateJoinTree(T1 *Tree, T2 *Tree, QG QueryGraph) *Tree {
	if QG.Connected(T1.Relations, T2.Relations) { // are the relations of T1 and T2 connected in QG?
		mincost := math.Inf(1) // positive infinity
		var costfuncidx costfunctionT
		var left bool // join(T1, T2)  or  join(T2, T1)
		for _, impl := range JTC.Costfunctions {
			if !JTC.RightDeepOnly {
				thiscost := impl(T1, T2, QG)
				if mincost > thiscost {
					mincost = thiscost
					costfuncidx = impl
					left = true
				}
			}
			if !JTC.LeftDeepOnly {
				thiscost := impl(T2, T1, QG)
				if mincost > thiscost {
					mincost = thiscost
					costfuncidx = impl
					left = false
				}
			}
		}
		card := T1.Cardinality * T2.Cardinality * QG.GetUnionSelTwoSets(T1.Relations, T2.Relations)
		rels := T1.Relations | T2.Relations
		if left {
			return &Tree{card, rels, T1, T2, mincost, costfuncidx}
		}
		return &Tree{card, rels, T2, T1, mincost, costfuncidx}
	}
	// Cross product
	card := T1.Cardinality * T2.Cardinality
	rels := T1.Relations | T2.Relations
	cost := T1.Cardinality*T2.Cardinality + T1.Cost + T2.Cost
	return &Tree{card, rels, T1, T2, cost, Ccross}
}

// --- cost functions ---

// Cout C_out cost function
func Cout(T1 *Tree, T2 *Tree, QG QueryGraph) float64 {
	joincost := T1.Cardinality * T2.Cardinality * QG.GetUnionSelTwoSets(T1.Relations, T2.Relations)
	childrencost := T1.Cost + T2.Cost
	return joincost + childrencost
}

// Cnlj C_nlj cost function
func Cnlj(T1 *Tree, T2 *Tree, QG QueryGraph) float64 {
	joincost := T1.Cardinality * T2.Cardinality
	childrencost := T1.Cost + T2.Cost
	return joincost + childrencost
}

// Chj C_hj cost function
func Chj(T1 *Tree, T2 *Tree, QG QueryGraph) float64 {
	h := 1.2
	joincost := h * T1.Cardinality
	childrencost := T1.Cost + T2.Cost
	return joincost + childrencost
}

// Csmj C_cmj cost function
func Csmj(T1 *Tree, T2 *Tree, QG QueryGraph) float64 {
	joincost := T1.Cardinality*math.Log2(T1.Cardinality) + T2.Cardinality*math.Log2(T2.Cardinality)
	childrencost := T1.Cost + T2.Cost
	return joincost + childrencost
}

// Ccross C_cross cost function
func Ccross(T1 *Tree, T2 *Tree, QG QueryGraph) float64 {
	joincost := T1.Cardinality * T2.Cardinality
	childrencost := T1.Cost + T2.Cost
	return joincost + childrencost
}

func applyCostfuncRec(T1 *Tree, T2 *Tree, QG QueryGraph, costfunc costfunctionT) float64 {
	// ToDo
	return -1.0
}

// GetJoinImplNameFromCostFunc helper method for convenience.
func GetJoinImplNameFromCostFunc(impl costfunctionT) string {
	// ValueOf returns a new Value initialized to the concrete value stored in the interface i. ValueOf(nil) returns the zero Value.
	var val reflect.Value = reflect.ValueOf(impl)
	// For Func arguments, Pointer returns a uintptr that is the underlying code pointer.
	var ptr uintptr = val.Pointer()
	// Package runtime contains operations that interact with Go's runtime system
	// FuncForPC returns a pointer to a Func object describing the function that contains the given program counter address
	var f *runtime.Func = runtime.FuncForPC(ptr)
	// Name() returns the name of the underlying function
	// result is: "packagename.functionname"
	var funcname string = f.Name()
	// split funcname, we are not interested in package name
	// convention: costfunction names are "C" followed by join implementation name
	split := strings.Split(funcname, ".C")
	joinimplname := split[1] // split[0] contains package name
	return joinimplname
}

// PowerSet Returns all subsets
func PowerSet(S uint) []uint {
	subsets := Subsets(S)
	if len(subsets) == 1 && subsets[0] == S {
		return subsets
	}
	return append(subsets, S)
}

func contains(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GetQueryGraph Gets a query graph with a graph type and a number of relations
func GetQueryGraph(graphType string, numberOfRelations uint) QueryGraph {
	return GetQueryGraphs([]string{graphType}, []uint{numberOfRelations})[0]
}
