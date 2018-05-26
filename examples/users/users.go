package users

import "context"

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//@rp(js:server, js:client)
//@implement
type UserService interface {
	Get(context.Context) (int, error)
	Create(context.Context, NewUser) (User, error)
	GetBy(context.Context, string) (int, error)
	CreateUser(NewUser) (User, error)
	GetUser(int) (User, error)
	GetUsers(context.Context) User
}
