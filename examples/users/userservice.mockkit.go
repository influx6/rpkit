package users

import (
	context "context"
)

// UserServiceImpl defines a concrete struct which implements the methods for the
// UserService interface. All methods will panic, so add necessary internal logic.
type UserServiceImpl struct {
	CreateFunc func(var1 context.Context, var2 NewUser) (User, error)
}

// Create implements the UserService.Create() method for UserServiceImpl.
func (impl UserServiceImpl) Create(var1 context.Context, var2 NewUser) (User, error) {
	ret1, ret2 := impl.CreateFunc(var1, var2)
	return ret1, ret2
}
