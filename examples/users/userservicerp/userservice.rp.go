package userservicerp

// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.
  
import (
	"context"
	"encoding/json"
	"bytes"
	"errors"
	"strings"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/gokit/rpkit/examples/users"
	
)

var (
    bytePool = sync.Pool{
        New: func() interface{} {
            return bytes.NewBuffer(make([]byte, 0, 1024))
        },
    }
)

const (
	// BaseServiceName defines the base name of service root.
	BaseServiceName = "users"

	// MethodServiceName defines the complete name of this giving API service.
	MethodServiceName = "users/UserService"

	// ServiceCodePath defines the path to this generated package which contains the implemented service methods.
	ServiceCodePath = "github.com/gokit/rpkit/examples/users/userservicerp"

	// ServiceCodePathName defines the name giving to
	// this package.
	ServiceCodePathName = "userservicerp"
)

// error type strings
const (
    URLError = "url_error"
    ActionError = "action_error"
    ActionPanicError = "action_panic_error"
    MethodTypeError = "method_type_error"
    AcceptTypeUnknownError = "accept_type_unknown_error"
    RequestDecodingError = "request_decoding_error"
    ResponseEncodingError = "response_encoding_error"
)

// errors ...
var (
	ErrNotInContext = errors.New("item not in context")
	ErrInvalidRequestURI = errors.New("invalid request uri")
	ErrBadRequest = errors.New("server rejected request as bad")
	ErrInvalidRequestMethod = errors.New("invalid request method")
	ErrInvalidContentType = errors.New("invalid request content type")
	ErrServerInternalProblem = errors.New("server had internal problems")
	ErrServerRejectedRequest = errors.New("server does not handle request type/method")
)

//****************************************************************************
// Context Keys, Setters And Getters
//****************************************************************************

type contextKeyType string

