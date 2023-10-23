package types

type User struct {
	Username     string
	PasswordHash string
	Currency     string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
