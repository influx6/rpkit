package tests

import (
	"context"
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

func TestRpkitGeneratedUserServer(t *testing.T) {
	var impl users.UserServiceImpl
	impl.CreateFunc = func(var1 context.Context, var2 users.NewUser) (users.User, error) {
		return users.User{ID: 1, Name: var2.Name}, nil
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	userEncoder := userservicerp.UserTypeEncoder{userservicerp.JSONEncoder{}}
	userDecoder := userservicerp.UserTypeDecoder{userservicerp.JSONTargetDecoder{}}
	newUserEncoder := userservicerp.NewUserTypeEncoder{userservicerp.JSONEncoder{}}
	newUserDecoder := userservicerp.NewUserTypeDecoder{userservicerp.JSONTargetDecoder{}}

	createService := userservicerp.ServeCreateMethod(impl, userEncoder, newUserDecoder)
	createServiceHandler := userservicerp.NewCreateServer(createService, nil, header)

	server := httptest.NewServer(createServiceHandler)
	defer server.Close()

	client, err := userservicerp.NewCreateMethodClient(server.URL, httpClient, newUserEncoder, userDecoder)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created CreateMethod client")
	}
	tests.Passed("Should have successfully created CreateMethod client")

	action := users.NewUser{Name: "Bob Squart"}

	ctx := userservicerp.WithCustomHeader(context.Background(), header)
	newUser, err := client.Create(ctx, action)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created new user")
	}
	tests.Passed("Should have successfully created new user")

	if newUser.Name != action.Name {
		tests.Info("Received: %+q", newUser.Name)
		tests.Info("Expected: %+q", action.Name)
		tests.Failed("Should have created user with same name")
	}
	tests.Passed("Should have created user with same name")
}

func TestRpkiResponseFailure(t *testing.T) {
	var impl users.UserServiceImpl
	impl.CreateFunc = func(var1 context.Context, var2 users.NewUser) (users.User, error) {
		return users.User{ID: 1, Name: var2.Name}, nil
	}

	header := http.Header{}
	header.Set("Accept", "data/json")

	userEncoder := userservicerp.UserTypeEncoder{userservicerp.JSONEncoder{}}
	userDecoder := userservicerp.UserTypeDecoder{userservicerp.JSONTargetDecoder{}}
	newUserEncoder := userservicerp.NewUserTypeEncoder{userservicerp.JSONEncoder{}}
	newUserDecoder := userservicerp.NewUserTypeDecoder{userservicerp.JSONTargetDecoder{}}

	createService := userservicerp.ServeCreateMethod(impl, userEncoder, newUserDecoder)
	createServiceHandler := userservicerp.NewCreateServer(createService, nil, header)

	server := httptest.NewServer(createServiceHandler)
	defer server.Close()

	client, err := userservicerp.NewCreateMethodClient(server.URL, httpClient, newUserEncoder, userDecoder)

	if err != nil {
		tests.FailedWithError(err, "Should have successfully created CreateMethod client")
	}
	tests.Passed("Should have successfully created CreateMethod client")

	header2 := http.Header{}
	header2.Set("Content-Type", "application/json")

	action := users.NewUser{Name: "Bob Squart"}
	ctx := userservicerp.WithCustomHeader(context.Background(), header2)
	_, err = client.Create(ctx, action)
	if err == nil {
		tests.Failed("Should have failed to create new user")
	}
	tests.PassedWithError(err, "Should have failed to create new user")

	jse, ok := err.(userservicerp.JSONErrorResponse)
	if !ok {
		tests.Failed("Should have received JSONErrorResponse error type")
	}
	tests.Passed("Should have received JSONErrorResponse error type")

	if jse.Code != http.StatusBadRequest {
		tests.Failed("Should have gotten http.StatusBadRequest")
	}
	tests.Passed("Should have gotten http.StatusBadRequest")
}

func TestRpkitGeneratedUserServerWithDeadline(t *testing.T) {
	var impl users.UserServiceImpl
	impl.CreateFunc = func(var1 context.Context, var2 users.NewUser) (users.User, error) {
		return users.User{ID: 1, Name: var2.Name}, nil
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	userEncoder := userservicerp.UserTypeEncoder{userservicerp.JSONEncoder{}}
	userDecoder := userservicerp.UserTypeDecoder{userservicerp.JSONTargetDecoder{}}
	newUserEncoder := userservicerp.NewUserTypeEncoder{userservicerp.JSONEncoder{}}
	newUserDecoder := userservicerp.NewUserTypeDecoder{userservicerp.JSONTargetDecoder{}}

	createService := userservicerp.ServeCreateMethod(impl, userEncoder, newUserDecoder)
	createServiceHandler := userservicerp.NewCreateServer(createService, nil, header)

	server := httptest.NewServer(createServiceHandler)
	defer server.Close()

	client, err := userservicerp.NewCreateMethodClient(server.URL, httpClient, newUserEncoder, userDecoder)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created CreateMethod client")
	}
	tests.Passed("Should have successfully created CreateMethod client")

	action := users.NewUser{Name: "Bob Squart"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*5)
	defer cancel()

	time.Sleep(time.Nanosecond * 10)

	ctx = userservicerp.WithCustomHeader(ctx, header)
	_, err = client.Create(ctx, action)
	if err == nil {
		tests.FailedWithError(err, "Should have failed to create new user due to deadline")
	}
	tests.PassedWithError(err, "Should have failed to create new user due to deadline")

}
