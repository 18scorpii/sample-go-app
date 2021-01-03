package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/18scorpii/sample-go-app/vals"
	"github.com/gorilla/mux"
)

/*
	Static files are served using : http://localhost:8080/files/

*/
func StartHttpServer() {
	router := mux.NewRouter()
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	router.HandleFunc("/employees", EmployeesGetHandler).Methods("GET")
	router.HandleFunc("/employees", EmployeesPostHandler).Methods("POST")
	router.HandleFunc("/employees/{id}", EmployeesPostHandler).Methods("PUT")
	router.HandleFunc("/employees/{id}", EmployeesDeleteHandler).Methods("DELETE")
	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatalf("Error in starting web server %v \n", server.ListenAndServe())
}

func EmployeesGetHandler(rw http.ResponseWriter, r *http.Request) {
	employeesMap, err := vals.LoadEmployeesFromFile("./files/emp.json")
	if err != nil {
		log.Fatalf("Error in reading employees from file %v \n", err)
	}
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

func EmployeesPostHandler(rw http.ResponseWriter, r *http.Request) {
	var employee vals.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(rw, fmt.Sprintf("Error in Parsing Post Data - %v", err), http.StatusInternalServerError)
		return
	}
	if err := employee.IsValid(); err != nil {
		http.Error(rw, fmt.Sprintf("Error in Validating Employee - %v", err), http.StatusInternalServerError)
		return
	}
	//Check if Employee is Already Present
	employeesMap, err := vals.LoadEmployeesFromFile("./files/emp.json")
	if err != nil {
		log.Fatalf("Error in reading employees from file %v \n", err)
	}
	//check if its a update call, by looking for Path Param for IDs
	vars := mux.Vars(r)
	//Update Call Block
	if employeeId, ok := vars["id"]; ok {
		if oldValue, ok := employeesMap[employeeId]; !ok {
			http.Error(rw, fmt.Sprintf("Employee Id Is Not Present - %v", employee.Id), http.StatusNotFound)
			return
		} else {
			if oldValue.UpdateEmployee(employee) != nil {
				http.Error(rw, fmt.Sprintf("Employee Id Doesnot Match - %v", employee.Id), http.StatusInternalServerError)
				return
			}
			employeesMap[employeeId] = oldValue
		}
		//Insert Call Block, employee is a new entry
	} else {
		if _, ok := employeesMap[employee.Id]; ok {
			http.Error(rw, fmt.Sprintf("Employee Id Already Present - %v", employee.Id), http.StatusInternalServerError)
			return
		}
		employeesMap[employee.Id] = employee
	}
	//Rewrite the emp JSON file with all entries agains
	err = vals.SaveEmployeesToFile("./files/emp.json", &employeesMap)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	//Now reload and send output to response
	EmployeesGetHandler(rw, r)
}

func EmployeesDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	//Check if Employee is Already Present
	employeesMap, err := vals.LoadEmployeesFromFile("./files/emp.json")
	if err != nil {
		log.Fatalf("Error in reading employees from file %v \n", err)
	}
	//check if its a update call, by looking for Path Param for IDs
	vars := mux.Vars(r)
	var employeeId string
	if _, ok := vars["id"]; ok {
		employeeId = vars["id"]
		if _, ok := employeesMap[employeeId]; !ok {
			http.Error(rw, fmt.Sprintf("Employee Id Is Not Present - %v", employeeId), http.StatusNotFound)
			return
		}
	}
	//Rewrite the emp JSON file with all entries agains
	delete(employeesMap, employeeId)
	err = vals.SaveEmployeesToFile("./files/emp.json", &employeesMap)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	//Now reload and send output to response
	EmployeesGetHandler(rw, r)
}
