package service

import (
	"fmt"
	"bytes"
	"testing"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/hellohttp/pkg/types"
	"github.com/stretchr/testify/assert"
)

type ClientMock struct {
	employee types.Employee
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	fmt.Println("Client do meothod invoked")
	payLoadInBytes, _ := json.Marshal(c.employee)
	w := ioutil.NopCloser(bytes.NewBuffer([]byte(payLoadInBytes)))
	res := &http.Response{Body: w, StatusCode: http.StatusOK}
	return res, nil
}

var expectedEmpResponse = &types.Employee{
	Fname : "John",
	Lname : "Smith",
	Age : 45,
	Gender: "M",
}

func TestGetEmployee(t *testing.T) {
	s := service {
		client: &ClientMock{employee: *expectedEmpResponse},
	}
	actualResponse, err := s.GetEmployee()
	assert.Nil(t, err)
	fmt.Println("Actual response ", actualResponse)
	assert.Equal(t, actualResponse.Fname, "John")
	assert.Equal(t, actualResponse.Lname, "Smith")
	assert.Equal(t, actualResponse.Age, 45)
	assert.Equal(t, actualResponse.Gender, "M")
}