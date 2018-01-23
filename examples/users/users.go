package users

import "context"

type NewUser struct {
	Name string `json:"name"`
}

type UpdateUser struct {
	Name string `json:"name"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// @rp
type UserService interface {
	Raise(context.Context) (string, error)
	MakeAdmin(context.Context)
	EnableSMS(context.Context) error
	Delete(context.Context, string) error
	Update(context.Context, UpdateUser) error
	Get(context.Context, string) (User, error)
	Create(context.Context, NewUser) (User, error)
	GetAll(context.Context, string) ([]User, error)
	Random(context.Context, string, User) (User, error) // this will not be used.
}
