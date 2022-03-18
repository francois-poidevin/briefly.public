# Briefly.public

## Introduction
This repository is part of the project briefly.
Briefly project is a URL shortcoder.

This repository concern the REST service to expose with gin-tonic golang library 

## Endpoints
### shortcode URL

GET method:  /shortcoder/:url

####  Path parameter

url: URL to shortcode

#### Return reponse

GET a id of the URL in return HTTP code 200

### unshortcode URL

GET method: /unshortcoder/:hash

#### Path parameter

hash: id of the URL

#### Return reponse

A redirect HTTP code 301 with the unshortcoded URL associated to the path parameter hash

## Compilation

### Dependancies
```shell
$ go mod vendor
```

### Bin generation
```shell
$ go build -o bin/briefly
```
## Usage
### Pre-requisite
The briefly gRPC server have to run to execute the backend work.

Please clone and follow documentation of [briefly gRPC](https://github.com/francois-poidevin/briefly), to have a running gRPC server.

### Run HTTP server
```shell
$ ./bin/briefly.public startHttp
```

### Call endpoints
#### shortcode URL endpoint
```shell
$ curl 127.0.0.1:8080/unshortcoder/test
2073697506084
```

#### unshortcode URL endpoint
```shell
$ curl 127.0.0.1:8080/shortcoder/test 
<a href="https://www.google.com">Moved Permanently</a>.
```



