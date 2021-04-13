package config

var AllowedOriginsCORS = []string{"http://localhost:80", "http://127.0.0.1:80",
	"http://localhost", "http://127.0.0.1",
	"http://localhost:3000", "http://127.0.0.1:3000"}
var AllowedHeadersCORS = []string{"Version", "Authorization", "Content-Type", "csrf_token"}
var AllowedMethodsCORS = []string{"GET", "POST", "PUT", "DELETE"}