package userservicerp

// All code below are auto-generated and should not be edited by handle.
// See https://github.com/gokit/rpkit for more info.
  
import (
	"context"
	"encoding/json"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gokit/rpkit/examples/users"
)

const (
	// BaseServiceName defines the base name of service root.
	BaseServiceName = "gokit.rpkit.examples.users"

	// MethodServiceName defines the complete name of this giving API service.
	MethodServiceName = "gokit.rpkit.examples.users.UserService"

	// ServiceCodePath defines the path to this generated package which contains the implemented service methods.
	ServiceCodePath = "github.com/gokit/rpkit/examples/users/userservicerp"

	// ServiceCodePathName defines the name giving to
	// this package.
	ServiceCodePathName = "userservicerp"
)

// errors ...
var (
	ErrNotInContext = errors.New("item not in context")
	ErrInvalidRequestURI = errors.New("invalid request uri")
	ErrBadRequest = errors.New("server rejected request as bad")
	ErrInvalidRequestMethod = errors.New("invalid request method")
	ErrServerInternalProblem = errors.New("server had internal problems")
	ErrServerRejectedRequest = errors.New("server does not handle request type/method")
)

//****************************************************************************
// Context Keys, Setters And Getters
//****************************************************************************

// sets of context keys used by the package.
const (
	contextCustomHeaderKey = "rp:custom:header"
	contextRequestKey = "rp:request"
	contextRequestMethodKey = "rp:request:method"
	contextRequestHeaderKey = "rp:request:header"
	contextResponseWriterKey = "rp:response:writer"
	contextServiceNameKey = "rp:service:name"
	contextServicePathKey = "rp:service:path"
	contextServicePackageKey = "rp:service:package"
	contextServicePackageNameKey = "rp:service:package:name"
	contextServiceSourcePackageKey = "rp:service:source:path"
	contextServiceSourcePackageNameKey = "rp:service:source:package"
	contextServiceMethodNameKey = "rp:service:method:name"
	contextServiceMethodPathKey = "rp:service:method:path"
	contextServiceMethodRouteKey = "rp:service:method:route"
	contextClientRequestURLKey = "rp:client:request:url"
	contextRequestTransportKey = "rp:request:transport"
)

// CtxRequestTransport retrieves if any set RequestTransport string from context.
func CtxRequestTransport(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextRequestTransportKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithRequestTransport sets the giving RequestTransport into context.
func WithRequestTransport(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextRequestTransportKey, value)
}

// CtxServiceSourcePackageName retrieves if any set ServiceSourcePackageName string from context.
func CtxServiceSourcePackageName(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServiceSourcePackageNameKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServiceSourcePackageName sets the giving ServiceSourcePackageName into context.
func WithServiceSourcePackageName(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServiceSourcePackageNameKey, value)
}

// CtxServiceSourcePackage retrieves if any set ServiceSourcePackage string from context.
func CtxServiceSourcePackage(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServiceSourcePackageKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServiceSourcePackage sets the giving ServiceSourcePackage into context.
func WithServiceSourcePackage(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServiceSourcePackageKey, value)
}

// CtxServiceMethodRoute retrieves if any set ServiceMethodRoute string from context.
func CtxServiceMethodRoute(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServiceMethodRouteKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServiceMethodRoute sets the giving ServiceMethodRoute into context.
func WithServiceMethodRoute(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServiceMethodRouteKey, value)
}

// CtxServiceMethodPath retrieves if any set ServiceMethodPath string from context.
func CtxServiceMethodPath(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServiceMethodPathKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServiceMethodPath sets the giving ServiceMethodPath into context.
func WithServiceMethodPath(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServiceMethodPathKey, value)
}

// CtxServiceMethodName retrieves if any set ServiceMethodName string from context.
func CtxServiceMethodName(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServiceMethodNameKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServiceMethodName sets the giving ServiceMethodName into context.
func WithServiceMethodName(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServiceMethodNameKey, value)
}

// CtxServicePackageName retrieves if any set ServicePackageName string from context.
func CtxServicePackageName(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServicePackageNameKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServicePackageName sets the giving ServicePackageName into context.
func WithServicePackageName(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServicePackageNameKey, value)
}

// CtxServicePackage retrieves if any set ServicePackage string from context.
func CtxServicePackage(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServicePackageKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServicePackage sets the giving ServicePackage into context.
func WithServicePackage(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServicePackageKey, value)
}

// CtxServicePath retrieves if any set ServicePath string from context.
func CtxServicePath(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServicePathKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServicePath sets the giving ServicePath into context.
func WithServicePath(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServicePathKey, value)
}

// CtxServiceName retrieves if any set ServiceName string from context.
func CtxServiceName(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextServiceNameKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithServiceName sets the giving ServiceName into context.
func WithServiceName(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextServiceNameKey, value)
}

// CtxCustomHeader retrieves if any set http.Header from context.
func CtxCustomHeader(ctx context.Context) (http.Header, error) {
	if item, ok := ctx.Value(contextCustomHeaderKey).(http.Header); ok {
		return item, nil
	}
	return nil, ErrNotInContext
}

// WithCustomHeader sets the giving http.Header into context.
func WithCustomHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, contextCustomHeaderKey, header)
}

// CtxResponseWriter retrieves if any set http.ResponseWriter from context.
func CtxResponseWriter(ctx context.Context) (http.ResponseWriter, error) {
	if item, ok := ctx.Value(contextResponseWriterKey).(http.ResponseWriter); ok {
		return item, nil
	}
	return nil, ErrNotInContext
}

// WithResponseWriter sets the giving http.ResponseWriter into context.
func WithResponseWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, contextResponseWriterKey, w)
}

// CtxRequestMethod retrieves if any set http.Request.Method string from context.
func CtxRequestMethod(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextRequestMethodKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithRequestMethod sets the giving http.RequestMethod into context.
func WithRequestMethod(ctx context.Context, method string) context.Context {
	return context.WithValue(ctx, contextRequestMethodKey, method)
}

// CtxRequestHeader retrieves if any set http.Request.Header http.Header from context.
func CtxRequestHeader(ctx context.Context) (http.Header, error) {
	if item, ok := ctx.Value(contextRequestHeaderKey).(http.Header); ok {
		return item, nil
	}
	return nil, ErrNotInContext
}

// WithRequestHeader sets the giving http.Request.Header into context.
func WithRequestHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, contextRequestHeaderKey, header)
}

// CtxRequest retrieves if any set *http.Request from context.
func CtxRequest(ctx context.Context) (*http.Request, error) {
	if item, ok := ctx.Value(contextRequestKey).(*http.Request); ok {
		return item, nil
	}
	return nil, ErrNotInContext
}

// WithRequest sets the giving http.Request into context.
func WithRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, contextRequestKey, r)
}

