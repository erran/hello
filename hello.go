package main

import (
	"fmt"
	"net/http"
  "encoding/json"
)

func main() {
	fmt.Printf("[hello] Starting server at 8080...\n")

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func badRequestHandler(resp http.ResponseWriter, req *http.Request, statusCode int) {
	header := resp.Header()
	header["Content-Type"] = []string{"application/json"}

  respMap := make(map[string]interface{})
  respMap["message"] = http.StatusText(statusCode)

  var scheme string
  if req.URL.Scheme == "" {
    scheme = "http"
  }

  // [todo] - link to written documentation
  respMap["documentation_url"] = fmt.Sprintf("%s://%s/", scheme, req.Host)

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

  respMap := make(map[string]interface{})
  respMap["path"] = req.URL.Path

  respJSON, _ := json.Marshal(respMap)
	fmt.Fprintf(resp, string(respJSON))
}
