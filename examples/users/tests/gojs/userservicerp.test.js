"use strict";

process.env.NODE_ENV = "test";

const http = require("http");
const userServiceServer = require("../../userservicerpjss");
const userServiceClient = require("../../userservicerpjsc");
const chai = require("chai");
const expect = chai.expect;

describe("UserService::GetUsers", () => {
    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.GetUsersClient({
            ServiceAddr: "http://localhost:7888",
            Timeout: 1990,
            Headers: {
                "Content-Type": "application/no-content",
                "Accept": "application/json",
            },
        });

        expect(rpc).to.not.equal(null);

        rpc().then((resources) => {
            expect(resources).to.not.be.null;
            expect(resources).to.not.be.undefined;
            expect(resources).to.be.an("array").that.is.length(1);

            const res = resources[0];
            expect(res.name).to.equal("Users");
            expect(res.cid).to.equal(0.1);
            expect(res.addr).to.equal("LA");
            expect(res.id).to.equal(1);
            done();
        }).catch((e) => {

            expect(e).to.equal(null);
            done();
        });
    });
});


describe("UserService::GetUser", () => {
    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.GetUserClient({
            ServiceAddr: "http://localhost:7888",
            Timeout: 1990,
            Headers: {
                "Content-Type": "application/no-content",
                "Accept": "application/json",
            },
        });

        expect(rpc).to.not.equal(null);

        rpc(1).then((res) => {
            expect(res).to.not.be.null;
            expect(res).to.not.be.undefined;
            expect(res.name).to.equal("Users");
            expect(res.cid).to.equal(0.1);
            expect(res.addr).to.equal("LA");
            expect(res.id).to.equal(1);
            done();
            return res;
        }).catch((e) => {

            expect(e).to.be.null;
            done();
        });
    });
});

describe("UserService::CreateUser", () => {
    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.CreateUserClient({
            ServiceAddr: "http://localhost:7888",
            Timeout: 1990,
            Headers: {
                "Content-Type": "application/no-content",
                "Accept": "application/json",
            },
        });

        expect(rpc).to.not.equal(null);

        rpc({
            name: "Simon"
        }).then((res) => {
            expect(res).to.not.be.null;
            expect(res).to.not.be.undefined;
            expect(res.name).to.equal("Simon");
            expect(res.cid).to.equal(0.1);
            expect(res.addr).to.equal("LA");
            expect(res.id).to.equal(1);
            done();
        }).catch((e) => {

            expect(e).to.equal(null);
            done();
        });
    });
});

describe("UserService::Create", () => {
    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.CreateClient({
            ServiceAddr: "http://localhost:7888",
            Timeout: 1990,
            Headers: {
                "Content-Type": "application/no-content",
                "Accept": "application/json",
            },
        });

        expect(rpc).to.not.equal(null);

        rpc({
            name: "Simon"
        }).then((res) => {
            expect(res).to.not.be.null;
            expect(res).to.not.be.undefined;
            expect(res.name).to.equal("Simon");
            expect(res.cid).to.equal(0.1);
            expect(res.addr).to.equal("LA");
            expect(res.id).to.equal(1);
            done();
        }).catch((e) => {

            expect(e).to.equal(null);
            done();
        });
    });
});


describe("UserService::GetBy", () => {
    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.GetByClient({
            ServiceAddr: "http://localhost:7888",
            Timeout: 1990,
            Headers: {
                "Content-Type": "application/no-content",
                "Accept": "application/json",
            },
        });

        expect(rpc).to.not.equal(null);

        rpc("Simon").then((res) => {
            expect(res).to.not.be.null;
            expect(res).to.not.be.undefined;
            expect(res).to.equal(1);
            done();
        }).catch((e) => {

            expect(e).to.equal(null);
            done();
        });
    });
});

describe("UserService::PokeAgain", () => {
    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.PokeAgainClient({
            ServiceAddr: "http://localhost:7888",
            Timeout: 1990,
            Headers: {
                "Content-Type": "application/no-content",
                "Accept": "application/json",
            },
        });

        expect(rpc).to.not.equal(null);

        rpc().catch((e) => {
            expect(e).to.not.be.null;
            done();
        });
    });
});

describe("UserService::Poke", () => {
    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.PokeClient({
            ServiceAddr: "http://localhost:7888",
            Timeout: 1990,
            Headers: {
                "Content-Type": "application/no-content",
                "Accept": "application/json",
            },
        });

        expect(rpc).to.not.equal(null);

        rpc().then((res) => {
            expect(res).to.be.null;
            done();
        }).catch((e) => {
            console.log("PokeError: ", e);
            expect(e).to.equal(null);
            done();
        });
    });
});

describe("UserService::Get", () => {
    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.GetClient({
            ServiceAddr: "http://localhost:7888",
            Timeout: 1990,
            Headers: {
                "Content-Type": "application/no-content",
                "Accept": "application/json",
            },
        });

        expect(rpc).to.not.equal(null);

        rpc().then((res) => {
            expect(res).to.not.be.null;
            expect(res).to.not.be.undefined;
            expect(res).to.equal(1);
            done();
        }).catch((e) => {

            expect(e).to.equal(null);
            done();
        });
    });
});