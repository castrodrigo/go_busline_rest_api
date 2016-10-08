# go_busline_rest_api
Draf Rest API with Go

I'm using this code to start some drafts to a Bus Lines Rest API with GO.
In this repo you'll find a Postman collection.

# Installing Go

Follow the instructions of oficial Go website https://golang.org/doc/install

# Application

## Dependencies

The API needs some dependencies:

```
github.com/gorilla/mux
github.com/twinj/uuid
```

Install each of them with:

```
go get github.com/gorilla/mux
```
```
go get github.com/twinj/uuid
```

## Installing

Create a folder "github.com/castrodrigo/go_simple_rest_api" on your GO Work directory, and clone this repo.
Install the application:

```
go install github.com/castrodrigo/go_simple_rest_api
```

On your bin directory. Run the application:

```
$GOPATH/bin/go_simple_rest_api
```

## Accessing API

http://localhost:9009/busline/
