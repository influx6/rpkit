package goside

import (
	"context"
	"errors"
	"testing"

	"net/http/httptest"

	"net/http"
	"time"

	"encoding/json"
	"io"

	"github.com/gokit/rpkit/examples/users"
	"github.com/gokit/rpkit/examples/users/mocks"
	"github.com/gokit/rpkit/examples/users/userservicerp"
	"github.com/stretchr/testify/assert"
)

var (
	httpClient  = &http.Client{Timeout: 10 * time.Second}
	userService = mocks.UserServiceImpl{
		PokeFunc: func() {

		},
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
				Name: "User",
				Addr: "LA",
				Cid:  0.1,
				ID:   1,
			}, nil
		},
		GetUsersFunc: func(ctx context.Context) []users.User {
			return []users.User{
				{
					Name: "User",
					Addr: "LA",
					Cid:  0.1,
					ID:   1,
				},
			}
		},
	}
)

func TestUserServiceRP_PokeAgain(t *testing.T) {
	service := userservicerp.ServePokeAgainMethod(userService)

	server := httptest.NewServer(userservicerp.NewPokeAgainServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewPokeAgainMethodClient(server.URL, httpClient)

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	err = getRCP.PokeAgain(context.Background())
	assert.Error(t, err)
}

func TestUserServiceRP_Poke(t *testing.T) {
	service := userservicerp.ServePokeMethod(userService)

	server := httptest.NewServer(userservicerp.NewPokeServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewPokeMethodClient(server.URL, httpClient)

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	err = getRCP.Poke(context.Background())
	assert.NoError(t, err, "Should have received no error response from server")
}

func TestUserServiceRP_Get(t *testing.T) {
	service := userservicerp.ServeGetMethod(userService, userservicerp.IntTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	})

	server := httptest.NewServer(userservicerp.NewGetServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewGetMethodClient(server.URL, httpClient, userservicerp.IntTargetDecoder{
		DecoderFunc: func(ctx context.Context, reader io.Reader) (int, error) {
			rContent, err := userservicerp.CtxResponseContentType(ctx)
			assert.NoError(t, err)
			assert.NotEmpty(t, rContent)

			var res int
			err = json.NewDecoder(reader).Decode(&res)
			return res, err
		},
	})

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	res, err := getRCP.Get(context.Background())
	assert.NoError(t, err, "Should have received no error response from server")
	assert.Equal(t, 1, res, "Should have received 1 from backend")
}

func TestUserServiceRP_GetBy(t *testing.T) {
	service := userservicerp.ServeGetByMethod(userService, userservicerp.IntTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	}, userservicerp.StringTypeDecoder{
		Decoder: userservicerp.JSONDecoder{},
	})

	server := httptest.NewServer(userservicerp.NewGetByServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewGetByMethodClient(server.URL, httpClient,
		userservicerp.StringTypeEncoder{
			Encoder: userservicerp.DefaultJSONEncoder,
		}, userservicerp.IntTargetDecoder{
			DecoderFunc: func(ctx context.Context, reader io.Reader) (int, error) {
				rContent, err := userservicerp.CtxResponseContentType(ctx)
				assert.NoError(t, err)
				assert.NotEmpty(t, rContent)

				var res int
				err = json.NewDecoder(reader).Decode(&res)
				return res, err
			},
		})

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	res, err := getRCP.GetBy(context.Background(), "User")
	assert.NoError(t, err, "Should have received no error response from server")
	assert.Equal(t, 1, res, "Should have received 1 from backend")
}

func TestUserServiceRP_GetUsers(t *testing.T) {
	service := userservicerp.ServeGetUsersMethod(userService,
		userservicerp.UserSliceTypeEncoder{
			Encoder: userservicerp.DefaultJSONEncoder,
		})

	server := httptest.NewServer(userservicerp.NewGetUsersServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewGetUsersMethodClient(
		server.URL,
		httpClient,
		userservicerp.UserSliceTypeDecoder{
			Decoder: userservicerp.DefaultJSONTargetDecoder,
		},
	)

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	list, err := getRCP.GetUsers(context.Background())
	assert.NoError(t, err, "Should have received no error response from server")

	assert.NotNil(t, list)
	assert.Len(t, list, 1)

	res := list[0]
	assert.Equal(t, "User", res.Name)
	assert.Equal(t, float64(0.1), res.Cid)
	assert.Equal(t, "LA", res.Addr)
	assert.Equal(t, 1, res.ID)
}

func TestUserServiceRP_GetUser(t *testing.T) {
	service := userservicerp.ServeGetUserMethod(userService,
		userservicerp.UserTypeEncoder{
			Encoder: userservicerp.DefaultJSONEncoder,
		}, userservicerp.IntTargetDecoder{
			DecoderFunc: func(ctx context.Context, reader io.Reader) (int, error) {
				var res int
				err := json.NewDecoder(reader).Decode(&res)
				return res, err
			},
		})

	server := httptest.NewServer(userservicerp.NewGetUserServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewGetUserMethodClient(
		server.URL,
		httpClient,
		userservicerp.IntTypeEncoder{Encoder: userservicerp.DefaultJSONEncoder},
		userservicerp.UserTypeDecoder{Decoder: userservicerp.DefaultJSONTargetDecoder},
	)

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	res, err := getRCP.GetUser(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "User", res.Name)
	assert.Equal(t, float64(0.1), res.Cid)
	assert.Equal(t, "LA", res.Addr)
	assert.Equal(t, 1, res.ID)
}

func TestUserServiceRP_Create(t *testing.T) {
	service := userservicerp.ServeCreateMethod(userService, userservicerp.UserTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	}, userservicerp.NewUserTypeDecoder{
		Decoder: userservicerp.JSONTargetDecoder{},
	})

	server := httptest.NewServer(userservicerp.NewCreateServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewCreateMethodClient(server.URL, httpClient,
		userservicerp.NewUserTypeEncoder{
			Encoder: userservicerp.DefaultJSONEncoder,
		}, userservicerp.UserTypeDecoder{
			Decoder: userservicerp.DefaultJSONTargetDecoder,
		})

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	res, err := getRCP.Create(context.Background(), users.NewUser{Name: "Simon"})
	assert.NoError(t, err)

	assert.Equal(t, "Simon", res.Name)
	assert.Equal(t, float64(0.1), res.Cid)
	assert.Equal(t, "LA", res.Addr)
	assert.Equal(t, 1, res.ID)
}

func TestUserServiceRP_CreateUser(t *testing.T) {
	service := userservicerp.ServeCreateUserMethod(userService, userservicerp.UserTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	}, userservicerp.NewUserTypeDecoder{
		Decoder: userservicerp.JSONTargetDecoder{},
	})

	server := httptest.NewServer(userservicerp.NewCreateUserServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewCreateUserMethodClient(server.URL, httpClient,
		userservicerp.NewUserTypeEncoder{
			Encoder: userservicerp.DefaultJSONEncoder,
		}, userservicerp.UserTypeDecoder{
			Decoder: userservicerp.DefaultJSONTargetDecoder,
		})

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	res, err := getRCP.CreateUser(context.Background(), users.NewUser{Name: "Simon"})
	assert.NoError(t, err)

	assert.Equal(t, "Simon", res.Name)
	assert.Equal(t, float64(0.1), res.Cid)
	assert.Equal(t, "LA", res.Addr)
	assert.Equal(t, 1, res.ID)
}

func TestUserServiceRP_CreateUserWithCreateMethodClient(t *testing.T) {
	service := userservicerp.ServeCreateUserMethod(userService, userservicerp.UserTypeEncoder{
		Encoder: userservicerp.DefaultJSONEncoder,
	}, userservicerp.NewUserTypeDecoder{
		Decoder: userservicerp.JSONTargetDecoder{},
	})

	server := httptest.NewServer(userservicerp.NewCreateUserServer(service, nil, http.Header{}))
	defer server.Close()

	getRCP, err := userservicerp.NewCreateMethodClient(server.URL, httpClient,
		userservicerp.NewUserTypeEncoder{
			Encoder: userservicerp.DefaultJSONEncoder,
		}, userservicerp.UserTypeDecoder{
			Decoder: userservicerp.DefaultJSONTargetDecoder,
		})

	assert.NoError(t, err, "Should have created GetMethod client")
	assert.NotNil(t, getRCP, "Should have created client for Get method")

	_, err = getRCP.Create(context.Background(), users.NewUser{Name: "Simon"})
	assert.Error(t, err)
}
