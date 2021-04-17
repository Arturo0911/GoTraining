package main

import (
	"fmt"
	"time"

	us "./user"
)

type userOne struct {
	us.User
}

func main() {

	newUser := new(userOne)
	newUser.HIghUser(1, "Arturo Negreiros", time.Now(), true)
	fmt.Println(newUser)
}