// CtxClientRequestURI retrieves if any set URL path from context.
func CtxClientRequestURI(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextClientRequestURLKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithClientRequestURI sets the giving http.Request into context.
func WithClientRequestURI(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, contextClientRequestURLKey, path)
}

//****************************************************************************
// Types
// Source: github.com/gokit/rpkit/examples/users
//****************************************************************************

// ActWithRequest defines a function type which is used by the client to provide
// access to the request object before it's used to make request.
type ActWithRequest func(*http.Request)

// ResponseValidation defines a function type which is used to validate an incoming
// request based on implemented logic. Use with clients to add custom validation
// steps.
type ResponseValidation func(*http.Response) error

// Hook defines a interface for having access to different areas of
// a call to the service of a method.
type Hook interface {
	ResponseSent(context.Context)
	ResponsePrepared(context.Context)
	RequestRejected(context.Context)
	RequestAccepted(context.Context)
	RequestReceived(context.Context)
	RequestProcessed(context.Context)
}

//****************************************************************************
// RP: No Arguments and No Return Methods
// Method: MakeAdmin
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.MakeAdmin
//****************************************************************************

// MakeAdminServiceRoute defines the route for the MakeAdmin method.
const MakeAdminServiceRoute = "gokit.rpkit.examples.users.UserService/MakeAdmin"

// MakeAdminServiceRoutePath defines the full method path for the MakeAdmin method.
const MakeAdminServiceRoutePath = "/gokit.rpkit.examples.users.UserService/MakeAdmin"

// makeadminServiceRoutePathURL defines a parsed path for the MakeAdmin, it
// ensures the created path is valid as a url.
var makeadminServiceRoutePathURL = mustParseURL(MakeAdminServiceRoutePath)

// MakeAdminContractSource contains the source version of expected method contract.
const MakeAdminContractSource = `type MakeAdminMethodContract interface {
	MakeAdmin(var1 context.Context) 
}
`

// MakeAdminMethodContract defines a contract interface for method "MakeAdmin"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type MakeAdminMethodContract interface{
	MakeAdmin(var1 context.Context) 
}

// MakeAdminMethodService defines the returned signature by the ServiceMakeAdmin
// which executes it's internal behaviour based off on it's MakeAdminMethodContract.
type MakeAdminMethodService func(context.Context) error

// ServeMakeAdminMethod implements the core contract behaviour to service "MakeAdmin"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users". It's job is to provide a pluggable function,
// that can then be used within any transport layer, to said execute behaviour of method as a service.
func ServeMakeAdminMethod(provider MakeAdminMethodContract) MakeAdminMethodService {
	return func(ctx context.Context) error {
		
		provider.MakeAdmin(ctx)
		
		return nil
	}
}

// MakeAdminClientContract defines a contract interface for clients to make request to MakeAdminServer.
type MakeAdminClientContract interface {
	MakeAdmin(var1 context.Context) error
}

// NewMakeAdminMethodClient returns a new MakeAdminMethodContract it relies on
// NewMakeAdminMethodContractClient.
func NewMakeAdminMethodClient(addr string, client *http.Client) (MakeAdminClientContract, error) {
	return NewMakeAdminMethodContractClient(addr, client, nil, nil)
}

// NewMakeAdminMethodContractClient returns a new MakeAdminMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a MakeAdminServer to
// perform action as specified by contract.
func NewMakeAdminMethodContractClient(addr string, client *http.Client, act ActWithRequest, resv ResponseValidation) (MakeAdminClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implMakeAdminClient{
		actor: act,
		resval: resv,
		rootURL: root,
		client: skipRedirects(client),
	}, nil
}

type implMakeAdminClient struct {
	rootURL *url.URL
	client *http.Client
	actor ActWithRequest
	resval ResponseValidation
}

// MakeAdmin makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implMakeAdminClient) MakeAdmin(var1 context.Context) error{
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(makeadminServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return err
	}

	if header, err := CtxCustomHeader(var1); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrServerRejectedRequest
	}
	

	return nil
}

// MakeAdminServer implements a http.Handler for servicing the method MakeAdmin
// from users.UserService.
type MakeAdminServer interface {
	http.Handler
}

type implMakeAdminHandler struct {
	hook Hook
	headers http.Header
	service MakeAdminMethodService
}

// NewMakeAdminServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.MakeAdmin.
func NewMakeAdminServer(service MakeAdminMethodService, hook Hook, headers http.Header) MakeAdminServer {
	return implMakeAdminHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "MakeAdmin".
func (impl implMakeAdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "MakeAdmin")
	ctx = WithServiceMethodPath(ctx, MakeAdminServiceRoute)
	ctx = WithServiceMethodRoute(ctx, MakeAdminServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "MakeAdmin",
			"api_service": MethodServiceName,
			"route": MakeAdminServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if r.Method == "HEAD" {
		for key, vals := range impl.headers {
			for _, item := range vals {
				w.Header().Add(key, item)
			}
		}

		w.Header().Add("X-Agent", "RPKIT:HTTP")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "MakeAdmin")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", MakeAdminServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != MakeAdminServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+MakeAdminServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "MakeAdmin",
			"api_service": MethodServiceName,
			"route": MakeAdminServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.RequestAccepted(ctx)
	}

	for key, vals := range impl.headers {
		for _, item := range vals {
			w.Header().Add(key, item)
		}
	}

	w.Header().Add("X-Agent", "RPKIT")
	w.Header().Add("X-Service", BaseServiceName)
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
	w.Header().Add("X-Method", "MakeAdmin")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", MakeAdminServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "MakeAdmin",
					"api_service": MethodServiceName,
					"route": MakeAdminServiceRoute,
					"api": "users.UserService",
				})
				panic(rerr)
			}
		}()

		if impl.hook != nil {
			impl.hook.RequestProcessed(ctx)
		}

		actionErr = impl.service(ctx)
	}()

	if actionErr != nil {
		jsonWriteError(w, http.StatusBadRequest, "Method returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "MakeAdmin",
			"api_service": MethodServiceName,
			"route": MakeAdminServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	w.WriteHeader(http.StatusOK)

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}


//****************************************************************************
// RP: Error Returning methods
// Method: EnableSMS
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.EnableSMS
//****************************************************************************

// EnableSMSServiceRoute defines the route for the EnableSMS method.
const EnableSMSServiceRoute = "gokit.rpkit.examples.users.UserService/EnableSMS"

// EnableSMSServiceRoutePath defines the full method path for the EnableSMS method.
const EnableSMSServiceRoutePath = "/gokit.rpkit.examples.users.UserService/EnableSMS"

// enablesmsServiceRoutePathURL defines a parsed path for the EnableSMS, it
// ensures the created path is valid as a url.
var enablesmsServiceRoutePathURL = mustParseURL(EnableSMSServiceRoutePath)

// EnableSMSContractSource contains the source version of expected method contract.
const EnableSMSContractSource = `type EnableSMSMethodContract interface {
	EnableSMS(var1 context.Context)  error  
}
`

// EnableSMSMethodContract defines a contract interface for method "EnableSMS"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type EnableSMSMethodContract interface{
	EnableSMS(var1 context.Context)  error  
}

