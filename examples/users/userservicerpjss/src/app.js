// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.
  
const URL = require("url-parse");
const lodash = require("lodash");
const httpStatus = require("statuses");

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

// FromPromise returns a new Promise with the
// function arguments passed to it. This
// is provided to allow clients use same
// promise as package has their are internal
// checks and use of `instanceof` with the
// Promise used.
export function FromPromise(rx, ry) {
    return new Promise(rx, ry);
}

// JSONEncoding implements a Object that possess methods
// able to encode and decode a json response.
const JSONEncoding = Object.freeze({
    Encode: function EncodeJSON(req, res, model) {
        return new Promise(function encodeJSONResolver(resolve, reject) {
            const accepts = req.headers.accept;
            if (accepts.indexOf("application/json") === -1) {
                reject(new Error("request does not accept json"));
                return;
            }

            let data = null;

            try{
                data = JSON.stringify(model);
            }catch(e){
                reject(e);
                return;
            }

            res.setHeader("Content-Type", "application/json");
            res.setHeader("Content-Length", Buffer.byteLength(data));
            res.write(data);
            resolve(true);
        });
    },
    Decode: function DecodeJSON(req) {
        return new Promise(function decodeJSONResolver(resolve, reject) {
            if (req.headers["content-type"].indexOf("application/json") === -1) {
                reject(new Error("request content type must be application/json"));
                return;
            }

            const data = [];
            req.on("data", (d) => { data.push(d); });
            req.on("end", () => {
                try {
                    const model = JSON.parse(data.join(""));
                    resolve(model);
                } catch (err) {
                    reject(err);
                }
            });
            req.on("error", (err) => {
                reject(err);
            });
        });
    },
});

// prefixSlash prefixes a forward slash to the beginning of provided
// string.
function prefixSlash(c) {
    if (c.startsWith("/")) {
        return c;
    }
    return `/${c}`;
}

