package bottlecellarrp

// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	cellar "github.com/gokit/rpkit/examples/bottles/cellar"
)

const (
	// BaseServiceName defines the base name of service root.
	BaseServiceName = "gokit.rpkit.examples.bottles"

	// MethodServiceName defines the complete name of this giving API service.
	MethodServiceName = "gokit.rpkit.examples.bottles.BottleCellar"

	// ServiceCodePath defines the path to this generated package which contains the implemented service methods.
	ServiceCodePath = "github.com/gokit/rpkit/examples/bottles/bottlecellarrp"

	// ServiceCodePathName defines the name giving to
	// this package.
	ServiceCodePathName = "bottlecellarrp"
)

// errors ...
var (
	ErrNotInContext          = errors.New("item not in context")
	ErrInvalidRequestURI     = errors.New("invalid request uri")
	ErrBadRequest            = errors.New("server rejected request as bad")
	ErrInvalidRequestMethod  = errors.New("invalid request method")
	ErrServerInternalProblem = errors.New("server had internal problems")
	ErrServerRejectedRequest = errors.New("server does not handle request type/method")
)

//****************************************************************************
// Context Keys, Setters And Getters
//****************************************************************************

// sets of context keys used by the package.
const (
	contextCustomHeaderKey             = "rp:custom:header"
	contextRequestKey                  = "rp:request"
	contextRequestMethodKey            = "rp:request:method"
	contextRequestHeaderKey            = "rp:request:header"
	contextResponseWriterKey           = "rp:response:writer"
	contextServiceNameKey              = "rp:service:name"
	contextServicePathKey              = "rp:service:path"
	contextServicePackageKey           = "rp:service:package"
	contextServicePackageNameKey       = "rp:service:package:name"
	contextServiceSourcePackageKey     = "rp:service:source:path"
	contextServiceSourcePackageNameKey = "rp:service:source:package"
	contextServiceMethodNameKey        = "rp:service:method:name"
	contextServiceMethodPathKey        = "rp:service:method:path"
	contextServiceMethodRouteKey       = "rp:service:method:route"
	contextClientRequestURLKey         = "rp:client:request:url"
	contextRequestTransportKey         = "rp:request:transport"
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
// Source: github.com/gokit/rpkit/examples/bottles
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
//
//@implement
type Hook interface {
	ResponseSent(context.Context)
	ResponsePrepared(context.Context)
	RequestRejected(context.Context)
	RequestAccepted(context.Context)
	RequestReceived(context.Context)
	RequestProcessed(context.Context)
}

//****************************************************************************
// RP: Output Returning No Error methods
// Method: Cellar
// Source: github.com/gokit/rpkit/examples/bottles
// Handler: bottles.BottleCellar.Cellar
//****************************************************************************

// CellarServiceRoute defines the route for the Cellar method.
const CellarServiceRoute = "gokit.rpkit.examples.bottles.BottleCellar/Cellar"

// CellarServiceRoutePath defines the full method path for the Cellar method.
const CellarServiceRoutePath = "/gokit.rpkit.examples.bottles.BottleCellar/Cellar"

// cellarServiceRoutePathURL defines a parsed path for the Cellar, it
// ensures the created path is valid as a url.
var cellarServiceRoutePathURL = mustParseURL(CellarServiceRoutePath)

// CellarContractSource contains the source version of expected method contract.
const CellarContractSource = `type CellarMethodContract interface {
	Cellar()  cellar.Cellar  
}
`

// CellarMethodContract defines a contract interface for method "Cellar"
// provided by "bottles.BottleCellar" in "github.com/gokit/rpkit/examples/bottles". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type CellarMethodContract interface {
	Cellar() cellar.Cellar
}

// CellarEncoder defines a interface which expose a single method to encode the response
// returned by CellarMethodContract.Cellar.
type CellarEncoder interface {
	Encode(io.Writer, cellar.Cellar) error
}

// CellarMethodService defines the returned signature by the ServiceCellar
// which executes it's internal behaviour based off on it's CellarMethodContract.
type CellarMethodService func(context.Context, io.Writer) error

