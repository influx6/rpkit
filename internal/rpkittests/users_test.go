package rpkittests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gokit/rpkit/examples/users/userservicerp"

	"github.com/gokit/rpkit/examples/users"
)

func TestRpkitGeneratedUserServer(t *testing.T) {
	var impl users.UserServiceImpl
	impl.CreateFunc = func(var1 context.Context, var2 users.NewUser) (users.User, error) {
		return users.User{ID: 1, Name: var2.Name}, nil
	}

	var header http.Header
	header.Set("Content-Type", "application/json")

	encoder := userservicerp.CreateEncoderWrapper{userservicerp.JSONEncoder{}}
	decoder := userservicerp.CreateDecoderWrapper{userservicerp.CreateJSONDecoder{}}

	createService := userservicerp.ServeCreateMethod(impl, encoder, decoder)
	createServiceHandler := userservicerp.NewCreateServer(createService, nil, header)

	server := httptest.NewServer(createServiceHandler)
	defer server.Close()

}
