package main

import "fmt"

/**
 * @author Arturo Negreiros
 * @description Migratory birds
 */

func isHere(value map[int32]int32, number int32) bool {

	var boolRes bool = false
	for element := range value {

		if element == number {
			boolRes = true
			break
		}
	}
	return boolRes

}

func migratoryBirds(arr []int32) int32 {

	hash := make(map[int32]int32)
	bigOne := 0
	maxId := 0

	for i := 0; i < len(arr); i++ {

		if isHere(hash, arr[i]) {
			hash[arr[i]] += 1
		} else {
			hash[arr[i]] = 1
		}

	}

	for key, value := range hash {

		if value > int32(bigOne) {
			bigOne = int(value)
			maxId = int(key)

		} else if value == int32(bigOne) {
			if key < int32(maxId) {
				maxId = int(key)
			}
		}

	}

	return int32(maxId)

}

func Testing() {

	//arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	fmt.Println(migratoryBirds(arr))
}

func main() {
	Testing()

}
