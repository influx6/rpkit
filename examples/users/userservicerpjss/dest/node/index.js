require("source-map-support").install();
(function webpackUniversalModuleDefinition(root, factory) {
	if(typeof exports === 'object' && typeof module === 'object')
		module.exports = factory(require("statuses"), require("url-parse"));
	else if(typeof define === 'function' && define.amd)
		define(["statuses", "url-parse"], factory);
	else if(typeof exports === 'object')
		exports["userservicerpjss"] = factory(require("statuses"), require("url-parse"));
	else
		root["userservicerpjss"] = factory(root["_"], root["_"]);
})(global, function(__WEBPACK_EXTERNAL_MODULE_statuses__, __WEBPACK_EXTERNAL_MODULE_url_parse__) {
return /******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/app.js");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./src/app.js":
/*!********************!*\
  !*** ./src/app.js ***!
  \********************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
    value: true
});

var _typeof = typeof Symbol === "function" && typeof Symbol.iterator === "symbol" ? function (obj) { return typeof obj; } : function (obj) { return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; };

exports.JSONErrorResponse = JSONErrorResponse;
exports.GetService = GetService;
exports.GetUsersMethodUser = GetUsersMethodUser;
exports.GetUsersService = GetUsersService;
exports.CreateMethodUser = CreateMethodUser;
exports.CreateMethodNewUser = CreateMethodNewUser;
exports.CreateService = CreateService;
exports.GetByService = GetByService;
exports.CreateUserMethodUser = CreateUserMethodUser;
exports.CreateUserMethodNewUser = CreateUserMethodNewUser;
exports.CreateUserService = CreateUserService;
exports.GetUserMethodUser = GetUserMethodUser;
exports.GetUserService = GetUserService;
// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.


var http = __webpack_require__(/*! http */ "http");
var URL = __webpack_require__(/*! url-parse */ "url-parse");
var httpStatus = __webpack_require__(/*! statuses */ "statuses");

// BaseServiceName defines the base name of service root.
var BaseServiceName = exports.BaseServiceName = "users";

// MethodServiceName defines the complete name of this giving API service.
var MethodServiceName = exports.MethodServiceName = "users/UserService";

// ServiceCodePath defines the path to this generated package which contains the implemented service methods.
var ServiceCodePath = exports.ServiceCodePath = "github.com/gokit/rpkit/examples/users/userservicerp";

// ServiceCodePathName defines the name giving to
// this package.
var ServiceCodePathName = exports.ServiceCodePathName = "userservicerp";

// error type strings
var URLError = exports.URLError = "url_error";
var ActionError = exports.ActionError = "action_error";
var ActionPanicError = exports.ActionPanicError = "action_panic_error";
var MethodTypeError = exports.MethodTypeError = "method_type_error";
var AcceptTypeUnknownError = exports.AcceptTypeUnknownError = "accept_type_unknown_error";
var RequestDecodingError = exports.RequestDecodingError = "request_decoding_error";
var ResponseEncodingError = exports.ResponseEncodingError = "response_encoding_error";

// JSONErrorResponse defines a structure to contain error message data
// delivered by the server.
function JSONErrorResponse(type, code, err, message, meta) {
    return {
        Type: type,
        Code: code,
        Err: err,
        Message: message,
        Meta: meta
    };
}

// JSONEncoding implements a Object that possess methods
// able to encode and decode a json response.
var JSONEncoding = Object.freeze({
    Encode: function EncodeJSON(req, res, model) {
        return new Promise(function encodeJSONResolver(resolve, reject) {
            var accepts = req.getHeader("accept");
            if (!accepts.indexOf("application/json")) {
                reject(new Error("request does not accept json"));
                return;
            }

            var data = null;

            try {
                data = JSON.stringify(model);
            } catch (e) {
                reject(e);
                return;
            }

            res.setHeader("Content-Type", "application/json");
            res.setHeader("Content-Length", data.length);
            res.Write(data);
            resolve(true);
        });
    },
    Decode: function DecodeJSON(req) {
        return new Promise(function decodeJSONResolver(resolve, reject) {
            if (req.header.contentType.indexOf("application/json") === -1) {
                reject(new Error("request content type must be application/json"));
                return;
            }

            var data = [];
            req.on("data", function (d) {
                return data.push(d);
            });
            req.on("end", function () {
                try {
                    var model = JSON.parse(data.join(""));
                    resolve(model);
                } catch (err) {
                    reject(err);
                }
            });
            req.on("error", function (err) {
                reject(err);
            });
        });
    }
});

function prefixSlash(c) {
    if (c.endsWith("/")) {
        return c;
    }
    return c + ";";
};

// /////////////////////////////////////////////////////////////////
// RP: Output with Returning Error methods
// Method: Get
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Get
// /////////////////////////////////////////////////////////////////

// GetServiceRoute defines the route for the Get method.
var GetServiceRoute = exports.GetServiceRoute = "users.UserService/Get";

// GetServiceRoutePath defines the full method path for the Get method.
var GetServiceRoutePath = exports.GetServiceRoutePath = "/rpkit/users.UserService/Get";

// GetContractSource contains the source version of expected method contract.
var GetContractSource = exports.GetContractSource = "type GetMethodContract interface {\n\tGet(var1 context.Context)  (int,error)  \n}\n";

// GetService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
function GetService(GetContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.BeforeResponse && typeof options.BeforeResponse !== "function") {
        throw new Error("options.BeforeResponse must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function GetServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["not allowed"], "method not served", "not servicing method, only posts or head", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Get",
                api_service: MethodServiceName,
                route: GetServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "Get");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", GetServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (var key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        var reqURL = URL(req.url);
        if (!prefixSlash(reqURL.pathname).endsWith(prefixSlash(GetUserServiceRoutePath))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["bad request"], "method can not be served", "not servicing method, only POST or HEAD", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Get",
                api_service: MethodServiceName,
                route: GetServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        Promise.resolve(true).then(function () {
            var actionResult = GetContract(req);
            if (!(actionResult instanceof Promise)) {
                res.writeHead(httpStatus["bad request"], {
                    "Content-Type": "application/json"
                });

                res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "method failed to return promise", new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Get",
                    api_service: MethodServiceName,
                    route: GetServiceRoute,
                    api: "users.UserService"
                })));

                res.end();
                return;
            }

            return actionResult;
        }).catch(function (err) {
            if (res.finished) {
                return;
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ActinError, httpStatus["bad request"], "method returned with an error", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Get",
                api_service: MethodServiceName,
                route: GetServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function (model) {
            return options.Encoder.Encode(req, res, model);
        }).catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError, httpStatus["bad request"], "failed to encode response", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Get",
                api_service: MethodServiceName,
                route: GetServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function () {
            res.end();
        });

        if (next && typeof next == "function") next();
    };
};

