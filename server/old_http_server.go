package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/18scorpii/sample-go-app/vals"
)

var employeesMap map[string]vals.Employee
var err error

func StartGoHttpServer() {
	//load all employees from the file system
	employeesMap, err = vals.LoadEmployeesFromFile("./files/emp.json")
	if err != nil {
		log.Fatalf("Error in reading employees from file %v \n", err)
	}
	http.HandleFunc("/employees", EmployeesHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error in starting the web server %v \n", err)
	}
}

func EmployeesHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		showAllEmployeesList(rw)
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var employee vals.Employee
		err := decoder.Decode(&employee)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error in Parsing Post Data - %v", err), http.StatusInternalServerError)
			return
		}
		err = employee.IsValid()
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error in Validating Employee - %v", err), http.StatusInternalServerError)
			return
		}
		if _, ok := employeesMap[employee.Id]; ok {
			http.Error(rw, fmt.Sprintf("Employee Id Already Present - %v", employee.Id), http.StatusInternalServerError)
			return
		}
		employeesMap[employee.Id] = employee
		err = vals.SaveEmployeesToFile("./files/emp.json", &employeesMap)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		} else {
			showAllEmployeesList(rw)
		}
	} else if r.Method == http.MethodPut {
		decoder := json.NewDecoder(r.Body)
		var employee vals.Employee
		err := decoder.Decode(&employee)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error in Parsing Post Data - %v", err), http.StatusInternalServerError)
			return
		}
		err = employee.IsValid()
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error in Validating Employee - %v", err), http.StatusInternalServerError)
			return
		}
		if _, ok := employeesMap[employee.Id]; !ok {
			http.Error(rw, fmt.Sprintf("Employee Id Is Absent - %v", employee.Id), http.StatusInternalServerError)
			return
		}
		employee.Version = employee.Version + 1
		employeesMap[employee.Id] = employee
		err = vals.SaveEmployeesToFile("./files/emp.json", &employeesMap)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		} else {
			showAllEmployeesList(rw)
		}
	} else {
		sendErrorResponse(rw, "Only GET,POST Methods are supported", errors.New("Unsupported Methods"))
	}
}

func showAllEmployeesList(rw http.ResponseWriter) {
	employeeList := make([]vals.Employee, 0)
	for _, v := range employeesMap {
		employeeList = append(employeeList, v)
	}
	employeesJson, err := json.Marshal(employeeList)
	if err != nil {
		log.Fatalf("Error in marshalling employees to JSON %v \n", err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(employeesJson)
}
func sendErrorResponse(rw http.ResponseWriter, message string, err error) {
	errObj := ErrorObject{
		Message: message,
		Error:   err.Error(),
		Time:    time.Now().String(),
	}
	errObjJson, err := json.Marshal(errObj)
	if err != nil {
		log.Fatalf("Error in marshalling Error Object to JSON %v \n", err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Write(errObjJson)
}

type ErrorObject struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Time    string `json:"times"`
}
