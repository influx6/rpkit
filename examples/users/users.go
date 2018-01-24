package users

import "context"

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//@rp
//@implement
type UserService interface {
	Create(context.Context, NewUser) (User, error)
}
