package grouter

import (
	"fmt"
	"net/http"
)

type ErrorPage struct {
	Code       int
	Title      string
	Heading    string
	Descripion string
}

func DisplayStatus(w http.ResponseWriter, ep ErrorPage) (int, error) {
	w.Header().Set("content-type", "text/html; charset='utf-8")
	w.WriteHeader(ep.Code)

	return w.Write([]byte(fmt.Sprintf(`
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <style>
        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
            background-color: #f8f9fa;
            color: #333333;
            display: flex;
            align-items: center;
            justify-content: center;
            min-height: 100vh;
            text-align: center;
            padding: 20px;
        }
        .container {
            max-width: 500px;
            background: #ffffff;
            padding: 40px 30px;
            border-radius: 8px;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
        }
        h1 {
            font-size: 72px;
            font-weight: 700;
            color: #e63946;
            margin-bottom: 10px;
        }
        h2 {
            font-size: 24px;
            margin-bottom: 15px;
            color: #1d3557;
        }
        p {
            font-size: 15px;
            color: #6c757d;
            line-height: 1.6;
            margin-bottom: 25px;
        }
        a.btn {
            display: inline-block;
            background-color: #1d3557;
            color: #ffffff;
            text-decoration: none;
            padding: 12px 24px;
            border-radius: 4px;
            font-weight: 600;
            transition: background-color 0.2s ease;
        }
        a.btn:hover {
            background-color: #457b9d;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>%s</h1>
        <h2>%s</h2>
    </div>
</body>
</html>
`, ep.Title, ep.Heading, ep.Descripion)))
}

// pre-defined html templates.

func Display404(w http.ResponseWriter) (int, error) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(404)
	return DisplayStatus(w, ErrorPage{
		Code:       http.StatusNotFound,
		Title:      "404 - not found",
		Heading:    "page not found",
		Descripion: "Sorry, the page you were looking for was not found.",
	})
}

func DisplayUnauthorized(w http.ResponseWriter) (int, error) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(401)
	return DisplayStatus(w, ErrorPage{
		Code:       http.StatusUnauthorized,
		Title:      fmt.Sprintf("status %d", http.StatusUnauthorized),
		Heading:    "status unauthorized",
		Descripion: "Sorry, you arent eligible to visit the page you were looking for.",
	})
}

func DisplayForbidden(w http.ResponseWriter) (int, error) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(403)
	return DisplayStatus(w, ErrorPage{
		Code:       http.StatusForbidden,
		Title:      fmt.Sprintf("status %d", http.StatusForbidden),
		Heading:    "status forbidden.",
		Descripion: "Sorry, you arent allowed to visit the page you were looking for.",
	})
}

func DisplayInternalServerError(w http.ResponseWriter) (int, error) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(500)
	return DisplayStatus(w, ErrorPage{
		Code:       http.StatusInternalServerError,
		Title:      fmt.Sprintf("status %d", http.StatusInternalServerError),
		Heading:    "status internal server error.",
		Descripion: "Sorry, the page you were looking for is facing internal issues, please try again later.",
	})
}
