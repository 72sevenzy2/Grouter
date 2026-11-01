package core

// this package will hold shared types from the router and the middlewares, to avoid import cycles upon instillation.

import (
	"net/http"
)

// Param is to store route parameters.
type Param struct {
	Key   string
	Value string
}
type Params []Param
type Request struct { // shared
	*http.Request
	Params Params // holding route parameters
}

// retrieving particular route parameters.
func (r *Request) Get(key string) (string, bool) {
	for _, param := range r.Params {
		if param.Key == key {
			return param.Value, true
		}
	}
	return "", false
}

type HandlerFunc func(http.ResponseWriter, *Request)
type Middleware func(HandlerFunc) HandlerFunc // the middleware type (takes in the current handler and returns a new one)

// Route struct to loop over dynamic routes
type Route struct {
	Handler HandlerFunc
	Parts   []string
	Method  string
}

// Router struct to hold all static/dynamic routes
type Grouter struct { // shared
	DynamicRoutes map[string][]Route // split dynamic routes by methods to reduce lookup time
	Middlewares   []Middleware       // storing our middlewares here (type is our Middleware function type)
	StaticRoutes  map[string]map[string]HandlerFunc
}
