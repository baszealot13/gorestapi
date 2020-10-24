# GO REST API

This is Framework concept with Golang, Gin, Gorm and etc

## Contents
- Rest API Framework
  - JWT Authentication
  - Redis + Session

## Install Golang
https://golang.org/doc/install

## Structural
```bash
|-- src
    |-- config
        |-- config.{environment}.yaml
        |-- database.go
    |-- controllers
        |-- ...
        |-- *.controller.go
    |-- entities
        |-- ...
        |-- *.entitiy.go
    |-- helpers
        |-- ...
        |-- *.helper.go
    |-- middlewares
       |-- ...
       |-- *.go
    |-- routes
       |-- ...
       |-- route.go
    |-- services
       |-- ...
       |-- *.service.go
    |-- main.go
```

## Run Project
#### Development
```
go run src/main.go
```

#### Hello world!
http://localhost:8080/public/api

#### Production
```
export ENV=prod; go run src/main.go
``` 