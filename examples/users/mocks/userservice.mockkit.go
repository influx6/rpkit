package mocks

import (
     "github.com/gokit/rpkit/examples/users"


    context "context"

)



// UserServiceImpl defines a lego block struct which services the methods defined
// by the UserService interface as assignable fields. These functions are called
// when the respective functions of the struct is called.
type UserServiceImpl struct{
    
    PokeFunc func() ()
    
    PokeAgainFunc func() (error)
    
    GetFunc func(var1 context.Context) (int,error)
    
    CreateFunc func(var1 context.Context,var2 users.NewUser) (users.User,error)
    
    GetByFunc func(var1 context.Context,var2 string) (int,error)
    
    CreateUserFunc func(var1 users.NewUser) (users.User,error)
    
    GetUserFunc func(var1 int) (users.User,error)
    
    GetUsersFunc func(var1 context.Context) ([]users.User)
    
}


// Poke implements the UserService.Poke method for UserService interface.
// It calls the UserServiceImpl.PokeFunc field underneath.
func (impl UserServiceImpl) Poke() { 
    impl.PokeFunc()
    
}

// PokeAgain implements the UserService.PokeAgain method for UserService interface.
// It calls the UserServiceImpl.PokeAgainFunc field underneath.
func (impl UserServiceImpl) PokeAgain() (error){ 
    ret1 := impl.PokeAgainFunc()
    return ret1
}

// Get implements the UserService.Get method for UserService interface.
// It calls the UserServiceImpl.GetFunc field underneath.
func (impl UserServiceImpl) Get(var1 context.Context) (int,error){ 
    ret1,ret2 := impl.GetFunc(var1)
    return ret1,ret2
}

// Create implements the UserService.Create method for UserService interface.
// It calls the UserServiceImpl.CreateFunc field underneath.
func (impl UserServiceImpl) Create(var1 context.Context,var2 users.NewUser) (users.User,error){ 
    ret1,ret2 := impl.CreateFunc(var1,var2)
    return ret1,ret2
}

// GetBy implements the UserService.GetBy method for UserService interface.
// It calls the UserServiceImpl.GetByFunc field underneath.
func (impl UserServiceImpl) GetBy(var1 context.Context,var2 string) (int,error){ 
    ret1,ret2 := impl.GetByFunc(var1,var2)
    return ret1,ret2
}

// CreateUser implements the UserService.CreateUser method for UserService interface.
// It calls the UserServiceImpl.CreateUserFunc field underneath.
func (impl UserServiceImpl) CreateUser(var1 users.NewUser) (users.User,error){ 
    ret1,ret2 := impl.CreateUserFunc(var1)
    return ret1,ret2
}

// GetUser implements the UserService.GetUser method for UserService interface.
// It calls the UserServiceImpl.GetUserFunc field underneath.
func (impl UserServiceImpl) GetUser(var1 int) (users.User,error){ 
    ret1,ret2 := impl.GetUserFunc(var1)
    return ret1,ret2
}

// GetUsers implements the UserService.GetUsers method for UserService interface.
// It calls the UserServiceImpl.GetUsersFunc field underneath.
func (impl UserServiceImpl) GetUsers(var1 context.Context) ([]users.User){ 
    ret1 := impl.GetUsersFunc(var1)
    return ret1
}