// ServeCellarMethod implements the core contract behaviour to service "Cellar"
// from "bottles.BottleCellar" in "github.com/gokit/rpkit/examples/bottles".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeCellarMethod(provider CellarMethodContract, encoder CellarEncoder) CellarMethodService {
	return func(ctx context.Context, w io.Writer) error {

		res := provider.Cellar()

		return encoder.Encode(w, res)
	}
}

// CellarClientContract defines a contract interface for clients to make request to CellarServer.
type CellarClientContract interface {
	Cellar(ctx context.Context) (cellar.Cellar, error)
}

// CellarClientDecoder defines a interface which expose a single method to decode the response
// returned by CellarClientContract.Cellar.
type CellarClientDecoder interface {
	Decode(io.Reader) (cellar.Cellar, error)
}

// NewCellarMethodClient returns a new CellarMethodContract it relies on
// NewCellarMethodContractClient.
func NewCellarMethodClient(addr string, client *http.Client, encoder CellarClientDecoder) (CellarClientContract, error) {
	return NewCellarMethodContractClient(addr, client, encoder, nil, nil)
}

// NewCellarMethodContractClient returns a new CellarMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a CellarServer to
// perform action as specified by contract.
func NewCellarMethodContractClient(addr string, client *http.Client, decoder CellarClientDecoder, act ActWithRequest, resv ResponseValidation) (CellarClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implCellarClient{
		actor:   act,
		resval:  resv,
		rootURL: root,
		decoder: decoder,
		client:  skipRedirects(client),
	}, nil
}

type implCellarClient struct {
	rootURL *url.URL
	client  *http.Client
	actor   ActWithRequest
	resval  ResponseValidation
	decoder CellarClientDecoder
}

// Cellar makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implCellarClient) Cellar(ctx context.Context) (cellar.Cellar, error) {
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(cellarServiceRoutePathURL)

	ctx = WithRequestMethod(ctx, "POST")
	ctx = WithClientRequestURI(ctx, targetURL.String())
	ctx = WithRequestTransport(ctx, "RPKIT:HTTP:CLIENT")

	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return cellar.Cellar{}, err
	}

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
		return cellar.Cellar{}, err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return cellar.Cellar{}, err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return cellar.Cellar{}, jsonErr
		}
		return cellar.Cellar{}, ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return cellar.Cellar{}, jsonErr
		}
		return cellar.Cellar{}, ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return cellar.Cellar{}, jsonErr
		}
		return cellar.Cellar{}, ErrServerRejectedRequest
	}

	rec, err := imp.decoder.Decode(res.Body)
	if err != nil {
		return cellar.Cellar{}, err
	}

	return rec, nil
}

// CellarServer implements a http.Handler for servicing the method Cellar
// from bottles.BottleCellar.
type CellarServer interface {
	http.Handler
}

type implCellarHandler struct {
	hook    Hook
	headers http.Header
	service CellarMethodService
}

