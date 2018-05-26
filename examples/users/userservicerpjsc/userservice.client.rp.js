// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.
  

const {URL} = require("universal-url");

// BaseServiceName defines the base name of service root.
export const BaseServiceName = "users"

// MethodServiceName defines the complete name of this giving API service.
export const MethodServiceName = "users/UserService"

// ServiceCodePath defines the path to this generated package which contains the implemented service methods.
export const ServiceCodePath = "github.com/gokit/rpkit/examples/users/userservicerp"

// ServiceCodePathName defines the name giving to
// this package.
export const ServiceCodePathName = "userservicerp"

// error type strings
export const URLError = "url_error"
export const ActionError = "action_error"
export const ActionPanicError = "action_panic_error"
export const MethodTypeError = "method_type_error"
export const AcceptTypeUnknownError = "accept_type_unknown_error"
export const RequestDecodingError = "request_decoding_error"
export const ResponseEncodingError = "response_encoding_error"

// JSONErrorResponse defines a structure to contain error message data
// delivered by the server.
export function JSONErrorResponse(type, code, err, message, meta){
    return {
        Type: type,
        Code: code,
        Err: err,
        Message: message,
        Meta: meta,
    }
}



///////////////////////////////////////////////////////////////////
// RP: Output with Returning Error methods
// Method: Get
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Get
///////////////////////////////////////////////////////////////////

// GetServiceRoute defines the route for the Get method.
export const GetServiceRoute = "users.UserService/Get"

// GetServiceRoutePath defines the full method path for the Get method.
export const GetServiceRoutePath = "/rpkit/users.UserService/Get"

// getServiceRoutePathURL defines a parsed path for the Get, it
// ensures the created path is valid as a url.
export const getServiceRoutePathURL = new URL(GetServiceRoutePath)

// GetContractSource contains the source version of expected method contract.
export const GetContractSource = `type GetMethodContract interface {
	Get(var1 context.Context)  (int,error)  
}
`




// GetContractSource contains the source version of expected method contract.
export const GetContractSource = `type GetMethodContract interface {
	Get(var1 context.Context)  (int,error)  
}
`

// GetEncoder defines a interface which expose a single method to encode the response
// returned by GetMethodContract.Get.
// type GetEncoder interface{
//	 Encode(io.Writer, int) error
// }



///////////////////////////////////////////////////////////////////
// RP: Output Returning No Error methods
// Method: GetUsers
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetUsers
///////////////////////////////////////////////////////////////////

// GetUsersServiceRoute defines the route for the GetUsers method.
export const GetUsersServiceRoute = "users.UserService/GetUsers"

// GetUsersServiceRoutePath defines the full method path for the GetUsers method.
export const GetUsersServiceRoutePath = "/rpkit/users.UserService/GetUsers"

// getusersServiceRoutePathURL defines a parsed path for the GetUsers, it
// ensures the created path is valid as a url.
export const getusersServiceRoutePathURL = new URL(GetUsersServiceRoutePath)

// GetUsersContractSource contains the source version of expected method contract.
export const GetUsersContractSource = `type GetUsersMethodContract interface {
	GetUsers(var1 context.Context)  users.User  
}
`


// GetUsersMethodUser defines a function to
// return a default object containing default field values of return value of
// GetUsers method.
function GetUsersMethodUser(){
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\"\n\n}");
}



// GetUsersMethodContract defines a contract interface for method "GetUsers"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
// type GetUsersMethodContract interface{
//  	GetUsers(var1 context.Context)  users.User  
// }

// GetUsersEncoder defines a interface which expose a single method to encode the response
// returned by GetUsersMethodContract.GetUsers.
// type GetUsersEncoder interface{
//  	Encode(io.Writer, users.User) error
// }





///////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: Create
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Create
///////////////////////////////////////////////////////////////////

// CreateServiceRoute defines the route for the Create method.
export const CreateServiceRoute = "users.UserService/Create"

// CreateServiceRoutePath defines the full method path for the Create method.
export const CreateServiceRoutePath = "/users/UserService/Create"

// createServiceRoutePathURL defines a parsed path for the Create, it
// ensures the created path is valid as a url.
export const createServiceRoutePathURL = new URL(CreateServiceRoutePath)

// CreateContractSource contains the source version of expected method contract.
export const CreateContractSource = `type CreateMethodContract interface {
	Create(var1 context.Context,var2 users.NewUser)  (users.User,error)  
}
`


// CreateMethodUser defines a function to
// return a default object containing default field values of return value of
// Create method.
function CreateMethodUser(){
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\"\n\n}");
}





// CreateMethodNewUser defines a function to
// return a default object containing default field values of argument of
// Create method.
function CreateMethodNewUser(){
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}



// CreateMethodContract defines a contract interface for method "Create"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
// type CreateMethodContract interface{
//  	Create(var1 context.Context,var2 users.NewUser)  (users.User,error)  
// }

// CreateDecoder defines a interface which expose a single method to decode the request data
// expected by CreateMethodContract.Create.
// type CreateDecoder interface{
//  	Decode(io.Reader) (users.NewUser, error)
// }

