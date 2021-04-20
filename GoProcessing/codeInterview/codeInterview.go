package main

import "fmt"

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
	//var usersAdd []string

	maping := make(map[string]int)
	for _, value := range users {

		if findElement(value.Following, userDestiny) {
			counter++
			whoFollow = value.User
			//usersAdd = append(usersAdd, whoFollow)
			maping[whoFollow] = counter
		}
	}

	if whoFollow == userOrigin {

		fmt.Println(maping)
		fmt.Println(counter)
	} else {
		findDistance(users, userOrigin, whoFollow, counter)
	}

}

func main() {

	loadUsersFollowers()

	//fmt.Println(findDistance(loadUsersFollowers(), "userA", "userE", 0))
	//fmt.Println(findDistance(loadUsersFollowers(), "userA", "userM", 0))
	//findDistance(loadUsersFollowers(), "userA", "userM", 0)
	//findDistance(loadUsersFollowers(), "userB", "userN", 0)
	findDistance(loadUsersFollowers(), "userC", "userI", 0)
}