// EnableSMSMethodService defines the returned signature by the ServiceEnableSMS
// which executes it's internal behaviour based off on it's EnableSMSMethodContract.
type EnableSMSMethodService func(context.Context) error

// ServeEnableSMSMethod implements the core contract behaviour to service "EnableSMS"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users". It's job is to provide a pluggable function,
// that can then be used within any transport layer, to said execute behaviour of method as a service.
func ServeEnableSMSMethod(provider EnableSMSMethodContract) EnableSMSMethodService {
	return func(ctx context.Context) error {
		
		return provider.EnableSMS(ctx)
		
	}
}

// EnableSMSClientContract defines a contract interface for clients to make request to EnableSMSServer.
type EnableSMSClientContract interface {
	EnableSMS(var1 context.Context) (error)
}

// NewEnableSMSMethodClient returns a new EnableSMSMethodContract it relies on
// NewEnableSMSMethodContractClient.
func NewEnableSMSMethodClient(addr string, client *http.Client) (EnableSMSClientContract, error) {
	return NewEnableSMSMethodContractClient(addr, client, nil, nil)
}

// NewEnableSMSMethodContractClient returns a new EnableSMSMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a EnableSMSServer to
// perform action as specified by contract.
func NewEnableSMSMethodContractClient(addr string, client *http.Client, act ActWithRequest, resv ResponseValidation) (EnableSMSClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implEnableSMSClient{
		actor: act,
		resval: resv,
		rootURL: root,
		client: skipRedirects(client),
	}, nil
}

type implEnableSMSClient struct {
	rootURL *url.URL
	client *http.Client
	actor ActWithRequest
	resval ResponseValidation
}

// EnableSMS makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implEnableSMSClient) EnableSMS(var1 context.Context) (error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(enablesmsServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return err
	}

	if header, err := CtxCustomHeader(var1); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrServerRejectedRequest
	}
	

	return nil
}

// EnableSMSServer implements a http.Handler for servicing the method EnableSMS
// from users.UserService.
type EnableSMSServer interface {
	http.Handler
}

type implEnableSMSHandler struct{
	hook Hook
	headers http.Header
	service EnableSMSMethodService
}

// NewEnableSMSServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.EnableSMS.
func NewEnableSMSServer(service EnableSMSMethodService, hook Hook, headers http.Header) EnableSMSServer {
	return implEnableSMSHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "EnableSMS".
func (impl implEnableSMSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "EnableSMS")
	ctx = WithServiceMethodPath(ctx, EnableSMSServiceRoute)
	ctx = WithServiceMethodRoute(ctx, EnableSMSServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "EnableSMS",
			"api_service": MethodServiceName,
			"route": EnableSMSServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if r.Method == "HEAD" {
		for key, vals := range impl.headers {
			for _, item := range vals {
				w.Header().Add(key, item)
			}
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "EnableSMS")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", EnableSMSServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != EnableSMSServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+EnableSMSServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "EnableSMS",
			"api_service": MethodServiceName,
			"route": EnableSMSServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.RequestAccepted(ctx)
	}

	for key, vals := range impl.headers {
		for _, item := range vals {
			w.Header().Add(key, item)
		}
	}

	w.Header().Add("X-Agent", "RPKIT")
	w.Header().Add("X-Service", BaseServiceName)
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
	w.Header().Add("X-Method", "EnableSMS")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", EnableSMSServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "EnableSMS",
					"api_service": MethodServiceName,
					"route": EnableSMSServiceRoute,
					"api": "users.UserService",
				})
				panic(rerr)
			}
		}()

		if impl.hook != nil {
			impl.hook.RequestProcessed(ctx)
		}

		actionErr = impl.service(ctx)
	}()

	if actionErr != nil {
		jsonWriteError(w, http.StatusBadRequest, "Method returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "EnableSMS",
			"api_service": MethodServiceName,
			"route": EnableSMSServiceRoute,
			"api": "users.UserService",
		})
		return
	}


	w.WriteHeader(http.StatusOK)

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}


//****************************************************************************
// RP: Output with Returning Error methods
// Method: Raise
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Raise
//****************************************************************************

// RaiseServiceRoute defines the route for the Raise method.
const RaiseServiceRoute = "gokit.rpkit.examples.users.UserService/Raise"

// RaiseServiceRoutePath defines the full method path for the Raise method.
const RaiseServiceRoutePath = "/gokit.rpkit.examples.users.UserService/Raise"

// raiseServiceRoutePathURL defines a parsed path for the Raise, it
// ensures the created path is valid as a url.
var raiseServiceRoutePathURL = mustParseURL(RaiseServiceRoutePath)

// RaiseContractSource contains the source version of expected method contract.
const RaiseContractSource = `type RaiseMethodContract interface {
	Raise(var1 context.Context)  (string,error)  
}
`

// RaiseMethodContract defines a contract interface for method "Raise"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type RaiseMethodContract interface{
	Raise(var1 context.Context)  (string,error)  
}

// RaiseEncoder defines a interface which expose a single method to encode the response
// returned by RaiseMethodContract.Raise.
type RaiseEncoder interface{
	Encode(io.Writer, string) error
}

// RaiseMethodService defines the returned signature by the ServiceRaise
// which executes it's internal behaviour based off on it's RaiseMethodContract.
type RaiseMethodService func(context.Context, io.Writer) error

// ServeRaiseMethod implements the core contract behaviour to service "Raise"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeRaiseMethod(provider RaiseMethodContract, encoder RaiseEncoder) RaiseMethodService {
	return func(ctx context.Context, w io.Writer) error {
		
		res, err := provider.Raise(ctx)
		
		if err != nil {
			return err
		}
		return encoder.Encode(w, res)
	}
}

// RaiseClientContract defines a contract interface for clients to make request to RaiseServer.
type RaiseClientContract interface {
	Raise(var1 context.Context) (string,error)
}

// RaiseClientDecoder defines a interface which expose a single method to decode the response
// returned by RaiseClientContract.Raise.
type RaiseClientDecoder interface{
	Decode(io.Reader) (string, error)
}

// NewRaiseMethodClient returns a new RaiseMethodContract it relies on
// NewRaiseMethodContractClient.
func NewRaiseMethodClient(addr string, client *http.Client, encoder RaiseClientDecoder) (RaiseClientContract, error) {
	return NewRaiseMethodContractClient(addr, client, encoder, nil, nil)
}