// CreateEncoder defines a interface which expose a single method to encode the response
// returned by CreateMethodContract.Create.
// type CreateEncoder interface{
//  	Encode(io.Writer, users.User) error
// }


///////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: GetBy
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetBy
///////////////////////////////////////////////////////////////////

// GetByServiceRoute defines the route for the GetBy method.
export const GetByServiceRoute = "users.UserService/GetBy"

// GetByServiceRoutePath defines the full method path for the GetBy method.
export const GetByServiceRoutePath = "/users/UserService/GetBy"

// getbyServiceRoutePathURL defines a parsed path for the GetBy, it
// ensures the created path is valid as a url.
export const getbyServiceRoutePathURL = new URL(GetByServiceRoutePath)

// GetByContractSource contains the source version of expected method contract.
export const GetByContractSource = `type GetByMethodContract interface {
	GetBy(var1 context.Context,var2 string)  (int,error)  
}
`








// GetByMethodContract defines a contract interface for method "GetBy"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
// type GetByMethodContract interface{
//  	GetBy(var1 context.Context,var2 string)  (int,error)  
// }

// GetByDecoder defines a interface which expose a single method to decode the request data
// expected by GetByMethodContract.GetBy.
// type GetByDecoder interface{
//  	Decode(io.Reader) (string, error)
// }

// GetByEncoder defines a interface which expose a single method to encode the response
// returned by GetByMethodContract.GetBy.
// type GetByEncoder interface{
//  	Encode(io.Writer, int) error
// }


///////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: CreateUser
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.CreateUser
///////////////////////////////////////////////////////////////////

// CreateUserServiceRoute defines the route for the CreateUser method.
export const CreateUserServiceRoute = "users.UserService/CreateUser"

// CreateUserServiceRoutePath defines the full method path for the CreateUser method.
export const CreateUserServiceRoutePath = "/users/UserService/CreateUser"

// createuserServiceRoutePathURL defines a parsed path for the CreateUser, it
// ensures the created path is valid as a url.
export const createuserServiceRoutePathURL = new URL(CreateUserServiceRoutePath)

// CreateUserContractSource contains the source version of expected method contract.
export const CreateUserContractSource = `type CreateUserMethodContract interface {
	CreateUser(var1 users.NewUser)  (users.User,error)  
}
`


// CreateUserMethodUser defines a function to
// return a default object containing default field values of return value of
// CreateUser method.
function CreateUserMethodUser(){
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\"\n\n}");
}



// CreateUserMethodNewUser defines a function to
// return a default object containing default field values of argument of
// CreateUser method.
function CreateUserMethodNewUser(){
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}





// CreateUserMethodContract defines a contract interface for method "CreateUser"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
// type CreateUserMethodContract interface{
//  	CreateUser(var1 users.NewUser)  (users.User,error)  
// }

// CreateUserDecoder defines a interface which expose a single method to decode the request data
// expected by CreateUserMethodContract.CreateUser.
// type CreateUserDecoder interface{
//  	Decode(io.Reader) (users.NewUser, error)
// }

// CreateUserEncoder defines a interface which expose a single method to encode the response
// returned by CreateUserMethodContract.CreateUser.
// type CreateUserEncoder interface{
//  	Encode(io.Writer, users.User) error
// }


///////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: GetUser
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetUser
///////////////////////////////////////////////////////////////////

// GetUserServiceRoute defines the route for the GetUser method.
export const GetUserServiceRoute = "users.UserService/GetUser"

// GetUserServiceRoutePath defines the full method path for the GetUser method.
export const GetUserServiceRoutePath = "/users/UserService/GetUser"

// getuserServiceRoutePathURL defines a parsed path for the GetUser, it
// ensures the created path is valid as a url.
export const getuserServiceRoutePathURL = new URL(GetUserServiceRoutePath)

// GetUserContractSource contains the source version of expected method contract.
export const GetUserContractSource = `type GetUserMethodContract interface {
	GetUser(var1 int)  (users.User,error)  
}
`


// GetUserMethodUser defines a function to
// return a default object containing default field values of return value of
// GetUser method.
function GetUserMethodUser(){
    return JSON.parse("{\n\n\n    \"name\":\t\"\",\n\n    \"id\":\t0\n\n}");
}







// GetUserMethodContract defines a contract interface for method "GetUser"
// provided by "users.UserService" in "github.com/gokit/rpkit/examples/users". It allows us
// establish a simple contract suitable for meeting the needs of said method.
// type GetUserMethodContract interface{
//  	GetUser(var1 int)  (users.User,error)  
// }

// GetUserDecoder defines a interface which expose a single method to decode the request data
// expected by GetUserMethodContract.GetUser.
// type GetUserDecoder interface{
//  	Decode(io.Reader) (int, error)
// }

// GetUserEncoder defines a interface which expose a single method to encode the response
// returned by GetUserMethodContract.GetUser.
// type GetUserEncoder interface{
//  	Encode(io.Writer, users.User) error
// }