// /////////////////////////////////////////////////////////////////
// RP: Output Returning No Error methods
// Method: GetUsers
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetUsers
// /////////////////////////////////////////////////////////////////

// GetUsersServiceRoute defines the route for the GetUsers method.
var GetUsersServiceRoute = exports.GetUsersServiceRoute = "users.UserService/GetUsers";

// GetUsersServiceRoutePath defines the full method path for the GetUsers method.
var GetUsersServiceRoutePath = exports.GetUsersServiceRoutePath = "/rpkit/users.UserService/GetUsers";

// GetUsersContractSource contains the source version of expected method contract.
var GetUsersContractSource = exports.GetUsersContractSource = "type GetUsersMethodContract interface {\n\tGetUsers(var1 context.Context)  users.User  \n}\n";

// GetUsersMethodUser defines a function to
// return a default object containing default field values of return value of
// GetUsers method.
function GetUsersMethodUser() {
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0\n\n}");
}

// GetUsersService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
function GetUsersService(GetUsersContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.BeforeResponse && typeof options.BeforeResponse !== "function") {
        throw new Error("options.BeforeResponse must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function GetUsersServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["not allowed"], "method not served", "not servicing method, only posts or head", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUsers",
                api_service: MethodServiceName,
                route: GetUsersServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "GetUsers");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", GetUsersServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (var key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        var reqURL = URL(req.url);
        if (!prefixSlash(reqURL.pathname).endsWith(prefixSlash(GetUserServiceRoutePath))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["bad request"], "method can not be served", "not servicing method, only POST or HEAD", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUsers",
                api_service: MethodServiceName,
                route: GetUsersServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        Promise.resolve(true).then(function () {
            var actionResult = GetUsersContract(req);
            if (!(actionResult instanceof Promise)) {
                res.writeHead(httpStatus["bad request"], {
                    "Content-Type": "application/json"
                });

                res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "method failed to return promise", new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUsers",
                    api_service: MethodServiceName,
                    route: GetUsersServiceRoute,
                    api: "users.UserService"
                })));

                res.end();
                return;
            }

            return actionResult;
        }).catch(function (err) {
            if (res.finished) {
                return;
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ActinError, httpStatus["bad request"], "method returned with an error", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUsers",
                api_service: MethodServiceName,
                route: GetUsersServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function (model) {
            return options.Encoder.Encode(req, res, model);
        }).catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError, httpStatus["bad request"], "failed to encode response", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUsers",
                api_service: MethodServiceName,
                route: GetUsersServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function () {
            res.end();
        });

        if (next && typeof next == "function") next();
    };
};

