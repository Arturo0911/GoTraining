package user

import "time"

type User struct {
	Id     int
	Name   string
	Date   time.Time
	Status bool
}

// some kind of constructor
func (this *User) HIghUser(id int, name string,
	highDate time.Time, status bool) {

	this.Id = id
	this.Name = name
	this.Date = highDate
	this.Status = status
}
