package grouter

import (
	"net/http"
)

// pre-defined html templates.

func Default404(w http.ResponseWriter) error {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(404)
	w.Write([]byte(`
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>404 - page not found</title>
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
        <h1>404</h1>
        <h2>page not found</h2>
        <p>Sorry, the page you are looking for doesn't exist, has been removed, or the URL was typed incorrectly.</p>
    </div>
</body>
</html>
	`))
	return nil
}