// /////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: Create
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Create
// /////////////////////////////////////////////////////////////////

// CreateServiceRoute defines the route for the Create method.
var CreateServiceRoute = exports.CreateServiceRoute = "users.UserService/Create";

// CreateServiceRoutePath defines the full method path for the Create method.
var CreateServiceRoutePath = exports.CreateServiceRoutePath = "/users/UserService/Create";

// CreateContractSource contains the source version of expected method contract.
var CreateContractSource = exports.CreateContractSource = "type CreateMethodContract interface {\n\tCreate(var1 context.Context,var2 users.NewUser)  (users.User,error)  \n}\n";

// CreateMethodUser defines a function to
// return a default object containing default field values of return value of
// Create method.
function CreateMethodUser() {
    return JSON.parse("{\n\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0,\n\n    \"id\":\t0\n\n}");
}

// CreateMethodNewUser defines a function to
// return a default object containing default field values of argument of
// Create method.
function CreateMethodNewUser() {
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}

// CreateService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
function CreateService(CreateContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.BeforeResponse && typeof options.BeforeResponse !== "function") {
        throw new Error("options.BeforeResponse must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function CreateServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["not allowed"], "method not served", "not servicing method, only posts or head", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Create",
                api_service: MethodServiceName,
                route: CreateServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "Create");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", CreateServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (var key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        var reqURL = URL(req.url);
        if (!prefixSlash(reqURL.pathname).endsWith(prefixSlash(GetUserServiceRoutePath))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["bad request"], "method can not be served", "not servicing method, only POST or HEAD", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Create",
                api_service: MethodServiceName,
                route: CreateServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        var decodePromise = options.Decoder.Decode(req);
        decodePromise.catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "failed to decode request body", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Create",
                api_service: MethodServiceName,
                route: CreateServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        });

        decodePromise.then(function (result) {
            var actionResult = CreateContract(req, result);
            if (!(actionResult instanceof Promise)) {
                res.writeHead(httpStatus["bad request"], {
                    "Content-Type": "application/json"
                });

                res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "method failed to return promise", new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "Create",
                    api_service: MethodServiceName,
                    route: CreateServiceRoute,
                    api: "users.UserService"
                })));

                res.end();
                return;
            }

            return actionResult;
        }).catch(function (err) {
            if (res.finished) {
                return;
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ActinError, httpStatus["bad request"], "method returned with an error", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Create",
                api_service: MethodServiceName,
                route: CreateServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function (model) {
            return options.Encoder.Encode(req, res, model);
        }).catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError, httpStatus["bad request"], "failed to encode response", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "Create",
                api_service: MethodServiceName,
                route: CreateServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function () {
            res.end();
        });

        if (next && typeof next == "function") next();
    };
};

// /////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: GetBy
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetBy
// /////////////////////////////////////////////////////////////////

// GetByServiceRoute defines the route for the GetBy method.
var GetByServiceRoute = exports.GetByServiceRoute = "users.UserService/GetBy";

// GetByServiceRoutePath defines the full method path for the GetBy method.
var GetByServiceRoutePath = exports.GetByServiceRoutePath = "/users/UserService/GetBy";

// GetByContractSource contains the source version of expected method contract.
var GetByContractSource = exports.GetByContractSource = "type GetByMethodContract interface {\n\tGetBy(var1 context.Context,var2 string)  (int,error)  \n}\n";

