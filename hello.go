package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("[hello] Starting server at 8080...\n")

	http.HandleFunc("/", rootHandler)

	// Capturing the trailing stating
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/test/", testHandler)

	http.ListenAndServe(":8080", nil)
}

func scheme(req *http.Request) (scheme string) {
	var schemeString string
	if req.URL.Scheme == "" {
		schemeString = "http"
	}

	return schemeString
}

func badRequestHandler(resp http.ResponseWriter, req *http.Request, statusCode int) {
	header := resp.Header()
	header["Content-Type"] = []string{"application/json"}
	resp.WriteHeader(statusCode)

	respMap := make(map[string]string)
	respMap["message"] = http.StatusText(statusCode)

	// [todo] - link to written documentation
	respMap["documentation_url"] = fmt.Sprintf("%s://%s/", scheme(req), req.Host)

	respJSON, _ := json.Marshal(respMap)
	fmt.Fprintf(resp, string(respJSON))
}

func rootHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		badRequestHandler(resp, req, http.StatusNotFound)
		return
	}

	header := resp.Header()
	header["Content-Type"] = []string{"application/json"}

	respMap := make(map[string]string)
	respMap["test_url"] = fmt.Sprintf("%s://%s/test", scheme(req), req.Host)

	respJSON, _ := json.Marshal(respMap)
	fmt.Fprintf(resp, string(respJSON))
}

func testHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/test" && req.URL.Path != "/test/" {
		badRequestHandler(resp, req, http.StatusNotFound)
		return
	}

	header := resp.Header()
	header["Content-Type"] = []string{"application/json"}

	respMap := make(map[string]string)
	respMap["foo"] = "bar"

	respJSON, _ := json.Marshal(respMap)
	fmt.Fprintf(resp, string(respJSON))
}
