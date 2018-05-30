require("source-map-support").install();
(function webpackUniversalModuleDefinition(root, factory) {
	if(typeof exports === 'object' && typeof module === 'object')
		module.exports = factory(require("axios"), require("lodash"), require("safe-buffer"));
	else if(typeof define === 'function' && define.amd)
		define(["axios", "lodash", "safe-buffer"], factory);
	else if(typeof exports === 'object')
		exports["userservicerpjsc"] = factory(require("axios"), require("lodash"), require("safe-buffer"));
	else
		root["userservicerpjsc"] = factory(root["_"], root["_"], root["_"]);
})(global, function(__WEBPACK_EXTERNAL_MODULE_axios__, __WEBPACK_EXTERNAL_MODULE_lodash__, __WEBPACK_EXTERNAL_MODULE_safe_buffer__) {
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
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
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
exports.FromPromise = FromPromise;
exports.NewRequest = NewRequest;
exports.PokeClient = PokeClient;
exports.PokeAgainClient = PokeAgainClient;
exports.GetClient = GetClient;
exports.GetUsersMethodUserSliceFactory = GetUsersMethodUserSliceFactory;
exports.GetUsersClient = GetUsersClient;
exports.CreateMethodUserFactory = CreateMethodUserFactory;
exports.CreateMethodNewUserFactory = CreateMethodNewUserFactory;
exports.CreateClient = CreateClient;
exports.GetByClient = GetByClient;
exports.CreateUserMethodUserFactory = CreateUserMethodUserFactory;
exports.CreateUserMethodNewUserFactory = CreateUserMethodNewUserFactory;
exports.CreateUserClient = CreateUserClient;
exports.GetUserMethodUserFactory = GetUserMethodUserFactory;
exports.GetUserClient = GetUserClient;
// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.


// const URL = require("url-parse");
// const httpStatus = require("statuses");
// const request = require("request");
var buffer = __webpack_require__(/*! safe-buffer */ "safe-buffer");
var lodash = __webpack_require__(/*! lodash */ "lodash");
var axios = __webpack_require__(/*! axios */ "axios");

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

// JSONEncoding defines a JSON encoding implementation
// meant to work with the Request from NewRequest and
// response from using https://github.com/axios/axios
// as transport.
var JSONEncoding = exports.JSONEncoding = Object.freeze({
    Decode: function Decode(req, res, body) {
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
    Encode: function Encode(req, model) {
        return new Promise(function EncodePromise(resolve, reject) {
            req.setHeader("Content-Type", "application/json");
            req.setHeader("Accept", "application/json");

            try {
                var encoded = JSON.stringify(model);
                req.write(encoded);
            } catch (e) {
                reject(e);
                return;
            }

            req.end();
            resolve(req);
        });
    }
});

// HTTPTransport implements a custom http transport for handling
// request processing and response returning. This allows user
// to plug a custom http request transport handling layer
// as desired. An HTTPTransport must always return;
// a promise. The http transport uses https://github.com/axios/axios
// underneath.
var HTTPTransport = exports.HTTPTransport = Object.freeze({
    Do: function Do(req, timeout) {
        var ops = axios.request({
            method: "POST",
            timeout: timeout,
            data: req.body,
            url: joinURL(req.base, req.url),
            headers: lodash.merge(req.headers, { "X-Requested-With-Lib": "axios (https://github.com/axios/axios)" })
        });

        return ops.then(function (res) {
            if (!res) {
                return Promise.reject(new Error("received null/nil response object"));
            }

            if (isError(res)) {
                return Promise.reject(res);
            }

            if (requestRedirected(res) || requestFailed(res) || requestFacedInternalIssues(res)) {
                return Promise.reject(getJSONError(res.data));
            }

            return {
                res: res,
                body: res.data
            };
        }).catch(function (e) {
            if (e.response && e.response.data) {
                var msg = getJSONError(e.response.data);
                msg.original_error = e;
                return Promise.reject(msg);
            }
            return Promise.reject(e);
        });
    }
});

// FromPromise returns a new Promise with the
// function arguments passed to it. This
// is provided to allow clients use same
// promise as package has their are internal
// checks and use of `instanceof` with the
// Promise used.
function FromPromise(rx, ry) {
    return new Promise(rx, ry);
}

// NewRequest returns a object containing basic data related
// to a outgoing request with associated headers to be sent along.
// It is to be used by http transport for fullfilling a request.
function NewRequest(base, url, headers) {
    var reqObj = {
        base: base,
        url: url,
        headers: headers || {},
        body: null,
        setHeader: function setHeader(key, value) {
            reqObj.headers[key] = value;
        },
        write: function write(data) {
            if (reqObj.body) {
                reqObj.push(data);
                return;
            }

            reqObj.body = [];
            reqObj.body.push(data);
        },
        end: function end() {
            if (reqObj.body) reqObj.body = reqObj.body.join("");
        }
    };

    return reqObj;
}

function isError(err) {
    if (lodash.isError(err)) {
        return true;
    }
    return err instanceof Error;
}

// joinURL joins provided string paths into a whole,
// ensuring to add the necessary slash in between parts.
function joinURL(base, to) {
    var toURL = to;
    if (toURL.startsWith("/")) {
        toURL = to.substring(1);
    }

    if (base.endsWith("/")) {
        return ("" + base + toURL).trim();
    }

    return (base + "/" + toURL).trim();
}

// prefixSlash prefixes a forward slash to the beginning of provided
// string.
function prefixSlash(c) {
    if (c.startsWith("/")) {
        return c;
    }
    return "/" + c;
}

// suffixSlash suffixes a forward slash to end of provided
// string.
function suffixSlash(c) {
    if (c.endsWith("/")) {
        return c;
    }
    return "" + c;
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
    if ((typeof body === "undefined" ? "undefined" : _typeof(body)) === "object") {
        return body;
    }

    var res = null;
    if (body instanceof buffer.Buffer) {
        try {
            var data = body.toString("utf8");
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
var PokeServiceRoute = exports.PokeServiceRoute = "users.UserService/Poke";

// PokeServiceRoutePath defines the full method path for the Poke method.
var PokeServiceRoutePath = exports.PokeServiceRoutePath = "/rpkit/users.UserService/Poke";

// PokeContractSource contains the source version of expected method contract.
var PokeContractSource = exports.PokeContractSource = "type PokeMethodContract interface {\n\tPoke() \n}";

// PokeClient returns a RPC method to be called to handle requests
// for the Poke method of "github.com/gokit/rpkit/examples/users".
function PokeClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function PokeRPC() {
        return new Promise(function GetUserPromise(resolve, reject) {
            var req = NewRequest(options.ServiceAddr, PokeServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "Poke");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", PokeServiceRoute);
            req.setHeader("X-API-Client-Route-Path", PokeServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            options.Transport.Do(req, options.Timeout).then(function (resObj) {
                resolve(null);
            }).catch(function (err) {
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
var PokeAgainServiceRoute = exports.PokeAgainServiceRoute = "users.UserService/PokeAgain";

// PokeAgainServiceRoutePath defines the full method path for the PokeAgain method.
var PokeAgainServiceRoutePath = exports.PokeAgainServiceRoutePath = "/rpkit/users.UserService/PokeAgain";

// PokeAgainContractSource contains the source version of expected method contract.
var PokeAgainContractSource = exports.PokeAgainContractSource = "type PokeAgainMethodContract interface {\n\tPokeAgain()  error  \n}";

// PokeAgainClient returns a RPC method to be called to handle requests
// for the PokeAgain method of "github.com/gokit/rpkit/examples/users".
function PokeAgainClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function PokeAgainRPC() {
        return new Promise(function GetUserPromise(resolve, reject) {
            var req = NewRequest(options.ServiceAddr, PokeAgainServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "PokeAgain");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", PokeAgainServiceRoute);
            req.setHeader("X-API-Client-Route-Path", PokeAgainServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            options.Transport.Do(req, options.Timeout).then(function (resObj) {
                resolve(null);
            }).catch(function (err) {
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
var GetServiceRoute = exports.GetServiceRoute = "users.UserService/Get";

// GetServiceRoutePath defines the full method path for the Get method.
var GetServiceRoutePath = exports.GetServiceRoutePath = "/rpkit/users.UserService/Get";

// GetContractSource contains the source version of expected method contract.
var GetContractSource = exports.GetContractSource = "type GetMethodContract interface {\n\tGet(var1 context.Context)  (int,error)  \n}";

// GetClient returns a RPC method to be called to handle requests
// for the Get method of "github.com/gokit/rpkit/examples/users".
function GetClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function GetRPC() {
        return new Promise(function GetUserPromise(resolve, reject) {
            var req = NewRequest(options.ServiceAddr, GetServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "Get");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", GetServiceRoute);
            req.setHeader("X-API-Client-Route-Path", GetServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            options.Transport.Do(req, options.Timeout).then(function (resObj) {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then(function (resModel) {
                resolve(resModel);
            }).catch(function (err) {
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
var GetUsersServiceRoute = exports.GetUsersServiceRoute = "users.UserService/GetUsers";

// GetUsersServiceRoutePath defines the full method path for the GetUsers method.
var GetUsersServiceRoutePath = exports.GetUsersServiceRoutePath = "/rpkit/users.UserService/GetUsers";

// GetUsersContractSource contains the source version of expected method contract.
var GetUsersContractSource = exports.GetUsersContractSource = "type GetUsersMethodContract interface {\n\tGetUsers(var1 context.Context)  []users.User  \n}";

// GetUsersMethodUserSliceFactory defines a function to
// return a default object containing default field values of return value of
// GetUsers method.
function GetUsersMethodUserSliceFactory() {
    return JSON.parse("{\n\n\n    \"cid\":\t0.0,\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\"\n\n}");
}

// GetUsersClient returns a RPC method to be called to handle requests
// for the GetUsers method of "github.com/gokit/rpkit/examples/users".
function GetUsersClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function GetUsersRPC() {
        return new Promise(function GetUserPromise(resolve, reject) {
            var req = NewRequest(options.ServiceAddr, GetUsersServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "GetUsers");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", GetUsersServiceRoute);
            req.setHeader("X-API-Client-Route-Path", GetUsersServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            options.Transport.Do(req, options.Timeout).then(function (resObj) {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then(function (resModel) {
                resolve(resModel);
            }).catch(function (err) {
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
var CreateServiceRoute = exports.CreateServiceRoute = "users.UserService/Create";

// CreateServiceRoutePath defines the full method path for the Create method.
var CreateServiceRoutePath = exports.CreateServiceRoutePath = "/users/UserService/Create";

// CreateContractSource contains the source version of expected method contract.
var CreateContractSource = exports.CreateContractSource = "type CreateMethodContract interface {\n\tCreate(var1 context.Context,var2 users.NewUser)  (users.User,error)  \n}";

// CreateMethodUserFactory defines a function to
// return a default object containing default field values of return value of
// Create method.
function CreateMethodUserFactory() {
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0\n\n}");
}

// CreateMethodNewUserFactory defines a function to
// return a default object containing default field values of argument of
// Create method. 
function CreateMethodNewUserFactory() {
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}

// CreateClient returns a RPC method to be called to handle requests
// for the Create method of "github.com/gokit/rpkit/examples/users".
function CreateClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function CreateRPC(model) {
        return new Promise(function GetUserPromise(resolve, reject) {
            var req = NewRequest(options.ServiceAddr, CreateServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "Create");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", CreateServiceRoute);
            req.setHeader("X-API-Client-Route-Path", CreateServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            var encoded = options.Encoder.Encode(req, model);
            if (!(encoded instanceof Promise)) {
                reject(new Error("Encoder.Encode does not return a Promise"));
                return;
            }

            encoded.then(function (req) {
                return options.Transport.Do(req, options.Timeout);
            }).then(function (resObj) {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then(function (resModel) {
                resolve(resModel);
            }).catch(function (err) {
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
var GetByServiceRoute = exports.GetByServiceRoute = "users.UserService/GetBy";

// GetByServiceRoutePath defines the full method path for the GetBy method.
var GetByServiceRoutePath = exports.GetByServiceRoutePath = "/users/UserService/GetBy";

// GetByContractSource contains the source version of expected method contract.
var GetByContractSource = exports.GetByContractSource = "type GetByMethodContract interface {\n\tGetBy(var1 context.Context,var2 string)  (int,error)  \n}";

// GetByClient returns a RPC method to be called to handle requests
// for the GetBy method of "github.com/gokit/rpkit/examples/users".
function GetByClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function GetByRPC(model) {
        return new Promise(function GetUserPromise(resolve, reject) {
            var req = NewRequest(options.ServiceAddr, GetByServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "GetBy");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", GetByServiceRoute);
            req.setHeader("X-API-Client-Route-Path", GetByServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            var encoded = options.Encoder.Encode(req, model);
            if (!(encoded instanceof Promise)) {
                reject(new Error("Encoder.Encode does not return a Promise"));
                return;
            }

            encoded.then(function (req) {
                return options.Transport.Do(req, options.Timeout);
            }).then(function (resObj) {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then(function (resModel) {
                resolve(resModel);
            }).catch(function (err) {
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
var CreateUserServiceRoute = exports.CreateUserServiceRoute = "users.UserService/CreateUser";

// CreateUserServiceRoutePath defines the full method path for the CreateUser method.
var CreateUserServiceRoutePath = exports.CreateUserServiceRoutePath = "/users/UserService/CreateUser";

// CreateUserContractSource contains the source version of expected method contract.
var CreateUserContractSource = exports.CreateUserContractSource = "type CreateUserMethodContract interface {\n\tCreateUser(var1 users.NewUser)  (users.User,error)  \n}";

// CreateUserMethodUserFactory defines a function to
// return a default object containing default field values of return value of
// CreateUser method.
function CreateUserMethodUserFactory() {
    return JSON.parse("{\n\n\n    \"cid\":\t0.0,\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\"\n\n}");
}

// CreateUserMethodNewUserFactory defines a function to
// return a default object containing default field values of argument of
// CreateUser method.
function CreateUserMethodNewUserFactory() {
    return JSON.parse("{\n\n\n    \"name\":\t\"\"\n\n}");
}

// CreateUserClient returns a RPC method to be called to handle requests
// for the CreateUser method of "github.com/gokit/rpkit/examples/users".
function CreateUserClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function CreateUserRPC(model) {
        return new Promise(function GetUserPromise(resolve, reject) {
            var req = NewRequest(options.ServiceAddr, CreateUserServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "CreateUser");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", CreateUserServiceRoute);
            req.setHeader("X-API-Client-Route-Path", CreateUserServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            var encoded = options.Encoder.Encode(req, model);
            if (!(encoded instanceof Promise)) {
                reject(new Error("Encoder.Encode does not return a Promise"));
                return;
            }

            encoded.then(function (req) {
                return options.Transport.Do(req, options.Timeout);
            }).then(function (resObj) {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then(function (resModel) {
                resolve(resModel);
            }).catch(function (err) {
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
var GetUserServiceRoute = exports.GetUserServiceRoute = "users.UserService/GetUser";

// GetUserServiceRoutePath defines the full method path for the GetUser method.
var GetUserServiceRoutePath = exports.GetUserServiceRoutePath = "/users/UserService/GetUser";

// GetUserContractSource contains the source version of expected method contract.
var GetUserContractSource = exports.GetUserContractSource = "type GetUserMethodContract interface {\n\tGetUser(var1 int)  (users.User,error)  \n}";

// GetUserMethodUserFactory defines a function to
// return a default object containing default field values of return value of
// GetUser method.
function GetUserMethodUserFactory() {
    return JSON.parse("{\n\n\n    \"id\":\t0,\n\n    \"name\":\t\"\",\n\n    \"addr\":\t\"\",\n\n    \"cid\":\t0.0\n\n}");
}

// GetUserClient returns a RPC method to be called to handle requests
// for the GetUser method of "github.com/gokit/rpkit/examples/users".
function GetUserClient(options) {
    if (options.BeforeRequest && typeof options.BeforeRequest !== "function") {
        throw new Error("options.BeforeRequest must be a function");
    }

    if (options.Headers && _typeof(options.Headers) !== "object") {
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

    if (!options.Headers) options.Headers = {};
    if (!options.Encoder) options.Encoder = JSONEncoding;
    if (!options.Decoder) options.Decoder = JSONEncoding;
    if (!options.Transport) options.Transport = HTTPTransport;

    return function GetUserRPC(model) {
        return new Promise(function GetUserPromise(resolve, reject) {
            var req = NewRequest(options.ServiceAddr, GetUserServiceRoutePath, options.Headers);
            req.setHeader("X-Client", "JS-RPKIT");
            req.setHeader("X-Service", BaseServiceName);
            req.setHeader("X-Package", "github.com/gokit/rpkit/examples/users");
            req.setHeader("X-Method-Client", "GetUser");
            req.setHeader("X-Method-ClientService", MethodServiceName);
            req.setHeader("X-API-Client-Route", GetUserServiceRoute);
            req.setHeader("X-API-Client-Route-Path", GetUserServiceRoutePath);
            req.setHeader("X-Client-Package-Interface", "users.UserService");

            if (options.BeforeRequest) options.BeforeRequest(req);

            var encoded = options.Encoder.Encode(req, model);
            if (!(encoded instanceof Promise)) {
                reject(new Error("Encoder.Encode does not return a Promise"));
                return;
            }

            encoded.then(function (req) {
                return options.Transport.Do(req, options.Timeout);
            }).then(function (resObj) {
                return options.Decoder.Decode(req, resObj.res, resObj.body);
            }).then(function (resModel) {
                resolve(resModel);
            }).catch(function (err) {
                reject(err);
            });
        });
    };
}

/***/ }),

/***/ "axios":
/*!**********************************************************************************!*\
  !*** external {"commonjs":"axios","commonjs2":"axios","amd":"axios","root":"_"} ***!
  \**********************************************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = __WEBPACK_EXTERNAL_MODULE_axios__;

/***/ }),

/***/ "lodash":
/*!*************************************************************************************!*\
  !*** external {"commonjs":"lodash","commonjs2":"lodash","amd":"lodash","root":"_"} ***!
  \*************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = __WEBPACK_EXTERNAL_MODULE_lodash__;

/***/ }),

/***/ "safe-buffer":
/*!****************************************************************************************************!*\
  !*** external {"commonjs":"safe-buffer","commonjs2":"safe-buffer","amd":"safe-buffer","root":"_"} ***!
  \****************************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = __WEBPACK_EXTERNAL_MODULE_safe_buffer__;

/***/ })

/******/ });
});
//# sourceMappingURL=index.js.map