package main

import "fmt"

func init() {
	fmt.Println("First Init")
}

func init() {
	fmt.Println("Second Init")
}

func secondHeighest(s1 []int) {

	max := 0
	secondMax := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] > max {
			secondMax = max
			max = s1[i]
		} else if secondMax < s1[i] {
			secondMax = s1[i]
		}

	}
	fmt.Println(secondMax)
}

s:= []int{1, 2, 3, 4, 5, 6, 7}
func whoAmI(s []int) {
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
                s[i], s[j] = s[j], s[i]
        }
}

func main() {

	s1 := []int{1, 2, 3, 4, 5, 6, -1, -19, -2, 34, 98, 100, -10}

	secondHeighest(s1)

}
