<h1 align="center">Grouter.</h1>
<h2 align="center">An http router built from scratch, ontop of the golang stdlib.</h2>

# Features:

- Route grouping for nested routes.
- RFC-Compliant route management.
- 0 external dependencies used (excluding ones made for Grouter).
- Includes middleware, and also supports route specific middleware.
- Route parameter compatibility.
- Wildcard route compatability.
- superiority with speed (ns/op) compared to golangs stdlib router (see benchmarks/results/).

  <h2>benchmark specifications:</h2>
  
	```
	CPU: cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
 	OS: windows
 	Arch: amd64
 	Frequency: 10 reps for both grouter and mux.
	```



<h2 align="center">pre-requisite:</h2>
<h3 align="center">when writing handlers, do make sure your using grouter's request struct instead of http.Request, which is grouter.Request</h3>
<br>
<h1 align="center">Getting started</h1>

<h1 align="center">instillation</h1>

```
go get github.com/72sevenzy2/grouter/grouter
```

<br>
<h1 align="center">General usage:</h1>

```
package main

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/grouter/grouter"
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

<h2 align="center">handling dynamic routes:</h2>

```
package main

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/grouter/grouter"
)

func main() {
	r := grouter.NewGrouter()

	// dynamic prefixes start with ":"
	r.Handle(http.MethodGet, "/test/:id", func(w http.ResponseWriter, r *grouter.Request) {
		val, ok := r.Get("id")
		if !ok {
			panic("route parameter does not exist.")
		}
		fmt.Println(val)
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
```

<h2 align="center">handling wildcard routes:</h2>

```
package main

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/grouter/grouter"
)

func main() {
	r := grouter.NewGrouter()

	// wildcard prefixes start with '*'
	r.Handle(http.MethodGet, "/test/*filePath", func(w http.ResponseWriter, r *grouter.Request) {
		// val -> whatever path was given after '/test/'.
		val, ok := r.Get("filePath")
		if !ok {
			panic("route parameter does not exist.")
		}
		fmt.Println(val)

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
```

<h2 align="center">route grouping with .Group():</h2>

```
package main

import (
	"net/http"

	"github.com/72sevenzy2/grouter/grouter"
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

<h2 align="center">lastly, to handle route-specific middleware:</h2>

```
package main

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/grouter/grouter"
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
<br>
<h2 align="center">for more info on proper middleware usage, see https://github.com/72sevenzy2/Grouter-middlewares/blob/main/README.md</h2>
