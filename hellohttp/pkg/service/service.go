package service

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/hellohttp/pkg/util"
	"github.com/hellohttp/pkg/types"
)

//EmployeeServiceHandler - Employee service handler to call the http method for the request and response
type EmployeeServiceHandler interface {
	GetEmployee()
}

type service struct {
	client util.HTTPClient
}

func (s *service) GetEmployee() (*types.Employee, error) {
	return s.invokeMounteBank()
}

func (s *service) invokeMounteBank() (*types.Employee, error) {
	fmt.Println("Invoke Mounte Bank API endpoint")
	request, err := http.NewRequest(http.MethodGet, "http://localhost:4600/api/employee", nil)
	if err != nil {
		fmt.Println("Unable to construc the http request ", err)
	}
	response, err := s.client.Do(request)
	if err != nil {
		fmt.Println("Unable to process the http request ", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {
		var data types.Employee
		unmarshallErr := json.Unmarshal(body, &data)
		return &data, unmarshallErr
	} 
	return nil, err
}