// sets of context keys used by the package.
const (
	contextCustomHeaderKey = contextKeyType("rp:custom:header")
	contextRequestKey = contextKeyType("rp:request")
	contextRequestMethodKey = contextKeyType("rp:request:method")
	contextRequestContentTypeKey = contextKeyType("rp:request:content:type")
	contextRequestAcceptsKey = contextKeyType("rp:request:accepts:types")
	contextRequestHeaderKey = contextKeyType("rp:request:header")
	contextResponseWriterKey = contextKeyType("rp:response:writer")
	contextResponseEncodedAsKey = contextKeyType("rp:response:encoded:as")
	contextServiceNameKey = contextKeyType("rp:service:name")
	contextServicePathKey = contextKeyType("rp:service:path")
	contextServicePackageKey = contextKeyType("rp:service:package")
	contextServicePackageNameKey = contextKeyType("rp:service:package:name")
	contextServiceSourcePackageKey = contextKeyType("rp:service:source:path")
	contextServiceSourcePackageNameKey = contextKeyType("rp:service:source:package")
	contextServiceMethodNameKey = contextKeyType("rp:service:method:name")
	contextServiceMethodPathKey = contextKeyType("rp:service:method:path")
	contextServiceMethodRouteKey = contextKeyType("rp:service:method:route")
	contextClientRequestURLKey = contextKeyType("rp:client:request:url")
	contextRequestTransportKey = contextKeyType("rp:request:transport")
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

// WithResponseEncodedAs sets the giving encoding type used to write to response into context.
func WithResponseEncodedAs(ctx context.Context, c string) context.Context {
	return context.WithValue(ctx, contextResponseEncodedAsKey, c)
}

// CtxResponseEncodedAs retrieves the Response encoding value from the context.
func CtxResponseEncodedAs(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextResponseEncodedAsKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithAcceptsType sets the giving accepts type for associated context.
func WithAcceptsType(ctx context.Context, c string) context.Context {
	return context.WithValue(ctx, contextRequestAcceptsKey, c)
}

// CtxAccepts retrieves the Accepts value from the http.Request.
func CtxAccepts(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextRequestAcceptsKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
}

// WithContentType sets the giving content type for associated context.
func WithContentType(ctx context.Context, c string) context.Context {
	return context.WithValue(ctx, contextRequestContentTypeKey, c)
}

// CtxRequestContentType retrieves if any set http.Request content type from context.
func CtxRequestContentType(ctx context.Context) (string, error) {
	if item, ok := ctx.Value(contextRequestContentTypeKey).(string); ok {
		return item, nil
	}
	return "", ErrNotInContext
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

// JSONErrorResponse defines a structure to contain error message data
// delivered by the server.
type JSONErrorResponse struct{
	Type string `json:"type"`
	Code int `json:"code"`
	Err []byte `json:"err"`
	Message string `json:"message"`
	Meta map[string]interface{} `json:"meta"`
}

// Error returns the underline error contents string.
func (jse JSONErrorResponse) Error() string {
	return jse.Message + " : " + string(jse.Err)
}

// Hook defines a interface for having access to different areas of
// a call to the service of a method.
//
//@implement
type Hook interface {
	ResponseSent(context.Context)
	ResponsePrepared(context.Context)
	RequestRejected(context.Context)
	RequestAccepted(context.Context)
	RequestReceived(context.Context)
	RequestProcessed(context.Context)
	RequestError(context.Context, error)
	RequestPanic(context.Context, interface{})
}

// HTTPClient defines an interface which is used as the processor for a
// raw http request.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}



//****************************************************************************
// RP: Output with Returning Error methods
// Method: Get
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Get
//****************************************************************************

// GetServiceRoute defines the route for the Get method.
const GetServiceRoute = "users.UserService/Get"

// GetServiceRoutePath defines the full method path for the Get method.
const GetServiceRoutePath = "/rpkit/users.UserService/Get"

// getServiceRoutePathURL defines a parsed path for the Get, it
// ensures the created path is valid as a url.
var getServiceRoutePathURL = mustSimpleParseURL(GetServiceRoutePath)

// GetContractSource contains the source version of expected method contract.
const GetContractSource = `type GetMethodContract interface {
	Get(var1 context.Context)  (int,error)  
}
`

// GetMethodContract defines a contract interface for method "Get"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type GetMethodContract interface{
	Get(var1 context.Context)  (int,error)  
}

// GetEncoder defines a interface which expose a single method to encode the response
// returned by GetMethodContract.Get.
type GetEncoder interface{
	Encode(context.Context, io.Writer, int) error
}

// GetMethodService defines the returned signature by the ServiceGet
// which executes it's internal behaviour based off on it's GetMethodContract.
type GetMethodService func(context.Context, io.Writer) error

// ServeGetMethod implements the core contract behaviour to service "Get"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeGetMethod(provider GetMethodContract, encoder GetEncoder) GetMethodService {
	return func(ctx context.Context, w io.Writer) error {
		
		res, err := provider.Get(ctx)
		
		if err != nil {
			return err
		}

		return encoder.Encode(ctx, w, res)
	}
}

// GetClientContract defines a contract interface for clients to make request to GetServer.
type GetClientContract interface {
	Get(var1 context.Context) (int,error)
}

// GetClientDecoder defines a interface which expose a single method to decode the response
// returned by GetClientContract.Get.
type GetClientDecoder interface{
	Decode(context.Context, io.Reader) (int, error)
}

// NewGetMethodClient returns a new GetMethodContract it relies on
// NewGetMethodContractClient.
func NewGetMethodClient(addr string, client HTTPClient, decoder GetClientDecoder) (GetClientContract, error) {
	return NewGetMethodContractClient(addr, client, decoder, nil, nil)
}

// NewGetMethodContractClient returns a new GetMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a GetServer to
// perform action as specified by contract.
func NewGetMethodContractClient(addr string, client HTTPClient, decoder GetClientDecoder, act ActWithRequest, resv ResponseValidation) (GetClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implGetClient{
		actor: act,
		resval: resv,
		rootURL: root,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implGetClient struct {
	rootURL *url.URL
	client HTTPClient
	actor ActWithRequest
	resval ResponseValidation
	decoder GetClientDecoder
}

// Get makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implGetClient) Get(var1 context.Context) (int,error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(getServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")
	
	var result int

	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return result, err
	}
	req = req.WithContext(var1)

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
		return result, err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return result, err
		}
	}

	if requestFacedInternalIssues(res) {
		jsonErr, err := loadJSONError(res.Body)
	    if err == nil {
		  return result, jsonErr
	     }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result, ErrServerInternalProblem
	}

	if requestFailed(res) {
		jsonErr, err := loadJSONError(res.Body)
	    if err == nil {
		  return result, jsonErr
	     }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result, ErrBadRequest
	}

	if requestRedirected(res) {
		jsonErr, err := loadJSONError(res.Body)
		if err == nil {
			return result, jsonErr
		}
		log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result, ErrServerRejectedRequest
	}

	result, err = imp.decoder.Decode(var1, res.Body)
	if err != nil {
		return result, err
	}

	return result, nil
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
	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithAcceptsType(ctx, r.Header.Get("Accept"))
	ctx = WithContentType(ctx, r.Header.Get("Content-Type"))
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

		jsonWriteError(w, MethodTypeError,  http.StatusMethodNotAllowed, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
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
		
		if impl.hook != nil {
			impl.hook.RequestAccepted(ctx)
			impl.hook.RequestProcessed(ctx)
			impl.hook.ResponsePrepared(ctx)
		}
		
		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "Get")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", GetServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		
		if impl.hook != nil {
			impl.hook.ResponseSent(ctx)
		}
		
		return
	}

	if !strings.HasSuffix(prefixIfNeed(r.URL.Path, "/"), GetServiceRoutePath) {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, URLError,  http.StatusBadRequest, "only POST request to "+GetServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Get",
			"api_service": MethodServiceName,
			"route": GetServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.headers != nil {
		if accepts := impl.headers.Get("Accept"); accepts != "" {
			ctype := r.Header.Get("Content-Type")
			if !strings.Contains(accepts, ctype){
				if impl.hook != nil {
					impl.hook.RequestRejected(ctx)
				}
		
				jsonWriteError(w, AcceptTypeUnknownError,  http.StatusBadRequest, "request content type not supported",
				ErrInvalidContentType, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "Get",
					"api_service": MethodServiceName,
					"route": GetServiceRoute,
					"api": "users.UserService",
				})
				return
			}
		}
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
	w.Header().Add("X-Method", "Get")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", GetServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")

	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				if impl.hook != nil {
					impl.hook.RequestPanic(ctx, rerr)
				}
				
				derr := fmt.Errorf("panic err: %#v", rerr)
				jsonWriteError(w, ActionPanicError,  http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
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

		actionErr = impl.service(ctx, w)
	}()

	if actionErr != nil {
		if impl.hook != nil {
			impl.hook.RequestError(ctx,actionErr)
		}
		
		jsonWriteError(w, ActionError,  http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
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
		impl.hook.ResponsePrepared(ctx)
	}

	w.WriteHeader(http.StatusOK)
	
	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}


//****************************************************************************
// RP: Output Returning No Error methods
// Method: GetUsers
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetUsers
//****************************************************************************

// GetUsersServiceRoute defines the route for the GetUsers method.
const GetUsersServiceRoute = "users.UserService/GetUsers"

// GetUsersServiceRoutePath defines the full method path for the GetUsers method.
const GetUsersServiceRoutePath = "/rpkit/users.UserService/GetUsers"

// getusersServiceRoutePathURL defines a parsed path for the GetUsers, it
// ensures the created path is valid as a url.
var getusersServiceRoutePathURL = mustSimpleParseURL(GetUsersServiceRoutePath)

// GetUsersContractSource contains the source version of expected method contract.
const GetUsersContractSource = `type GetUsersMethodContract interface {
	GetUsers(var1 context.Context)  users.User  
}
`

// GetUsersMethodContract defines a contract interface for method "GetUsers"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type GetUsersMethodContract interface{
	GetUsers(var1 context.Context)  users.User  
}

// GetUsersEncoder defines a interface which expose a single method to encode the response
// returned by GetUsersMethodContract.GetUsers.
type GetUsersEncoder interface{
	Encode(context.Context, io.Writer, users.User) error
}

