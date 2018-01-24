package iboxcellarrp

// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gokit/rpkit/examples/bottles"
)

const (
	// BaseServiceName defines the base name of service root.
	BaseServiceName = "gokit.rpkit.examples.bottles"

	// MethodServiceName defines the complete name of this giving API service.
	MethodServiceName = "gokit.rpkit.examples.bottles.IBoxCellar"

	// ServiceCodePath defines the path to this generated package which contains the implemented service methods.
	ServiceCodePath = "github.com/gokit/rpkit/examples/bottles/iboxcellarrp"

	// ServiceCodePathName defines the name giving to
	// this package.
	ServiceCodePathName = "iboxcellarrp"
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

// HTTPClient defines an interface which is used as the processor for a
// raw http request.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

//****************************************************************************
// RP: Input Returning Only Error methods
// Method: StoreBottle
// Source: github.com/gokit/rpkit/examples/bottles
// Handler: bottles.IBoxCellar.StoreBottle
//****************************************************************************

// StoreBottleServiceRoute defines the route for the StoreBottle method.
const StoreBottleServiceRoute = "gokit.rpkit.examples.bottles.IBoxCellar/StoreBottle"

// StoreBottleServiceRoutePath defines the full method path for the StoreBottle method.
const StoreBottleServiceRoutePath = "/gokit.rpkit.examples.bottles.IBoxCellar/StoreBottle"

// storebottleServiceRoutePathURL defines a parsed path for the StoreBottle, it
// ensures the created path is valid as a url.
var storebottleServiceRoutePathURL = mustSimpleParseURL(StoreBottleServiceRoutePath)

// StoreBottleContractSource contains the source version of expected method contract.
const StoreBottleContractSource = `type StoreBottleMethodContract interface {
	StoreBottle(var1 context.Context,var2 bottles.IBox)  error  
}
`

// StoreBottleMethodContract defines a contract interface for method "StoreBottle"
// provided by "bottles.IBoxCellar" in "github.com/gokit/rpkit/examples/bottles". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type StoreBottleMethodContract interface {
	StoreBottle(var1 context.Context, var2 bottles.IBox) error
}

// StoreBottleDecoder defines a interface which expose a single method to decode the request data
// expected by StoreBottleMethodContract.StoreBottle.
type StoreBottleDecoder interface {
	Decode(io.Reader) (bottles.IBox, error)
}

// StoreBottleMethodService defines the returned signature by the ServiceStoreBottle
// which executes it's internal behaviour based off on it's StoreBottleMethodContract.
type StoreBottleMethodService func(context.Context, io.Reader) error

// ServeStoreBottleMethod implements the core contract behaviour to service "StoreBottle"
// from "bottles.IBoxCellar" in "github.com/gokit/rpkit/examples/bottles".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeStoreBottleMethod(provider StoreBottleMethodContract, decoder StoreBottleDecoder) StoreBottleMethodService {
	return func(ctx context.Context, r io.Reader) error {
		input, err := decoder.Decode(r)
		if err != nil {
			return err
		}

		return provider.StoreBottle(ctx, input)

	}
}

// StoreBottleClientEncoder defines a interface which expose a single method to encode the request
// data sent by StoreBottleClientContract.StoreBottle.
type StoreBottleClientEncoder interface {
	Encode(io.Writer, bottles.IBox) error
}

// StoreBottleClientContract defines a contract interface for clients to make request to StoreBottleServer.
type StoreBottleClientContract interface {
	StoreBottle(var1 context.Context, var2 bottles.IBox) error
}

// NewStoreBottleMethodClient returns a new StoreBottleMethodContract it relies on
// NewStoreBottleMethodContractClient.
func NewStoreBottleMethodClient(addr string, client HTTPClient, encoder StoreBottleClientEncoder) (StoreBottleClientContract, error) {
	return NewStoreBottleMethodContractClient(addr, client, encoder, nil, nil)
}

// NewStoreBottleMethodContractClient returns a new StoreBottleMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a StoreBottleServer to
// perform action as specified by contract.
func NewStoreBottleMethodContractClient(addr string, client HTTPClient, encoder StoreBottleClientEncoder, act ActWithRequest, resv ResponseValidation) (StoreBottleClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implStoreBottleClient{
		actor:   act,
		resval:  resv,
		rootURL: root,
		encoder: encoder,
		client:  skipRedirects(client),
	}, nil
}

