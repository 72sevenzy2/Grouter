package grouter

import (
	"net/http"
	"strings"

	"github.com/72sevenzy2/http-router/core"
)

// func to check whether route is dynamic or static
func isDynamic(path string) bool {
	return strings.Contains(path, ":") // strings.Contains() returns a boolean
}

type pathIter struct {
	path string
	pos  int
}

func NewPathIter(p string) *pathIter {
	return &pathIter{
		path: strings.Trim(p, "/"),
		pos:  0,
	}
}

func (i *pathIter) next() (string, bool) {
	if i.pos >= len(i.path) {
		return "", false
	}
	s := i.pos
	for i.pos < len(i.path) && i.path[i.pos] != '/' {
		i.pos++
	}
	routeSegmt := i.path[s:i.pos]
	if i.pos < len(i.path) {
		i.pos++
	}
	return routeSegmt, true
}

// func to split route path
func splitPath(path string) []string {
	// manual string indexing to avoid strings.genSplit overhead.
	path = strings.Trim(path, "/") // normalise from any other outter slashes.
	if path == "" {
		return nil
	}

	n := 1
	for v := 0; v < len(path); v++ {
		if path[v] == '/' {
			n++
		}
	}

	// preallocate according to n (number of parts in the request)
	parts := make([]string, 0, n)
	start := 0
	for v := 0; v < len(path); v++ {
		if path[v] == '/' {
			parts = append(parts, path[start:v])
			start = v + 1 // update start position after each append
		}
	}

	parts = append(parts, path[start:]) // ensures the last part URL (after final occurence of '/') is appended.
	return parts
}

// func to match request parts and route parts
func match(routeP []string, reqP string) (core.Params, error) {
	iter := NewPathIter(reqP)
	var params core.Params
	for _, rp := range routeP { // can use both reqP and routeP to loop over
		reqp, ok := iter.next()
		if !ok {
			return nil, InvalidPathError
		}
		if len(rp) > 0 && rp[0] == ':' {
			params = append(params, &core.Param{
				Key:   rp[1:],
				Value: reqp,
			})
			continue
		}

		// verifies whether request path matches registered route path
		// RFC-3986 states that path components must be treated case-sensitive, so we dont normalise them here.
		if rp != reqp {
			return nil, InvalidPathError
		}
	}
	if _, ok := iter.next(); ok {
		return nil, InvalidPathError
	}
	return params, nil
}

// path joining helper (for Grouping)
func Join(p1, p2 string) string {
	switch {
	case p1 == "/":
		return "/" + strings.TrimLeft(p2, "/")
	case p2 == "/":
		return strings.TrimRight(p1, "/") + "/"
	default:
		return strings.TrimRight(p1, "/") + "/" + strings.TrimLeft(p2, "/")
	}
}

// IsValidHTTPMethod determines whether a given method via router.Handle() or router.Group.Handle() is an valid http method.
func IsValidHTTPMethod(m string) bool {
	switch m {
	case http.MethodDelete,
		http.MethodGet,
		http.MethodConnect,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
		http.MethodTrace:
		return true

	default:
		return false
	}
}