// GetUsersMethodService defines the returned signature by the ServiceGetUsers
// which executes it's internal behaviour based off on it's GetUsersMethodContract.
type GetUsersMethodService func(context.Context, io.Writer) error

// ServeGetUsersMethod implements the core contract behaviour to service "GetUsers"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeGetUsersMethod(provider GetUsersMethodContract, encoder GetUsersEncoder) GetUsersMethodService {
	return func(ctx context.Context, w io.Writer) error {
		
		res := provider.GetUsers(ctx)
		
		return encoder.Encode(ctx, w, res)
	}
}

// GetUsersClientContract defines a contract interface for clients to make request to GetUsersServer.
type GetUsersClientContract interface {
	GetUsers(var1 context.Context) (users.User, error)
}

// GetUsersClientDecoder defines a interface which expose a single method to decode the response
// returned by GetUsersClientContract.GetUsers.
type GetUsersClientDecoder interface{
	Decode(context.Context, io.Reader) (users.User, error)
}

// NewGetUsersMethodClient returns a new GetUsersMethodContract it relies on
// NewGetUsersMethodContractClient.
func NewGetUsersMethodClient(addr string, client HTTPClient, decoder GetUsersClientDecoder) (GetUsersClientContract, error) {
	return NewGetUsersMethodContractClient(addr, client, decoder, nil, nil)
}

// NewGetUsersMethodContractClient returns a new GetUsersMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a GetUsersServer to
// perform action as specified by contract.
func NewGetUsersMethodContractClient(addr string, client HTTPClient, decoder GetUsersClientDecoder, act ActWithRequest, resv ResponseValidation) (GetUsersClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implGetUsersClient{
		actor: act,
		resval: resv,
		rootURL: root,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implGetUsersClient struct {
	rootURL *url.URL
	client HTTPClient
	actor ActWithRequest
	resval ResponseValidation
	decoder GetUsersClientDecoder
}

// GetUsers makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implGetUsersClient) GetUsers(var1 context.Context) (users.User, error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(getusersServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	var result users.User
	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return result, err
	}
	req = req.WithContext(var1)

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
		return result, err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return result, err
		}
	}

	if requestFacedInternalIssues(res) {
		jsonErr, err := loadJSONError(res.Body) 
	    if err == nil {
		  return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result, ErrServerInternalProblem
	}

	if requestFailed(res) {
		jsonErr, err := loadJSONError(res.Body)
	    if err == nil {
		  return result, jsonErr
	     }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result, ErrBadRequest
	}

	if requestRedirected(res) {
		jsonErr, err := loadJSONError(res.Body)
	    if err == nil {
		  return result, jsonErr
	     }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result, ErrServerRejectedRequest
	}

	result, err = imp.decoder.Decode(var1, res.Body)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetUsersServer implements a http.Handler for servicing the method GetUsers
// from users.UserService.
type GetUsersServer interface {
	http.Handler
}

type implGetUsersHandler struct{
	hook Hook
	headers http.Header
	service GetUsersMethodService
}

// NewGetUsersServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.GetUsers.
func NewGetUsersServer(service GetUsersMethodService, hook Hook, headers http.Header) GetUsersServer {
	return implGetUsersHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "GetUsers".
func (impl implGetUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithAcceptsType(ctx, r.Header.Get("Accept"))
	ctx = WithContentType(ctx, r.Header.Get("Content-Type"))
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "GetUsers")
	ctx = WithServiceMethodPath(ctx, GetUsersServiceRoute)
	ctx = WithServiceMethodRoute(ctx, GetUsersServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, MethodTypeError,  http.StatusMethodNotAllowed, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetUsers",
			"api_service": MethodServiceName,
			"route": GetUsersServiceRoute,
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

		if impl.hook != nil {
			impl.hook.RequestAccepted(ctx)
			impl.hook.RequestProcessed(ctx)
			impl.hook.ResponsePrepared(ctx)
		}
		
		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "GetUsers")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", GetUsersServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)
		
		if impl.hook != nil {
			impl.hook.ResponseSent(ctx)
		}
		
		return
	}

	if !strings.HasSuffix(prefixIfNeed(r.URL.Path, "/"), GetUsersServiceRoutePath) {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, URLError,  http.StatusBadRequest, "only POST request to "+GetUsersServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetUsers",
			"api_service": MethodServiceName,
			"route": GetUsersServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.headers != nil {
		if accepts := impl.headers.Get("Accept"); accepts != "" {
			ctype := r.Header.Get("Content-Type")
			if !strings.Contains(accepts, ctype){
				if impl.hook != nil {
					impl.hook.RequestRejected(ctx)
				}
		
				jsonWriteError(w, AcceptTypeUnknownError,  http.StatusBadRequest, "request content type not supported",
				ErrInvalidContentType, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "GetUsers",
					"api_service": MethodServiceName,
					"route": GetUsersServiceRoute,
					"api": "users.UserService",
				})
				return
			}
		}
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
	w.Header().Add("X-Method", "GetUsers")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", GetUsersServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")


	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				if impl.hook != nil {
					impl.hook.RequestPanic(ctx, rerr)
				}
				
				derr := fmt.Errorf("panic err: %#v", rerr)
				jsonWriteError(w, ActionPanicError,  http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "GetUsers",
					"api_service": MethodServiceName,
					"route": GetUsersServiceRoute,
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
		if impl.hook != nil {
			impl.hook.RequestError(ctx, actionErr)
		}
		
		jsonWriteError(w, ActionError,  http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetUsers",
			"api_service": MethodServiceName,
			"route": GetUsersServiceRoute,
			"api": "users.UserService",
		})
		return
	}
	
	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	w.WriteHeader(http.StatusOK)
	
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
const CreateServiceRoute = "users.UserService/Create"

// CreateServiceRoutePath defines the full method path for the Create method.
const CreateServiceRoutePath = "/users/UserService/Create"

// createServiceRoutePathURL defines a parsed path for the Create, it
// ensures the created path is valid as a url.
var createServiceRoutePathURL = mustSimpleParseURL(CreateServiceRoutePath)

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
	Decode(context.Context, io.Reader) (users.NewUser, error)
}

