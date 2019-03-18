/**
count of all set bits in byte. Understanding init()
Example from the Donovan and Kernighan book
*/

package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//Popcount counts set bits in given int argument
func Popcount(x uint64) int {
	var num int = 0
	for i := 0; i < 7; i++ {
		num += int(pc[byte(x>>(uint(i)*8))])
	}
	return num
}
