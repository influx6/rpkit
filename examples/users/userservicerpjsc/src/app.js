// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.
  

// const URL = require("url-parse");
// const httpStatus = require("statuses");
// const request = require("request");
const buffer = require("safe-buffer");
const lodash = require("lodash");
const axios = require("axios");

// BaseServiceName defines the base name of service root.
export const BaseServiceName = "users"

// MethodServiceName defines the complete name of this giving API service.
export const MethodServiceName = "users/UserService";

// ServiceCodePath defines the path to this generated package which contains the implemented service methods.
export const ServiceCodePath = "github.com/gokit/rpkit/examples/users/userservicerp";

// ServiceCodePathName defines the name giving to
// this package.
export const ServiceCodePathName = "userservicerp";

// error type strings
export const URLError = "url_error";
export const ActionError = "action_error";
export const ActionPanicError = "action_panic_error";
export const MethodTypeError = "method_type_error";
export const AcceptTypeUnknownError = "accept_type_unknown_error";
export const RequestDecodingError = "request_decoding_error";
export const ResponseEncodingError = "response_encoding_error";

// JSONErrorResponse defines a structure to contain error message data
// delivered by the server.
export function JSONErrorResponse(type, code, err, message, meta) {
    return {
        Type: type,
        Code: code,
        Err: err,
        Message: message,
        Meta: meta,
    };
}

// JSONEncoding defines a JSON encoding implementation
// meant to work with the Request from NewRequest and
// response from using https://github.com/axios/axios
// as transport.
export const JSONEncoding = Object.freeze({
    Decode: function(req, res, body) {
        return new Promise(function EncodePromise(resolve, reject) {
            if (res.headers["content-type"].indexOf("application/json") === -1) {
                reject(new Error("request content type must be application/json"));
                return;
            }

            // Since we are using axios, any response that is valid JSON
            // will immediately get parsed, to just return.
            resolve(body);
        });
    },
    Encode: function(req, model) {
        return new Promise(function EncodePromise(resolve, reject) {
            req.setHeader("Content-Type", "application/json");
            req.setHeader("Accept", "application/json");

            try {
                const encoded = JSON.stringify(model);
                req.write(encoded);
            } catch (e) {
                reject(e);
                return;
            }

            req.end();
            resolve(req);
        });
    },
});

// HTTPTransport implements a custom http transport for handling
// request processing and response returning. This allows user
// to plug a custom http request transport handling layer
// as desired. An HTTPTransport must always return;
// a promise. The http transport uses https://github.com/axios/axios
// underneath.
export const HTTPTransport = Object.freeze({
    Do: function(req, timeout) {
        const ops = axios.request({
            method: "POST",
            timeout: timeout,
            data: req.body,
            url: joinURL(req.base, req.url),
            headers: lodash.merge(req.headers, {"X-Requested-With-Lib": "axios (https://github.com/axios/axios)"}),
        })

        return ops.then((res) => {
            if(!res){
                return Promise.reject(new Error("received null/nil response object"));
            }

            if(isError(res)){
                return Promise.reject(res)
            }

            if (requestRedirected(res) || requestFailed(res) || requestFacedInternalIssues(res)) {
                return Promise.reject(getJSONError(res.data));
            }

            return {
                res: res,
                body: res.data,
            };
        }).catch((e) =>{
            if(e.response && e.response.data){
                const msg = getJSONError(e.response.data);
                msg.original_error = e;
                return Promise.reject(msg);
            }
            return Promise.reject(e);
        });
    },
});


// FromPromise returns a new Promise with the
// function arguments passed to it. This
// is provided to allow clients use same
// promise as package has their are internal
// checks and use of `instanceof` with the
// Promise used.
export function FromPromise(rx, ry) {
    return new Promise(rx, ry);
}

// NewRequest returns a object containing basic data related
// to a outgoing request with associated headers to be sent along.
// It is to be used by http transport for fullfilling a request.
export function NewRequest(base, url, headers) {
    const reqObj = {
        base: base,
        url: url,
        headers: headers || {},
        body: null,
        setHeader: (key, value) => {
            reqObj.headers[key] = value;
        },
        write: (data) => {
            if (reqObj.body) {
                reqObj.push(data);
                return;
            }

            reqObj.body = [];
            reqObj.body.push(data);
        },
        end: () => {
            if (reqObj.body) reqObj.body = reqObj.body.join("");
        },
    };

    return reqObj;
}

