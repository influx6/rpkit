package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/influx6/faux/tests"

	"github.com/gokit/rpkit/examples/users/userservicerp"
	"github.com/gokit/rpkit/internal/rp"

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

// helper type for the requestRecorder to keep track of which hooks have been
// called.
type hookCall string

const (
	sent      hookCall = "ResponseSent"
	processed hookCall = "RequestProcessed"
	received  hookCall = "RequestReceived"
	accepted  hookCall = "RequestAccepted"
	prepared  hookCall = "ResponsePrepared"
	rejected  hookCall = "RequestRejected"
	errored   hookCall = "RequestErrored"
	paniced   hookCall = "RequestPaniced"
)

type requestRecorder struct {
	sync.Mutex
	calls []hookCall
}

func (r *requestRecorder) reset() {
	r.Lock()
	r.calls = nil
	r.Unlock()
}

func (r *requestRecorder) assertHookCalls(t *testing.T, want []hookCall) {
	r.Lock()
	defer r.Unlock()

	if len(r.calls) != len(want) {
		tests.Info("Wanted: %v", want)
		tests.Info("Received: %v", r.calls)
		tests.Failed("Hook calls are wrong")
	}

	for i, haveCall := range r.calls {
		wantCall := want[i]
		if haveCall != wantCall {
			tests.Info("Wanted: %v", want)
			tests.Info("Received: %v", r.calls)
			tests.Failed("Hook calls are wrong")
		}
	}
}

func recorderHooks() (*rp.ServerHooks, *requestRecorder) {
	recs := &requestRecorder{}

	hooks := &rp.ServerHooks{
		RequestReceivedFunc: func(ctx context.Context) {
			recs.Lock()
			recs.calls = append(recs.calls, received)
			recs.Unlock()
		},
		RequestAcceptedFunc: func(ctx context.Context) {
			recs.Lock()
			recs.calls = append(recs.calls, accepted)
			recs.Unlock()
		},
		RequestProcessedFunc: func(ctx context.Context) {
			recs.Lock()
			recs.calls = append(recs.calls, processed)
			recs.Unlock()
		},
		ResponsePreparedFunc: func(ctx context.Context) {
			recs.Lock()
			recs.calls = append(recs.calls, prepared)
			recs.Unlock()
		},
		ResponseSentFunc: func(ctx context.Context) {
			recs.Lock()
			recs.calls = append(recs.calls, sent)
			recs.Unlock()
		},
		RequestRejectedFunc: func(ctx context.Context) {
			recs.Lock()
			recs.calls = append(recs.calls, rejected)
			recs.Unlock()
		},
		RequestErrorFunc: func(ctx context.Context, err error) {
			recs.Lock()
			recs.calls = append(recs.calls, errored)
			recs.Unlock()
		},
		RequestPanicFunc: func(ctx context.Context, err interface{}) {
			recs.Lock()
			recs.calls = append(recs.calls, paniced)
			recs.Unlock()
		},
	}
	return hooks, recs
}

func TestServerHooks(t *testing.T) {
	var impl users.UserServiceImpl
	impl.CreateFunc = func(var1 context.Context, var2 users.NewUser) (users.User, error) {
		return users.User{ID: 1, Name: var2.Name}, nil
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	hooks, recorder := recorderHooks()

	userEncoder := userservicerp.UserTypeEncoder{userservicerp.JSONEncoder{}}
	userDecoder := userservicerp.UserTypeDecoder{userservicerp.JSONTargetDecoder{}}
	newUserEncoder := userservicerp.NewUserTypeEncoder{userservicerp.JSONEncoder{}}
	newUserDecoder := userservicerp.NewUserTypeDecoder{userservicerp.JSONTargetDecoder{}}

	createService := userservicerp.ServeCreateMethod(impl, userEncoder, newUserDecoder)
	createServiceHandler := userservicerp.NewCreateServer(createService, hooks, header)

	server := httptest.NewServer(createServiceHandler)
	defer server.Close()

	client, err := userservicerp.NewCreateMethodClient(server.URL, httpClient, newUserEncoder, userDecoder)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created CreateMethod client")
	}
	tests.Passed("Should have successfully created CreateMethod client")

	action := users.NewUser{Name: "Bob Squart"}
	ctx := userservicerp.WithCustomHeader(context.Background(), header)
	client.Create(ctx, action)

	recorder.assertHookCalls(t, []hookCall{received, accepted, processed, prepared, sent})
	recorder.reset()
}

func TestServerHooksWithFailure(t *testing.T) {
	var impl users.UserServiceImpl
	impl.CreateFunc = func(var1 context.Context, var2 users.NewUser) (users.User, error) {
		return users.User{ID: 1, Name: var2.Name}, nil
	}

	header := http.Header{}
	header.Set("Accept", "application/json")

	hooks, recorder := recorderHooks()

	userEncoder := userservicerp.UserTypeEncoder{userservicerp.JSONEncoder{}}
	userDecoder := userservicerp.UserTypeDecoder{userservicerp.JSONTargetDecoder{}}
	newUserEncoder := userservicerp.NewUserTypeEncoder{userservicerp.JSONEncoder{}}
	newUserDecoder := userservicerp.NewUserTypeDecoder{userservicerp.JSONTargetDecoder{}}

	createService := userservicerp.ServeCreateMethod(impl, userEncoder, newUserDecoder)
	createServiceHandler := userservicerp.NewCreateServer(createService, hooks, header)

	server := httptest.NewServer(createServiceHandler)
	defer server.Close()

	client, err := userservicerp.NewCreateMethodClient(server.URL, httpClient, newUserEncoder, userDecoder)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created CreateMethod client")
	}
	tests.Passed("Should have successfully created CreateMethod client")

	header2 := http.Header{}
	header2.Set("Content-Type", "application/dart")
	action := users.NewUser{Name: "Bob Squart"}
	ctx := userservicerp.WithCustomHeader(context.Background(), header2)
	client.Create(ctx, action)

	recorder.assertHookCalls(t, []hookCall{received, rejected})
	recorder.reset()
}