// CreateEncoder defines a interface which expose a single method to encode the response
// returned by CreateMethodContract.Create.
type CreateEncoder interface{
	Encode(context.Context, io.Writer, users.User) error
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
		input, err := decoder.Decode(ctx, r)
		if err != nil {
			return err
		}

		
		res, err := provider.Create(ctx, input)
		
		if err != nil {
			return err
		}

		return encoder.Encode(ctx, w, res)
	}
}

// CreateClientEncoder defines a interface which expose a single method to encode the request
// data sent by CreateClientContract.Create.
type CreateClientEncoder interface{
	Encode(context.Context, io.Writer, users.NewUser) error
}

// CreateClientDecoder defines a interface which expose a single method to decode the response
// returned by CreateClientContract.Create.
type CreateClientDecoder interface{
	Decode(context.Context, io.Reader) (users.User, error)
}

// CreateClientContract defines a contract interface for clients to make request to CreateServer.
type CreateClientContract interface {
	Create(var1 context.Context,var2 users.NewUser) (users.User,error)
}

// NewCreateMethodClient returns a new CreateMethodContract it relies on
// NewCreateMethodContractClient.
func NewCreateMethodClient(addr string, client HTTPClient, encoder CreateClientEncoder, decoder CreateClientDecoder) (CreateClientContract, error) {
	return NewCreateMethodContractClient(addr, client, encoder, decoder, nil, nil)
}

// NewCreateMethodContractClient returns a new CreateMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a CreateServer to
// perform action as specified by contract.
func NewCreateMethodContractClient(addr string, client HTTPClient, encoder CreateClientEncoder, decoder CreateClientDecoder, act ActWithRequest, resv ResponseValidation) (CreateClientContract, error) {
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
	client HTTPClient
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

	var result users.User
	body := bytePool.Get().(*bytes.Buffer)
	defer bytePool.Put(body)
	if err := imp.encoder.Encode(var1, body, var2); err != nil {
		return result, err
	}

     req, err := http.NewRequest("POST", targetURL.String(), body)
      if err != nil {
           return result, err
      }
      req = req.WithContext(var1)

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
		return result,err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return result,err
		}
	}

	if requestFacedInternalIssues(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrServerInternalProblem
	}

	if requestFailed(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrBadRequest
	}

	if requestRedirected(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrServerRejectedRequest
	}

	result, err = imp.decoder.Decode(var1, res.Body)
	if err != nil {
		return result, err
	}

	return result, nil
	
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
	ctx = WithAcceptsType(ctx, r.Header.Get("Accept"))
	ctx = WithContentType(ctx, r.Header.Get("Content-Type"))
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

		jsonWriteError(w, MethodTypeError,  http.StatusMethodNotAllowed, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
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

		if impl.hook != nil {
			impl.hook.RequestAccepted(ctx)
			impl.hook.RequestProcessed(ctx)
			impl.hook.ResponsePrepared(ctx)
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "Create")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", CreateServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)

		if impl.hook != nil {
			impl.hook.ResponseSent(ctx)
		}
		return
	}

	if !strings.HasSuffix(prefixIfNeed(r.URL.Path, "/"), CreateServiceRoutePath) {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, URLError,  http.StatusBadRequest, "only POST request to "+CreateServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "Create",
			"api_service": MethodServiceName,
			"route": CreateServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.headers != nil {
		if accepts := impl.headers.Get("Accept"); accepts != "" {
			ctype := r.Header.Get("Content-Type")
			if !strings.Contains(accepts, ctype){
				if impl.hook != nil {
					impl.hook.RequestRejected(ctx)
				}

				jsonWriteError(w, AcceptTypeUnknownError,  http.StatusBadRequest, "request content type not supported",
				ErrInvalidContentType, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "Create",
					"api_service": MethodServiceName,
					"route": CreateServiceRoute,
					"api": "users.UserService",
				})
				return
			}
		}
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
	w.Header().Add("X-Method", "Create")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", CreateServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")


	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				if impl.hook != nil {
					impl.hook.RequestPanic(ctx, rerr)
				}

				derr := fmt.Errorf("panic err: %#v", rerr)
				jsonWriteError(w, ActionPanicError,  http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
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
		if impl.hook != nil {
			impl.hook.RequestError(ctx, actionErr)
		}

		jsonWriteError(w, ActionError,  http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
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
		impl.hook.ResponsePrepared(ctx)
	}

	w.WriteHeader(http.StatusOK)

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// RP: Input And Output Returning Error methods
// Method: GetBy
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetBy
//****************************************************************************

// GetByServiceRoute defines the route for the GetBy method.
const GetByServiceRoute = "users.UserService/GetBy"

// GetByServiceRoutePath defines the full method path for the GetBy method.
const GetByServiceRoutePath = "/users/UserService/GetBy"

// getbyServiceRoutePathURL defines a parsed path for the GetBy, it
// ensures the created path is valid as a url.
var getbyServiceRoutePathURL = mustSimpleParseURL(GetByServiceRoutePath)

// GetByContractSource contains the source version of expected method contract.
const GetByContractSource = `type GetByMethodContract interface {
	GetBy(var1 context.Context,var2 string)  (int,error)  
}
`

// GetByMethodContract defines a contract interface for method "GetBy"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type GetByMethodContract interface{
	GetBy(var1 context.Context,var2 string)  (int,error)  
}

// GetByDecoder defines a interface which expose a single method to decode the request data
// expected by GetByMethodContract.GetBy.
type GetByDecoder interface{
	Decode(context.Context, io.Reader) (string, error)
}

// GetByEncoder defines a interface which expose a single method to encode the response
// returned by GetByMethodContract.GetBy.
type GetByEncoder interface{
	Encode(context.Context, io.Writer, int) error
}

// GetByMethodService defines the returned signature by the ServiceGetBy
// which executes it's internal behaviour based off on it's GetByMethodContract.
type GetByMethodService func(context.Context, io.Writer, io.Reader) error

