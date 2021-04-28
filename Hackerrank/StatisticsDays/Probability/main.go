package main

import "fmt"

// Complete the birthday function below.
func birthday(s []int32, d int32, m int32) int32 {

	if len(s) > 1 {
		line := 0

		counter := 0
		for {
			sum := 0
			if line+int(m) <= len(s) {
				for i := line; i < int(m)+line; i++ {
					sum += int(s[i])

				}
				fmt.Println(sum)
				if sum == int(d) {
					counter++
				}
			} else {
				break
			}

			line++
		}

		return int32(counter)
	} else {

		if s[0] == d {
			return 1
		} else {
			return 0
		}
	}

}

func main() {

	d := 18
	m := 7

	arr := []int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}

	fmt.Println(birthday(arr, int32(d), int32(m)))
}