// GetByService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
function GetByService(GetByContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.BeforeResponse && typeof options.BeforeResponse !== "function") {
        throw new Error("options.BeforeResponse must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function GetByServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["not allowed"], "method not served", "not servicing method, only posts or head", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetBy",
                api_service: MethodServiceName,
                route: GetByServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "GetBy");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", GetByServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (var key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        var reqURL = URL(req.url);
        if (!prefixSlash(reqURL.pathname).endsWith(prefixSlash(GetUserServiceRoutePath))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["bad request"], "method can not be served", "not servicing method, only POST or HEAD", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetBy",
                api_service: MethodServiceName,
                route: GetByServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        var decodePromise = options.Decoder.Decode(req);
        decodePromise.catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "failed to decode request body", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetBy",
                api_service: MethodServiceName,
                route: GetByServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        });

        decodePromise.then(function (result) {
            var actionResult = GetByContract(req, result);
            if (!(actionResult instanceof Promise)) {
                res.writeHead(httpStatus["bad request"], {
                    "Content-Type": "application/json"
                });

                res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "method failed to return promise", new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetBy",
                    api_service: MethodServiceName,
                    route: GetByServiceRoute,
                    api: "users.UserService"
                })));

                res.end();
                return;
            }

            return actionResult;
        }).catch(function (err) {
            if (res.finished) {
                return;
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ActinError, httpStatus["bad request"], "method returned with an error", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetBy",
                api_service: MethodServiceName,
                route: GetByServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function (model) {
            return options.Encoder.Encode(req, res, model);
        }).catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError, httpStatus["bad request"], "failed to encode response", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetBy",
                api_service: MethodServiceName,
                route: GetByServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function () {
            res.end();
        });

        if (next && typeof next == "function") next();
    };
};

// /////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: CreateUser
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.CreateUser
// /////////////////////////////////////////////////////////////////

// CreateUserServiceRoute defines the route for the CreateUser method.
var CreateUserServiceRoute = exports.CreateUserServiceRoute = "users.UserService/CreateUser";

// CreateUserServiceRoutePath defines the full method path for the CreateUser method.
var CreateUserServiceRoutePath = exports.CreateUserServiceRoutePath = "/users/UserService/CreateUser";

// CreateUserContractSource contains the source version of expected method contract.
var CreateUserContractSource = exports.CreateUserContractSource = "type CreateUserMethodContract interface {\n\tCreateUser(var1 users.NewUser)  (users.User,error)  \n}\n";

// CreateUserMethodUser defines a function to
// return a default object containing default field values of return value of
// CreateUser method.
function CreateUserMethodUser() {
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0\n\n}");
}

// CreateUserMethodNewUser defines a function to
// return a default object containing default field values of argument of
// CreateUser method.
function CreateUserMethodNewUser() {
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}

// CreateUserService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
function CreateUserService(CreateUserContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.BeforeResponse && typeof options.BeforeResponse !== "function") {
        throw new Error("options.BeforeResponse must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function CreateUserServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["not allowed"], "method not served", "not servicing method, only posts or head", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "CreateUser",
                api_service: MethodServiceName,
                route: CreateUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "CreateUser");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", CreateUserServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (var key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        var reqURL = URL(req.url);
        if (!prefixSlash(reqURL.pathname).endsWith(prefixSlash(GetUserServiceRoutePath))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["bad request"], "method can not be served", "not servicing method, only POST or HEAD", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "CreateUser",
                api_service: MethodServiceName,
                route: CreateUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        var decodePromise = options.Decoder.Decode(req);
        decodePromise.catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "failed to decode request body", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "CreateUser",
                api_service: MethodServiceName,
                route: CreateUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        });

        decodePromise.then(function (result) {
            var actionResult = CreateUserContract(req, result);
            if (!(actionResult instanceof Promise)) {
                res.writeHead(httpStatus["bad request"], {
                    "Content-Type": "application/json"
                });

                res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "method failed to return promise", new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "CreateUser",
                    api_service: MethodServiceName,
                    route: CreateUserServiceRoute,
                    api: "users.UserService"
                })));

                res.end();
                return;
            }

            return actionResult;
        }).catch(function (err) {
            if (res.finished) {
                return;
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ActinError, httpStatus["bad request"], "method returned with an error", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "CreateUser",
                api_service: MethodServiceName,
                route: CreateUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function (model) {
            return options.Encoder.Encode(req, res, model);
        }).catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError, httpStatus["bad request"], "failed to encode response", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "CreateUser",
                api_service: MethodServiceName,
                route: CreateUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function () {
            res.end();
        });

        if (next && typeof next == "function") next();
    };
};

// /////////////////////////////////////////////////////////////////
// RP: Input And Output Returning Error methods
// Method: GetUser
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetUser
// /////////////////////////////////////////////////////////////////

// GetUserServiceRoute defines the route for the GetUser method.
var GetUserServiceRoute = exports.GetUserServiceRoute = "users.UserService/GetUser";

// GetUserServiceRoutePath defines the full method path for the GetUser method.
var GetUserServiceRoutePath = exports.GetUserServiceRoutePath = "/users/UserService/GetUser";

// GetUserContractSource contains the source version of expected method contract.
var GetUserContractSource = exports.GetUserContractSource = "type GetUserMethodContract interface {\n\tGetUser(var1 int)  (users.User,error)  \n}\n";