// ServeGetByMethod implements the core contract behaviour to service "GetBy"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeGetByMethod(provider GetByMethodContract, encoder GetByEncoder, decoder GetByDecoder) GetByMethodService {
	return func(ctx context.Context, w io.Writer, r io.Reader) error {
		input, err := decoder.Decode(ctx, r)
		if err != nil {
			return err
		}

		
		res, err := provider.GetBy(ctx, input)
		
		if err != nil {
			return err
		}

		return encoder.Encode(ctx, w, res)
	}
}

// GetByClientEncoder defines a interface which expose a single method to encode the request
// data sent by GetByClientContract.GetBy.
type GetByClientEncoder interface{
	Encode(context.Context, io.Writer, string) error
}

// GetByClientDecoder defines a interface which expose a single method to decode the response
// returned by GetByClientContract.GetBy.
type GetByClientDecoder interface{
	Decode(context.Context, io.Reader) (int, error)
}

// GetByClientContract defines a contract interface for clients to make request to GetByServer.
type GetByClientContract interface {
	GetBy(var1 context.Context,var2 string) (int,error)
}

// NewGetByMethodClient returns a new GetByMethodContract it relies on
// NewGetByMethodContractClient.
func NewGetByMethodClient(addr string, client HTTPClient, encoder GetByClientEncoder, decoder GetByClientDecoder) (GetByClientContract, error) {
	return NewGetByMethodContractClient(addr, client, encoder, decoder, nil, nil)
}

// NewGetByMethodContractClient returns a new GetByMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a GetByServer to
// perform action as specified by contract.
func NewGetByMethodContractClient(addr string, client HTTPClient, encoder GetByClientEncoder, decoder GetByClientDecoder, act ActWithRequest, resv ResponseValidation) (GetByClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implGetByClient{
		actor: act,
		resval: resv,
		rootURL: root,
		encoder: encoder,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implGetByClient struct {
	rootURL *url.URL
	client HTTPClient
	actor ActWithRequest
	resval ResponseValidation
	encoder GetByClientEncoder
	decoder GetByClientDecoder
}

// GetBy makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implGetByClient) GetBy(var1 context.Context,var2 string) (int,error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(getbyServiceRoutePathURL)
	
	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	var result int
	body := bytePool.Get().(*bytes.Buffer)
	defer bytePool.Put(body)
	if err := imp.encoder.Encode(var1, body, var2); err != nil {
		return result, err
	}

     req, err := http.NewRequest("POST", targetURL.String(), body)
      if err != nil {
           return result, err
      }
      req = req.WithContext(var1)

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
		return result,err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return result,err
		}
	}

	if requestFacedInternalIssues(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrServerInternalProblem
	}

	if requestFailed(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrBadRequest
	}

	if requestRedirected(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrServerRejectedRequest
	}

	result, err = imp.decoder.Decode(var1, res.Body)
	if err != nil {
		return result, err
	}

	return result, nil
	
}

// GetByServer implements a http.Handler for servicing the method GetBy
// from users.UserService.
type GetByServer interface {
	http.Handler
}

type implGetByHandler struct{
	hook Hook
	headers http.Header
	service GetByMethodService
}

// NewGetByServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.GetBy.
func NewGetByServer(service GetByMethodService, hook Hook, headers http.Header) GetByServer {
	return implGetByHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "GetBy".
func (impl implGetByHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithAcceptsType(ctx, r.Header.Get("Accept"))
	ctx = WithContentType(ctx, r.Header.Get("Content-Type"))
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "GetBy")
	ctx = WithServiceMethodPath(ctx, GetByServiceRoute)
	ctx = WithServiceMethodRoute(ctx, GetByServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, MethodTypeError,  http.StatusMethodNotAllowed, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetBy",
			"api_service": MethodServiceName,
			"route": GetByServiceRoute,
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

		if impl.hook != nil {
			impl.hook.RequestAccepted(ctx)
			impl.hook.RequestProcessed(ctx)
			impl.hook.ResponsePrepared(ctx)
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "GetBy")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", GetByServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)

		if impl.hook != nil {
			impl.hook.ResponseSent(ctx)
		}
		return
	}

	if !strings.HasSuffix(prefixIfNeed(r.URL.Path, "/"), GetByServiceRoutePath) {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, URLError,  http.StatusBadRequest, "only POST request to "+GetByServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetBy",
			"api_service": MethodServiceName,
			"route": GetByServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.headers != nil {
		if accepts := impl.headers.Get("Accept"); accepts != "" {
			ctype := r.Header.Get("Content-Type")
			if !strings.Contains(accepts, ctype){
				if impl.hook != nil {
					impl.hook.RequestRejected(ctx)
				}

				jsonWriteError(w, AcceptTypeUnknownError,  http.StatusBadRequest, "request content type not supported",
				ErrInvalidContentType, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "GetBy",
					"api_service": MethodServiceName,
					"route": GetByServiceRoute,
					"api": "users.UserService",
				})
				return
			}
		}
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
	w.Header().Add("X-Method", "GetBy")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", GetByServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")


	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				if impl.hook != nil {
					impl.hook.RequestPanic(ctx, rerr)
				}

				derr := fmt.Errorf("panic err: %#v", rerr)
				jsonWriteError(w, ActionPanicError,  http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "GetBy",
					"api_service": MethodServiceName,
					"route": GetByServiceRoute,
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
		if impl.hook != nil {
			impl.hook.RequestError(ctx, actionErr)
		}

		jsonWriteError(w, ActionError,  http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetBy",
			"api_service": MethodServiceName,
			"route": GetByServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	w.WriteHeader(http.StatusOK)

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// RP: Input And Output Returning Error methods
// Method: CreateUser
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.CreateUser
//****************************************************************************

// CreateUserServiceRoute defines the route for the CreateUser method.
const CreateUserServiceRoute = "users.UserService/CreateUser"

// CreateUserServiceRoutePath defines the full method path for the CreateUser method.
const CreateUserServiceRoutePath = "/users/UserService/CreateUser"

// createuserServiceRoutePathURL defines a parsed path for the CreateUser, it
// ensures the created path is valid as a url.
var createuserServiceRoutePathURL = mustSimpleParseURL(CreateUserServiceRoutePath)

// CreateUserContractSource contains the source version of expected method contract.
const CreateUserContractSource = `type CreateUserMethodContract interface {
	CreateUser(var1 users.NewUser)  (users.User,error)  
}
`

// CreateUserMethodContract defines a contract interface for method "CreateUser"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type CreateUserMethodContract interface{
	CreateUser(var1 users.NewUser)  (users.User,error)  
}

// CreateUserDecoder defines a interface which expose a single method to decode the request data
// expected by CreateUserMethodContract.CreateUser.
type CreateUserDecoder interface{
	Decode(context.Context, io.Reader) (users.NewUser, error)
}

// CreateUserEncoder defines a interface which expose a single method to encode the response
// returned by CreateUserMethodContract.CreateUser.
type CreateUserEncoder interface{
	Encode(context.Context, io.Writer, users.User) error
}

// CreateUserMethodService defines the returned signature by the ServiceCreateUser
// which executes it's internal behaviour based off on it's CreateUserMethodContract.
type CreateUserMethodService func(context.Context, io.Writer, io.Reader) error

// ServeCreateUserMethod implements the core contract behaviour to service "CreateUser"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeCreateUserMethod(provider CreateUserMethodContract, encoder CreateUserEncoder, decoder CreateUserDecoder) CreateUserMethodService {
	return func(ctx context.Context, w io.Writer, r io.Reader) error {
		input, err := decoder.Decode(ctx, r)
		if err != nil {
			return err
		}

		
		res, err := provider.CreateUser(input)
		
		if err != nil {
			return err
		}

		return encoder.Encode(ctx, w, res)
	}
}

// CreateUserClientEncoder defines a interface which expose a single method to encode the request
// data sent by CreateUserClientContract.CreateUser.
type CreateUserClientEncoder interface{
	Encode(context.Context, io.Writer, users.NewUser) error
}

// CreateUserClientDecoder defines a interface which expose a single method to decode the response
// returned by CreateUserClientContract.CreateUser.
type CreateUserClientDecoder interface{
	Decode(context.Context, io.Reader) (users.User, error)
}

// CreateUserClientContract defines a contract interface for clients to make request to CreateUserServer.
type CreateUserClientContract interface {
	CreateUser(ctx context.Context, var1 users.NewUser) (users.User,error)
}

// NewCreateUserMethodClient returns a new CreateUserMethodContract it relies on
// NewCreateUserMethodContractClient.
func NewCreateUserMethodClient(addr string, client HTTPClient, encoder CreateUserClientEncoder, decoder CreateUserClientDecoder) (CreateUserClientContract, error) {
	return NewCreateUserMethodContractClient(addr, client, encoder, decoder, nil, nil)
}

// NewCreateUserMethodContractClient returns a new CreateUserMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a CreateUserServer to
// perform action as specified by contract.
func NewCreateUserMethodContractClient(addr string, client HTTPClient, encoder CreateUserClientEncoder, decoder CreateUserClientDecoder, act ActWithRequest, resv ResponseValidation) (CreateUserClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implCreateUserClient{
		actor: act,
		resval: resv,
		rootURL: root,
		encoder: encoder,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implCreateUserClient struct {
	rootURL *url.URL
	client HTTPClient
	actor ActWithRequest
	resval ResponseValidation
	encoder CreateUserClientEncoder
	decoder CreateUserClientDecoder
}

// CreateUser makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implCreateUserClient) CreateUser(ctx context.Context, var1 users.NewUser) (users.User,error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(createuserServiceRoutePathURL)
	
	ctx = WithRequestMethod(ctx, "POST")
	ctx = WithClientRequestURI(ctx, targetURL.String())
	ctx = WithRequestTransport(ctx, "RPKIT:HTTP:CLIENT")

	var result users.User
	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return result,err
	}
	req = req.WithContext(ctx)

	if header, err := CtxCustomHeader(ctx); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return result, err
		}
	}

	if requestFacedInternalIssues(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrServerInternalProblem
	}

	if requestFailed(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrBadRequest
	}

	if requestRedirected(res) {
		jsonErr, err := loadJSONError(res.Body)
		if err == nil {
			return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrServerRejectedRequest
	}

	result, err = imp.decoder.Decode(ctx, res.Body)
	if err != nil {
		return result, err
	}

	return result, nil
}

// CreateUserServer implements a http.Handler for servicing the method CreateUser
// from users.UserService.
type CreateUserServer interface {
	http.Handler
}

type implCreateUserHandler struct{
	hook Hook
	headers http.Header
	service CreateUserMethodService
}

// NewCreateUserServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.CreateUser.
func NewCreateUserServer(service CreateUserMethodService, hook Hook, headers http.Header) CreateUserServer {
	return implCreateUserHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "CreateUser".
func (impl implCreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithAcceptsType(ctx, r.Header.Get("Accept"))
	ctx = WithContentType(ctx, r.Header.Get("Content-Type"))
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "CreateUser")
	ctx = WithServiceMethodPath(ctx, CreateUserServiceRoute)
	ctx = WithServiceMethodRoute(ctx, CreateUserServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, MethodTypeError,  http.StatusMethodNotAllowed, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "CreateUser",
			"api_service": MethodServiceName,
			"route": CreateUserServiceRoute,
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

		if impl.hook != nil {
			impl.hook.RequestAccepted(ctx)
			impl.hook.RequestProcessed(ctx)
			impl.hook.ResponsePrepared(ctx)
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "CreateUser")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", CreateUserServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)

		if impl.hook != nil {
			impl.hook.ResponseSent(ctx)
		}
		return
	}

	if !strings.HasSuffix(prefixIfNeed(r.URL.Path, "/"), CreateUserServiceRoutePath) {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, URLError,  http.StatusBadRequest, "only POST request to "+CreateUserServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "CreateUser",
			"api_service": MethodServiceName,
			"route": CreateUserServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.headers != nil {
		if accepts := impl.headers.Get("Accept"); accepts != "" {
			ctype := r.Header.Get("Content-Type")
			if !strings.Contains(accepts, ctype){
				if impl.hook != nil {
					impl.hook.RequestRejected(ctx)
				}

				jsonWriteError(w, AcceptTypeUnknownError,  http.StatusBadRequest, "request content type not supported",
				ErrInvalidContentType, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "CreateUser",
					"api_service": MethodServiceName,
					"route": CreateUserServiceRoute,
					"api": "users.UserService",
				})
				return
			}
		}
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
	w.Header().Add("X-Method", "CreateUser")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", CreateUserServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")


	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				if impl.hook != nil {
					impl.hook.RequestPanic(ctx, rerr)
				}

				derr := fmt.Errorf("panic err: %#v", rerr)
				jsonWriteError(w, ActionPanicError,  http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "CreateUser",
					"api_service": MethodServiceName,
					"route": CreateUserServiceRoute,
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
		if impl.hook != nil {
			impl.hook.RequestError(ctx, actionErr)
		}

		jsonWriteError(w, ActionError,  http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "CreateUser",
			"api_service": MethodServiceName,
			"route": CreateUserServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	w.WriteHeader(http.StatusOK)

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// RP: Input And Output Returning Error methods
// Method: GetUser
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetUser
//****************************************************************************

// GetUserServiceRoute defines the route for the GetUser method.
const GetUserServiceRoute = "users.UserService/GetUser"

// GetUserServiceRoutePath defines the full method path for the GetUser method.
const GetUserServiceRoutePath = "/users/UserService/GetUser"

// getuserServiceRoutePathURL defines a parsed path for the GetUser, it
// ensures the created path is valid as a url.
var getuserServiceRoutePathURL = mustSimpleParseURL(GetUserServiceRoutePath)

// GetUserContractSource contains the source version of expected method contract.
const GetUserContractSource = `type GetUserMethodContract interface {
	GetUser(var1 int)  (users.User,error)  
}
`

// GetUserMethodContract defines a contract interface for method "GetUser"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type GetUserMethodContract interface{
	GetUser(var1 int)  (users.User,error)  
}

// GetUserDecoder defines a interface which expose a single method to decode the request data
// expected by GetUserMethodContract.GetUser.
type GetUserDecoder interface{
	Decode(context.Context, io.Reader) (int, error)
}

// GetUserEncoder defines a interface which expose a single method to encode the response
// returned by GetUserMethodContract.GetUser.
type GetUserEncoder interface{
	Encode(context.Context, io.Writer, users.User) error
}

// GetUserMethodService defines the returned signature by the ServiceGetUser
// which executes it's internal behaviour based off on it's GetUserMethodContract.
type GetUserMethodService func(context.Context, io.Writer, io.Reader) error

// ServeGetUserMethod implements the core contract behaviour to service "GetUser"
// from "users.UserService" in "github.com/gokit/rpkit/examples/users".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeGetUserMethod(provider GetUserMethodContract, encoder GetUserEncoder, decoder GetUserDecoder) GetUserMethodService {
	return func(ctx context.Context, w io.Writer, r io.Reader) error {
		input, err := decoder.Decode(ctx, r)
		if err != nil {
			return err
		}

		
		res, err := provider.GetUser(input)
		
		if err != nil {
			return err
		}

		return encoder.Encode(ctx, w, res)
	}
}

// GetUserClientEncoder defines a interface which expose a single method to encode the request
// data sent by GetUserClientContract.GetUser.
type GetUserClientEncoder interface{
	Encode(context.Context, io.Writer, int) error
}

// GetUserClientDecoder defines a interface which expose a single method to decode the response
// returned by GetUserClientContract.GetUser.
type GetUserClientDecoder interface{
	Decode(context.Context, io.Reader) (users.User, error)
}

// GetUserClientContract defines a contract interface for clients to make request to GetUserServer.
type GetUserClientContract interface {
	GetUser(ctx context.Context, var1 int) (users.User,error)
}

// NewGetUserMethodClient returns a new GetUserMethodContract it relies on
// NewGetUserMethodContractClient.
func NewGetUserMethodClient(addr string, client HTTPClient, encoder GetUserClientEncoder, decoder GetUserClientDecoder) (GetUserClientContract, error) {
	return NewGetUserMethodContractClient(addr, client, encoder, decoder, nil, nil)
}

// NewGetUserMethodContractClient returns a new GetUserMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a GetUserServer to
// perform action as specified by contract.
func NewGetUserMethodContractClient(addr string, client HTTPClient, encoder GetUserClientEncoder, decoder GetUserClientDecoder, act ActWithRequest, resv ResponseValidation) (GetUserClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implGetUserClient{
		actor: act,
		resval: resv,
		rootURL: root,
		encoder: encoder,
		decoder: decoder,
		client: skipRedirects(client),
	}, nil
}

type implGetUserClient struct {
	rootURL *url.URL
	client HTTPClient
	actor ActWithRequest
	resval ResponseValidation
	encoder GetUserClientEncoder
	decoder GetUserClientDecoder
}

// GetUser makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implGetUserClient) GetUser(ctx context.Context, var1 int) (users.User,error){
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(getuserServiceRoutePathURL)
	
	ctx = WithRequestMethod(ctx, "POST")
	ctx = WithClientRequestURI(ctx, targetURL.String())
	ctx = WithRequestTransport(ctx, "RPKIT:HTTP:CLIENT")

	var result users.User
	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return result,err
	}
	req = req.WithContext(ctx)

	if header, err := CtxCustomHeader(ctx); err == nil {
		for key, list := range header {
			req.Header[key] = append(req.Header[key], list...)
		}
	}

	if imp.actor != nil {
		imp.actor(req)
	}

	res, err := imp.client.Do(req)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return result, err
		}
	}

	if requestFacedInternalIssues(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrServerInternalProblem
	}

	if requestFailed(res) {
		jsonErr, err := loadJSONError(res.Body)
	   if err == nil {
		return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrBadRequest
	}

	if requestRedirected(res) {
		jsonErr, err := loadJSONError(res.Body)
		if err == nil {
			return result, jsonErr
	    }
	    log.Printf("Failed to load JSONErrorResponse from server: %+q\n", err)
		return result,ErrServerRejectedRequest
	}

	result, err = imp.decoder.Decode(ctx, res.Body)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetUserServer implements a http.Handler for servicing the method GetUser
// from users.UserService.
type GetUserServer interface {
	http.Handler
}

type implGetUserHandler struct{
	hook Hook
	headers http.Header
	service GetUserMethodService
}

// NewGetUserServer returns a new instance of the HTTPHandler which services all
// http requests for giving method users.UserService.GetUser.
func NewGetUserServer(service GetUserMethodService, hook Hook, headers http.Header) GetUserServer {
	return implGetUserHandler{
		hook: hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "GetUser".
func (impl implGetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	ctx = WithRequest(ctx, r)
	ctx = WithResponseWriter(ctx, w)
	ctx = WithCustomHeader(ctx, impl.headers)
	ctx = WithRequestMethod(ctx, r.Method)
	ctx = WithRequestHeader(ctx, r.Header)
	ctx = WithAcceptsType(ctx, r.Header.Get("Accept"))
	ctx = WithContentType(ctx, r.Header.Get("Content-Type"))
	ctx = WithServiceName(ctx, BaseServiceName)
	ctx = WithServicePath(ctx, MethodServiceName)
	ctx = WithServicePackage(ctx, ServiceCodePath)
	ctx = WithServicePackageName(ctx, ServiceCodePathName)
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/users")
	ctx = WithServiceSourcePackageName(ctx, "users")
	ctx = WithServiceMethodName(ctx, "GetUser")
	ctx = WithServiceMethodPath(ctx, GetUserServiceRoute)
	ctx = WithServiceMethodRoute(ctx, GetUserServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, MethodTypeError,  http.StatusMethodNotAllowed, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetUser",
			"api_service": MethodServiceName,
			"route": GetUserServiceRoute,
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

		if impl.hook != nil {
			impl.hook.RequestAccepted(ctx)
			impl.hook.RequestProcessed(ctx)
			impl.hook.ResponsePrepared(ctx)
		}

		w.Header().Add("X-Agent", "RPKIT")
		w.Header().Add("X-Service", BaseServiceName)
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/users")
		w.Header().Add("X-Method", "GetUser")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", GetUserServiceRoute)
		w.Header().Add("X-Package-Interface", "users.UserService")

		w.WriteHeader(http.StatusNoContent)

		if impl.hook != nil {
			impl.hook.ResponseSent(ctx)
		}
		return
	}

	if !strings.HasSuffix(prefixIfNeed(r.URL.Path, "/"), GetUserServiceRoutePath) {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, URLError,  http.StatusBadRequest, "only POST request to "+GetUserServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetUser",
			"api_service": MethodServiceName,
			"route": GetUserServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.headers != nil {
		if accepts := impl.headers.Get("Accept"); accepts != "" {
			ctype := r.Header.Get("Content-Type")
			if !strings.Contains(accepts, ctype){
				if impl.hook != nil {
					impl.hook.RequestRejected(ctx)
				}

				jsonWriteError(w, AcceptTypeUnknownError,  http.StatusBadRequest, "request content type not supported",
				ErrInvalidContentType, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "GetUser",
					"api_service": MethodServiceName,
					"route": GetUserServiceRoute,
					"api": "users.UserService",
				})
				return
			}
		}
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
	w.Header().Add("X-Method", "GetUser")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", GetUserServiceRoute)
	w.Header().Add("X-Package-Interface", "users.UserService")


	var actionErr error
	func(){
		defer func(){
			if rerr := recover(); rerr != nil {
				if impl.hook != nil {
					impl.hook.RequestPanic(ctx, rerr)
				}

				derr := fmt.Errorf("panic err: %#v", rerr)
				jsonWriteError(w, ActionPanicError,  http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package": "github.com/gokit/rpkit/examples/users",
					"api_base": BaseServiceName,
					"method": "GetUser",
					"api_service": MethodServiceName,
					"route": GetUserServiceRoute,
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
		if impl.hook != nil {
			impl.hook.RequestError(ctx, actionErr)
		}

		jsonWriteError(w, ActionError,  http.StatusBadRequest, "method call returned err", actionErr, map[string]interface{}{
			"package": "github.com/gokit/rpkit/examples/users",
			"api_base": BaseServiceName,
			"method": "GetUser",
			"api_service": MethodServiceName,
			"route": GetUserServiceRoute,
			"api": "users.UserService",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponsePrepared(ctx)
	}

	w.WriteHeader(http.StatusOK)

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

// mustSimpleParseURL attempts to parse provided access using simpleParseURL, it panics
// if an error occured.
func mustSimpleParseURL(addr string) *url.URL {
	parsed, err := simpleParseURL(addr)
	if err == nil { return parsed }
	panic(`failed to parse url: `+addr+` , error occurred: `+err.Error())
}

// simpleParseURL parses giving address returning *url.URL instance for address if
// it's a valid address. Returns an error if address is an invalid URL.
func simpleParseURL(addr string) (*url.URL, error) {
	parsed, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return parsed, nil
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
func skipRedirects(in HTTPClient) HTTPClient {
	if rin, ok := in.(*http.Client); ok {
		former := rin.CheckRedirect
		copy := *rin
		copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if former != nil {
				err := former(req, via)
				_ = err // to fix issue with warning about not checking error.
			}
			return http.ErrUseLastResponse
		}
		return &copy
	}
	return in
}

func loadJSONError(r io.Reader) (JSONErrorResponse, error) {
	var res JSONErrorResponse
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}

func jsonWriteError(w http.ResponseWriter, itype string, code int, message string, err error, meta map[string]interface{}){
	var res JSONErrorResponse
	res.Type = itype
	res.Code = code
	res.Meta = meta
	res.Message = message

	if err != nil {
	  if data, badErr := json.Marshal(err); badErr == nil {
	    res.Err = data
	  }else{
	    res.Err = []byte(err.Error())
	  }
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	data, err2 := json.Marshal(res);
	if err2 == nil {
		w.Write(data)
		return
	}

	log.Printf("unable to send error response for error %+q: %+q", err, err2)
}

func prefixIfNeed(s string, p string) string {
    if strings.HasPrefix(s, p){
        return s
    }
    return p+s
}