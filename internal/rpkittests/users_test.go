package rpkittests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/influx6/faux/tests"

	"github.com/gokit/rpkit/examples/users/userservicerp"

	"github.com/gokit/rpkit/examples/users"
)

var (
	httpClient = &http.Client{
		Timeout: 5 * time.Second,
	}
)

type NewUserDecoder struct{}

func (NewUserDecoder) Decode(r io.Reader) (users.NewUser, error) {

}

func TestRpkitGeneratedUserServer(t *testing.T) {
	var impl users.UserServiceImpl
	impl.CreateFunc = func(var1 context.Context, var2 users.NewUser) (users.User, error) {
		return users.User{ID: 1, Name: var2.Name}, nil
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	encoder := userservicerp.UserTypeEncoder{userservicerp.JSONEncoder{}}
	decoder := userservicerp.NewUserTypeDecoder{userservicerp.CreateJSONDecoder{}}
	clientEncoder := userservicerp.CreateClientEncoderWrapper{userservicerp.JSONEncoder{}}
	clientDecoder := userservicerp.CreateClientDecoderWrapper{userservicerp.JSONDecoder{}}

	createService := userservicerp.ServeCreateMethod(impl, encoder, decoder)
	createServiceHandler := userservicerp.NewCreateServer(createService, nil, header)

	server := httptest.NewServer(createServiceHandler)
	defer server.Close()

	tests.Info("Servicing CreateService at %q", server.URL)

	if server.URL == "" {
		tests.Failed("Should have received test server with a valid URL")
	}

	client, err := userservicerp.NewCreateMethodClient(server.URL, httpClient, clientEncoder, clientDecoder)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created CreateMethod client")
	}
	tests.Passed("Should have successfully created CreateMethod client")

	action := users.NewUser{Name: "Bob Squart"}
	tests.Info("Creating New User: %#v\n", action)

	ctx := userservicerp.WithCustomHeader(context.Background(), header)
	newUser, err := client.Create(ctx, action)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created new user")
	}
	tests.Passed("Should have successfully created new user")

	tests.Info("Created New User: %#v\n", newUser)

	if newUser.Name != action.Name {
		tests.Info("Received: %+q", newUser.Name)
		tests.Info("Expected: %+q", action.Name)
		tests.Failed("Should have created user with same name")
	}
	tests.Passed("Should have created user with same name")
}
