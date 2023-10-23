package types

import "github.com/lib/pq"

type User struct {
	Id        int
	Username  string
	Password  string
	ChangedAt pq.NullTime
	CreatedAt pq.NullTime
	Currency  string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