function isError(err) {
    if (lodash.isError(err)) {
        return true;
    }
    return (err instanceof Error);
}

// joinURL joins provided string paths into a whole,
// ensuring to add the necessary slash in between parts.
function joinURL(base, to) {
    let toURL = to;
    if (toURL.startsWith("/")) {
        toURL = to.substring(1);
    }

    if (base.endsWith("/")) {
        return (`${base}${toURL}`).trim();
    }

    return (`${base}/${toURL}`).trim();
}

// prefixSlash prefixes a forward slash to the beginning of provided
// string.
function prefixSlash(c) {
    if (c.startsWith("/")) {
        return c;
    }
    return `/${c}`;
}

// suffixSlash suffixes a forward slash to end of provided
// string.
function suffixSlash(c) {
    if (c.endsWith("/")) {
        return c;
    }
    return `${c}`;
}

function requestRedirected(res) {
    if (res.status >= 300 && res.status <= 308) {
        return true;
    }
    return false;
}

function requestFailed(res) {
    if (res.status >= 400 && res.status < 500) {
        return true;
    }
    return false;
}

function requestFacedInternalIssues(res) {
    if (res.status >= 500 && res.status < 600) {
        return true;
    }
    return false;
}

function requestSucceeded(res) {
    if (res.status >= 200 && res.status <= 299) {
        return true;
    }
    return false;
}

function getJSONError(body) {
    if (typeof body === "object") {
        return body;
    }

    let res = null;
    if (body instanceof buffer.Buffer) {
        try {
            const data = body.toString("utf8");
            res = JSON.parse(data);
        } catch (e) {
            return e;
        }
        return res;
    }

    try {
        res = JSON.parse(body);
    } catch (e) {
        return e;
    }

    return res;
}

// /////////////////////////////////////////////////////////////////
// RP: No Arguments and No Return Methods
// Method: Poke
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Poke
// /////////////////////////////////////////////////////////////////

// PokeServiceRoute defines the route for the Poke method.
export const PokeServiceRoute = "users.UserService/Poke";

// PokeServiceRoutePath defines the full method path for the Poke method.
export const PokeServiceRoutePath = "/rpkit/users.UserService/Poke";

// PokeContractSource contains the source version of expected method contract.
export const PokeContractSource = `type PokeMethodContract interface {
	Poke() 
}`;