// NewRaiseMethodContractClient returns a new RaiseMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a RaiseServer to
// perform action as specified by contract.
func NewRaiseMethodContractClient(addr string, client *http.Client, decoder RaiseClientDecoder, act ActWithRequest, resv ResponseValidation) (RaiseClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implRaiseClient{
		actor: act,
		resval: resv,
		rootURL: root,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implRaiseClient struct {
	rootURL *url.URL
	client *http.Client
	actor ActWithRequest
	resval ResponseValidation
	decoder RaiseClientDecoder
}

// Raise makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implRaiseClient) Raise(var1 context.Context) (string,error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(raiseServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return "", err
	}

	if header, err := CtxCustomHeader(var1); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return "", err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return "", jsonErr
		}
		return "", ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return "", jsonErr
		}
		return "", ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return "", jsonErr
		}
		return "", ErrServerRejectedRequest
	}

	rec, err := imp.decoder.Decode(res.Body)
	if err != nil {
		return "", err
	}

	return rec, nil
}

// RaiseServer implements a http.Handler for servicing the method Raise
// from users.UserService.
type RaiseServer interface {
	http.Handler
}

type implRaiseHandler struct{
	hook Hook
	headers http.Header
	service RaiseMethodService
}

// NewRaiseServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.Raise.
func NewRaiseServer(service RaiseMethodService, hook Hook, headers http.Header) RaiseServer {
	return implRaiseHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "Raise".
func (impl implRaiseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "Raise")
	ctx = WithServiceMethodPath(ctx, RaiseServiceRoute)
	ctx = WithServiceMethodRoute(ctx, RaiseServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Raise",
			"api_service": MethodServiceName,
			"route": RaiseServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if r.Method == "HEAD" {
		for key, vals := range impl.headers {
			for _, item := range vals {
				w.Header().Add(key, item)
			}
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "Raise")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", RaiseServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != RaiseServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+RaiseServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Raise",
			"api_service": MethodServiceName,
			"route": RaiseServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.RequestAccepted(ctx)
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	for key, vals := range impl.headers {
		for _, item := range vals {
			w.Header().Add(key, item)
		}
	}

	w.Header().Add("X-Agent", "RPKIT")
	w.Header().Add("X-Service", BaseServiceName)
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
	w.Header().Add("X-Method", "Raise")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", RaiseServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "Raise",
					"api_service": MethodServiceName,
					"route": RaiseServiceRoute,
					"api": "users.UserService",
				})
				panic(rerr)
			}
		}()

		if impl.hook != nil {
			impl.hook.RequestProcessed(ctx)
		}

		actionErr = impl.service(ctx, w)
	}()

	if actionErr != nil {
		jsonWriteError(w, http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Raise",
			"api_service": MethodServiceName,
			"route": RaiseServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}



//****************************************************************************
// RP: Input Returning Only Error methods
// Method: Delete
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Delete
//****************************************************************************

// DeleteServiceRoute defines the route for the Delete method.
const DeleteServiceRoute = "gokit.rpkit.examples.users.UserService/Delete"

// DeleteServiceRoutePath defines the full method path for the Delete method.
const DeleteServiceRoutePath = "/gokit.rpkit.examples.users.UserService/Delete"

// deleteServiceRoutePathURL defines a parsed path for the Delete, it
// ensures the created path is valid as a url.
var deleteServiceRoutePathURL = mustParseURL(DeleteServiceRoutePath)

// DeleteContractSource contains the source version of expected method contract.
const DeleteContractSource = `type DeleteMethodContract interface {
	Delete(var1 context.Context,var2 string)  error  
}
`

// DeleteMethodContract defines a contract interface for method "Delete"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type DeleteMethodContract interface{
	Delete(var1 context.Context,var2 string)  error  
}

// DeleteDecoder defines a interface which expose a single method to decode the request data
// expected by DeleteMethodContract.Delete.
type DeleteDecoder interface{
	Decode(io.Reader) (string, error)
}

// DeleteMethodService defines the returned signature by the ServiceDelete
// which executes it's internal behaviour based off on it's DeleteMethodContract.
type DeleteMethodService func(context.Context, io.Reader) error

// ServeDeleteMethod implements the core contract behaviour to service "Delete"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeDeleteMethod(provider DeleteMethodContract, decoder DeleteDecoder) DeleteMethodService {
	return func(ctx context.Context, r io.Reader) error {
		input, err := decoder.Decode(r)
		if err != nil {
			return err
		}

		
		return provider.Delete(ctx, input)
		
	}
}

// DeleteClientEncoder defines a interface which expose a single method to encode the request
// data sent by DeleteClientContract.Delete.
type DeleteClientEncoder interface{
	Encode(io.Writer, string) error
}

// DeleteClientContract defines a contract interface for clients to make request to DeleteServer.
type DeleteClientContract interface {
	Delete(var1 context.Context,var2 string) (error)
}

// NewDeleteMethodClient returns a new DeleteMethodContract it relies on
// NewDeleteMethodContractClient.
func NewDeleteMethodClient(addr string, client *http.Client, encoder DeleteClientEncoder) (DeleteClientContract, error) {
	return NewDeleteMethodContractClient(addr, client, encoder, nil, nil)
}

// NewDeleteMethodContractClient returns a new DeleteMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a DeleteServer to
// perform action as specified by contract.
func NewDeleteMethodContractClient(addr string, client *http.Client, encoder DeleteClientEncoder, act ActWithRequest, resv ResponseValidation) (DeleteClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implDeleteClient{
		actor: act,
		resval: resv,
		rootURL: root,
		encoder: encoder,
		client: skipRedirects(client),
	}, nil
}

type implDeleteClient struct {
	rootURL *url.URL
	client *http.Client
	actor ActWithRequest
	resval ResponseValidation
	encoder DeleteClientEncoder
}

// Delete makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implDeleteClient) Delete(var1 context.Context,var2 string) (error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(deleteServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	var body bytes.Buffer
	if err := imp.encoder.Encode(&body, var2); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", targetURL.String(), &body)
	if err != nil {
		return err
	}

	if header, err := CtxCustomHeader(var1); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrServerRejectedRequest
	}
	

	return nil
}

// DeleteServer implements a http.Handler for servicing the method Delete
// from users.UserService.
type DeleteServer interface {
	http.Handler
}

type implDeleteHandler struct{
	hook Hook
	headers http.Header
	service DeleteMethodService
}

// NewDeleteServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.Delete.
func NewDeleteServer(service DeleteMethodService, hook Hook, headers http.Header) DeleteServer {
	return implDeleteHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "Delete".
func (impl implDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "Delete")
	ctx = WithServiceMethodPath(ctx, DeleteServiceRoute)
	ctx = WithServiceMethodRoute(ctx, DeleteServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Delete",
			"api_service": MethodServiceName,
			"route": DeleteServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if r.Method == "HEAD" {
		for key, vals := range impl.headers {
			for _, item := range vals {
				w.Header().Add(key, item)
			}
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "Delete")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", DeleteServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != DeleteServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+DeleteServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Delete",
			"api_service": MethodServiceName,
			"route": DeleteServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.RequestAccepted(ctx)
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	for key, vals := range impl.headers {
		for _, item := range vals {
			w.Header().Add(key, item)
		}
	}

	w.Header().Add("X-Agent", "RPKIT")
	w.Header().Add("X-Service", BaseServiceName)
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
	w.Header().Add("X-Method", "Delete")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", DeleteServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "Delete",
					"api_service": MethodServiceName,
					"route": DeleteServiceRoute,
					"api": "users.UserService",
				})
				panic(rerr)
			}
		}()

		if impl.hook != nil {
			impl.hook.RequestProcessed(ctx)
		}

		actionErr = impl.service(ctx, r.Body)
	}()

	if actionErr != nil {
		jsonWriteError(w, http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Delete",
			"api_service": MethodServiceName,
			"route": DeleteServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// RP: Input Returning Only Error methods
// Method: Update
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Update
//****************************************************************************

// UpdateServiceRoute defines the route for the Update method.
const UpdateServiceRoute = "gokit.rpkit.examples.users.UserService/Update"

// UpdateServiceRoutePath defines the full method path for the Update method.
const UpdateServiceRoutePath = "/gokit.rpkit.examples.users.UserService/Update"

// updateServiceRoutePathURL defines a parsed path for the Update, it
// ensures the created path is valid as a url.
var updateServiceRoutePathURL = mustParseURL(UpdateServiceRoutePath)

// UpdateContractSource contains the source version of expected method contract.
const UpdateContractSource = `type UpdateMethodContract interface {
	Update(var1 context.Context,var2 users.UpdateUser)  error  
}
`

// UpdateMethodContract defines a contract interface for method "Update"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type UpdateMethodContract interface{
	Update(var1 context.Context,var2 users.UpdateUser)  error  
}

// UpdateDecoder defines a interface which expose a single method to decode the request data
// expected by UpdateMethodContract.Update.
type UpdateDecoder interface{
	Decode(io.Reader) (users.UpdateUser, error)
}

// UpdateMethodService defines the returned signature by the ServiceUpdate
// which executes it's internal behaviour based off on it's UpdateMethodContract.
type UpdateMethodService func(context.Context, io.Reader) error

// ServeUpdateMethod implements the core contract behaviour to service "Update"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeUpdateMethod(provider UpdateMethodContract, decoder UpdateDecoder) UpdateMethodService {
	return func(ctx context.Context, r io.Reader) error {
		input, err := decoder.Decode(r)
		if err != nil {
			return err
		}

		
		return provider.Update(ctx, input)
		
	}
}

// UpdateClientEncoder defines a interface which expose a single method to encode the request
// data sent by UpdateClientContract.Update.
type UpdateClientEncoder interface{
	Encode(io.Writer, users.UpdateUser) error
}

// UpdateClientContract defines a contract interface for clients to make request to UpdateServer.
type UpdateClientContract interface {
	Update(var1 context.Context,var2 users.UpdateUser) (error)
}

// NewUpdateMethodClient returns a new UpdateMethodContract it relies on
// NewUpdateMethodContractClient.
func NewUpdateMethodClient(addr string, client *http.Client, encoder UpdateClientEncoder) (UpdateClientContract, error) {
	return NewUpdateMethodContractClient(addr, client, encoder, nil, nil)
}

// NewUpdateMethodContractClient returns a new UpdateMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a UpdateServer to
// perform action as specified by contract.
func NewUpdateMethodContractClient(addr string, client *http.Client, encoder UpdateClientEncoder, act ActWithRequest, resv ResponseValidation) (UpdateClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implUpdateClient{
		actor: act,
		resval: resv,
		rootURL: root,
		encoder: encoder,
		client: skipRedirects(client),
	}, nil
}

type implUpdateClient struct {
	rootURL *url.URL
	client *http.Client
	actor ActWithRequest
	resval ResponseValidation
	encoder UpdateClientEncoder
}

// Update makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implUpdateClient) Update(var1 context.Context,var2 users.UpdateUser) (error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(updateServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	var body bytes.Buffer
	if err := imp.encoder.Encode(&body, var2); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", targetURL.String(), &body)
	if err != nil {
		return err
	}

	if header, err := CtxCustomHeader(var1); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return jsonErr
		}
		return ErrServerRejectedRequest
	}
	

	return nil
}

// UpdateServer implements a http.Handler for servicing the method Update
// from users.UserService.
type UpdateServer interface {
	http.Handler
}

type implUpdateHandler struct{
	hook Hook
	headers http.Header
	service UpdateMethodService
}

// NewUpdateServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.Update.
func NewUpdateServer(service UpdateMethodService, hook Hook, headers http.Header) UpdateServer {
	return implUpdateHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "Update".
func (impl implUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "Update")
	ctx = WithServiceMethodPath(ctx, UpdateServiceRoute)
	ctx = WithServiceMethodRoute(ctx, UpdateServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Update",
			"api_service": MethodServiceName,
			"route": UpdateServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if r.Method == "HEAD" {
		for key, vals := range impl.headers {
			for _, item := range vals {
				w.Header().Add(key, item)
			}
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "Update")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", UpdateServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != UpdateServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+UpdateServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Update",
			"api_service": MethodServiceName,
			"route": UpdateServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.RequestAccepted(ctx)
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	for key, vals := range impl.headers {
		for _, item := range vals {
			w.Header().Add(key, item)
		}
	}

	w.Header().Add("X-Agent", "RPKIT")
	w.Header().Add("X-Service", BaseServiceName)
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
	w.Header().Add("X-Method", "Update")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", UpdateServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "Update",
					"api_service": MethodServiceName,
					"route": UpdateServiceRoute,
					"api": "users.UserService",
				})
				panic(rerr)
			}
		}()

		if impl.hook != nil {
			impl.hook.RequestProcessed(ctx)
		}

		actionErr = impl.service(ctx, r.Body)
	}()

	if actionErr != nil {
		jsonWriteError(w, http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Update",
			"api_service": MethodServiceName,
			"route": UpdateServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}



//****************************************************************************
// RP: Input And Output Returning Error methods
// Method: Get
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Get
//****************************************************************************

// GetServiceRoute defines the route for the Get method.
const GetServiceRoute = "gokit.rpkit.examples.users.UserService/Get"

// GetServiceRoutePath defines the full method path for the Get method.
const GetServiceRoutePath = "/gokit.rpkit.examples.users.UserService/Get"

// getServiceRoutePathURL defines a parsed path for the Get, it
// ensures the created path is valid as a url.
var getServiceRoutePathURL = mustParseURL(GetServiceRoutePath)

// GetContractSource contains the source version of expected method contract.
const GetContractSource = `type GetMethodContract interface {
	Get(var1 context.Context,var2 string)  (users.User,error)  
}
`

// GetMethodContract defines a contract interface for method "Get"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type GetMethodContract interface{
	Get(var1 context.Context,var2 string)  (users.User,error)  
}

