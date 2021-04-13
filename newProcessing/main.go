package main

import (
	"fmt"
	"time"

	"./user"
)

/*type userOne struct {
	us.User
}*/

func main() {

	newUser := new(user.User)
	newUser.HIghUser(1, "Arturo Negreiros", time.Now(), true)
	fmt.Println(newUser)
}
