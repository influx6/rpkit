package goside

import (
	"context"
	"testing"

	"net/http/httptest"

	"net/http"
	"time"

	"github.com/gokit/rpkit/examples/users"
	"github.com/gokit/rpkit/examples/users/mocks"
	"github.com/gokit/rpkit/examples/users/userservicerp"
)

var (
	httpClient  = &http.Client{Timeout: 10 * time.Second}
	userService = mocks.UserServiceImpl{
		GetByFunc: func(ctx context.Context, var2 string) (int, error) {
			return 1, nil
		},
		GetFunc: func(ctx context.Context) (int, error) {
			return 1, nil
		},
		CreateFunc: func(ctx context.Context, nu users.NewUser) (users.User, error) {
			return users.User{
				Name: nu.Name,
				Addr: "LA",
				Cid:  0.1,
				ID:   1,
			}, nil
		},
		CreateUserFunc: func(nu users.NewUser) (users.User, error) {
			return users.User{
				Name: nu.Name,
				Addr: "LA",
				Cid:  0.1,
				ID:   1,
			}, nil
		},
		GetUserFunc: func(var1 int) (users.User, error) {
			return users.User{
				Name: "User",
				Addr: "LA",
				Cid:  0.1,
				ID:   1,
			}, nil
		},
		GetUsersFunc: func(ctx context.Context) users.User {
			return users.User{
				Name: "User",
				Addr: "LA",
				Cid:  0.1,
				ID:   1,
			}
		},
	}
)

func TestUserServiceRP_GetFunc(t *testing.T) {
	service := userservicerp.ServeGetMethod(userService, userservicerp.IntTypeEncoder{
		Encoder: &userservicerp.JSONEncoder{},
	})

	server := httptest.NewServer(userservicerp.NewGetServer(service, nil, http.Header{}))

	getRCP, err := userservicerp.NewGetMethodClient(server.URL, httpClient, userservicerp.IntTypeDecoder{
		Decoder: &
	})
}
