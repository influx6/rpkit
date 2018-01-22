package userservicerp

// All code below are auto-generated and should not be edited by handle.
// See https://github.com/gokit/rpkit for more info.
  
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"errors"
	"log"

	"github.com/gokit/rpkit/examples/users"
)

const (
	// BaseServiceName defines the base name of service root.
	BaseServiceName = "gokit.rpkit.examples.users"

	// MethodServiceName defines the complete name of this giving API service.
	MethodServiceName = "gokit.rpkit.examples.users.UserService"
)

// errors ...
var (
	ErrInvalidRequestURI = errors.New("invalid request uri")
	ErrInvalidRequestMethod = errors.New("invalid request method")
)

//****************************************************************************
// Types
// Source: github.com/gokit/rpkit/examples/users
//****************************************************************************

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

// MakeAdminServer implements a http.Handler for servicing the method MakeAdmin
// from users.UserService.
type MakeAdminServer interface {
	http.Handler
}

type implMakeAdminHandler struct{
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

		w.Header().Add("X-Agent", "RPKIT")
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
				jsonWriteError(w, http.StatusBadRequest, "panic occured with method run", derr, map[string]interface{}{
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
				jsonWriteError(w, http.StatusBadRequest, "panic occured with method run", derr, map[string]interface{}{
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
// RP: Input Returning Only Error methods
// Method: Delete
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Delete
//****************************************************************************

// DeleteServiceRoute defines the route for the Delete method.
const DeleteServiceRoute = "gokit.rpkit.examples.users.UserService/Delete"

// DeleteServiceRoutePath defines the full method path for the Delete method.
const DeleteServiceRoutePath = "/gokit.rpkit.examples.users.UserService/Delete"

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
				jsonWriteError(w, http.StatusBadRequest, "panic occured with method run", derr, map[string]interface{}{
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
				jsonWriteError(w, http.StatusBadRequest, "panic occured with method run", derr, map[string]interface{}{
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
				jsonWriteError(w, http.StatusBadRequest, "panic occured with method run", derr, map[string]interface{}{
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
				jsonWriteError(w, http.StatusBadRequest, "panic occured with method run", derr, map[string]interface{}{
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
				jsonWriteError(w, http.StatusBadRequest, "panic occured with method run", derr, map[string]interface{}{
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

// willRedirect returns true/false if giving code is a redirect status.
func willRedirect(code int) bool {
	return code >= 301 && code <= 308
}

// isRedirect returns true/false if giving code is a redirect status.
func isRedirect(code int) bool {
	return code == http.StatusTemporaryRedirect || code == http.StatusPermanentRedirect
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
