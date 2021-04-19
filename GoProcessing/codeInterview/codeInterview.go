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

func findDistance(users []UsersFollowers, userOrigin string, userDestiny string) {
	
	var minDistance int = 0
	var position int = 0
	var whoFollow string = ""

	for _, value := range users {
		
		if findElement(value.Following, userDestiny) {

			whoFollow =  value.User
			minDistance ++
			break
	
		}
		
		position ++

	}

	/*if whoFollow == userorigin {
		
		return minDistance
		// finish the algrotihm

	}*/

	for _,element := range users {

		if findElement(element.Following, whoFollow){
			
			if element.User == userOrigin{
				
				minDistance ++
			}

		}

		
	}






	fmt.Println(whoFollow)
	fmt.Println(minDistance)
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

func findElement(slice []string, element string) (bool){
	
	var found bool = false
	for _, value := range slice {
		if value == element{
			found = true
			break
		}else{
			continue
		}
	}

	return found

}




func main() {

	loadUsersFollowers()

	findDistance(loadUsersFollowers(), "userA", "userM")
}
