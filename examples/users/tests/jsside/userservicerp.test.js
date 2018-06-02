"use strict";

process.env.NODE_ENV = "test";

const http = require("http");
const userServiceServer = require("../../userservicerpjss");
const userServiceClient = require("../../userservicerpjsc");
const chai = require("chai");
const expect = chai.expect;

describe("UserService::GetUsers", () => {
    const server = http.createServer(userServiceServer.GetUsersService((req) => {
        return Promise.resolve([{
            name: "Users",
            addr: "LA",
            cid: 0.1,
            id: 1,
        }]);
    }, {}));

    before(() => {
        server.listen(8050);
    });

    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.GetUsersClient({
            ServiceAddr: "http://localhost:8050",
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

    after(() => {
        server.close();
    });
});


describe("UserService::GetUser", () => {
    const server = http.createServer(userServiceServer.GetUserService((req, model) => {
        expect(model).to.equal(1);
        return Promise.resolve({
            name: "Users",
            addr: "LA",
            cid: 0.1,
            id: 1,
        });
    }, {}));

    before(() => {
        server.listen(8060);
    });

    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.GetUserClient({
            ServiceAddr: "http://localhost:8060",
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

    after(() => {
        server.close();
    });
});

describe("UserService::CreateUser", () => {
    const server = http.createServer(userServiceServer.CreateUserService((req, model) => {
        expect(model.name).to.equal("Simon");
        return Promise.resolve({
            name: model.name,
            addr: "LA",
            cid: 0.1,
            id: 1,
        });
    }, {}));

    before(() => {
        server.listen(8040);
    });

    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.CreateUserClient({
            ServiceAddr: "http://localhost:8040",
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

    after(() => {
        server.close();
    });
});

describe("UserService::Create", () => {
    const server = http.createServer(userServiceServer.CreateService((req, model) => {
        expect(model.name).to.equal("Simon");
        return Promise.resolve({
            name: model.name,
            addr: "LA",
            cid: 0.1,
            id: 1,
        });
    }, {}));

    before(() => {
        server.listen(8070);
    });

    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.CreateClient({
            ServiceAddr: "http://localhost:8070",
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

    after(() => {
        server.close();
    });
});


describe("UserService::GetBy", () => {
    const server = http.createServer(userServiceServer.GetByService((req, model) => {
        expect(model).to.equal("Simon");
        return Promise.resolve(1);
    }, {}));

    before(() => {
        server.listen(8030);
    });

    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.GetByClient({
            ServiceAddr: "http://localhost:8030",
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

    after(() => {
        server.close();
    });
});

describe("UserService::PokeAgain", () => {
    const server = http.createServer(userServiceServer.PokeAgainService((req) => {
        return Promise.reject(new Error("don't poke me again"));
    }, {}));

    before(() => {
        server.listen(8000);
    });

    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.PokeAgainClient({
            ServiceAddr: "http://localhost:8000",
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

    after(() => {
        server.close();
    });
});

describe("UserService::Poke", () => {
    const server = http.createServer(userServiceServer.PokeService((req) => {
        return Promise.resolve(1);
    }, {}));

    before(() => {
        server.listen(8000);
    });

    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.PokeClient({
            ServiceAddr: "http://localhost:8000",
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

    after(() => {
        server.close();
    });
});
describe("UserService::Get", () => {
    const server = http.createServer(userServiceServer.GetService((req) => {
        return Promise.resolve(1);
    }, {}));

    before(() => {
        server.listen(8000);
    });

    it("Should be able to get a response with RPC client", (done) => {
        const rpc = userServiceClient.GetClient({
            ServiceAddr: "http://localhost:8000",
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

    after(() => {
        server.close();
    });
});