type implStoreBottleClient struct {
	rootURL *url.URL
	client  HTTPClient
	actor   ActWithRequest
	resval  ResponseValidation
	encoder StoreBottleClientEncoder
}

// StoreBottle makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implStoreBottleClient) StoreBottle(var1 context.Context, var2 bottles.IBox) error {
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(storebottleServiceRoutePathURL)

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

// StoreBottleServer implements a http.Handler for servicing the method StoreBottle
// from bottles.IBoxCellar.
type StoreBottleServer interface {
	http.Handler
}

type implStoreBottleHandler struct {
	hook    Hook
	headers http.Header
	service StoreBottleMethodService
}

// NewStoreBottleServer returns a new instance of the HTTPHandler which services all
// http requests for giving method bottles.IBoxCellar.StoreBottle.
func NewStoreBottleServer(service StoreBottleMethodService, hook Hook, headers http.Header) StoreBottleServer {
	return implStoreBottleHandler{
		hook:    hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "StoreBottle".
func (impl implStoreBottleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	ctx = WithServiceMethodName(ctx, "StoreBottle")
	ctx = WithServiceMethodPath(ctx, StoreBottleServiceRoute)
	ctx = WithServiceMethodRoute(ctx, StoreBottleServiceRoutePath)

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
			"method":      "StoreBottle",
			"api_service": MethodServiceName,
			"route":       StoreBottleServiceRoute,
			"api":         "bottles.IBoxCellar",
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
		w.Header().Add("X-Method", "StoreBottle")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", StoreBottleServiceRoute)
		w.Header().Add("X-Package-Interface", "bottles.IBoxCellar")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != StoreBottleServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+StoreBottleServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "StoreBottle",
			"api_service": MethodServiceName,
			"route":       StoreBottleServiceRoute,
			"api":         "bottles.IBoxCellar",
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
	w.Header().Add("X-Method", "StoreBottle")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", StoreBottleServiceRoute)
	w.Header().Add("X-Package-Interface", "bottles.IBoxCellar")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func() {
		defer func() {
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package":     "github.com/gokit/rpkit/examples/bottles",
					"api_base":    BaseServiceName,
					"method":      "StoreBottle",
					"api_service": MethodServiceName,
					"route":       StoreBottleServiceRoute,
					"api":         "bottles.IBoxCellar",
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
			"method":      "StoreBottle",
			"api_service": MethodServiceName,
			"route":       StoreBottleServiceRoute,
			"api":         "bottles.IBoxCellar",
		})
		return
	}

	if impl.hook != nil {
		impl.hook.ResponseSent(ctx)
	}
}

//****************************************************************************
// RP: Input And Output Returning Error methods
// Method: GetBox
// Source: github.com/gokit/rpkit/examples/bottles
// Handler: bottles.IBoxCellar.GetBox
//****************************************************************************

// GetBoxServiceRoute defines the route for the GetBox method.
const GetBoxServiceRoute = "gokit.rpkit.examples.bottles.IBoxCellar/GetBox"

// GetBoxServiceRoutePath defines the full method path for the GetBox method.
const GetBoxServiceRoutePath = "/gokit.rpkit.examples.bottles.IBoxCellar/GetBox"

// getboxServiceRoutePathURL defines a parsed path for the GetBox, it
// ensures the created path is valid as a url.
var getboxServiceRoutePathURL = mustSimpleParseURL(GetBoxServiceRoutePath)

// GetBoxContractSource contains the source version of expected method contract.
const GetBoxContractSource = `type GetBoxMethodContract interface {
	GetBox(var1 context.Context,var2 string)  (bottles.IBox,error)  
}
`

// GetBoxMethodContract defines a contract interface for method "GetBox"
// provided by "bottles.IBoxCellar" in "github.com/gokit/rpkit/examples/bottles". It allows us
// establish a simple contract suitable for meeting the needs of said method.
type GetBoxMethodContract interface {
	GetBox(var1 context.Context, var2 string) (bottles.IBox, error)
}

// GetBoxDecoder defines a interface which expose a single method to decode the request data
// expected by GetBoxMethodContract.GetBox.
type GetBoxDecoder interface {
	Decode(io.Reader) (string, error)
}

// GetBoxEncoder defines a interface which expose a single method to encode the response
// returned by GetBoxMethodContract.GetBox.
type GetBoxEncoder interface {
	Encode(io.Writer, bottles.IBox) error
}

// GetBoxMethodService defines the returned signature by the ServiceGetBox
// which executes it's internal behaviour based off on it's GetBoxMethodContract.
type GetBoxMethodService func(context.Context, io.Writer, io.Reader) error

// ServeGetBoxMethod implements the core contract behaviour to service "GetBox"
// from "bottles.IBoxCellar" in "github.com/gokit/rpkit/examples/bottles".
// It returns a function that can be used within any transport layer to process, to said execute
// behaviour of method as a service.
func ServeGetBoxMethod(provider GetBoxMethodContract, encoder GetBoxEncoder, decoder GetBoxDecoder) GetBoxMethodService {
	return func(ctx context.Context, w io.Writer, r io.Reader) error {
		input, err := decoder.Decode(r)
		if err != nil {
			return err
		}

		res, err := provider.GetBox(ctx, input)

		if err != nil {
			return err
		}

		return encoder.Encode(w, res)
	}
}

// GetBoxClientEncoder defines a interface which expose a single method to encode the request
// data sent by GetBoxClientContract.GetBox.
type GetBoxClientEncoder interface {
	Encode(io.Writer, string) error
}

// GetBoxClientDecoder defines a interface which expose a single method to decode the response
// returned by GetBoxClientContract.GetBox.
type GetBoxClientDecoder interface {
	Decode(io.Reader) (bottles.IBox, error)
}

// GetBoxClientContract defines a contract interface for clients to make request to GetBoxServer.
type GetBoxClientContract interface {
	GetBox(var1 context.Context, var2 string) (bottles.IBox, error)
}

// NewGetBoxMethodClient returns a new GetBoxMethodContract it relies on
// NewGetBoxMethodContractClient.
func NewGetBoxMethodClient(addr string, client HTTPClient, encoder GetBoxClientEncoder, decoder GetBoxClientDecoder) (GetBoxClientContract, error) {
	return NewGetBoxMethodContractClient(addr, client, encoder, decoder, nil, nil)
}

// NewGetBoxMethodContractClient returns a new GetBoxMethodContract implementation, which
// will make it's call to provided address with the provided http.Client to a GetBoxServer to
// perform action as specified by contract.
func NewGetBoxMethodContractClient(addr string, client HTTPClient, encoder GetBoxClientEncoder, decoder GetBoxClientDecoder, act ActWithRequest, resv ResponseValidation) (GetBoxClientContract, error) {
	root, err := parseURL(addr)
	if err != nil {
		return nil, err
	}

	return implGetBoxClient{
		actor:   act,
		resval:  resv,
		rootURL: root,
		encoder: encoder,
		decoder: decoder,
		client:  skipRedirects(client),
	}, nil
}

type implGetBoxClient struct {
	rootURL *url.URL
	client  HTTPClient
	actor   ActWithRequest
	resval  ResponseValidation
	encoder GetBoxClientEncoder
	decoder GetBoxClientDecoder
}

// GetBox makes a request to the server's address with provided arguments and returns
// response received from server.
func (imp implGetBoxClient) GetBox(var1 context.Context, var2 string) (bottles.IBox, error) {
	// targetURL for the requests to be made.
	targetURL := imp.rootURL.ResolveReference(getboxServiceRoutePathURL)

	var1 = WithRequestMethod(var1, "POST")
	var1 = WithClientRequestURI(var1, targetURL.String())
	var1 = WithRequestTransport(var1, "RPKIT:HTTP:CLIENT")

	var result bottles.IBox
	var body bytes.Buffer
	if err := imp.encoder.Encode(&body, var2); err != nil {
		return result, err
	}

	req, err := http.NewRequest("POST", targetURL.String(), &body)
	if err != nil {
		return result, err
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
		return result, err
	}

	defer res.Body.Close()

	if imp.resval != nil {
		if err := imp.resval(res); err != nil {
			return result, err
		}
	}

	if requestFacedInternalIssues(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return result, jsonErr
		}
		return result, ErrServerInternalProblem
	}

	if requestFailed(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return result, jsonErr
		}
		return result, ErrBadRequest
	}

	if requestRedirected(res) {
		if jsonErr, err := loadJSONError(res.Body); err == nil {
			return result, jsonErr
		}
		return result, ErrServerRejectedRequest
	}

	result, err = imp.decoder.Decode(res.Body)
	if err != nil {
		return result, err
	}

	return result, nil

}