function suffixSlash(c) {
    if (c.endsWith("/")) {
        return c;
    }
    return `${c};`;
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

// PokeService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
export function PokeService(PokeContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder function");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode function");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function PokeServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["not allowed"],
                "method not served",
                "not servicing method, only posts or head", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Poke",
                    api_service: MethodServiceName,
                    route: PokeServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        if(options.BeforeRequest) options.BeforeRequest(req);

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "Poke");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", PokeServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (let key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        const reqURL = URL(req.url);
        if (!(suffixSlash(reqURL.pathname).endsWith(suffixSlash(PokeServiceRoutePath)))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["bad request"],
                "method can not be served",
                "not servicing method, only POST or HEAD", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Poke",
                    api_service: MethodServiceName,
                    route: PokeServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        const actionResult = PokeContract(req);
        if (!(actionResult instanceof Promise)) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method failed to return promise",
                new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Poke",
                    api_service: MethodServiceName,
                    route: PokeServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        actionResult.then((model) => {
            res.writeHead(httpStatus["no content"], {});
            res.end();
            return model;
        }).catch((err) => {
            if (res.finished) {
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method returned with an error",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Poke",
                    api_service: MethodServiceName,
                    route: PokeServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        if (next && typeof next == "function") next();
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

// PokeAgainService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
export function PokeAgainService(PokeAgainContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder function");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode function");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function PokeAgainServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["not allowed"],
                "method not served",
                "not servicing method, only posts or head", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "PokeAgain",
                    api_service: MethodServiceName,
                    route: PokeAgainServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        if(options.BeforeRequest) options.BeforeRequest(req);

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "PokeAgain");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", PokeAgainServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (let key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        const reqURL = URL(req.url);
        if (!(suffixSlash(reqURL.pathname).endsWith(suffixSlash(PokeAgainServiceRoutePath)))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["bad request"],
                "method can not be served",
                "not servicing method, only POST or HEAD", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "PokeAgain",
                    api_service: MethodServiceName,
                    route: PokeAgainServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        const actionResult = PokeAgainContract(req);
        if (!(actionResult instanceof Promise)) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method failed to return promise",
                new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "PokeAgain",
                    api_service: MethodServiceName,
                    route: PokeAgainServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        actionResult.then((model) => {
            res.writeHead(httpStatus["no content"], {});
            res.end();
            return model;
        }).catch((err) => {
            if (res.finished) {
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method returned with an error",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "PokeAgain",
                    api_service: MethodServiceName,
                    route: PokeAgainServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        if (next && typeof next == "function") next();
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



// GetService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
export function GetService(GetContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder function");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode function");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function GetServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["not allowed"],
                "method not served",
                "not servicing method, only posts or head", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Get",
                    api_service: MethodServiceName,
                    route: GetServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        if(options.BeforeRequest) options.BeforeRequest(req);

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "Get");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", GetServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (let key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        const reqURL = URL(req.url);
        if (!(suffixSlash(reqURL.pathname).endsWith(suffixSlash(GetServiceRoutePath)))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["bad request"],
                "method can not be served",
                "not servicing method, only POST or HEAD", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Get",
                    api_service: MethodServiceName,
                    route: GetServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        const actionResult = GetContract(req);
        if (!(actionResult instanceof Promise)) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method failed to return promise",
                new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Get",
                    api_service: MethodServiceName,
                    route: GetServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        actionResult.catch((err) => {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method returned with an error",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Get",
                    api_service: MethodServiceName,
                    route: GetServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        actionResult.then((model) => {
            return options.Encoder.Encode(req, res, model);
        }).then((m) => {
            res.end();
            return m;
        }).catch((err) => {
            if (res.finished) {
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError,
                httpStatus["bad request"],
                "failed to encode response",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Get",
                    api_service: MethodServiceName,
                    route: GetServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        if (next && typeof next == "function") next();
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


// GetUsersService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
export function GetUsersService(GetUsersContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder function");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode function");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function GetUsersServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["not allowed"],
                "method not served",
                "not servicing method, only posts or head", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUsers",
                    api_service: MethodServiceName,
                    route: GetUsersServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }
        
        if(options.BeforeRequest) options.BeforeRequest(req);

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "GetUsers");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", GetUsersServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (let key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        const reqURL = URL(req.url);
        if (!(suffixSlash(reqURL.pathname).endsWith(suffixSlash(GetUsersServiceRoutePath)))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["bad request"],
                "method can not be served",
                "not servicing method, only POST or HEAD", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUsers",
                    api_service: MethodServiceName,
                    route: GetUsersServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        const actionResult = GetUsersContract(req);
        if (!(actionResult instanceof Promise)) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method failed to return promise",
                new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUsers",
                    api_service: MethodServiceName,
                    route: GetUsersServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        actionResult.catch((err) => {
            if (res.finished) {
                return;
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method returned with an error",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUsers",
                    api_service: MethodServiceName,
                    route: GetUsersServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        actionResult.then((model) => {
            return options.Encoder.Encode(req, res, model);
        }).then((m) => {
            res.end();
            return m;
        }).catch((err) => {
            if (res.finished) {
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError,
                httpStatus["bad request"],
                "failed to encode response",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUsers",
                    api_service: MethodServiceName,
                    route: GetUsersServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        if (next && typeof next == "function") next();
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


// CreateService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
export function CreateService(CreateContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder function");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode function");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function CreateServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["not allowed"],
                "method not served",
                "not servicing method, only posts or head", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Create",
                    api_service: MethodServiceName,
                    route: CreateServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        if(options.BeforeRequest) options.BeforeRequest(req);
        
        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "Create");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", CreateServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (let key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        const reqURL = URL(req.url);
        if (!(suffixSlash(reqURL.pathname).endsWith(suffixSlash(CreateServiceRoutePath)))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["bad request"],
                "method can not be served",
                "not servicing method, only POST or HEAD", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Create",
                    api_service: MethodServiceName,
                    route: CreateServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        const decodePromise = options.Decoder.Decode(req);
        if (!(decodePromise instanceof Promise)) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError,
                httpStatus["bad request"],
                "decoder.Decode method failed to return promise",
                new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Create",
                    api_service: MethodServiceName,
                    route: CreateServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        decodePromise.catch((err) => {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError,
                httpStatus["bad request"],
                "failed to decode request body",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Create",
                    api_service: MethodServiceName,
                    route: CreateServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        const actionRes = decodePromise.then((result) => {
            const actionResult = CreateContract(req, result);
            if (!(actionResult instanceof Promise)) {
                return Promise.reject(new Error("failed action: action does not return a Promise"));
            }

            return actionResult;
        });

        actionRes.catch((err) => {
            if (res.finished) {
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method returned with an error",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Create",
                    api_service: MethodServiceName,
                    route: CreateServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        actionRes.then((model) => {
            return options.Encoder.Encode(req, res, model);
        }).then((m) => {
            res.end();
            return m;
        }).catch((err) => {
            if(res.finished){
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError,
                httpStatus["bad request"],
                "failed to encode response",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Create",
                    api_service: MethodServiceName,
                    route: CreateServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        if (next && typeof next == "function") next();
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





// GetByService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
export function GetByService(GetByContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder function");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode function");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function GetByServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["not allowed"],
                "method not served",
                "not servicing method, only posts or head", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetBy",
                    api_service: MethodServiceName,
                    route: GetByServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        if(options.BeforeRequest) options.BeforeRequest(req);
        
        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "GetBy");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", GetByServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (let key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        const reqURL = URL(req.url);
        if (!(suffixSlash(reqURL.pathname).endsWith(suffixSlash(GetByServiceRoutePath)))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["bad request"],
                "method can not be served",
                "not servicing method, only POST or HEAD", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetBy",
                    api_service: MethodServiceName,
                    route: GetByServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        const decodePromise = options.Decoder.Decode(req);
        if (!(decodePromise instanceof Promise)) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError,
                httpStatus["bad request"],
                "decoder.Decode method failed to return promise",
                new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetBy",
                    api_service: MethodServiceName,
                    route: GetByServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        decodePromise.catch((err) => {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError,
                httpStatus["bad request"],
                "failed to decode request body",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetBy",
                    api_service: MethodServiceName,
                    route: GetByServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        const actionRes = decodePromise.then((result) => {
            const actionResult = GetByContract(req, result);
            if (!(actionResult instanceof Promise)) {
                return Promise.reject(new Error("failed action: action does not return a Promise"));
            }

            return actionResult;
        });

        actionRes.catch((err) => {
            if (res.finished) {
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method returned with an error",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetBy",
                    api_service: MethodServiceName,
                    route: GetByServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        actionRes.then((model) => {
            return options.Encoder.Encode(req, res, model);
        }).then((m) => {
            res.end();
            return m;
        }).catch((err) => {
            if(res.finished){
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError,
                httpStatus["bad request"],
                "failed to encode response",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetBy",
                    api_service: MethodServiceName,
                    route: GetByServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        if (next && typeof next == "function") next();
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
    return JSON.parse("{\n\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0,\n\n    \"id\":\t0,\n\n    \"name\":\t\"\"\n\n}");
}


// CreateUserMethodNewUserFactory defines a function to
// return a default object containing default field values of argument of
// CreateUser method.
export function CreateUserMethodNewUserFactory(){
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}



// CreateUserService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
export function CreateUserService(CreateUserContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder function");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode function");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function CreateUserServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["not allowed"],
                "method not served",
                "not servicing method, only posts or head", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "CreateUser",
                    api_service: MethodServiceName,
                    route: CreateUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        if(options.BeforeRequest) options.BeforeRequest(req);
        
        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "CreateUser");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", CreateUserServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (let key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        const reqURL = URL(req.url);
        if (!(suffixSlash(reqURL.pathname).endsWith(suffixSlash(CreateUserServiceRoutePath)))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["bad request"],
                "method can not be served",
                "not servicing method, only POST or HEAD", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "CreateUser",
                    api_service: MethodServiceName,
                    route: CreateUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        const decodePromise = options.Decoder.Decode(req);
        if (!(decodePromise instanceof Promise)) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError,
                httpStatus["bad request"],
                "decoder.Decode method failed to return promise",
                new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "CreateUser",
                    api_service: MethodServiceName,
                    route: CreateUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        decodePromise.catch((err) => {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError,
                httpStatus["bad request"],
                "failed to decode request body",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "CreateUser",
                    api_service: MethodServiceName,
                    route: CreateUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        const actionRes = decodePromise.then((result) => {
            const actionResult = CreateUserContract(req, result);
            if (!(actionResult instanceof Promise)) {
                return Promise.reject(new Error("failed action: action does not return a Promise"));
            }

            return actionResult;
        });

        actionRes.catch((err) => {
            if (res.finished) {
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method returned with an error",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "CreateUser",
                    api_service: MethodServiceName,
                    route: CreateUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        actionRes.then((model) => {
            return options.Encoder.Encode(req, res, model);
        }).then((m) => {
            res.end();
            return m;
        }).catch((err) => {
            if(res.finished){
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError,
                httpStatus["bad request"],
                "failed to encode response",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "CreateUser",
                    api_service: MethodServiceName,
                    route: CreateUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        if (next && typeof next == "function") next();
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
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0\n\n}");
}




// GetUserService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
export function GetUserService(GetUserContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && typeof options.Headers !== "object") {
        throw new Error("options.Headers must be a object map");
    }

    // encoders must be functions that returns promises from their Encode methods.
    if (options.Encoder && typeof options.Encoder.Encode !== "function") {
        throw new Error("Encoder must provide Encoder function");
    }

    // decoders must be functions that returns promises from their Encode methods.
    if (options.Decoder && typeof options.Decoder.Decode !== "function") {
        throw new Error("Decoder must provide Decode function");
    }

    if (!options.Headers) options.Headers = {}
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function GetUserServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["not allowed"],
                "method not served",
                "not servicing method, only posts or head", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUser",
                    api_service: MethodServiceName,
                    route: GetUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        if(options.BeforeRequest) options.BeforeRequest(req);
        
        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "GetUser");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", GetUserServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (let key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        const reqURL = URL(req.url);
        if (!(suffixSlash(reqURL.pathname).endsWith(suffixSlash(GetUserServiceRoutePath)))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError,
                httpStatus["bad request"],
                "method can not be served",
                "not servicing method, only POST or HEAD", {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUser",
                    api_service: MethodServiceName,
                    route: GetUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        const decodePromise = options.Decoder.Decode(req);
        if (!(decodePromise instanceof Promise)) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError,
                httpStatus["bad request"],
                "decoder.Decode method failed to return promise",
                new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUser",
                    api_service: MethodServiceName,
                    route: GetUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
            return;
        }

        decodePromise.catch((err) => {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError,
                httpStatus["bad request"],
                "failed to decode request body",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUser",
                    api_service: MethodServiceName,
                    route: GetUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        const actionRes = decodePromise.then((result) => {
            const actionResult = GetUserContract(req, result);
            if (!(actionResult instanceof Promise)) {
                return Promise.reject(new Error("failed action: action does not return a Promise"));
            }

            return actionResult;
        });

        actionRes.catch((err) => {
            if (res.finished) {
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ActionError,
                httpStatus["bad request"],
                "method returned with an error",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUser",
                    api_service: MethodServiceName,
                    route: GetUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        actionRes.then((model) => {
            return options.Encoder.Encode(req, res, model);
        }).then((m) => {
            res.end();
            return m;
        }).catch((err) => {
            if(res.finished){
                return Promise.reject(err);
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json",
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError,
                httpStatus["bad request"],
                "failed to encode response",
                err, {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUser",
                    api_service: MethodServiceName,
                    route: GetUserServiceRoute,
                    api: "users.UserService",
                })));

            res.end();
        });

        if (next && typeof next == "function") next();
    };
}



