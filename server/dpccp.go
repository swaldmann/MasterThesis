package main

import "fmt"

// EnumerateCsg Enumerate Csg pairs
func EnumerateCsg(QG QueryGraph) []uint {
	emits := []uint{}
	n := uint(len(QG.R))
	𝔅 := uint(1<<n - 1)

	//fmt.Println("=============")
	//fmt.Println("𝔅")
	//fmt.Printf("%08b", 𝔅)
	//fmt.Println("")

	for i := n - 1; i < n; i-- {
		//fmt.Println("v")
		v := uint(1 << i)
		//fmt.Printf("%08b", v)
		emits = append(emits, v)
		//fmt.Println("")
		//fmt.Println(len(emits))
		//fmt.Println(emits)
		//fmt.Println("Yo")
		//fmt.Println(i)
		EnumerateCsgRec(QG, v, 𝔅)
		fmt.Println("_______________")
		𝔅 = SetMinus(𝔅, v, n)
	}
	fmt.Println("=======================")
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
func EnumerateCsgRec(QG QueryGraph, S uint, X uint) uint {
	n := uint(len(QG.R))
	//fmt.Println(QG.N)
	ℕ := ℕ(QG, S)
	N := SetMinus(ℕ, X, n)
	emit := uint(0)
	//fmt.Println("______")
	// fmt.Println("S, X, N")
	// fmt.Printf("%08b", S)
	// fmt.Print(", ")
	// fmt.Printf("%08b", X)
	// fmt.Print(", ")
	// fmt.Printf("%08b", N)
	// fmt.Println("")

	if N == 0 {
		//fmt.Print("Emit 1, ")
		//fmt.Printf("%08b", emit)
		//fmt.Println("")
		return emit
	}

	for _, SPrime := range Subsets(N) {
		if SPrime == 0 {
			//fmt.Print("Emit 2, ")
			//fmt.Printf("%08b", emit)
			//fmt.Println("")
			return emit
		}
		SuSPrime := S | SPrime
		//fmt.Println("SuSPrime")
		fmt.Printf("%08b", SuSPrime)
		fmt.Print(" ")
		fmt.Print(IdxsOfSetBits(SuSPrime))
		fmt.Println("")
		// fmt.Println("Length")
		// fmt.Println(len(N))
		emit = emit | SuSPrime
	}
	for _, SPrime := range Subsets(N) {
		if SPrime == 0 {
			//fmt.Print("Emit 3, ")
			//fmt.Printf("%08b", emit)
			//fmt.Println("")
			return emit
		}
		SuSPrime := S | SPrime
		XuN := X | N
		fmt.Println("-->")
		EnumerateCsgRec(QG, SuSPrime, XuN)
	}
	//fmt.Print("Emit 4, ")
	//fmt.Printf("%08b", emit)
	//fmt.Println("")
	return emit
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