// GetBoxServer implements a http.Handler for servicing the method GetBox
// from bottles.IBoxCellar.
type GetBoxServer interface {
	http.Handler
}

type implGetBoxHandler struct {
	hook    Hook
	headers http.Header
	service GetBoxMethodService
}

// NewGetBoxServer returns a new instance of the HTTPHandler which services all
// http requests for giving method bottles.IBoxCellar.GetBox.
func NewGetBoxServer(service GetBoxMethodService, hook Hook, headers http.Header) GetBoxServer {
	return implGetBoxHandler{
		hook:    hook,
		headers: headers,
		service: service,
	}
}

// ServeHTTP implements the http.Handler.ServeHTTP method and services requests for giving method "GetBox".
func (impl implGetBoxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	ctx = WithServiceMethodName(ctx, "GetBox")
	ctx = WithServiceMethodPath(ctx, GetBoxServiceRoute)
	ctx = WithServiceMethodRoute(ctx, GetBoxServiceRoutePath)

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
			"method":      "GetBox",
			"api_service": MethodServiceName,
			"route":       GetBoxServiceRoute,
			"api":         "bottles.IBoxCellar",
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
		w.Header().Add("X-Method", "GetBox")
		w.Header().Add("X-Method-Service", MethodServiceName)
		w.Header().Add("X-API-Route", GetBoxServiceRoute)
		w.Header().Add("X-Package-Interface", "bottles.IBoxCellar")

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path != GetBoxServiceRoutePath {
		if impl.hook != nil {
			impl.hook.RequestRejected(ctx)
		}

		jsonWriteError(w, http.StatusBadRequest, "only POST request to "+GetBoxServiceRoutePath+" allowed", ErrInvalidRequestURI, map[string]interface{}{
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "GetBox",
			"api_service": MethodServiceName,
			"route":       GetBoxServiceRoute,
			"api":         "bottles.IBoxCellar",
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
	w.Header().Add("X-Method", "GetBox")
	w.Header().Add("X-Method-Service", MethodServiceName)
	w.Header().Add("X-API-Route", GetBoxServiceRoute)
	w.Header().Add("X-Package-Interface", "bottles.IBoxCellar")

	w.WriteHeader(http.StatusOK)

	var actionErr error
	func() {
		defer func() {
			if rerr := recover(); rerr != nil {
				derr := fmt.Errorf("panic err: %+q", rerr)
				jsonWriteError(w, http.StatusInternalServerError, "panic occured with method run", derr, map[string]interface{}{
					"package":     "github.com/gokit/rpkit/examples/bottles",
					"api_base":    BaseServiceName,
					"method":      "GetBox",
					"api_service": MethodServiceName,
					"route":       GetBoxServiceRoute,
					"api":         "bottles.IBoxCellar",
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
			"package":     "github.com/gokit/rpkit/examples/bottles",
			"api_base":    BaseServiceName,
			"method":      "GetBox",
			"api_service": MethodServiceName,
			"route":       GetBoxServiceRoute,
			"api":         "bottles.IBoxCellar",
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

// mustSimpleParseURL attempts to parse provided access using simpleParseURL, it panics
// if an error occured.
func mustSimpleParseURL(addr string) *url.URL {
	parsed, err := simpleParseURL(addr)
	if err == nil {
		return parsed
	}
	panic(`failed to parse url: ` + addr + ` , error occurred: ` + err.Error())
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

	if parsed.IsAbs() {
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
