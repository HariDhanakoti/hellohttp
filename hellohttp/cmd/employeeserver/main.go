package main

import (
	"fmt"
	"github.com/hellohttp/pkg/service"
)

func main() {
	fmt.Println("Main program started")
	fmt.Println("Invoking the /api/employee endpoint")
	employee, err := service.GetEmployee()
	fmt.Println("Response from /api/employee endpoint ", employee)
}