// GetDecoder defines a interface which expose a single method to decode the request data
// expected by GetMethodContract.Get.
type GetDecoder interface{
	Decode(io.Reader) (string, error)
}

// GetEncoder defines a interface which expose a single method to encode the response
// returned by GetMethodContract.Get.
type GetEncoder interface{
	Encode(io.Writer, users.User) error
}

// GetMethodService defines the returned signature by the ServiceGet
// which executes it's internal behaviour based off on it's GetMethodContract.
type GetMethodService func(context.Context, io.Writer, io.Reader) error

// ServeGetMethod implements the core contract behaviour to service "Get"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeGetMethod(provider GetMethodContract, encoder GetEncoder, decoder GetDecoder) GetMethodService {
	return func(ctx context.Context, w io.Writer, r io.Reader) error {
		input, err := decoder.Decode(r)
		if err != nil {
			return err
		}

		
		res, err := provider.Get(ctx, input)
		
		if err != nil {
			return err
		}

		return encoder.Encode(w, res)
	}
}

// GetClientEncoder defines a interface which expose a single method to encode the request
// data sent by GetClientContract.Get.
type GetClientEncoder interface{
	Encode(io.Writer, string) error
}

// GetClientDecoder defines a interface which expose a single method to decode the response
// returned by GetClientContract.Get.
type GetClientDecoder interface{
	Decode(io.Reader) (users.User, error)
}

// GetClientContract defines a contract interface for clients to make request to GetServer.
type GetClientContract interface {
	Get(var1 context.Context,var2 string) (users.User,error)
}

// NewGetMethodClient returns a new GetMethodContract it relies on
// NewGetMethodContractClient.
func NewGetMethodClient(addr string, client *http.Client, encoder GetClientEncoder, decoder GetClientDecoder) (GetClientContract, error) {
	return NewGetMethodContractClient(addr, client, encoder, decoder, nil, nil)
}

// NewGetMethodContractClient returns a new GetMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a GetServer to
// perform action as specified by contract.
func NewGetMethodContractClient(addr string, client *http.Client, encoder GetClientEncoder, decoder GetClientDecoder, act ActWithRequest, resv ResponseValidation) (GetClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implGetClient{
		actor: act,
		resval: resv,
		rootURL: root,
		encoder: encoder,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implGetClient struct {
	rootURL *url.URL
	client *http.Client
	actor ActWithRequest
	resval ResponseValidation
	encoder GetClientEncoder
	decoder GetClientDecoder
}

// Get makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implGetClient) Get(var1 context.Context,var2 string) (users.User,error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(getServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	var body bytes.Buffer
	if err := imp.encoder.Encode(&body, var2); err != nil {
		return users.User{}, err
	}

	req, err := http.NewRequest("POST", targetURL.String(), &body)
	if err != nil {
		return users.User{}, err
	}

	if header, err := CtxCustomHeader(var1); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return users.User{},err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return users.User{},err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return users.User{},jsonErr
		}
		return users.User{},ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return users.User{},jsonErr
		}
		return users.User{},ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return users.User{},jsonErr
		}
		return users.User{},ErrServerRejectedRequest
	}

	rec, err := imp.decoder.Decode(res.Body)
	if err != nil {
		return users.User{}, err
	}

	return rec, nil
	
}

// GetServer implements a http.Handler for servicing the method Get
// from users.UserService.
type GetServer interface {
	http.Handler
}

type implGetHandler struct{
	hook Hook
	headers http.Header
	service GetMethodService
}

// NewGetServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.Get.
func NewGetServer(service GetMethodService, hook Hook, headers http.Header) GetServer {
	return implGetHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "Get".
func (impl implGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "Get")
	ctx = WithServiceMethodPath(ctx, GetServiceRoute)
	ctx = WithServiceMethodRoute(ctx, GetServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Get",
			"api_service": MethodServiceName,
			"route": GetServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if r.Method == "HEAD" {
		for key, vals := range impl.headers {
			for _, item := range vals {
				w.Header().Add(key, item)
			}
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "Get")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", GetServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != GetServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+GetServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Get",
			"api_service": MethodServiceName,
			"route": GetServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.RequestAccepted(ctx)
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	for key, vals := range impl.headers {
		for _, item := range vals {
			w.Header().Add(key, item)
		}
	}

	w.Header().Add("X-Agent", "RPKIT")
	w.Header().Add("X-Service", BaseServiceName)
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
	w.Header().Add("X-Method", "Get")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", GetServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "Get",
					"api_service": MethodServiceName,
					"route": GetServiceRoute,
					"api": "users.UserService",
				})
				panic(rerr)
			}
		}()

		if impl.hook != nil {
			impl.hook.RequestProcessed(ctx)
		}

		actionErr = impl.service(ctx, w, r.Body)
	}()

	if actionErr != nil {
		jsonWriteError(w, http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Get",
			"api_service": MethodServiceName,
			"route": GetServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// RP: Input And Output Returning Error methods
// Method: Create
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Create
//****************************************************************************

// CreateServiceRoute defines the route for the Create method.
const CreateServiceRoute = "gokit.rpkit.examples.users.UserService/Create"

// CreateServiceRoutePath defines the full method path for the Create method.
const CreateServiceRoutePath = "/gokit.rpkit.examples.users.UserService/Create"

// createServiceRoutePathURL defines a parsed path for the Create, it
// ensures the created path is valid as a url.
var createServiceRoutePathURL = mustParseURL(CreateServiceRoutePath)

// CreateContractSource contains the source version of expected method contract.
const CreateContractSource = `type CreateMethodContract interface {
	Create(var1 context.Context,var2 users.NewUser)  (users.User,error)  
}
`

// CreateMethodContract defines a contract interface for method "Create"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type CreateMethodContract interface{
	Create(var1 context.Context,var2 users.NewUser)  (users.User,error)  
}

// CreateDecoder defines a interface which expose a single method to decode the request data
// expected by CreateMethodContract.Create.
type CreateDecoder interface{
	Decode(io.Reader) (users.NewUser, error)
}

// CreateEncoder defines a interface which expose a single method to encode the response
// returned by CreateMethodContract.Create.
type CreateEncoder interface{
	Encode(io.Writer, users.User) error
}

// CreateMethodService defines the returned signature by the ServiceCreate
// which executes it's internal behaviour based off on it's CreateMethodContract.
type CreateMethodService func(context.Context, io.Writer, io.Reader) error

// ServeCreateMethod implements the core contract behaviour to service "Create"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeCreateMethod(provider CreateMethodContract, encoder CreateEncoder, decoder CreateDecoder) CreateMethodService {
	return func(ctx context.Context, w io.Writer, r io.Reader) error {
		input, err := decoder.Decode(r)
		if err != nil {
			return err
		}

		
		res, err := provider.Create(ctx, input)
		
		if err != nil {
			return err
		}

		return encoder.Encode(w, res)
	}
}

