<h1 align="center">Grouter.</h1>
<h2 align="center">An http router built from scratch, ontop of the golang stdlib.</h2>

# Features:

- Route grouping for nested routes.
- RFC-Compliant route management.
- 0 external dependencies used (excluding ones made for Grouter).
- Includes middleware, and also supports route specific middleware.
- Route parameter compatibility.
- proven superiority in speed compared to golangs stdlib router (see benchmarks/results/).

  <h2>benchmark specifications:</h2>
  
	```
	CPU: cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
 	OS: windows
 	Arch: amd64
 	Frequency: 10 reps for both grouter and mux.
	```
<br>
 <h1 align="center">Getting started</h1>
<h2>pre-requisites:</h2>
<h4>when writing handlers, do make sure your using the custom Request struct instead of http.Request, which is grouter.Request</h4>
<h4>When handling dynamic routes, the route parameter values are stored in inside the grouter.Request struct ('Params' field), that can be retrieved like so:</h4>

```
 package main

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/Grouter/grouter"
)

func main() {
	r := grouter.NewGrouter()

	r.Get("/user/:64", func(w http.ResponseWriter, r *grouter.Request) {
		HandleThisFunc()...
		param, ok := r.Params["user"]
		if !ok {
			handleErr()
		}
	})

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

```
<br>

<h1 align="center">General usage:</h1>

```
 package main

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/Grouter/grouter"
)

func main() {
	r := grouter.NewGrouter()

	r.Get("/resp", func(w http.ResponseWriter, r *grouter.Request) {
		w.Write([]byte("responded"))
	})

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

```
<h2 align="center">Grouping routes:</h2>

```
package main

import (
	"net/http"

	"github.com/72sevenzy2/Grouter/grouter"
)

func main() {
	r := grouter.NewGrouter()

	api := r.Group("/parent")
	/* 
	for inline nesting (without having to create seperate child routes individually):
	
	api := r.Group("/parent", "child1", "child2")
	(in which the registered route would be /parent/child1/child2 )

	though this prevents configuring middleware for a specific route
	*/

	v1 := api.Group("/child1")
	v1.Use(router.Logger(0)) // you can use middleware with specific route inside nested routes.

	v2 := v1.Group("/child2")
	v2.Use(grouter.BasicAuth("username", "password")) // would indicate only this route having this middleware

	v2.Get("/test", func(w http.ResponseWriter, r *grouter.Request) {
		w.Write([]byte("pong"))
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
```

<br>
<h2 align="center" style.Background="gray">Middleware(s) usage:</h2>

```

package main

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/Grouter/grouter"
)

func main() {
	r := grouter.NewGrouter()

	limiter := grouter.NewLimiter(100, 1) // NewLimiter() takes in a maximum number of requests, and how often each token is refilled for      // the client.

	r.Use(limiter.RateLimiter())
	
	r.Use(grouter.Recoverer()) // recoverer middleware always goes first, prevents server crashes when a bug has occured.
	r.Use(grouter.Logger(0)) // standard logging, in which param 0 indicates default body size logging.

	r.Use(grouter.BearerAuth("secretKey")) // can be any token (which has to be a string),
	r.Use(grouter.BasicAuth("user1", "password1234")) // parameters username and password need to be included when using.
	r.Use(grouter.Timeout(time.Second * 10)) //  depending on how long you want the time limit on every request.

	r.Get("/resp", func(w http.ResponseWriter, r *grouter.Request) {
		w.Write([]byte("responded"))
	})

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

```
^ Request body read from the logging middleware can be capped, the default size being 1KiB, but can be configured via passing in an limit in grouter.Logger(..), like such: "r.Use(grouter.Logger(1024 * 2)" (2 KiB cap), though be sure the limit passed in is of appropriate type (uint32).
<br>

<br>
<h2 align="center">Route-specific middleware:</h2>

```
package main

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/Grouter/grouter"
)

func main() {
	r := grouter.NewGrouter()

	r.Get("/greet", func(w http.ResponseWriter, r *grouter.Request) {
		w.Write([]byte("hello"))
	}, grouter.Recoverer(), grouter.Logger(0)) // you can do route-specific middleware(s) like this (can be 1 or many).

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

```
