# A HTTP Client for a API REST

## Intro
HTTP Client is a challenge that aims to abstract the use of an API REST covering the methods GET, POST, and DELETE. 
It covers how to use the HTTP and Context packages and how to deal with errors. Also, it includes an implementation of Option Parttern for creating the client.

This challenge was proposed by [DevGym](https://app.devgym.com.br/challenges/9bcad7c4-a809-4ef5-929d-a000aede5b25)

## How to run
First, clone this repository
```
    git clone https://github.com/mauriciomd/http-client-lib.git
```


You'll need Docker to run the API. Once you have it installed, use the command bellow to start the API. 
```
    docker-compose up
```

Then, just run the tests by running
```
    go test ./...
```

