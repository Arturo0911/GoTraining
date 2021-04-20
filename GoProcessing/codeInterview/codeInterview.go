package main

import (
	"fmt"
	"sort"
)

/*

				CODE TEST INTERVIEW

Implement an algorithm to solve the min distance from the users and followers

{ 'user': 'userA',  'Following': ['userB', 'userD','userE', 'userG'] }
{ 'user': 'userB',  'Following': ['userC', 'userJ','userI', 'userE'] }
{ 'user': 'userC',  'Following': ['userM', 'userN','userJ', 'userI', 'userE'] }

For example: from the userA to userM there 3 users of distances ->
	userA follow userB (1), userB follow userC (2) and userC follow userM(3


	return 3


*/

type UsersFollowers struct {
	User      string
	Following []string
}

func loadUsersFollowers() []UsersFollowers {

	userA := []string{"userB", "userD", "userE", "userG"}
	userB := []string{"userC", "userJ", "userI", "userE"}
	userC := []string{"userM", "userN", "userJ", "userI", "userE"}

	var user = []UsersFollowers{
		{
			User:      "userA",
			Following: userA,
		},
		{
			User:      "userB",
			Following: userB,
		},
		{
			User:      "userC",
			Following: userC,
		},
	}

	return user

}

func findElement(slice []string, element string) bool {

	var found bool = false
	for _, value := range slice {
		if value == element {
			found = true
			break
		} else {
			continue
		}
	}

	return found

}

func findDistance(users []UsersFollowers, userOrigin string, userDestiny string, minDistance int) {

	var whoFollow string = ""
	var counter int = 0 + minDistance

	maping := make(map[int]string)
	for _, value := range users {

		if findElement(value.Following, userDestiny) {
			counter++
			whoFollow = value.User
			maping[counter] = whoFollow
		}
	}
	sortArray := make([]int, len(maping))
	if len(maping) > 1 {
		i := 0
		for k := range maping {
			sortArray[i] = k
			i++
		}

		sort.Ints(sortArray)
		counter = sortArray[0]
		if whoFollow == userOrigin {
			fmt.Println("mininmun distance: ", counter)

		} else {
			findDistance(users, userOrigin, whoFollow, sortArray[0])
		}
	} else {
		if whoFollow == userOrigin {
			fmt.Println("mininmun distance: ", counter)
		} else {
			findDistance(users, userOrigin, whoFollow, counter)
		}
	}

}

func main() {

	loadUsersFollowers()

	//fmt.Println(findDistance(loadUsersFollowers(), "userA", "userE", 0))
	//fmt.Println(findDistance(loadUsersFollowers(), "userA", "userM", 0))
	findDistance(loadUsersFollowers(), "userA", "userM", 0)
	findDistance(loadUsersFollowers(), "userB", "userN", 0)
	findDistance(loadUsersFollowers(), "userC", "userI", 0)
}