// GetUserMethodUser defines a function to
// return a default object containing default field values of return value of
// GetUser method.
function GetUserMethodUser() {
    return JSON.parse("{\n\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0,\n\n    \"id\":\t0,\n\n    \"name\":\t\"\"\n\n}");
}

// GetUserService returns a middleware-aware function which services
// server-based requests for the GetByContract method. The GetByContract
// must return a promise, which must be used for receiving a response.
function GetUserService(GetUserContract, options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.BeforeResponse && typeof options.BeforeResponse !== "function") {
        throw new Error("options.BeforeResponse must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;

    return function GetUserServiceMiddleware(req, res, next) {
        if (["HEAD", "POST"].indexOf(req.method) === -1) {
            res.writeHead(httpStatus["not allowed"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["not allowed"], "method not served", "not servicing method, only posts or head", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUser",
                api_service: MethodServiceName,
                route: GetUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        res.setHeader("X-Agent", "RPKIT");
        res.setHeader("X-Service", BaseServiceName);
        res.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
        res.setHeader("X-Method", "GetUser");
        res.setHeader("X-Method-Service", MethodServiceName);
        res.setHeader("X-API-Route", GetUserServiceRoute);
        res.setHeader("X-Package-Interface", "users.UserService");

        for (var key in options.Headers) {
            res.setHeader(key, options.Headers[key]);
        }

        if (req.method === "HEAD") {
            res.writeHead(httpStatus["no content"], {});
            return;
        }

        var reqURL = URL(req.url);
        if (!prefixSlash(reqURL.pathname).endsWith(prefixSlash(GetUserServiceRoutePath))) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(MethodTypeError, httpStatus["bad request"], "method can not be served", "not servicing method, only POST or HEAD", {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUser",
                api_service: MethodServiceName,
                route: GetUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
            return;
        }

        var decodePromise = options.Decoder.Decode(req);
        decodePromise.catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "failed to decode request body", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUser",
                api_service: MethodServiceName,
                route: GetUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        });

        decodePromise.then(function (result) {
            var actionResult = GetUserContract(req, result);
            if (!(actionResult instanceof Promise)) {
                res.writeHead(httpStatus["bad request"], {
                    "Content-Type": "application/json"
                });

                res.write(JSON.stringify(JSONErrorResponse(RequestDecodingError, httpStatus["bad request"], "method failed to return promise", new Error("Promise expected as return from method"), {
                    url: req.url,
                    headers: req.headers,
                    http_method: req.method,
                    package: "github.com/gokit/rpkit/examples/users",
                    api_base: BaseServiceName,
                    method: "GetUser",
                    api_service: MethodServiceName,
                    route: GetUserServiceRoute,
                    api: "users.UserService"
                })));

                res.end();
                return;
            }

            return actionResult;
        }).catch(function (err) {
            if (res.finished) {
                return;
            }

            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ActinError, httpStatus["bad request"], "method returned with an error", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUser",
                api_service: MethodServiceName,
                route: GetUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function (model) {
            return options.Encoder.Encode(req, res, model);
        }).catch(function (err) {
            res.writeHead(httpStatus["bad request"], {
                "Content-Type": "application/json"
            });

            res.write(JSON.stringify(JSONErrorResponse(ResponseEncodingError, httpStatus["bad request"], "failed to encode response", err, {
                url: req.url,
                headers: req.headers,
                http_method: req.method,
                package: "github.com/gokit/rpkit/examples/users",
                api_base: BaseServiceName,
                method: "GetUser",
                api_service: MethodServiceName,
                route: GetUserServiceRoute,
                api: "users.UserService"
            })));

            res.end();
        }).then(function () {
            res.end();
        });

        if (next && typeof next == "function") next();
    };
};

/***/ }),

/***/ "http":
/*!***********************!*\
  !*** external "http" ***!
  \***********************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("http");

/***/ }),

/***/ "statuses":
/*!*******************************************************************************************!*\
  !*** external {"commonjs":"statuses","commonjs2":"statuses","amd":"statuses","root":"_"} ***!
  \*******************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = __WEBPACK_EXTERNAL_MODULE_statuses__;

/***/ }),

/***/ "url-parse":
/*!**********************************************************************************************!*\
  !*** external {"commonjs":"url-parse","commonjs2":"url-parse","amd":"url-parse","root":"_"} ***!
  \**********************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = __WEBPACK_EXTERNAL_MODULE_url_parse__;

/***/ })

/******/ });
});
//# sourceMappingURL=index.js.map