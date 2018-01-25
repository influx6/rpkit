package users

import (
	context "context"
)

// UserServiceImpl defines a lego block struct which services the methods defined
// by the UserService interface as assignable fields. These functions are called when the respective
// functions of the struct is called.
type UserServiceImpl struct {
	CreateFunc func(var1 context.Context, var2 NewUser) (User, error)
}

// Create implements the UserService.Create method for UserService interface.
// It calls the UserServiceImpl.CreateFunc field underneath.
func (impl UserServiceImpl) Create(var1 context.Context, var2 NewUser) (User, error) {
	ret1, ret2 := impl.CreateFunc(var1, var2)
	return ret1, ret2
}