// CreateClientEncoder defines a interface which expose a single method to encode the request
// data sent by CreateClientContract.Create.
type CreateClientEncoder interface{
	Encode(io.Writer, users.NewUser) error
}

// CreateClientDecoder defines a interface which expose a single method to decode the response
// returned by CreateClientContract.Create.
type CreateClientDecoder interface{
	Decode(io.Reader) (users.User, error)
}

// CreateClientContract defines a contract interface for clients to make request to CreateServer.
type CreateClientContract interface {
	Create(var1 context.Context,var2 users.NewUser) (users.User,error)
}

// NewCreateMethodClient returns a new CreateMethodContract it relies on
// NewCreateMethodContractClient.
func NewCreateMethodClient(addr string, client *http.Client, encoder CreateClientEncoder, decoder CreateClientDecoder) (CreateClientContract, error) {
	return NewCreateMethodContractClient(addr, client, encoder, decoder, nil, nil)
}

// NewCreateMethodContractClient returns a new CreateMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a CreateServer to
// perform action as specified by contract.
func NewCreateMethodContractClient(addr string, client *http.Client, encoder CreateClientEncoder, decoder CreateClientDecoder, act ActWithRequest, resv ResponseValidation) (CreateClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implCreateClient{
		actor: act,
		resval: resv,
		rootURL: root,
		encoder: encoder,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implCreateClient struct {
	rootURL *url.URL
	client *http.Client
	actor ActWithRequest
	resval ResponseValidation
	encoder CreateClientEncoder
	decoder CreateClientDecoder
}

// Create makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implCreateClient) Create(var1 context.Context,var2 users.NewUser) (users.User,error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(createServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	var body bytes.Buffer
	if err := imp.encoder.Encode(&body, var2); err != nil {
		return users.User{}, err
	}

	req, err := http.NewRequest("POST", targetURL.String(), &body)
	if err != nil {
		return users.User{}, err
	}

	if header, err := CtxCustomHeader(var1); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return users.User{},err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return users.User{},err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return users.User{},jsonErr
		}
		return users.User{},ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return users.User{},jsonErr
		}
		return users.User{},ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return users.User{},jsonErr
		}
		return users.User{},ErrServerRejectedRequest
	}

	rec, err := imp.decoder.Decode(res.Body)
	if err != nil {
		return users.User{}, err
	}

	return rec, nil
	
}

// CreateServer implements a http.Handler for servicing the method Create
// from users.UserService.
type CreateServer interface {
	http.Handler
}

type implCreateHandler struct{
	hook Hook
	headers http.Header
	service CreateMethodService
}

// NewCreateServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.Create.
func NewCreateServer(service CreateMethodService, hook Hook, headers http.Header) CreateServer {
	return implCreateHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "Create".
func (impl implCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "Create")
	ctx = WithServiceMethodPath(ctx, CreateServiceRoute)
	ctx = WithServiceMethodRoute(ctx, CreateServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Create",
			"api_service": MethodServiceName,
			"route": CreateServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if r.Method == "HEAD" {
		for key, vals := range impl.headers {
			for _, item := range vals {
				w.Header().Add(key, item)
			}
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "Create")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", CreateServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != CreateServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+CreateServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Create",
			"api_service": MethodServiceName,
			"route": CreateServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.RequestAccepted(ctx)
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	for key, vals := range impl.headers {
		for _, item := range vals {
			w.Header().Add(key, item)
		}
	}

	w.Header().Add("X-Agent", "RPKIT")
	w.Header().Add("X-Service", BaseServiceName)
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
	w.Header().Add("X-Method", "Create")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", CreateServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "Create",
					"api_service": MethodServiceName,
					"route": CreateServiceRoute,
					"api": "users.UserService",
				})
				panic(rerr)
			}
		}()

		if impl.hook != nil {
			impl.hook.RequestProcessed(ctx)
		}

		actionErr = impl.service(ctx, w, r.Body)
	}()

	if actionErr != nil {
		jsonWriteError(w, http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Create",
			"api_service": MethodServiceName,
			"route": CreateServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// RP: Input And Output Returning Error methods
// Method: GetAll
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetAll
//****************************************************************************

// GetAllServiceRoute defines the route for the GetAll method.
const GetAllServiceRoute = "gokit.rpkit.examples.users.UserService/GetAll"

// GetAllServiceRoutePath defines the full method path for the GetAll method.
const GetAllServiceRoutePath = "/gokit.rpkit.examples.users.UserService/GetAll"

// getallServiceRoutePathURL defines a parsed path for the GetAll, it
// ensures the created path is valid as a url.
var getallServiceRoutePathURL = mustParseURL(GetAllServiceRoutePath)

// GetAllContractSource contains the source version of expected method contract.
const GetAllContractSource = `type GetAllMethodContract interface {
	GetAll(var1 context.Context,var2 string)  ([]users.User,error)  
}
`

// GetAllMethodContract defines a contract interface for method "GetAll"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type GetAllMethodContract interface{
	GetAll(var1 context.Context,var2 string)  ([]users.User,error)  
}

// GetAllDecoder defines a interface which expose a single method to decode the request data
// expected by GetAllMethodContract.GetAll.
type GetAllDecoder interface{
	Decode(io.Reader) (string, error)
}

// GetAllEncoder defines a interface which expose a single method to encode the response
// returned by GetAllMethodContract.GetAll.
type GetAllEncoder interface{
	Encode(io.Writer, []users.User) error
}

// GetAllMethodService defines the returned signature by the ServiceGetAll
// which executes it's internal behaviour based off on it's GetAllMethodContract.
type GetAllMethodService func(context.Context, io.Writer, io.Reader) error

// ServeGetAllMethod implements the core contract behaviour to service "GetAll"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeGetAllMethod(provider GetAllMethodContract, encoder GetAllEncoder, decoder GetAllDecoder) GetAllMethodService {
	return func(ctx context.Context, w io.Writer, r io.Reader) error {
		input, err := decoder.Decode(r)
		if err != nil {
			return err
		}

		
		res, err := provider.GetAll(ctx, input)
		
		if err != nil {
			return err
		}

		return encoder.Encode(w, res)
	}
}