// NewCellarServer returns a new instance of the HTTPHandler which services all
// http requests for giving method bottles.BottleCellar.Cellar.
func NewCellarServer(service CellarMethodService, hook Hook, headers http.Header) CellarServer {
	return implCellarHandler{
		hook:    hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "Cellar".
func (impl implCellarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/bottles")
	ctx = WithServiceSourcePackageName(ctx, "bottles")
	ctx = WithServiceMethodName(ctx, "Cellar")
	ctx = WithServiceMethodPath(ctx, CellarServiceRoute)
	ctx = WithServiceMethodRoute(ctx, CellarServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "Cellar",
			"api_service": MethodServiceName,
			"route":       CellarServiceRoute,
			"api":         "bottles.BottleCellar",
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
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/bottles")
		w.Header().Add("X-Method", "Cellar")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", CellarServiceRoute)
		w.Header().Add("X-Package-Interface", "bottles.BottleCellar")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != CellarServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+CellarServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "Cellar",
			"api_service": MethodServiceName,
			"route":       CellarServiceRoute,
			"api":         "bottles.BottleCellar",
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
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/bottles")
	w.Header().Add("X-Method", "Cellar")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", CellarServiceRoute)
	w.Header().Add("X-Package-Interface", "bottles.BottleCellar")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func() {
		defer func() {
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package":     "github.com/gokit/rpkit/examples/bottles",
					"api_base":    BaseServiceName,
					"method":      "Cellar",
					"api_service": MethodServiceName,
					"route":       CellarServiceRoute,
					"api":         "bottles.BottleCellar",
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
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "Cellar",
			"api_service": MethodServiceName,
			"route":       CellarServiceRoute,
			"api":         "bottles.BottleCellar",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// RP: Input Returning Only Error methods
// Method: AddBottle
// Source: github.com/gokit/rpkit/examples/bottles
// Handler: bottles.BottleCellar.AddBottle
//****************************************************************************

// AddBottleServiceRoute defines the route for the AddBottle method.
const AddBottleServiceRoute = "gokit.rpkit.examples.bottles.BottleCellar/AddBottle"

// AddBottleServiceRoutePath defines the full method path for the AddBottle method.
const AddBottleServiceRoutePath = "/gokit.rpkit.examples.bottles.BottleCellar/AddBottle"

// addbottleServiceRoutePathURL defines a parsed path for the AddBottle, it
// ensures the created path is valid as a url.
var addbottleServiceRoutePathURL = mustParseURL(AddBottleServiceRoutePath)

// AddBottleContractSource contains the source version of expected method contract.
const AddBottleContractSource = `type AddBottleMethodContract interface {
	AddBottle(var1 string)  error  
}
`

// AddBottleMethodContract defines a contract interface for method "AddBottle"
// provided by "bottles.BottleCellar" in "github.com/gokit/rpkit/examples/bottles". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type AddBottleMethodContract interface {
	AddBottle(var1 string) error
}

// AddBottleDecoder defines a interface which expose a single method to decode the request data
// expected by AddBottleMethodContract.AddBottle.
type AddBottleDecoder interface {
	Decode(io.Reader) (string, error)
}

// AddBottleMethodService defines the returned signature by the ServiceAddBottle
// which executes it's internal behaviour based off on it's AddBottleMethodContract.
type AddBottleMethodService func(context.Context, io.Reader) error

// ServeAddBottleMethod implements the core contract behaviour to service "AddBottle"
// from "bottles.BottleCellar" in "github.com/gokit/rpkit/examples/bottles".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeAddBottleMethod(provider AddBottleMethodContract, decoder AddBottleDecoder) AddBottleMethodService {
	return func(ctx context.Context, r io.Reader) error {
		input, err := decoder.Decode(r)
		if err != nil {
			return err
		}

		return provider.AddBottle(input)

	}
}

// AddBottleClientEncoder defines a interface which expose a single method to encode the request
// data sent by AddBottleClientContract.AddBottle.
type AddBottleClientEncoder interface {
	Encode(io.Writer, string) error
}

// AddBottleClientContract defines a contract interface for clients to make request to AddBottleServer.
type AddBottleClientContract interface {
	AddBottle(ctx context.Context, var1 string) error
}

// NewAddBottleMethodClient returns a new AddBottleMethodContract it relies on
// NewAddBottleMethodContractClient.
func NewAddBottleMethodClient(addr string, client *http.Client, encoder AddBottleClientEncoder) (AddBottleClientContract, error) {
	return NewAddBottleMethodContractClient(addr, client, encoder, nil, nil)
}

// NewAddBottleMethodContractClient returns a new AddBottleMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a AddBottleServer to
// perform action as specified by contract.
func NewAddBottleMethodContractClient(addr string, client *http.Client, encoder AddBottleClientEncoder, act ActWithRequest, resv ResponseValidation) (AddBottleClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implAddBottleClient{
		actor:   act,
		resval:  resv,
		rootURL: root,
		encoder: encoder,
		client:  skipRedirects(client),
	}, nil
}

type implAddBottleClient struct {
	rootURL *url.URL
	client  *http.Client
	actor   ActWithRequest
	resval  ResponseValidation
	encoder AddBottleClientEncoder
}

// AddBottle makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implAddBottleClient) AddBottle(ctx context.Context, var1 string) error {
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(addbottleServiceRoutePathURL)

	ctx = WithRequestMethod(ctx, "POST")
	ctx = WithClientRequestURI(ctx, targetURL.String())
	ctx = WithRequestTransport(ctx, "RPKIT:HTTP:CLIENT")

	req, err := http.NewRequest("POST", targetURL.String(), nil)
	if err != nil {
		return err
	}

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

// AddBottleServer implements a http.Handler for servicing the method AddBottle
// from bottles.BottleCellar.
type AddBottleServer interface {
	http.Handler
}

type implAddBottleHandler struct {
	hook    Hook
	headers http.Header
	service AddBottleMethodService
}

// NewAddBottleServer returns a new instance of the HTTPHandler which services all
// http requests for giving method bottles.BottleCellar.AddBottle.
func NewAddBottleServer(service AddBottleMethodService, hook Hook, headers http.Header) AddBottleServer {
	return implAddBottleHandler{
		hook:    hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "AddBottle".
func (impl implAddBottleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	ctx = WithServiceSourcePackage(ctx, "github.com/gokit/rpkit/examples/bottles")
	ctx = WithServiceSourcePackageName(ctx, "bottles")
	ctx = WithServiceMethodName(ctx, "AddBottle")
	ctx = WithServiceMethodPath(ctx, AddBottleServiceRoute)
	ctx = WithServiceMethodRoute(ctx, AddBottleServiceRoutePath)

	if impl.hook != nil {
		impl.hook.RequestReceived(ctx)
	}

	if r.Method != "POST" && r.Method != "HEAD" {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST or HEAD request allowed", ErrInvalidRequestMethod, map[string]interface{}{
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "AddBottle",
			"api_service": MethodServiceName,
			"route":       AddBottleServiceRoute,
			"api":         "bottles.BottleCellar",
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
		w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/bottles")
		w.Header().Add("X-Method", "AddBottle")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", AddBottleServiceRoute)
		w.Header().Add("X-Package-Interface", "bottles.BottleCellar")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != AddBottleServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+AddBottleServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "AddBottle",
			"api_service": MethodServiceName,
			"route":       AddBottleServiceRoute,
			"api":         "bottles.BottleCellar",
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
	w.Header().Add("X-Package", "github.com/gokit/rpkit/examples/bottles")
	w.Header().Add("X-Method", "AddBottle")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", AddBottleServiceRoute)
	w.Header().Add("X-Package-Interface", "bottles.BottleCellar")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func() {
		defer func() {
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package":     "github.com/gokit/rpkit/examples/bottles",
					"api_base":    BaseServiceName,
					"method":      "AddBottle",
					"api_service": MethodServiceName,
					"route":       AddBottleServiceRoute,
					"api":         "bottles.BottleCellar",
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
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "AddBottle",
			"api_service": MethodServiceName,
			"route":       AddBottleServiceRoute,
			"api":         "bottles.BottleCellar",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// Utils
// Source: github.com/gokit/rpkit/examples/bottles
//****************************************************************************

func requestRedirected(res *http.Response) bool {
	if res.StatusCode >= 300 && res.StatusCode <= 308 {
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
	if err == nil {
		return parsed
	}
	panic(`failed to parse url: ` + addr + ` , error occurred: ` + err.Error())
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

	if parsed.IsAbs() {
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

func jsonWriteError(w http.ResponseWriter, code int, message string, err error, meta map[string]interface{}) {
	var res jsonErrorResponse
	res.Code = code
	res.Err = err
	res.Meta = meta
	res.Message = message

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	data, err2 := json.Marshal(res)
	if err2 == nil {
		w.Write(data)
		return
	}

	log.Printf("unable to send error response for error %+q: %+q", err, err2)
}

type jsonErrorResponse struct {
	Code    int                    `json:"code"`
	Err     error                  `json:"err"`
	Message string                 `json:"message"`
	Meta    map[string]interface{} `json:"meta"`
}

func (jse jsonErrorResponse) Error() string {
	return jse.Message + " " + jse.Err.Error()
}
