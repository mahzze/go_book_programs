package popcount

var pc [256]byte

func init() {
	for i := range pc {
		// & is a bitwise operator. AKA something I currently know nothing about
		// Due to that, I won't do anything requiring bitwise operations in chapter 2
		// If it appears again, I'll learn it and do it.
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
