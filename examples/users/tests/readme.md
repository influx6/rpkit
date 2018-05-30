## Users Tests

## Run Go Tests

Simply navigate to this directory and run the following:

```bash
> go test -v ./goside/...
```

You should get:

```bash
λ go test .\examples\users\tests\... -v
=== RUN   TestUserServiceRP_PokeAgain
--- PASS: TestUserServiceRP_PokeAgain (0.01s)
=== RUN   TestUserServiceRP_Poke
--- PASS: TestUserServiceRP_Poke (0.00s)
=== RUN   TestUserServiceRP_Get
2018/05/30 15:46:08 http: multiple response.WriteHeader calls
--- PASS: TestUserServiceRP_Get (0.00s)
=== RUN   TestUserServiceRP_GetBy
2018/05/30 15:46:08 http: multiple response.WriteHeader calls
--- PASS: TestUserServiceRP_GetBy (0.00s)
=== RUN   TestUserServiceRP_GetUsers
2018/05/30 15:46:08 http: multiple response.WriteHeader calls
--- PASS: TestUserServiceRP_GetUsers (0.01s)
=== RUN   TestUserServiceRP_GetUser
2018/05/30 15:46:09 http: multiple response.WriteHeader calls
--- PASS: TestUserServiceRP_GetUser (0.00s)
=== RUN   TestUserServiceRP_Create
2018/05/30 15:46:09 http: multiple response.WriteHeader calls
--- PASS: TestUserServiceRP_Create (0.00s)
=== RUN   TestUserServiceRP_CreateUser
2018/05/30 15:46:09 http: multiple response.WriteHeader calls
--- PASS: TestUserServiceRP_CreateUser (0.00s)
=== RUN   TestUserServiceRP_CreateUserWithCreateMethodClient
--- PASS: TestUserServiceRP_CreateUserWithCreateMethodClient (0.00s)
PASS
ok      github.com/gokit/rpkit/examples/users/tests/goside      0.111s
```

## Run Go Server and Javascript Client Tests with Nodejs

Simple ensure [Nodejs](https://nodejs.org/) has being installed, then navigate to this directory and run:

```bash
> npm install -g mocha
```

Open another terminal/command-line and navigate it to this same directory and run the go server:

```bash
go run main.go
```

Open another terminal and run:

```bash
> mocha ./jsside/userservicerp.test.js
```

You should get:

```bash
λ mocha userservericerp.test.js


  UserService::GetUsers
    √ Should be able to get a response with RPC client (39ms)

  UserService::GetUser
    √ Should be able to get a response with RPC client

  UserService::CreateUser
    √ Should be able to get a response with RPC client

  UserService::Create
    √ Should be able to get a response with RPC client

  UserService::GetBy
    √ Should be able to get a response with RPC client

  UserService::PokeAgain
    √ Should be able to get a response with RPC client

  UserService::Poke
    √ Should be able to get a response with RPC client

  UserService::Get
    √ Should be able to get a response with RPC client


  8 passing (132ms)
```

## Run Javascript Server and Client Tests with Nodejs

Simple ensure [Nodejs](https://nodejs.org/) has being installed, then navigate to this directory and run:

```bash
> npm install -g mocha
> mocha ./jsside/userservicerp.test.js
```

You should get:

```bash
λ mocha userservericerp.test.js


  UserService::GetUsers
    √ Should be able to get a response with RPC client (39ms)

  UserService::GetUser
    √ Should be able to get a response with RPC client

  UserService::CreateUser
    √ Should be able to get a response with RPC client

  UserService::Create
    √ Should be able to get a response with RPC client

  UserService::GetBy
    √ Should be able to get a response with RPC client

  UserService::PokeAgain
    √ Should be able to get a response with RPC client

  UserService::Poke
    √ Should be able to get a response with RPC client

  UserService::Get
    √ Should be able to get a response with RPC client


  8 passing (132ms)
```
