package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"encoding/json"
	"io"

	"github.com/gokit/rpkit/examples/users"
	"github.com/gokit/rpkit/examples/users/mocks"
	"github.com/gokit/rpkit/examples/users/userservicerp"
)

var (
	userService = mocks.UserServiceImpl{
		PokeFunc: func() {},
		PokeAgainFunc: func() error {
			return errors.New("stop poking")
		},
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
				Name: "Users",
				Addr: "LA",
				Cid:  0.1,
				ID:   1,
			}, nil
		},
		GetUsersFunc: func(ctx context.Context) []users.User {
			return []users.User{
				{
					Name: "Users",
					Addr: "LA",
					Cid:  0.1,
					ID:   1,
				},
			}
		},
	}
)

func main() {
	header := http.Header{
		"Content-Type": []string{"application/json"},
	}

	pokeService := userservicerp.NewPokeServer(userservicerp.ServePokeMethod(userService), nil, header)
	pokeAgainService := userservicerp.NewPokeAgainServer(userservicerp.ServePokeAgainMethod(userService), nil, header)

	getService := userservicerp.NewGetServer(userservicerp.ServeGetMethod(userService, userservicerp.IntTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	}), nil, header)

	getByService := userservicerp.NewGetByServer(userservicerp.ServeGetByMethod(userService, userservicerp.IntTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	}, userservicerp.StringTypeDecoder{
		Decoder: userservicerp.JSONDecoder{},
	}), nil, header)

	getUsersService := userservicerp.NewGetUsersServer(userservicerp.ServeGetUsersMethod(userService,
		userservicerp.UserSliceTypeEncoder{
			Encoder: userservicerp.DefaultJSONEncoder,
		}), nil, header)

	getUserService := userservicerp.NewGetUserServer(userservicerp.ServeGetUserMethod(userService,
		userservicerp.UserTypeEncoder{
			Encoder: userservicerp.DefaultJSONEncoder,
		}, userservicerp.IntTargetDecoder{
			DecoderFunc: func(ctx context.Context, reader io.Reader) (int, error) {
				var res int
				err := json.NewDecoder(reader).Decode(&res)
				return res, err
			},
		}), nil, header)

	createService := userservicerp.NewCreateServer(userservicerp.ServeCreateMethod(userService, userservicerp.UserTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	}, userservicerp.NewUserTypeDecoder{
		Decoder: userservicerp.JSONTargetDecoder{},
	}), nil, header)

	createUserService := userservicerp.NewCreateUserServer(userservicerp.ServeCreateUserMethod(userService, userservicerp.UserTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	}, userservicerp.NewUserTypeDecoder{
		Decoder: userservicerp.JSONTargetDecoder{},
	}), nil, header)

	mux := http.NewServeMux()
	mux.Handle(userservicerp.PokeServiceRoutePath, pokeService)
	mux.Handle(userservicerp.PokeAgainServiceRoutePath, pokeAgainService)
	mux.Handle(userservicerp.GetServiceRoutePath, getService)
	mux.Handle(userservicerp.GetByServiceRoutePath, getByService)
	mux.Handle(userservicerp.GetUsersServiceRoutePath, getUsersService)
	mux.Handle(userservicerp.GetUserServiceRoutePath, getUserService)
	mux.Handle(userservicerp.CreateServiceRoutePath, createService)
	mux.Handle(userservicerp.CreateUserServiceRoutePath, createUserService)

	addr := "localhost:7888"

	log.Printf("HTTP Server started on: %q\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
