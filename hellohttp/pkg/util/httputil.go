package util

import (
   "net/http"
   "time"
)

//HTTPClient - For the request and response client
type HTTPClient interface {
   Do(req *http.Request) (*http.Response, error)
}

type httpclient struct {
   timeout int
}

//NewHTTPClient - To create a new instance of the http client
func NewHTTPClient() *httpclient{
	return &httpclient{}
}

//Do - 
func (httpClient *httpclient) Do(request *http.Request) (*http.Response, error) {
   client := &http.Client{
      Timeout: time.Second * time.Duration(httpClient.timeout),
   }
   return client.Do(request)
}