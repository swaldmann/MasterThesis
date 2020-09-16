package main

// log2_64 CPU-abstract DeBruijn-like algorithm to calculate log_2(i)
// adapted from C from https://stackoverflow.com/a/11398748
func log2_64(value uint64) int {
	if value == 0 {
		return 0
	}
	value |= value >> 1
	value |= value >> 2
	value |= value >> 4
	value |= value >> 8
	value |= value >> 16
	value |= value >> 32
	return tab64[(uint64(((value - (value >> 1)) * 0x07EDD5E59A4E28C2)))>>58]
}

var tab64 = [64]int{
	63, 0, 58, 1, 59, 47, 53, 2,
	60, 39, 48, 27, 54, 33, 42, 3,
	61, 51, 37, 40, 49, 18, 28, 20,
	55, 30, 34, 11, 43, 14, 22, 4,
	62, 57, 46, 52, 38, 26, 32, 41,
	50, 36, 17, 19, 29, 10, 13, 21,
	56, 45, 25, 31, 35, 16, 9, 12,
	44, 24, 15, 8, 23, 7, 6, 5,
}

// MinUintSlice Returns the minimum element of a uint slice
func MinUintSlice(slice []uint) uint {
	min := slice[0]
	for _, value := range slice {
		if min > value {
			min = value
		}
	}
	return min
}

// MinUintSetBitIndex Returns the minimum set index of a uint
func MinUintSetBitIndex(S uint) uint {
	SetIndexes := IdxsOfSetBits(S)
	MinIndex := MinUintSlice(SetIndexes)
	return MinIndex
}