// GetAllClientEncoder defines a interface which expose a single method to encode the request
// data sent by GetAllClientContract.GetAll.
type GetAllClientEncoder interface{
	Encode(io.Writer, string) error
}

// GetAllClientDecoder defines a interface which expose a single method to decode the response
// returned by GetAllClientContract.GetAll.
type GetAllClientDecoder interface{
	Decode(io.Reader) ([]users.User, error)
}

// GetAllClientContract defines a contract interface for clients to make request to GetAllServer.
type GetAllClientContract interface {
	GetAll(var1 context.Context,var2 string) ([]users.User,error)
}

// NewGetAllMethodClient returns a new GetAllMethodContract it relies on
// NewGetAllMethodContractClient.
func NewGetAllMethodClient(addr string, client *http.Client, encoder GetAllClientEncoder, decoder GetAllClientDecoder) (GetAllClientContract, error) {
	return NewGetAllMethodContractClient(addr, client, encoder, decoder, nil, nil)
}

// NewGetAllMethodContractClient returns a new GetAllMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a GetAllServer to
// perform action as specified by contract.
func NewGetAllMethodContractClient(addr string, client *http.Client, encoder GetAllClientEncoder, decoder GetAllClientDecoder, act ActWithRequest, resv ResponseValidation) (GetAllClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implGetAllClient{
		actor: act,
		resval: resv,
		rootURL: root,
		encoder: encoder,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implGetAllClient struct {
	rootURL *url.URL
	client *http.Client
	actor ActWithRequest
	resval ResponseValidation
	encoder GetAllClientEncoder
	decoder GetAllClientDecoder
}

// GetAll makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implGetAllClient) GetAll(var1 context.Context,var2 string) ([]users.User,error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(getallServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	var body bytes.Buffer
	if err := imp.encoder.Encode(&body, var2); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", targetURL.String(), &body)
	if err != nil {
		return nil, err
	}

	if header, err := CtxCustomHeader(var1); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return nil,err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return nil,err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return nil,jsonErr
		}
		return nil,ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return nil,jsonErr
		}
		return nil,ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return nil,jsonErr
		}
		return nil,ErrServerRejectedRequest
	}

	rec, err := imp.decoder.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return rec, nil
	
}

// GetAllServer implements a http.Handler for servicing the method GetAll
// from users.UserService.
type GetAllServer interface {
	http.Handler
}

type implGetAllHandler struct{
	hook Hook
	headers http.Header
	service GetAllMethodService
}

// NewGetAllServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.GetAll.
func NewGetAllServer(service GetAllMethodService, hook Hook, headers http.Header) GetAllServer {
	return implGetAllHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "GetAll".
func (impl implGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "GetAll")
	ctx = WithServiceMethodPath(ctx, GetAllServiceRoute)
	ctx = WithServiceMethodRoute(ctx, GetAllServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetAll",
			"api_service": MethodServiceName,
			"route": GetAllServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if r.Method == "HEAD" {
		for key, vals := range impl.headers {
			for _, item := range vals {
				w.Header().Add(key, item)
			}
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "GetAll")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", GetAllServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != GetAllServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+GetAllServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetAll",
			"api_service": MethodServiceName,
			"route": GetAllServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.RequestAccepted(ctx)
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	for key, vals := range impl.headers {
		for _, item := range vals {
			w.Header().Add(key, item)
		}
	}

	w.Header().Add("X-Agent", "RPKIT")
	w.Header().Add("X-Service", BaseServiceName)
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
	w.Header().Add("X-Method", "GetAll")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", GetAllServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "GetAll",
					"api_service": MethodServiceName,
					"route": GetAllServiceRoute,
					"api": "users.UserService",
				})
				panic(rerr)
			}
		}()

		if impl.hook != nil {
			impl.hook.RequestProcessed(ctx)
		}

		actionErr = impl.service(ctx, w, r.Body)
	}()

	if actionErr != nil {
		jsonWriteError(w, http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetAll",
			"api_service": MethodServiceName,
			"route": GetAllServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}


//****************************************************************************
// Utils
// Source: github.com/gokit/rpkit/examples/users
//****************************************************************************

func requestRedirected(res *http.Response) bool {
	if res.StatusCode >= 300 && res.StatusCode <= 308  {
		 return true
	}
	return false
}

func requestFailed(res *http.Response) bool {
	if res.StatusCode >= 400 && res.StatusCode < 500 {
		 return true
	}
	return false
}

func requestFacedInternalIssues(res *http.Response) bool {
	if res.StatusCode >= 500 && res.StatusCode < 600 {
		 return true
	}
	return false
}

func requestSucceeded(res *http.Response) bool {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		 return true
	}
	return false
}

// mustParseURL attempts to parse provided access using parseURL, it panics
// if an error occured.
func mustParseURL(addr string) *url.URL {
	parsed, err := parseURL(addr)
	if err == nil { return parsed }
	panic(`failed to parse url: `+addr+` , error occurred: `+err.Error())
}

// parseURL parses giving address returning *url.URL instance for address if
// it's a valid address, ensuring the scheme is provided or setting 'http'
// if not set. Returns an error if address is an invalid URL.
func parseURL(addr string) (*url.URL, error) {
	parsed, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	if parsed.Scheme == "" {
		parsed.Scheme = "http"
	}

	return parsed, nil
}

// joinURL will attempt to join the provided address (addr) to the root URL
// provided if the address is not an absolute URL address path. An error will
// be returned if addr is not a valid URL path.
func joinURL(root *url.URL, addr string) (*url.URL, error) {
	parsed, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	if parsed.IsAbs(){
		parsed = root.ResolveReference(parsed)
	}

	return parsed, nil
}


// skipRedirects copies giving http.Client and fix issue with possible redirects
// in go1.8 when a post receives a status code which is undesired. Such has 302, 303,
// and 301s.
func skipRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			err := in.CheckRedirect(req, via)
			_ = err // to fix issue with warning about not checking error.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

func loadJSONError(r io.Reader) (jsonErrorResponse, error) {
	var res jsonErrorResponse
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}

func jsonWriteError(w http.ResponseWriter, code int, message string, err error, meta map[string]interface{}){
	var res jsonErrorResponse
	res.Code = code
	res.Err = err
	res.Meta = meta
	res.Message = message

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	data, err2 := json.Marshal(res);
	if err2 == nil {
		w.Write(data)
		return
	}

	log.Printf("unable to send error response for error %+q: %+q", err, err2)
}

type jsonErrorResponse struct{
	Code int `json:"code"`
	Err error `json:"err"`
	Message string `json:"message"`
	Meta map[string]interface{} `json:"meta"`
}

func (jse jsonErrorResponse) Error() string {
	return jse.Message + " " + jse.Err.Error()
}

