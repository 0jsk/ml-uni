package user

type Id int64

type User struct {
	Id
	Name     string
	IsClosed bool
}

type UserDTO struct {
	Id
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	CanAccessClosed bool   `json:"can_access_closed"`
	IsClosed        bool   `json:"is_closed"`
}