// PokeClient returns a RPC method to be called to handle requests
// for the Poke method of "github.com/gokit/rpkit/examples/users".
export function PokeClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder method");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode method");
    }

    // Transport must be function that returns promises from their calls.
    if (options.Transport && typeof options.Transport.Do !== "function") {
        throw new Error("Transport must provide Do method");
    }

    if (!options.ServiceAddr) {
        throw new Error("options must provide a ServiceAddr pointing to service http server");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function PokeRPC() {
        return new Promise(function GetUserPromise(resolve, reject) {
            const req = NewRequest(options.ServiceAddr, PokeServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "Poke");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", PokeServiceRoute);
            req.setHeader("X-API-Client-Route-Path", PokeServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            options.Transport.Do(req, options.Timeout).then((resObj) => {
                resolve(null);
            }).catch((err) => {
                reject(err);
            });
        });
    };
}



// /////////////////////////////////////////////////////////////////
// RP: Error Returning methods
// Method: PokeAgain
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.PokeAgain
// /////////////////////////////////////////////////////////////////

// PokeAgainServiceRoute defines the route for the PokeAgain method.
export const PokeAgainServiceRoute = "users.UserService/PokeAgain";

// PokeAgainServiceRoutePath defines the full method path for the PokeAgain method.
export const PokeAgainServiceRoutePath = "/rpkit/users.UserService/PokeAgain";

// PokeAgainContractSource contains the source version of expected method contract.
export const PokeAgainContractSource = `type PokeAgainMethodContract interface {
	PokeAgain()  error  
}`;

// PokeAgainClient returns a RPC method to be called to handle requests
// for the PokeAgain method of "github.com/gokit/rpkit/examples/users".
export function PokeAgainClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder method");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode method");
    }

    // Transport must be function that returns promises from their calls.
    if (options.Transport && typeof options.Transport.Do !== "function") {
        throw new Error("Transport must provide Do method");
    }

    if (!options.ServiceAddr) {
        throw new Error("options must provide a ServiceAddr pointing to service http server");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function PokeAgainRPC() {
        return new Promise(function GetUserPromise(resolve, reject) {
            const req = NewRequest(options.ServiceAddr, PokeAgainServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "PokeAgain");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", PokeAgainServiceRoute);
            req.setHeader("X-API-Client-Route-Path", PokeAgainServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            options.Transport.Do(req, options.Timeout).then((resObj) => {
                resolve(null);
            }).catch((err) => {
                reject(err);
            });
        });
    };
}



// /////////////////////////////////////////////////////////////////
// RP: Output with Returning Error methods
// Method: Get
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Get
// /////////////////////////////////////////////////////////////////

// GetServiceRoute defines the route for the Get method.
export const GetServiceRoute = "users.UserService/Get";

// GetServiceRoutePath defines the full method path for the Get method.
export const GetServiceRoutePath = "/rpkit/users.UserService/Get";

// GetContractSource contains the source version of expected method contract.
export const GetContractSource = `type GetMethodContract interface {
	Get(var1 context.Context)  (int,error)  
}`;


// GetClient returns a RPC method to be called to handle requests
// for the Get method of "github.com/gokit/rpkit/examples/users".
export function GetClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder method");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode method");
    }

    // Transport must be function that returns promises from their calls.
    if (options.Transport && typeof options.Transport.Do !== "function") {
        throw new Error("Transport must provide Do method");
    }

    if (!options.ServiceAddr) {
        throw new Error("options must provide a ServiceAddr pointing to service http server");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function GetRPC() {
        return new Promise(function GetUserPromise(resolve, reject) {
            const req = NewRequest(options.ServiceAddr, GetServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "Get");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", GetServiceRoute);
            req.setHeader("X-API-Client-Route-Path", GetServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            options.Transport.Do(req, options.Timeout).then((resObj) => {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then((resModel) => {
                resolve(resModel);
            }).catch((err) => {
                reject(err);
            });
        });
    };
}



// /////////////////////////////////////////////////////////////////
// RP: Output Returning No Error methods
// Method: GetUsers
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetUsers
// /////////////////////////////////////////////////////////////////

// GetUsersServiceRoute defines the route for the GetUsers method.
export const GetUsersServiceRoute = "users.UserService/GetUsers";

// GetUsersServiceRoutePath defines the full method path for the GetUsers method.
export const GetUsersServiceRoutePath = "/rpkit/users.UserService/GetUsers";

// GetUsersContractSource contains the source version of expected method contract.
export const GetUsersContractSource = `type GetUsersMethodContract interface {
	GetUsers(var1 context.Context)  []users.User  
}`;


// GetUsersMethodUserSliceFactory defines a function to
// return a default object containing default field values of return value of
// GetUsers method.
export function GetUsersMethodUserSliceFactory(){
    return JSON.parse("{\n\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0,\n\n    \"id\":\t0\n\n}");
}


// GetUsersClient returns a RPC method to be called to handle requests
// for the GetUsers method of "github.com/gokit/rpkit/examples/users".
export function GetUsersClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder method");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode method");
    }

    // Transport must be function that returns promises from their calls.
    if (options.Transport && typeof options.Transport.Do !== "function") {
        throw new Error("Transport must provide Do method");
    }

    if (!options.ServiceAddr) {
        throw new Error("options must provide a ServiceAddr pointing to service http server");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function GetUsersRPC() {
        return new Promise(function GetUserPromise(resolve, reject) {
            const req = NewRequest(options.ServiceAddr, GetUsersServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "GetUsers");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", GetUsersServiceRoute);
            req.setHeader("X-API-Client-Route-Path", GetUsersServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            options.Transport.Do(req, options.Timeout).then((resObj) => {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then((resModel) => {
                resolve(resModel);
            }).catch((err) => {
                reject(err);
            });
        });
    };
}





// /////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: Create
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Create
// /////////////////////////////////////////////////////////////////

// CreateServiceRoute defines the route for the Create method.
export const CreateServiceRoute = "users.UserService/Create";

// CreateServiceRoutePath defines the full method path for the Create method.
export const CreateServiceRoutePath = "/users/UserService/Create";

// CreateContractSource contains the source version of expected method contract.
export const CreateContractSource = `type CreateMethodContract interface {
	Create(var1 context.Context,var2 users.NewUser)  (users.User,error)  
}`;


// CreateMethodUserFactory defines a function to
// return a default object containing default field values of return value of
// Create method.
export function CreateMethodUserFactory(){
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0\n\n}");
}



// CreateMethodNewUserFactory defines a function to
// return a default object containing default field values of argument of
// Create method. 
export function CreateMethodNewUserFactory(){
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}


// CreateClient returns a RPC method to be called to handle requests
// for the Create method of "github.com/gokit/rpkit/examples/users".
export function CreateClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder method");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode method");
    }

    // Transport must be function that returns promises from their calls.
    if (options.Transport && typeof options.Transport.Do !== "function") {
        throw new Error("Transport must provide Do method");
    }

    if (!options.ServiceAddr) {
        throw new Error("options must provide a ServiceAddr pointing to service http server");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function CreateRPC(model) {
        return new Promise(function GetUserPromise(resolve, reject) {
            const req = NewRequest(options.ServiceAddr, CreateServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "Create");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", CreateServiceRoute);
            req.setHeader("X-API-Client-Route-Path", CreateServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            const encoded = options.Encoder.Encode(req, model);
            if (!(encoded instanceof Promise)) {
                reject(new Error("Encoder.Encode does not return a Promise"));
                return;
            }

            encoded.then((req) => {
                return options.Transport.Do(req, options.Timeout);
            }).then((resObj) => {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then((resModel) => {
                resolve(resModel);
            }).catch((err) => {
                reject(err);
            });
        });
    };
}

// /////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: GetBy
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetBy
// /////////////////////////////////////////////////////////////////

// GetByServiceRoute defines the route for the GetBy method.
export const GetByServiceRoute = "users.UserService/GetBy";

// GetByServiceRoutePath defines the full method path for the GetBy method.
export const GetByServiceRoutePath = "/users/UserService/GetBy";

// GetByContractSource contains the source version of expected method contract.
export const GetByContractSource = `type GetByMethodContract interface {
	GetBy(var1 context.Context,var2 string)  (int,error)  
}`;





// GetByClient returns a RPC method to be called to handle requests
// for the GetBy method of "github.com/gokit/rpkit/examples/users".
export function GetByClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder method");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode method");
    }

    // Transport must be function that returns promises from their calls.
    if (options.Transport && typeof options.Transport.Do !== "function") {
        throw new Error("Transport must provide Do method");
    }

    if (!options.ServiceAddr) {
        throw new Error("options must provide a ServiceAddr pointing to service http server");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function GetByRPC(model) {
        return new Promise(function GetUserPromise(resolve, reject) {
            const req = NewRequest(options.ServiceAddr, GetByServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "GetBy");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", GetByServiceRoute);
            req.setHeader("X-API-Client-Route-Path", GetByServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            const encoded = options.Encoder.Encode(req, model);
            if (!(encoded instanceof Promise)) {
                reject(new Error("Encoder.Encode does not return a Promise"));
                return;
            }

            encoded.then((req) => {
                return options.Transport.Do(req, options.Timeout);
            }).then((resObj) => {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then((resModel) => {
                resolve(resModel);
            }).catch((err) => {
                reject(err);
            });
        });
    };
}

// /////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: CreateUser
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.CreateUser
// /////////////////////////////////////////////////////////////////

// CreateUserServiceRoute defines the route for the CreateUser method.
export const CreateUserServiceRoute = "users.UserService/CreateUser";

// CreateUserServiceRoutePath defines the full method path for the CreateUser method.
export const CreateUserServiceRoutePath = "/users/UserService/CreateUser";

// CreateUserContractSource contains the source version of expected method contract.
export const CreateUserContractSource = `type CreateUserMethodContract interface {
	CreateUser(var1 users.NewUser)  (users.User,error)  
}`;


// CreateUserMethodUserFactory defines a function to
// return a default object containing default field values of return value of
// CreateUser method.
export function CreateUserMethodUserFactory(){
    return JSON.parse("{\n\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0,\n\n    \"id\":\t0\n\n}");
}


// CreateUserMethodNewUserFactory defines a function to
// return a default object containing default field values of argument of
// CreateUser method.
export function CreateUserMethodNewUserFactory(){
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}



// CreateUserClient returns a RPC method to be called to handle requests
// for the CreateUser method of "github.com/gokit/rpkit/examples/users".
export function CreateUserClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder method");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode method");
    }

    // Transport must be function that returns promises from their calls.
    if (options.Transport && typeof options.Transport.Do !== "function") {
        throw new Error("Transport must provide Do method");
    }

    if (!options.ServiceAddr) {
        throw new Error("options must provide a ServiceAddr pointing to service http server");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function CreateUserRPC(model) {
        return new Promise(function GetUserPromise(resolve, reject) {
            const req = NewRequest(options.ServiceAddr, CreateUserServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "CreateUser");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", CreateUserServiceRoute);
            req.setHeader("X-API-Client-Route-Path", CreateUserServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            const encoded = options.Encoder.Encode(req, model);
            if (!(encoded instanceof Promise)) {
                reject(new Error("Encoder.Encode does not return a Promise"));
                return;
            }

            encoded.then((req) => {
                return options.Transport.Do(req, options.Timeout);
            }).then((resObj) => {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then((resModel) => {
                resolve(resModel);
            }).catch((err) => {
                reject(err);
            });
        });
    };
}

// /////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: GetUser
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetUser
// /////////////////////////////////////////////////////////////////

// GetUserServiceRoute defines the route for the GetUser method.
export const GetUserServiceRoute = "users.UserService/GetUser";

// GetUserServiceRoutePath defines the full method path for the GetUser method.
export const GetUserServiceRoutePath = "/users/UserService/GetUser";

// GetUserContractSource contains the source version of expected method contract.
export const GetUserContractSource = `type GetUserMethodContract interface {
	GetUser(var1 int)  (users.User,error)  
}`;


// GetUserMethodUserFactory defines a function to
// return a default object containing default field values of return value of
// GetUser method.
export function GetUserMethodUserFactory(){
    return JSON.parse("{\n\n\n    \"cid\":\t0.0,\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\"\n\n}");
}




// GetUserClient returns a RPC method to be called to handle requests
// for the GetUser method of "github.com/gokit/rpkit/examples/users".
export function GetUserClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder method");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode method");
    }

    // Transport must be function that returns promises from their calls.
    if (options.Transport && typeof options.Transport.Do !== "function") {
        throw new Error("Transport must provide Do method");
    }

    if (!options.ServiceAddr) {
        throw new Error("options must provide a ServiceAddr pointing to service http server");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function GetUserRPC(model) {
        return new Promise(function GetUserPromise(resolve, reject) {
            const req = NewRequest(options.ServiceAddr, GetUserServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "GetUser");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", GetUserServiceRoute);
            req.setHeader("X-API-Client-Route-Path", GetUserServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            const encoded = options.Encoder.Encode(req, model);
            if (!(encoded instanceof Promise)) {
                reject(new Error("Encoder.Encode does not return a Promise"));
                return;
            }

            encoded.then((req) => {
                return options.Transport.Do(req, options.Timeout);
            }).then((resObj) => {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then((resModel) => {
                resolve(resModel);
            }).catch((err) => {
                reject(err);
            });
        });
    };
}


