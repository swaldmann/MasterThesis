package main

// DPsize Generate best plan using DPsize.
func DPsize(QG QueryGraph, JTC JoinTreeCreator) *Tree {
	n := uint(len(QG.R))
	BestTree := make([]*Tree, 1<<n)
	PlansSizeK := make([][]uint, n+1)

	for i := uint(0); i < n; i++ {
		BestTree[1<<i] = &Tree{float64(QG.R[i]), 1 << i, nil, nil, 0, nil}
		PlansSizeK[1] = append(PlansSizeK[1], 1<<i)
	}
	for s := uint(2); s <= n; s++ { // size of plan
		for s1 := uint(1); s1 <= s/2; s1++ { // size of one subplan
			s2 := s - s1 // size of other subplan
			for _, S1 := range PlansSizeK[s1] {
				for _, S2 := range PlansSizeK[s2] {
					if (S1&S2) != 0 || !(QG.Connected(S1, S2)) {
						continue
					}
					p1 := BestTree[S1]
					p2 := BestTree[S2]
					CurrTree := JTC.CreateJoinTree(p1, p2, QG)
					S1uS2 := S1 | S2

					variableState := VariableTable{}
					variableState["S1"] = IdxsOfSetBits(S1)
					variableState["S2"] = IdxsOfSetBits(S2)
					visualizeRelations(QG, variableState, stack)

					if BestTree[S1uS2] == nil {
						PlansSizeK[s] = append(PlansSizeK[s], S1uS2)
						BestTree[S1uS2] = CurrTree
					} else if BestTree[S1uS2].Cost > CurrTree.Cost {
						BestTree[S1uS2] = CurrTree
					}
				}
			}
		}
	}
	return BestTree[(1<<n)-1]
}

/*
// innerLoop Implementation of DPsize's inner loop
func innerLoop(S1 uint, S2 uint, s uint, QG *QueryGraph, JTC *JoinTreeCreator, innerCounter *uint, csgCmpPairCounter *uint, PlansSizeK [][]uint, BestTree []*Tree) {
	*innerCounter++
	if (S1 & S2) != 0 {
		return
	}
	if !(QG.Connected(S1, S2)) {
		return
	}
	*csgCmpPairCounter++
	p1 := BestTree[S1]
	p2 := BestTree[S2]
	CurrTree := *JTC.CreateJoinTree(p1, p2, *QG)
	S1uS2 := S1 | S2
	if BestTree[S1uS2] == nil {
		PlansSizeK[s] = append(PlansSizeK[s], S1uS2)
		BestTree[S1uS2] = &CurrTree
	} else if BestTree[S1uS2].Cost > CurrTree.Cost {
		BestTree[S1uS2] = &CurrTree
	}
}

// DPsize Implementation of DPsize
func DPsize(QG QueryGraph, JTC JoinTreeCreator, counterTrace bool, optimized bool) *Tree {
	n := uint(len(QG.R))
	BestTree := make([]*Tree, 1<<n) // 1<<n == 2^n
	PlansSizeK := make([][]uint, n+1)

	for i := uint(0); i < n; i++ {
		BestTree[1<<i] = &Tree{float64(QG.R[i]), 1 << i, nil, nil, 0, nil}
		PlansSizeK[1] = append(PlansSizeK[1], 1<<i)
	}

	innerCounter := uint(0)
	csgCmpPairCounter := uint(0)
	for s := uint(2); s <= n; s++ { // size of plan
		for s1 := uint(1); s1 <= s/2; s1++ { // size of one subplan
			s2 := s - s1 // size of other subplan
			if !optimized || s1 != s2 {
				for _, S1 := range PlansSizeK[s1] {
					for _, S2 := range PlansSizeK[s2] {
						innerLoop(S1, S2, s, &QG, &JTC, &innerCounter, &csgCmpPairCounter, PlansSizeK, BestTree)
					}
				}
			} else { // s1 == s2
				for i := 1; i < len(PlansSizeK[s1]); i++ {
					for j := 0; j < i; j++ {
						S1 := PlansSizeK[s1][i]
						S2 := PlansSizeK[s1][j]
						innerLoop(S1, S2, s, &QG, &JTC, &innerCounter, &csgCmpPairCounter, PlansSizeK, BestTree)
					}
				}
			}
		}
	}
	if counterTrace {
		OnoLohmanCounter := csgCmpPairCounter / 2
		fmt.Println("innerCounter:", innerCounter)
		fmt.Println("csgCmpPairCounter:", csgCmpPairCounter)
		fmt.Println("OnoLohmanCounter:", OnoLohmanCounter)
	}
	return BestTree[(1<<n)-1]
}
*/
