# Briefly.public

## Introduction
This repository is part of the project briefly.
Briefly project is a URL shortcoder.

This repository concern the REST service to expose with gin-tonic golang library 

## Endpoints
### shortcode URL

POST method:  /shortcoder

####  Body content

POST json body: 

```json
{"url": "http://www.google.fr"}
```
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
$ go build -o bin/briefly.public
```
## Usage
### Pre-requisite
The briefly gRPC server have to run to execute the backend work.

Please clone and follow documentation of [briefly gRPC](https://github.com/francois-poidevin/briefly), to have a running gRPC server.

### Generate configuration file
```shell
$ ./bin/briefly.public config new
```

will generate a default configuration file

```toml
###############################
# Briefly.public Settings 
###############################
[Briefly_public]

  ###############################
  # Briefly gRPC API settings 
  ###############################
  [Briefly_public.Briefly]

    # URL and Port for the Briefly gRPC server
    adress = "localhost:5556"

  ###############################
  # REST API settings 
  ###############################
  [Briefly_public.REST]

    # On which port REST HTTP service will listen
    listenPort = ":8080"

###############################
# Logs Settings 
###############################
[Log]

  # Allow to display logs in Json format if true
  jsonformatter = false

  # Log level: trace, debug, info, warning, error, panic, and fatal
  level = "warning"

```

store it as toml file in example 'conf' folder in './conf/briefly.public.conf.toml' file

### Run HTTP server
```shell
$ ./bin/briefly.public startHttp --config ./conf/briefly.public.conf.toml
```

### Call endpoints
#### shortcode URL endpoint
```shell
$ curl -X POST  http://127.0.0.1:8080/shortcoder/ -H 'content-type: application/json' -d '{ "url": "https://www.google.com" }'
2073697506084
```

#### unshortcode URL endpoint
```shell
$ curl 127.0.0.1:8080/unshortcoder/2073697506084 
<a href="https://www.google.com">Moved Permanently</a>.
```



